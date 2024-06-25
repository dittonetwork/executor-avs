package executor

import (
	"context"
	"errors"
	"fmt"
	"math/big"
	"sync"
	"time"

	portdep "github.com/dittonetwork/executor-avs/cmd/operator/internal/ports/dep"

	"github.com/dittonetwork/executor-avs/cmd/operator/internal/models"
	"github.com/dittonetwork/executor-avs/pkg/log"
	"github.com/dittonetwork/executor-avs/pkg/primitives"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
)

var (
	ErrBlockIsNil           = errors.New("block is nil")
	ErrUnregisteredExecutor = errors.New("executor is not registered")
)

//go:generate mockery --name ethereumClient --output ./mocks --outpkg mocks
type ethereumClient interface {
	SubscribeNewHead(ctx context.Context) (chan *types.Header, ethereum.Subscription, error)
	BlockByHash(ctx context.Context, hash common.Hash) (*types.Block, error)
	SimulateTransaction(ctx context.Context, tx *types.Transaction, blockNum *big.Int) (string, error)
	SendTransaction(ctx context.Context, tx *types.Transaction) error
	GetBalance(ctx context.Context) (*big.Int, error)
	TransactionReceipt(ctx context.Context, txHash common.Hash) (*types.Receipt, error)
	CodeAt(ctx context.Context, account common.Address, blockNumber *big.Int) ([]byte, error)
}

//go:generate mockery --name dittoEntryPoint --output ./mocks --outpkg mocks
type dittoEntryPoint interface {
	GetAllActiveWorkflows(ctx context.Context) ([]models.Workflow, error)
	IsExecutor(ctx context.Context) (bool, error)
	IsValidExecutor(ctx context.Context, blockNumber *big.Int) (bool, error)
	GetRunWorkflowTx(ctx context.Context, vaultAddr common.Address, workflowID *big.Int) (*types.Transaction, error)
	RunMultipleWorkflows(ctx context.Context, workflows []models.Workflow, estimatedGasMultiplier float64) (*types.Transaction, error)
	DeactivateExecutor(ctx context.Context) (*types.Transaction, error)
	ActivateExecutor(ctx context.Context) (*types.Transaction, error)
	GetSucceededWorkflows(logs []*types.Log) ([]models.Workflow, error)
}

type Executor struct {
	Client     ethereumClient
	EntryPoint dittoEntryPoint

	metrics        *Metrics
	workflowsCache *WorkflowsCache
}

type Options func(*Executor)

