package executor

import (
	"context"
	"errors"
	"fmt"
	"sync"

	"github.com/ethereum/go-ethereum/core/types"

	"github.com/dittonetwork/executor-avs/cmd/operator/config"
	"github.com/dittonetwork/executor-avs/cmd/operator/internal/models"
	"github.com/dittonetwork/executor-avs/pkg/log"
)

var (
	ErrBlockIsNil = errors.New("block is nil")
)

type Runner struct {
	cfg        config.Config
	client     EthereumClient
	entryPoint DittoEntryPoint
}

func NewRunner(cfg config.Config, client EthereumClient, entryPoint DittoEntryPoint) *Runner {
	return &Runner{
		cfg:        cfg,
		client:     client,
		entryPoint: entryPoint,
	}
}

func (r *Runner) Handle(ctx context.Context, block *types.Block) error {
	if block == nil {
		return ErrBlockIsNil
	}

	// log.Info("block", log.String("number", block.Number().String()))

	isExecutor, err := r.entryPoint.IsExecutor(ctx, r.cfg.DittoEntryPoint.ContractAddress)
	if err != nil {
		return fmt.Errorf("check if is executor: %w", err)
	}
	if !isExecutor {
		if err = r.entryPoint.RegisterExecutor(ctx); err != nil {
			return fmt.Errorf("register executor: %w", err)
		}
	}

	var isValidExecutor bool
	isValidExecutor, err = r.entryPoint.IsValidExecutor(ctx, block.Number().Int64(), r.cfg.DittoEntryPoint.ContractAddress)
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
			canRun, err = r.simulate(ctx, workflow)
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
		if err = r.run(ctx, workflow); err != nil {
			return fmt.Errorf("execute workflow: %w", err)
		}
	}

	return nil
}

func (r *Runner) simulate(ctx context.Context, workflow models.Workflow) (bool, error) {
	return false, nil
}

func (r *Runner) run(ctx context.Context, workflow models.Workflow) error {
	return nil
}
