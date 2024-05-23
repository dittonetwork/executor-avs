package dittoentrypoint

import (
	"context"
	"crypto/ecdsa"
	"fmt"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
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
	opts, err := d.makeTransacOpts(ctx)
	if err != nil {
		return fmt.Errorf("make transac opts: %w", err)
	}

	tx, err := d.dep.RegisterExecutor(opts)
	if err != nil {
		return fmt.Errorf("call registerExecutor: %w", err)
	}

	log.With(log.String("tx_hash", tx.Hash().String())).Info("register as operator")

	return nil
}

func (d *DittoEntryPoint) UnregisterExecutor(ctx context.Context) error {
	opts, err := d.makeTransacOpts(ctx)
	if err != nil {
		return fmt.Errorf("make transac opts: %w", err)
	}

	tx, err := d.dep.UnregisterExecutor(opts)
	if err != nil {
		return fmt.Errorf("call registerExecutor: %w", err)
	}

	log.With(log.String("tx_hash", tx.Hash().String())).Info("unregister as operator")

	return nil
}

func (d *DittoEntryPoint) IsExecutor(ctx context.Context) (bool, error) {
	address := crypto.PubkeyToAddress(d.privateKey.PublicKey)

	opts := &bind.CallOpts{
		Context: ctx,
		From:    address,
	}

	isExecutor, err := d.dep.IsExecutor(opts)
	if err != nil {
		return false, fmt.Errorf("call isExecutor: %w", err)
	}

	return isExecutor, nil
}

func (d *DittoEntryPoint) IsValidExecutor(ctx context.Context, blockNumber *big.Int) (bool, error) {
	address := crypto.PubkeyToAddress(d.privateKey.PublicKey)

	opts := &bind.CallOpts{
		Context:     ctx,
		From:        address,
		BlockNumber: blockNumber,
	}

	isValidExecutor, err := d.dep.IsValidExecutor(opts, blockNumber, address)
	if err != nil {
		return false, fmt.Errorf("call isValidExecutor: %w", err)
	}

	return isValidExecutor, nil
}

func (d *DittoEntryPoint) GetAllActiveWorkflows(ctx context.Context) ([]models.Workflow, error) {
	workflows, err := d.dep.GetAllActiveWorkflows(&bind.CallOpts{
		Context: ctx,
	})
	if err != nil {
		return nil, fmt.Errorf("call getAllActiveWorkflows: %w", err)
	}

	result := make([]models.Workflow, 0, len(workflows))
	for _, workflow := range workflows {
		result = append(result, models.Workflow{
			VaultAddress: workflow.VaultAddress,
			WorkflowID:   workflow.WorkflowId,
		})
	}

	return result, nil
}

func (d *DittoEntryPoint) RunWorkflow(ctx context.Context, vaultAddr common.Address, workflowID *big.Int) (
	*types.Transaction, error,
) {
	tx, err := d.dep.RunWorkflow(&bind.TransactOpts{Context: ctx}, vaultAddr, workflowID)
	if err != nil {
		return nil, fmt.Errorf("call runWorkflow: %w", err)
	}

	return tx, nil
}

func (d *DittoEntryPoint) ArrangeExecutors(ctx context.Context) error {
	opts, err := d.makeTransacOpts(ctx)
	if err != nil {
		return fmt.Errorf("make transac opts: %w", err)
	}

	tx, err := d.dep.ArrangeExecutors(opts)
	if err != nil {
		return fmt.Errorf("call arrange executors: %w", err)
	}

	log.With(log.String("tx_hash", tx.Hash().String())).Info("arrange executors")

	return nil
}

func (d *DittoEntryPoint) GetAmountExecutors(ctx context.Context) (*big.Int, error) {
	amount, err := d.dep.GetAmountExecutors(&bind.CallOpts{Context: ctx})
	if err != nil {
		return nil, fmt.Errorf("call arrange executors: %w", err)
	}

	return amount, nil
}

func (d *DittoEntryPoint) makeTransacOpts(ctx context.Context) (*bind.TransactOpts, error) {
	chainID, err := d.client.NetworkID(ctx)
	if err != nil {
		return nil, fmt.Errorf("get network id: %w", err)
	}

	opts, err := bind.NewKeyedTransactorWithChainID(d.privateKey, chainID)
	if err != nil {
		return nil, fmt.Errorf("new keyd transactor with chain id: %w", err)
	}

	return opts, nil
}
