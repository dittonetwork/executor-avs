package dittoentrypoint

import (
	"context"

	"github.com/dittonetwork/executor-avs/cmd/operator/internal/adapters/node/ethclient"
	"github.com/dittonetwork/executor-avs/cmd/operator/internal/models"
)

type DittoEntryPoint struct {
	ethClient *ethclient.Client
}

func New(ethClient *ethclient.Client) *DittoEntryPoint {
	return &DittoEntryPoint{
		ethClient: ethClient,
	}
}

func (d *DittoEntryPoint) RegisterExecutor(ctx context.Context) error {
	return nil
}

func (d *DittoEntryPoint) IsExecutor(ctx context.Context, executorAddr string) (bool, error) {
	return false, nil
}

func (d *DittoEntryPoint) IsValidExecutor(ctx context.Context, blockNumber int64, executorAddr string) (bool, error) {
	return false, nil
}

func (d *DittoEntryPoint) GetAllActiveWorkflows(_ context.Context) ([]models.Workflow, error) {
	return make([]models.Workflow, 0), nil
}

func (d *DittoEntryPoint) RunWorkflow(ctx context.Context, vaultAddr string, workflowID uint64) error {
	return nil
}
