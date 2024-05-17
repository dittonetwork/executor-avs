package executor

import (
	"context"
	"errors"
	"fmt"
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

	contractAddress string
}

func NewExecutor(client EthereumClient, entryPoint DittoEntryPoint, contractAddress string) *Executor {
	return &Executor{
		contractAddress: contractAddress,
		client:          client,
		entryPoint:      entryPoint,
	}
}

func (r *Executor) Handle(ctx context.Context, block *types.Block) error {
	if block == nil {
		return ErrBlockIsNil
	}

	// log.Info("block", log.String("number", block.Number().String()))

	isExecutor, err := r.entryPoint.IsExecutor(ctx, r.contractAddress)
	if err != nil {
		return fmt.Errorf("check if is executor: %w", err)
	}

	if !isExecutor {
		return ErrUnregisteredExecutor
	}

	var isValidExecutor bool
	isValidExecutor, err = r.entryPoint.IsValidExecutor(ctx, block.Number().Int64(), r.contractAddress)
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

	executableWorkflows := make([]models.Workflow, 0, len(workflows))

	for _, workflow := range workflows {
		go func(workflow models.Workflow) {
			defer wg.Done()

			var canRun bool
			canRun, err = r.Simulate(ctx, workflow)
			if err != nil {
				errCh <- fmt.Errorf("simulate workflow: %w", err)
			}

			if canRun {
				executableWorkflows = append(executableWorkflows, workflow)
			}
		}(workflow)
	}

	wg.Wait()
	close(errCh)

	for err := range errCh {
		if err != nil {
			log.With(log.Err(err)).Error("simulate workflow error")
		}
	}

	for _, workflow := range executableWorkflows {
		if err = r.Run(ctx, workflow); err != nil {
			return fmt.Errorf("execute workflow: %w", err)
		}
	}

	return nil
}

func (r *Executor) Simulate(ctx context.Context, workflow models.Workflow) (bool, error) {
	return false, nil
}

func (r *Executor) Run(ctx context.Context, workflow models.Workflow) error {
	return nil
}

func (r *Executor) CheckIsExecutor(ctx context.Context) (bool, error) {
	isValid, err := r.entryPoint.IsExecutor(ctx, r.contractAddress)
	if err != nil {
		return false, fmt.Errorf("check is executor: %w", err)
	}

	return isValid, nil
}
