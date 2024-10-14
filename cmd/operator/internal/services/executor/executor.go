package executor

import (
	"context"
	"errors"
	"fmt"
	"math/big"
	"sync"
	"time"

	"github.com/dittonetwork/executor-avs/cmd/operator/internal/adapters/dittoentrypoint"
	"github.com/dittonetwork/executor-avs/cmd/operator/internal/models"
	"github.com/dittonetwork/executor-avs/pkg/log"
	"github.com/dittonetwork/executor-avs/pkg/primitives"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	lru "github.com/hashicorp/golang-lru/v2"
	"github.com/prometheus/client_golang/prometheus"
)

var (
	ErrBlockIsNil           = errors.New("block is nil")
	ErrUnregisteredExecutor = errors.New("executor is not registered")
)

//go:generate mockery --name Executor --output ./mocks --outpkg mocks
type Executor interface {
	LongPollingLoop(ctx context.Context) error
	IsAutoDeactivate() bool
	ProcessBlock(ctx context.Context, blockHash common.Hash) error
	Activate(ctx context.Context) error
	Deactivate(ctx context.Context) error
}

//go:generate mockery --name EthClient --output ./mocks --outpkg mocks
type EthClient interface {
	BlockByHash(ctx context.Context, hash common.Hash) (*types.Block, error)
	HeaderByNumber(ctx context.Context, number *big.Int) (*types.Header, error)
	TransactionReceipt(ctx context.Context, txHash common.Hash) (*types.Receipt, error)
	CodeAt(ctx context.Context, account common.Address, blockNumber *big.Int) ([]byte, error)
}

type Impl struct {
	Client     EthClient
	EntryPoint dittoentrypoint.DittoEntryPoint

	metrics              *Metrics
	workflowsCache       *WorkflowsCache
	autoDeactivate       bool
	ignoreEndBlocks      *big.Int
	blockHandlingTimeout time.Duration
	longPollInterval     time.Duration
}

var _ Executor = new(Impl)

type Options func(*Impl)

func WithMetrics() Options {
	return func(e *Impl) {
		e.metrics.Register()
	}
}

func WithCustomLiveCycle(autoDeactivate bool) Options {
	return func(e *Impl) {
		e.autoDeactivate = autoDeactivate
	}
}

type WorkflowsCache struct {
	data map[common.Address]map[string]bool
	mu   sync.Mutex
}

func NewWorkflowsCache() *WorkflowsCache {
	return &WorkflowsCache{
		data: make(map[common.Address]map[string]bool),
	}
}

// ConditionalAcquireWorkflowsLock Atomically checks and adds non-existing workflows to the cache.
// It is done to prevent concurrent checks of the same workflow by different blocks handlers. Since every handler
// waits for receipts, lack of such exclution could lead to significant duplication.
func (wc *WorkflowsCache) ConditionalAcquireWorkflowsLock(workflows []models.Workflow) []models.Workflow {
	// we could lock mutex for each workflow insertion/deletion, but current implementation won't
	//   degrade execution anyway, and it is easier to debug if concurrent calls had place.
	wc.mu.Lock()
	defer wc.mu.Unlock()

	var inserted []models.Workflow
	for _, workflow := range workflows {
		address := workflow.VaultAddress
		workflowID := workflow.WorkflowID.String()

		if _, exists := wc.data[address]; !exists {
			wc.data[address] = make(map[string]bool)
		}

		if _, exists := wc.data[address][workflowID]; !exists {
			wc.data[address][workflowID] = true
			inserted = append(inserted, workflow)
		}
	}
	return inserted
}

func (wc *WorkflowsCache) ReleaseWorkflowsLock(workflows []models.Workflow) {
	wc.mu.Lock()
	defer wc.mu.Unlock()

	for _, workflow := range workflows {
		address := workflow.VaultAddress
		workflowID := workflow.WorkflowID.String()

		if workflows, exists := wc.data[address]; exists {
			delete(workflows, workflowID)
			if len(workflows) == 0 {
				delete(wc.data, address)
			}
		}
	}
}

