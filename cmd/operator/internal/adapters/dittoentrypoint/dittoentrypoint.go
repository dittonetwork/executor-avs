package dittoentrypoint

import (
	"context"
	"crypto/ecdsa"
	"fmt"
	"math/big"
	"strings"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"

	"github.com/dittonetwork/executor-avs/cmd/operator/internal/contracts/gen/dittoentrypoint"
	"github.com/dittonetwork/executor-avs/cmd/operator/internal/models"
	"github.com/dittonetwork/executor-avs/pkg/log"
)

//go:generate mockery --name DittoEntryPoint --output ./mocks --outpkg mocks
type DittoEntryPoint interface {
	ActivateExecutor(ctx context.Context) (*types.Transaction, error)
	DeactivateExecutor(ctx context.Context) (*types.Transaction, error)
	SetDelegatedSigner(ctx context.Context, signerAddress string) (*types.Transaction, error)
	IsExecutor(ctx context.Context) (bool, error)
	IsValidExecutor(ctx context.Context, blockNumber *big.Int) (bool, error)
	GetAllActiveWorkflows(ctx context.Context) ([]models.Workflow, error)
	SimulateMutlipleWorkflows(ctx context.Context, workflows []models.Workflow) ([]bool, error)
	RunMultipleWorkflows(
		ctx context.Context,
		workflows []models.Workflow,
		estimatedGasMultiplier float64,
	) (*types.Transaction, error)
	ArrangeExecutors(ctx context.Context) (*types.Transaction, error)
	GetAmountExecutors(ctx context.Context) (*big.Int, error)
	CalculateOperatorAVSRegistrationDigestHash(
		ctx context.Context,
		address common.Address,
		salt [32]byte,
		expiry *big.Int,
	) ([32]byte, error)
	GetSucceededWorkflows(logs []*types.Log) ([]models.Workflow, error)
	// makeTransacOpts(ctx context.Context) (*bind.TransactOpts, error)
}

type Impl struct {
	client       *ethclient.Client
	dep          *dittoentrypoint.Dittoentrypoint
	privateKey   *ecdsa.PrivateKey
	contractAddr common.Address
}

var _ DittoEntryPoint = new(Impl)

func New(ethClient *ethclient.Client, contractAddress, privateKey string) (*Impl, error) {
	dep, err := dittoentrypoint.NewDittoentrypoint(common.HexToAddress(contractAddress), ethClient)
	if err != nil {
		return nil, fmt.Errorf("new ditto entry point: %w", err)
	}

	var privateKeyECDSA *ecdsa.PrivateKey
	privateKeyECDSA, err = crypto.HexToECDSA(privateKey)
	if err != nil {
		return nil, fmt.Errorf("private key hex to ECDSA: %w", err)
	}

	return &Impl{
		client:       ethClient,
		dep:          dep,
		privateKey:   privateKeyECDSA,
		contractAddr: common.HexToAddress(contractAddress),
	}, nil
}

func (d *Impl) ActivateExecutor(ctx context.Context) (*types.Transaction, error) {
	opts, err := d.makeTransacOpts(ctx)
	if err != nil {
		return nil, fmt.Errorf("make transac opts: %w", err)
	}

	tx, err := d.dep.ActivateExecutor(opts)
	if err != nil {
		return nil, fmt.Errorf("call ActivateExecutor: %w", err)
	}

	log.With(log.String("tx_hash", tx.Hash().String())).Info("register as operator")

	return tx, nil
}

func (d *Impl) DeactivateExecutor(ctx context.Context) (*types.Transaction, error) {
	opts, err := d.makeTransacOpts(ctx)
	if err != nil {
		return nil, fmt.Errorf("make transac opts: %w", err)
	}

	tx, err := d.dep.DeactivateExecutor(opts)
	if err != nil {
		return nil, fmt.Errorf("call DeactivateExecutor: %w", err)
	}

	log.With(log.String("tx_hash", tx.Hash().String())).Info("DeactivateExecutor")

	return tx, nil
}

func (d *Impl) SetDelegatedSigner(ctx context.Context, signerAddress string) (*types.Transaction, error) {
	opts, err := d.makeTransacOpts(ctx)
	if err != nil {
		return nil, fmt.Errorf("make transac opts: %w", err)
	}

	tx, err := d.dep.SetDelegatedSigner(opts, common.HexToAddress(signerAddress))
	if err != nil {
		return nil, fmt.Errorf("call SetDelegatedSigner: %w", err)
	}

	log.With(log.String("tx_hash", tx.Hash().String())).Info("SetDelegatedSigner")

	return tx, nil
}

