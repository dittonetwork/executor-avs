package dittoentrypoint

import (
	"context"
	"fmt"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"

	"github.com/dittonetwork/executor-avs/cmd/operator/internal/models"
	"github.com/dittonetwork/executor-avs/contracts/gen/dittoentrypoint"
)

type DittoEntryPoint struct {
	dep *dittoentrypoint.Dittoentrypoint
}

func New(ethClient *ethclient.Client, contractAddress string) (*DittoEntryPoint, error) {
	dittoEntryPoint := common.HexToAddress(contractAddress)

	dep, err := dittoentrypoint.NewDittoentrypoint(dittoEntryPoint, ethClient)
	if err != nil {
		return nil, fmt.Errorf("new ditto entry point: %w", err)
	}

	return &DittoEntryPoint{
		dep: dep,
	}, nil
}

func (d *DittoEntryPoint) RegisterExecutor(ctx context.Context) error {
	return nil
}

func (d *DittoEntryPoint) UnregisterExecutor(ctx context.Context) error {
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
