package executor

import (
	"context"
	"errors"
	"fmt"
	"math/big"
	"sync"

	"github.com/dittonetwork/executor-avs/cmd/operator/internal/models"
	"github.com/dittonetwork/executor-avs/pkg/log"
	"github.com/ethereum/go-ethereum"
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
}

//go:generate mockery --name dittoEntryPoint --output ./mocks --outpkg mocks
type dittoEntryPoint interface {
	GetAllActiveWorkflows(ctx context.Context) ([]models.Workflow, error)
	IsExecutor(ctx context.Context) (bool, error)
	IsValidExecutor(ctx context.Context, blockNumber *big.Int) (bool, error)
	GetRunWorkflowTx(ctx context.Context, vaultAddr common.Address, workflowID *big.Int) (*types.Transaction, error)
	RunMultipleWorkflows(ctx context.Context, workflows []models.Workflow) (*types.Transaction, error)
}

type Executor struct {
	Client     ethereumClient
	EntryPoint dittoEntryPoint
}

func NewExecutor(client ethereumClient, entryPoint dittoEntryPoint) *Executor {
	return &Executor{
		Client:     client,
		EntryPoint: entryPoint,
	}
}

func (r *Executor) SubscribeToNewBlocks(ctx context.Context) (chan *types.Header, ethereum.Subscription, error) {
	return r.Client.SubscribeNewHead(ctx)
}

func (r *Executor) Handle(ctx context.Context, blockHash common.Hash) error {
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
		log.Info("Nothing to execute")

		return nil
	}

	if err = r.executeWorkflows(ctx, executableWorkflows); err != nil {
		return fmt.Errorf("execute workflows: %w", err)
	}

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

	log.With(log.Bool("result", isSuccess)).Debug("Simulation done")

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

func (r *Executor) executeWorkflows(ctx context.Context, workflows []models.Workflow) error {
	tx, err := r.EntryPoint.RunMultipleWorkflows(ctx, workflows)
	if err != nil {
		return fmt.Errorf("run multiple workflows: %w", err)
	}

	if err = r.Client.SendTransaction(ctx, tx); err != nil {
		return fmt.Errorf("send transaction: %w", err)
	}

	log.With(log.String("tx_hash", tx.Hash().String())).Info("run multiple workflows")

	return nil
}
