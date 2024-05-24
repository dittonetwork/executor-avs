package executor

import (
	"context"
	"errors"
	"fmt"
	"math/big"
	"sync"

	"github.com/ethereum/go-ethereum/core/types"

	"github.com/dittonetwork/executor-avs/cmd/operator/internal/models"
	"github.com/dittonetwork/executor-avs/pkg/log"
)

var (
	ErrBlockIsNil           = errors.New("block is nil")
	ErrUnregisteredExecutor = errors.New("executor is not registered")
)

type Executor struct {
	client     EthereumClient
	entryPoint DittoEntryPoint
}

func NewExecutor(client EthereumClient, entryPoint DittoEntryPoint) *Executor {
	return &Executor{
		client:     client,
		entryPoint: entryPoint,
	}
}

func (r *Executor) Handle(ctx context.Context, block *types.Block) error {
	if block == nil {
		return ErrBlockIsNil
	}

	isExecutor, err := r.entryPoint.IsExecutor(ctx)
	if err != nil {
		return fmt.Errorf("check if is executor: %w", err)
	}

	if !isExecutor {
		return ErrUnregisteredExecutor
	}

	var isValidExecutor bool
	isValidExecutor, err = r.entryPoint.IsValidExecutor(ctx, block.Number())
	if err != nil {
		return fmt.Errorf("check if executor is valid: %w", err)
	}

	if !isValidExecutor {
		log.Info("isValidExecutor is false")

		return nil
	}

	workflows, err := r.entryPoint.GetAllActiveWorkflows(ctx)
	if err != nil {
		return fmt.Errorf("get all active workflows: %w", err)
	}

	var wg sync.WaitGroup
	errCh := make(chan error, len(workflows))
	wg.Add(len(workflows))

	for _, workflow := range workflows {
		log.With(
			log.String("vault_addr", workflow.VaultAddress.String()),
			log.String("workflow_id", workflow.WorkflowID.String()),
		).Info("active workflow")

		go func(workflow models.Workflow) {
			defer wg.Done()

			var canRun bool
			canRun, err = r.Simulate(ctx, workflow, block.Number())
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

	if err = r.ExecuteWorkflows(ctx, executableWorkflows); err != nil {
		return fmt.Errorf("execute workflows: %w", err)
	}

	return nil
}

func hexStringToBool(hexStr string) (bool, error) {
	value := new(big.Int)
	_, isSuccess := value.SetString(hexStr, 0)

	if !isSuccess {
		return false, fmt.Errorf("invalid hex format: %s", hexStr)
	}

	return value.Cmp(big.NewInt(0)) != 0, nil
}

func (r *Executor) Simulate(ctx context.Context, workflow models.Workflow, blockNum *big.Int) (bool, error) {
	tx, err := r.entryPoint.GetRunWorkflowTx(ctx, workflow.VaultAddress, workflow.WorkflowID)
	if err != nil {
		return false, fmt.Errorf("run workflow: %w", err)
	}

	var resp string
	err = r.client.SimulateTransaction(ctx, tx, blockNum, &resp)
	if err != nil {
		return false, fmt.Errorf("simulate transaction: %w", err)
	}

	isSuccess, err := hexStringToBool(resp)
	if err != nil {
		return false, fmt.Errorf("interpreting result: %w", err)
	}
	log.With(
		log.Bool("result", isSuccess),
	).Debug("Simulation done")

	return isSuccess, nil
}

func (r *Executor) ExecuteWorkflows(ctx context.Context, workflows []models.Workflow) error {
	tx, err := r.entryPoint.RunMultipleWorkflows(ctx, workflows)
	if err != nil {
		return fmt.Errorf("run multiple workflows: %w", err)
	}

	if err = r.client.SendTransaction(ctx, tx); err != nil {
		return fmt.Errorf("send transaction: %w", err)
	}

	log.With(log.String("tx_hash", tx.Hash().String())).Info("run multiple workflows")

	return nil
}

func (r *Executor) CheckIsExecutor(ctx context.Context) (bool, error) {
	isValid, err := r.entryPoint.IsExecutor(ctx)
	if err != nil {
		return false, fmt.Errorf("check is executor: %w", err)
	}

	return isValid, nil
}