func WithMetrics() Options {
	return func(e *Executor) {
		e.metrics.Register()
		go e.metrics.CollectBackgroundMetrics(e.Client)
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

// Atomically checks and adds non-existing workflows to the cache.
func (awc *WorkflowsCache) ConditionalAcquireWorkflowsLock(workflows []models.Workflow) []models.Workflow {
	// we could lock mutex for each workflow insertion/deletion, but current implementation won't
	//   degrade execution anyway, and it is easier to debug if concurrent calls had place.
	awc.mu.Lock()
	defer awc.mu.Unlock()

	var inserted []models.Workflow
	for _, workflow := range workflows {
		address := workflow.VaultAddress
		workflowID := workflow.WorkflowID.String()

		if _, exists := awc.data[address]; !exists {
			awc.data[address] = make(map[string]bool)
		}

		if _, exists := awc.data[address][workflowID]; !exists {
			awc.data[address][workflowID] = true
			inserted = append(inserted, workflow)
		}
	}
	return inserted
}

func (awc *WorkflowsCache) ReleaseWorkflowsLock(workflows []models.Workflow) {
	awc.mu.Lock()
	defer awc.mu.Unlock()

	for _, workflow := range workflows {
		address := workflow.VaultAddress
		workflowID := workflow.WorkflowID.String()

		if workflows, exists := awc.data[address]; exists {
			delete(workflows, workflowID)
			if len(workflows) == 0 {
				delete(awc.data, address)
			}
		}
	}
}

func NewExecutor(client ethereumClient, entryPoint dittoEntryPoint, opts ...Options) *Executor {
	e := &Executor{
		Client:         client,
		EntryPoint:     entryPoint,
		metrics:        NewMetrics(),
		workflowsCache: NewWorkflowsCache(),
	}

	for _, opt := range opts {
		opt(e)
	}

	return e
}

func (r *Executor) SubscribeToNewBlocks(ctx context.Context) (chan *types.Header, ethereum.Subscription, error) {
	return r.Client.SubscribeNewHead(ctx)
}

func (r *Executor) Handle(ctx context.Context, blockHash common.Hash) error {
	if err := r.handle(ctx, blockHash); err != nil {
		r.metrics.CountErrorsTotal(1)

		return err
	}

	return nil
}

func (r *Executor) Activate(ctx context.Context) error {
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

func (r *Executor) Deactivate(ctx context.Context) error {
	tx, err := r.EntryPoint.DeactivateExecutor(ctx)
	if err != nil {
		return fmt.Errorf("deactivate executor: %w", err)
	}

	log.With(log.String("tx_hash", tx.Hash().Hex())).Info("Deactivate transaction created")

	r.waitForTransaction(ctx, tx)

	return nil
}

func (r *Executor) waitForTransaction(_ context.Context, tx *types.Transaction) {
	const pollIntervalSecs = 5

	log.Info("Waiting for transaction to complete...")
	for {
		receipt, err := r.Client.TransactionReceipt(context.Background(), tx.Hash())
		if receipt != nil {
			log.With(
				log.String("tx_hash", receipt.TxHash.Hex()),
				log.String("block_hash", receipt.BlockHash.Hex()),
			).Info("Transaction included into block ")
		}
		if err != nil {
			log.Info("Transaction receipt not available yet, waiting...")
			time.Sleep(pollIntervalSecs * time.Second)
		}
	}
}

func (r *Executor) handle(ctx context.Context, blockHash common.Hash) error {
	block, err := r.Client.BlockByHash(ctx, blockHash)
	if err != nil {
		return fmt.Errorf("get block by hash: %w", err)
	}

	if block.Number() == nil {
		return ErrBlockIsNil
	}

	log.With(log.Int64("block_number", block.Number().Int64())).Info("checking if it is executor...")

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

	executableWorkflows, err := r.getExecutableWorkflows(ctx, block.Number())
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

	r.metrics.CountExecutedWorkflowsAmountTotal(len(succeededWorkflows))

	return nil
}

func (r *Executor) getExecutableWorkflows(ctx context.Context, blockNum *big.Int) ([]models.Workflow, error) {
	workflows, err := r.EntryPoint.GetAllActiveWorkflows(ctx)
	if err != nil {
		return nil, fmt.Errorf("get all active workflows: %w", err)
	}

	var wg sync.WaitGroup
	errCh := make(chan error, len(workflows))
	wg.Add(len(workflows))

	for i := range workflows {
		workflow := &workflows[i]
		log.With(
			log.String("vault_addr", workflow.VaultAddress.String()),
			log.String("workflow_id", workflow.WorkflowID.String()),
		).Debug("active workflow")

		go func(workflow *models.Workflow) {
			defer wg.Done()

			var canRun bool
			canRun, err = r.simulate(ctx, *workflow, blockNum)
			if err != nil {
				errCh <- fmt.Errorf("simulate workflow %d: %w", workflow.WorkflowID.Int64(), err)
				return
			}
			log.With(
				log.String("vault_addr", workflow.VaultAddress.Hex()),
				log.String("workflow_id", workflow.WorkflowID.String()),
				log.Bool("result", canRun),
			).Debug("Simulation done")
			workflow.CouldBeExecuted = canRun
		}(workflow)
	}

	wg.Wait()
	close(errCh)

	for err := range errCh {
		if err != nil {
			log.With(log.Err(err)).Error("simulate workflow error")
		}
	}

	var executableWorkflows []models.Workflow
	for _, workflow := range workflows {
		if workflow.CouldBeExecuted {
			executableWorkflows = append(executableWorkflows, workflow)
		}
	}

	return executableWorkflows, nil
}

func (r *Executor) simulate(ctx context.Context, workflow models.Workflow, blockNum *big.Int) (bool, error) {
	tx, err := r.EntryPoint.GetRunWorkflowTx(ctx, workflow.VaultAddress, workflow.WorkflowID)
	if err != nil {
		if errors.Is(err, portdep.ErrExecutionReverted) {
			return false, nil
		}

		return false, fmt.Errorf("run workflow: %w", err)
	}

	var resp string
	resp, err = r.Client.SimulateTransaction(ctx, tx, blockNum)
	if err != nil {
		return false, fmt.Errorf("simulate transaction: %w", err)
	}

	isSuccess, err := hexStringToBool(resp)
	if err != nil {
		return false, fmt.Errorf("interpreting result: %w", err)
	}

	return isSuccess, nil
}

func hexStringToBool(hexStr string) (bool, error) {
	value := new(big.Int)
	_, isSuccess := value.SetString(hexStr, 0)

	if !isSuccess {
		return false, fmt.Errorf("invalid hex format: %s", hexStr)
	}

	return value.Cmp(big.NewInt(0)) != 0, nil
}

func (r *Executor) executeWorkflows(ctx context.Context, workflows []models.Workflow) ([]models.Workflow, error) {
	tx, err := r.EntryPoint.RunMultipleWorkflows(ctx, workflows, 1.2)
	if err != nil {
		return nil, fmt.Errorf("run multiple workflows: %w", err)
	}

	log.With(
		log.String("tx_hash", tx.Hash().String()),
		log.Uint64("gas_used", tx.Gas()),
		log.Uint64("gas_price", tx.GasPrice().Uint64()),
		log.Uint64("native amount", tx.Gas()*tx.GasPrice().Uint64()),
	).Debug("RunMultipleWorkflows")

	spentAmount := new(big.Int).Mul(new(big.Int).SetUint64(tx.Gas()), tx.GasPrice())

	r.metrics.CountNativeTokenSpentAmountTotal(primitives.WeiToETH(spentAmount))

	receipt, err := bind.WaitMined(ctx, r.Client, tx)
	if err != nil {
		return nil, fmt.Errorf("waiting for RunMultipleWorkflows tx: %w", err)
	}
	log.With(log.String("tx_hash", tx.Hash().String())).Debug("waiting for transaction to appear on chain")
	if receipt.Status == types.ReceiptStatusFailed {
		return nil, fmt.Errorf("RunMultipleWorkflows failed on chain, status: %d", receipt.Status)
	}
	log.With(log.String("tx_hash", tx.Hash().String())).Debug("transaction is on chain")

	succeededWorkflows, err := r.EntryPoint.GetSucceededWorkflows(receipt.Logs)
	if err != nil {
		return nil, fmt.Errorf("GetSucceededWorkflows: %w", err)
	}

	return succeededWorkflows, nil
}