func (d *Impl) IsExecutor(ctx context.Context) (bool, error) {
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

func (d *Impl) IsValidExecutor(ctx context.Context, blockNumber *big.Int) (bool, error) {
	address := crypto.PubkeyToAddress(d.privateKey.PublicKey)

	opts := &bind.CallOpts{
		Context: ctx,
		From:    address,
	}

	isValidExecutor, err := d.dep.IsValidExecutor(opts, blockNumber, address)
	if err != nil {
		return false, fmt.Errorf("call isValidExecutor: %w", err)
	}

	return isValidExecutor, nil
}

func (d *Impl) GetAllActiveWorkflows(ctx context.Context) ([]models.Workflow, error) {
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

func (d *Impl) SimulateMutlipleWorkflows(
	ctx context.Context,
	workflows []models.Workflow,
) ([]bool, error) {
	wfs := make([]dittoentrypoint.IDittoEntryPointWorkflow, 0, len(workflows))

	for _, workflow := range workflows {
		wfs = append(wfs, dittoentrypoint.IDittoEntryPointWorkflow{
			VaultAddress: workflow.VaultAddress,
			WorkflowId:   workflow.WorkflowID,
		})
	}

	var result []interface{}
	raw := dittoentrypoint.DittoentrypointCallerRaw{Contract: &d.dep.DittoentrypointCaller}
	err := raw.Call(
		&bind.CallOpts{
			// Pending: false // Whether to operate on the pending state or the last known one
			From:    crypto.PubkeyToAddress(d.privateKey.PublicKey),
			Context: ctx,
		},
		&result,
		"runMultipleWorkflows",
		wfs,
	)
	if err != nil {
		return nil, err
	}
	convertedValue, ok := result[0].([]bool)
	if !ok {
		log.Fatal("Failed to cast result[0] to []bool")
	}

	return convertedValue, nil
}

func (d *Impl) RunMultipleWorkflows(ctx context.Context,
	workflows []models.Workflow, estimatedGasMultiplier float64) (
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
	opts.NoSend = true
	tx, err := d.dep.RunMultipleWorkflows(opts, wfs)
	if err != nil {
		return nil, fmt.Errorf("gas estimation RunMultipleWorkflows: %w", err)
	}

	opts.NoSend = false
	// opts.GasPrice = tx.GasPrice()
	// opts.GasFeeCap = tx.GasFeeCap()
	// opts.GasTipCap = tx.GasTipCap()
	opts.GasLimit = uint64(float64(tx.Gas()) * estimatedGasMultiplier)

	tx, err = d.dep.RunMultipleWorkflows(opts, wfs)
	if err != nil {
		return nil, fmt.Errorf("call runMultipleWorkflows: %w", err)
	}

	return tx, nil
}

func (d *Impl) ArrangeExecutors(ctx context.Context) (*types.Transaction, error) {
	opts, err := d.makeTransacOpts(ctx)
	if err != nil {
		return nil, fmt.Errorf("make transac opts: %w", err)
	}

	tx, err := d.dep.ArrangeExecutors(opts)
	if err != nil {
		return nil, fmt.Errorf("call arrange executors: %w", err)
	}

	log.With(log.String("tx_hash", tx.Hash().String())).Info("arrange executors")

	return tx, nil
}

func (d *Impl) GetAmountExecutors(ctx context.Context) (*big.Int, error) {
	amount, err := d.dep.GetAmountExecutors(&bind.CallOpts{Context: ctx})
	if err != nil {
		return nil, fmt.Errorf("call arrange executors: %w", err)
	}

	return amount, nil
}

func (d *Impl) CalculateOperatorAVSRegistrationDigestHash(
	ctx context.Context,
	address common.Address,
	salt [32]byte,
	expiry *big.Int,
) ([32]byte, error) {
	opts := &bind.CallOpts{
		Context: ctx,
	}
	stringToSign, err := d.dep.CalculateOperatorAVSRegistrationDigestHash(opts, address, salt, expiry)
	if err != nil {
		return [32]byte{}, fmt.Errorf("call CalculateOperatorAVSRegistrationDigestHash: %w", err)
	}

	return stringToSign, nil
}

func (d *Impl) GetSucceededWorkflows(logs []*types.Log) ([]models.Workflow, error) {
	succeededWorkflows := make([]models.Workflow, 0, len(logs))
	for _, vLog := range logs {
		event, err := d.dep.ParseDittoEntryPointWorkflowSuccess(*vLog)

		if err != nil {
			if strings.Contains(err.Error(), "no event signature") || strings.Contains(err.Error(), "event signature mismatch") {
				// These aren't logs you're looking for
				continue
			}
			return nil, fmt.Errorf("ParseDittoEntryPointWorkflowSuccess: %w", err)
		}
		succeededWorkflows = append(succeededWorkflows, models.Workflow{
			VaultAddress: event.Vault,
			WorkflowID:   event.WorkflowId,
		})
	}

	return succeededWorkflows, nil
}

func (d *Impl) makeTransacOpts(ctx context.Context) (*bind.TransactOpts, error) {
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