func NewExecutor(
	client EthClient,
	entryPoint dittoentrypoint.DittoEntryPoint,
	ignoreEndBlocks *big.Int,
	blockHandlingTimeout time.Duration,
	longPollInterval time.Duration,
	opts ...Options,
) *Impl {
	e := &Impl{
		Client:               client,
		EntryPoint:           entryPoint,
		metrics:              NewMetrics(),
		workflowsCache:       NewWorkflowsCache(),
		ignoreEndBlocks:      ignoreEndBlocks,
		blockHandlingTimeout: blockHandlingTimeout,
		longPollInterval:     longPollInterval,
	}

	for _, opt := range opts {
		opt(e)
	}

	return e
}

func (r *Impl) LongPollingLoop(ctx context.Context) error {
	// Since we use long polling, the same block will be read multiple times between tick. Keep blocks
	// currently being processed in cache.
	const cacheSize = 128
	cache, cacheErr := lru.New[common.Hash, bool](cacheSize)
	if cacheErr != nil {
		return cacheErr
	}

	var wg sync.WaitGroup
	ticker := time.NewTicker(r.longPollInterval)
	defer ticker.Stop()

	func() {
		for {
			select {
			case <-ticker.C:
				processingCtx, cancel := context.WithTimeout(ctx, r.blockHandlingTimeout)
				defer cancel()

				header, err := r.Client.HeaderByNumber(processingCtx, nil)
				if err != nil {
					log.With(log.Err(err)).Error("getting last header")
					continue
				}

				_, ok := cache.Get(header.Hash())
				if ok {
					log.Debug("block already in processing cache, skipping")
					continue
				}

				cache.Add(header.Hash(), true)
				wg.Add(1)
				go func(b *types.Header) {
					defer wg.Done()
					if err = r.ProcessBlock(processingCtx, b.Hash()); err != nil {
						log.With(log.Err(err)).Error("processing block")
						cache.Remove(b.Hash())
					}
				}(header)

			case <-ctx.Done():
				return
			}
		}
	}()

	wg.Wait()
	return nil
}

func (r *Impl) IsAutoDeactivate() bool {
	return r.autoDeactivate
}

func (r *Impl) ProcessBlock(ctx context.Context, blockHash common.Hash) error {
	processingTimer := prometheus.NewTimer(r.metrics.blockProcessingDurationSeconds)
	defer processingTimer.ObserveDuration()

	if err := r.processBlock(ctx, blockHash); err != nil {
		r.metrics.errorsTotal.Add(1)
		return err
	}

	return nil
}

func (r *Impl) Activate(ctx context.Context) error {
	if isExecutor, err := r.EntryPoint.IsExecutor(ctx); err != nil {
		return fmt.Errorf("check if is executor: %w", err)
	} else if isExecutor {
		log.Info("Executor is already activated")
		return nil
	}

	tx, err := r.EntryPoint.ActivateExecutor(ctx)
	if err != nil {
		return fmt.Errorf("activate executor: %w", err)
	}

	log.With(log.String("tx_hash", tx.Hash().Hex())).Info("Activate transaction created")

	r.waitForTransaction(ctx, tx)

	return nil
}

func (r *Impl) Deactivate(ctx context.Context) error {
	tx, err := r.EntryPoint.DeactivateExecutor(ctx)
	if err != nil {
		return fmt.Errorf("deactivate executor: %w", err)
	}

	log.With(log.String("tx_hash", tx.Hash().Hex())).Info("Deactivate transaction created")

	r.waitForTransaction(ctx, tx)

	return nil
}

func (r *Impl) waitForTransaction(ctx context.Context, tx *types.Transaction) {
	const pollIntervalSecs = 5

	log.Info("Waiting for transaction to complete...")
	for {
		receipt, err := r.Client.TransactionReceipt(ctx, tx.Hash())
		if receipt != nil {
			log.With(
				log.String("tx_hash", receipt.TxHash.Hex()),
				log.String("block_hash", receipt.BlockHash.Hex()),
			).Info("Transaction included into block ")
			return
		}
		if err != nil {
			log.Info("Transaction receipt not available yet, waiting...")
			time.Sleep(pollIntervalSecs * time.Second)
		}
	}
}

