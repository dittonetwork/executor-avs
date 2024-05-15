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
	client       *ethclient.Client
	dep          *dittoentrypoint.Dittoentrypoint
	privateKey   *ecdsa.PrivateKey
	contractAddr common.Address
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
		client:       ethClient,
		dep:          dep,
		privateKey:   privateKeyECDSA,
		contractAddr: common.HexToAddress(contractAddress),
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

func (d *DittoEntryPoint) GetRunWorkflowTx(ctx context.Context, vaultAddr common.Address, workflowID *big.Int) (
	*types.Transaction, error,
) {
	// TODO: figure out what is StorageTransactor
	// https://geth.ethereum.org/docs/developers/dapp-developer/native-bindings)
	// and do we need it (auth.Signer, auth.From)
	dummySigner := func(_ common.Address, tx *types.Transaction) (*types.Transaction, error) {
		return tx, nil
	}
	tx, err := d.dep.RunWorkflowWithoutRevert(&bind.TransactOpts{
		Context:   ctx,
		NoSend:    true,
		GasTipCap: big.NewInt(0), // To prevent it from calling eth_maxPriorityFeePerGas
		GasLimit:  1,             // To prevent it from calling eth_estimateGas
		Signer:    dummySigner,   // No need to sign for simulation
	}, vaultAddr, workflowID)
	if err != nil {
		return nil, fmt.Errorf("call runWorkflowWithoutRevert: %w", err)
	}

	return tx, nil
}

func (d *DittoEntryPoint) RunMultipleWorkflows(ctx context.Context, workflows []models.Workflow) (
	*types.Transaction, error,
) {
	wfs := make([]dittoentrypoint.IDittoEntryPointWorkflow, 0, len(workflows))

	for _, workflow := range workflows {
		wfs = append(wfs, dittoentrypoint.IDittoEntryPointWorkflow{
			VaultAddress: workflow.VaultAddress,
			WorkflowId:   workflow.WorkflowID,
		})
	}

	opts, err := d.makeTransacOpts(ctx)
	if err != nil {
		return nil, fmt.Errorf("make transac opts: %w", err)
	}

	tx, err := d.dep.RunMultipleWorkflows(opts, wfs)
	if err != nil {
		return nil, fmt.Errorf("call runMultipleWorkflows: %w", err)
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
