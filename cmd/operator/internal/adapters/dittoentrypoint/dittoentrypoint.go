package dittoentrypoint

import (
	"context"
	"crypto/ecdsa"
	"fmt"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"

	"github.com/dittonetwork/executor-avs/cmd/operator/internal/models"
	"github.com/dittonetwork/executor-avs/contracts/gen/dittoentrypoint"
	"github.com/dittonetwork/executor-avs/pkg/hex"
	"github.com/dittonetwork/executor-avs/pkg/log"
)

type DittoEntryPoint struct {
	client     *ethclient.Client
	dep        *dittoentrypoint.Dittoentrypoint
	privateKey *ecdsa.PrivateKey
}

func New(ethClient *ethclient.Client, contractAddress, privateKey string) (*DittoEntryPoint, error) {
	dep, err := dittoentrypoint.NewDittoentrypoint(common.HexToAddress(contractAddress), ethClient)
	if err != nil {
		return nil, fmt.Errorf("new ditto entry point: %w", err)
	}

	var privateKeyECDSA *ecdsa.PrivateKey
	privateKeyECDSA, err = crypto.HexToECDSA(hex.ConvertTo16Bit(privateKey))
	if err != nil {
		return nil, fmt.Errorf("private key hex to ECDSA: %w", err)
	}

	return &DittoEntryPoint{
		client:     ethClient,
		dep:        dep,
		privateKey: privateKeyECDSA,
	}, nil
}

func (d *DittoEntryPoint) RegisterExecutor(ctx context.Context) error {
	chainID, err := d.client.NetworkID(ctx)
	if err != nil {
		return fmt.Errorf("get network id: %w", err)
	}

	opts, err := bind.NewKeyedTransactorWithChainID(d.privateKey, chainID)
	if err != nil {
		return fmt.Errorf("new keyd transactor with chain id: %w", err)
	}

	tx, err := d.dep.RegisterExecutor(opts)
	if err != nil {
		return fmt.Errorf("call register executor: %w", err)
	}

	log.With(log.String("tx_hash", tx.Hash().String())).Info("Debug: SUCCESS")

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