func (r *Impl) processBlock(ctx context.Context, blockHash common.Hash) error {
	block, err := r.Client.BlockByHash(ctx, blockHash)
	if err != nil {
		return fmt.Errorf("get block by hash: %w", err)
	}

	if block.Number() == nil {
		return ErrBlockIsNil
	}

	log.With(
		log.Int64("block_number", block.Number().Int64()),
		log.String("block_hash", blockHash.Hex()),
	).Info("processing block")

	isExecutor, err := r.EntryPoint.IsExecutor(ctx)
	if err != nil {
		return fmt.Errorf("check if is executor: %w", err)
	}

	if !isExecutor {
		log.Info("not executor ❌")

		return ErrUnregisteredExecutor
	}

	var isValidExecutor bool
	isValidExecutor, err = r.EntryPoint.IsValidExecutor(ctx, block.Number())
	if err != nil {
		return fmt.Errorf("check if executor is valid: %w", err)
	}

	if !isValidExecutor {
		log.Info("Not my turn to execute")
		return nil
	}

	willBeValidExecutor, err := r.EntryPoint.IsValidExecutor(ctx, new(big.Int).Add(block.Number(), r.ignoreEndBlocks))
	if err != nil {
		return fmt.Errorf("check if executor is valid in future: %w", err)
	}

	if !willBeValidExecutor {
		log.Info("Almost end of my turn, not gonna execute")
		return nil
	}

	executableWorkflows, err := r.getExecutableWorkflows(ctx)
	if err != nil {
		return fmt.Errorf("get executable workflows: %w", err)
	}
	if len(executableWorkflows) == 0 {
		log.Info("No workflows met execution condition")
		return nil
	}

	acquiredWorkflows := r.workflowsCache.ConditionalAcquireWorkflowsLock(executableWorkflows)
	defer r.workflowsCache.ReleaseWorkflowsLock(acquiredWorkflows)

	log.With(
		log.Int("execution_candidates", len(executableWorkflows)),
		log.Int("after_filtering", len(acquiredWorkflows)),
	).Info("Filtered already cached worflows from execution candidates")

	if len(acquiredWorkflows) == 0 {
		log.Info("No workflows to execute after filtering")
		return nil
	}

	succeededWorkflows, err := r.executeWorkflows(ctx, acquiredWorkflows)
	if err != nil {
		return fmt.Errorf("execute workflows: %w", err)
	}
	log.With(
		log.Int("succeeded_workflows", len(succeededWorkflows)),
		log.Int("were_sent_workflows", len(acquiredWorkflows)),
	).Debug("Workflows executed")

	r.metrics.sentWorkflowsAmountTotal.Add(float64(len(acquiredWorkflows)))
	r.metrics.executedWorkflowsAmountTotal.Add(float64(len(succeededWorkflows)))

	return nil
}

func (r *Impl) getExecutableWorkflows(ctx context.Context) ([]models.Workflow, error) {
	workflows, getWorkflowsError := r.EntryPoint.GetAllActiveWorkflows(ctx)
	if getWorkflowsError != nil {
		return nil, getWorkflowsError
	}

	const batchSize = 150 // block gas limit / aprox wf cost = 30.000.000 / 200.000 = 150
	var successfulWorkflows []models.Workflow

	for i := 0; i < len(workflows); i += batchSize {
		end := i + batchSize
		if end > len(workflows) {
			end = len(workflows)
		}
		batch := workflows[i:end]

		executableBatch, err := r.findExecutableBatch(ctx, batch)
		if err != nil {
			return nil, err
		}

		successfulWorkflows = append(successfulWorkflows, executableBatch...)
	}

	// Mb after merging into a single batch transactions will fail. Final check.
	executionResults, err := r.EntryPoint.SimulateMutlipleWorkflows(ctx, successfulWorkflows)
	if err != nil {
		return nil, err
	}

	successfulWorkflows = successfulWorkflows[:0]
	for i, result := range executionResults {
		if result {
			successfulWorkflows = append(successfulWorkflows, workflows[i])
		}
	}

	for _, workflow := range successfulWorkflows {
		log.With(
			log.String("vault_addr", workflow.VaultAddress.String()),
			log.String("workflow_id", workflow.WorkflowID.String()),
		).Debug("succeeded workflow")
	}

	return successfulWorkflows, nil
}

func (r *Impl) findExecutableBatch(ctx context.Context, batch []models.Workflow) ([]models.Workflow, error) {
	executionResults, simulationError := r.EntryPoint.SimulateMutlipleWorkflows(ctx, batch)
	if simulationError != nil {
		return r.handleSimulationError(ctx, batch, simulationError)
	}

	return r.filterExecutableBatch(batch, executionResults), nil
}

func (r *Impl) handleSimulationError(
	ctx context.Context,
	batch []models.Workflow,
	simulationError error,
) ([]models.Workflow, error) {
	log.With(log.Err(simulationError)).Debug("error during batch simulation")

	if !isGasLimitError(simulationError) {
		return nil, simulationError
	}

	if len(batch) == 1 {
		return nil, nil // Skip if we can't execute even a single workflow
	}

	return r.splitAndRecurseBatch(ctx, batch)
}

func (r *Impl) splitAndRecurseBatch(ctx context.Context, batch []models.Workflow) ([]models.Workflow, error) {
	const divisorForMidpoint = 2
	mid := len(batch) / divisorForMidpoint

	left, err := r.findExecutableBatch(ctx, batch[:mid])
	if err != nil {
		return nil, err
	}

	right, err := r.findExecutableBatch(ctx, batch[mid:])
	if err != nil {
		return nil, err
	}

	return append(left, right...), nil
}

func (r *Impl) filterExecutableBatch(batch []models.Workflow, executionResults []bool) []models.Workflow {
	var executableBatch []models.Workflow
	for i, result := range executionResults {
		if result {
			executableBatch = append(executableBatch, batch[i])
		}
	}
	return executableBatch
}

func isGasLimitError(_ error) bool {
	// manually created huge batch, error was: "invalid character 'c' looking for beginning of value",
	// since wss RPC returns "content length too large (6291774>5242880)".
	// TODO: figure out which error is actually thrown in case of gas limit
	//   (mb it is not rised during simulation calls at all)
	return false
}

func (r *Impl) executeWorkflows(ctx context.Context, workflows []models.Workflow) ([]models.Workflow, error) {
	gasLimitMultiplier := 1.5
	tx, err := r.EntryPoint.RunMultipleWorkflows(ctx, workflows, gasLimitMultiplier)
	if err != nil {
		return nil, fmt.Errorf("run multiple workflows: %w", err)
	}
	miningTimer := prometheus.NewTimer(r.metrics.miningLatencySeconds)

	log.With(
		log.String("tx_hash", tx.Hash().String()),
		log.Uint64("gas_used", tx.Gas()),
		log.Uint64("gas_price", tx.GasPrice().Uint64()),
		log.Uint64("native amount", tx.Gas()*tx.GasPrice().Uint64()),
	).Debug("RunMultipleWorkflows")

	spentAmount := new(big.Int).Mul(new(big.Int).SetUint64(tx.Gas()), tx.GasPrice())

	r.metrics.nativeTokenSpentAmountTotal.Add(primitives.WeiToETH(spentAmount))

	log.With(log.String("tx_hash", tx.Hash().String())).Debug("waiting for transaction to appear on chain")
	receipt, err := bind.WaitMined(ctx, r.Client, tx)
	if err != nil {
		return nil, fmt.Errorf("waiting for RunMultipleWorkflows tx: %w", err)
	}
	miningTimer.ObserveDuration()
	log.With(log.String("tx_hash", tx.Hash().String())).Debug("transaction is on chain")

	if receipt.Status == types.ReceiptStatusFailed {
		return nil, fmt.Errorf("RunMultipleWorkflows failed on chain, status: %d", receipt.Status)
	}

	succeededWorkflows, err := r.EntryPoint.GetSucceededWorkflows(receipt.Logs)
	if err != nil {
		return nil, fmt.Errorf("GetSucceededWorkflows: %w", err)
	}

	return succeededWorkflows, nil
}
