package app

import (
	"context"
	"crypto/ecdsa"
	"crypto/rand"
	"encoding/hex"
	"errors"
	"fmt"
	"math/big"
	"time"

	"github.com/dittonetwork/executor-avs/cmd/operator/internal/adapters/dittoentrypoint"
	"github.com/dittonetwork/executor-avs/cmd/operator/internal/contracts"
	"github.com/dittonetwork/executor-avs/cmd/operator/internal/contracts/gen/iservicemanager"
	"github.com/dittonetwork/executor-avs/pkg/log"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
)

func RegisterOperator(ctx context.Context, cfg *Config) error {
	if len(cfg.OperatorPrivateKey) == 0 {
		return errors.New("OPERATOR_PRIVATE_KEY env in not set")
	}

	contractHelper, err := contracts.NewContractHelper(ctx, cfg.NodeURL, cfg.OperatorPrivateKey)
	if err != nil {
		return fmt.Errorf("ContractHelper creation: %w", err)
	}

	iServiceManager, err := iservicemanager.NewIservicemanager(
		common.HexToAddress(cfg.ContractAddress),
		contractHelper.Client,
	)
	if err != nil {
		return fmt.Errorf("IServiceManager creation: %w", err)
	}

	dep, err := initDittoEntryPoint(cfg.NodeURL, cfg.ContractAddress, cfg.OperatorPrivateKey)
	if err != nil {
		return fmt.Errorf("init ditto entry point: %w", err)
	}

	registrationStructure, err := createRegistrationStructure(dep, contractHelper.PrivateKey)
	if err != nil {
		return fmt.Errorf("registration structure creation: %w", err)
	}

	tx, err := iServiceManager.RegisterOperatorToAVS(
		contractHelper.KeyedTransactor,
		crypto.PubkeyToAddress(contractHelper.PrivateKey.PublicKey),
		registrationStructure,
	)
	if err != nil {
		return fmt.Errorf("register operator: %w", err)
	}

	log.With(log.String("tx_hash", tx.Hash().Hex())).Info("Register transaction created")
	if waitForTransaction(ctx, cfg, tx) != nil {
		return fmt.Errorf("wait for tx receipt: %w", err)
	}
	return nil
}

func DeregisterOperator(ctx context.Context, cfg *Config) error {
	if len(cfg.OperatorPrivateKey) == 0 {
		return errors.New("OPERATOR_PRIVATE_KEY env in not set")
	}

	contractHelper, err := contracts.NewContractHelper(ctx, cfg.NodeURL, cfg.OperatorPrivateKey)
	if err != nil {
		return fmt.Errorf("ContractHelper creation: %w", err)
	}

	iServiceManager, err := iservicemanager.NewIservicemanager(
		common.HexToAddress(cfg.ContractAddress),
		contractHelper.Client,
	)
	if err != nil {
		return fmt.Errorf("IServiceManager creation: %w", err)
	}

	tx, err := iServiceManager.DeregisterOperatorFromAVS(
		contractHelper.KeyedTransactor,
		crypto.PubkeyToAddress(contractHelper.PrivateKey.PublicKey),
	)
	if err != nil {
		return fmt.Errorf("register executor transaction creation: %w", err)
	}

	log.With(log.String("tx_hash", tx.Hash().Hex())).Info("Register transaction created")
	if waitForTransaction(ctx, cfg, tx) != nil {
		return fmt.Errorf("wait for tx receipt: %w", err)
	}
	return nil
}

func ActivateExecutor(ctx context.Context, cfg *Config) error {
	if len(cfg.ExecutorPrivateKey) == 0 {
		return errors.New("EXECUTOR_PRIVATE_KEY env in not set")
	}

	dep, err := initDittoEntryPoint(cfg.NodeURL, cfg.ContractAddress, cfg.ExecutorPrivateKey)
	if err != nil {
		return fmt.Errorf("init ditto entry point: %w", err)
	}

	tx, err := dep.ActivateExecutor(context.Background())
	if err != nil {
		return fmt.Errorf("ActivateExecutor: %w", err)
	}

	log.With(log.String("tx_hash", tx.Hash().Hex())).Info("Register transaction created")
	if waitForTransaction(ctx, cfg, tx) != nil {
		return fmt.Errorf("wait for tx receipt: %w", err)
	}
	return nil
}

func SetDelegatedSigner(ctx context.Context, cfg *Config, signerAddress string) error {
	if len(cfg.OperatorPrivateKey) == 0 {
		return errors.New("OPERATOR_PRIVATE_KEY env in not set")
	}

	dep, err := initDittoEntryPoint(cfg.NodeURL, cfg.ContractAddress, cfg.OperatorPrivateKey)
	if err != nil {
		return fmt.Errorf("init ditto entry point: %w", err)
	}

	tx, err := dep.SetDelegatedSigner(context.Background(), signerAddress)
	if err != nil {
		return fmt.Errorf("SetDelegatedSigner: %w", err)
	}

	log.With(log.String("tx_hash", tx.Hash().Hex())).Info("Register transaction created")
	if waitForTransaction(ctx, cfg, tx) != nil {
		return fmt.Errorf("wait for tx receipt: %w", err)
	}
	return nil
}

func DeactivateExecutor(ctx context.Context, cfg *Config) error {
	if len(cfg.ExecutorPrivateKey) == 0 {
		return errors.New("EXECUTOR_PRIVATE_KEY env in not set")
	}

	dep, err := initDittoEntryPoint(cfg.NodeURL, cfg.ContractAddress, cfg.ExecutorPrivateKey)
	if err != nil {
		return fmt.Errorf("init ditto entry point: %w", err)
	}

	tx, err := dep.DeactivateExecutor(context.Background())
	if err != nil {
		return fmt.Errorf("DeactivateExecutor: %w", err)
	}

	log.With(log.String("tx_hash", tx.Hash().Hex())).Info("Register transaction created")
	if waitForTransaction(ctx, cfg, tx) != nil {
		return fmt.Errorf("wait for tx receipt: %w", err)
	}
	return nil
}

func ArrangeExecutors(ctx context.Context, cfg *Config) error {
	if len(cfg.ExecutorPrivateKey) == 0 {
		return errors.New("EXECUTOR_PRIVATE_KEY env in not set")
	}

	dep, err := initDittoEntryPoint(cfg.NodeURL, cfg.ContractAddress, cfg.ExecutorPrivateKey)
	if err != nil {
		return fmt.Errorf("init ditto entry point: %w", err)
	}

	executorsAmount, err := dep.GetAmountExecutors(ctx)
	if err != nil {
		return fmt.Errorf("GetAmountExecutors: %w", err)
	}

	log.With(log.Int64("executors_amount", executorsAmount.Int64())).Info("executors amount")

	tx, err := dep.ArrangeExecutors(ctx)
	if err != nil {
		return fmt.Errorf("ArrangeExecutors: %w", err)
	}

	log.With(log.String("tx_hash", tx.Hash().Hex())).Info("Register transaction created")
	if waitForTransaction(ctx, cfg, tx) != nil {
		return fmt.Errorf("wait for tx receipt: %w", err)
	}
	return nil
}

func GenerateKey() error {
	privateKey, err := crypto.GenerateKey()
	if err != nil {
		return err
	}

	// Obtain the public key in the uncompressed form.
	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		log.Fatal("error casting public key to ECDSA")
	}

	// Print the private key in hexadecimal format.
	privateKeyBytes := crypto.FromECDSA(privateKey)
	fmt.Println("Private Key:", hex.EncodeToString(privateKeyBytes))

	// Generate the Ethereum address from the public key.
	address := crypto.PubkeyToAddress(*publicKeyECDSA).Hex()
	fmt.Println("Address:", address)

	return nil
}

func initDittoEntryPoint(
	nodeURL string,
	contractAddr string,
	privateKey string,
) (*dittoentrypoint.DittoEntryPoint, error) {
	conn, err := ethclient.Dial(nodeURL)
	if err != nil {
		return nil, fmt.Errorf("ethereum client dial: %w", err)
	}

	entryPoint, err := dittoentrypoint.New(conn, contractAddr, privateKey)
	if err != nil {
		return nil, fmt.Errorf("dittoentrypoint client init: %w", err)
	}

	return entryPoint, nil
}

func waitForTransaction(ctx context.Context, cfg *Config, tx *types.Transaction) error {
	const pollIntervalSecs = 5
	ethClient, err := ethclient.DialContext(ctx, cfg.NodeURL)
	if err != nil {
		return fmt.Errorf("ethereum client dial: %w", err)
	}

	log.Info("Waiting for transaction to complete...")
	for {
		receipt, innerErr := ethClient.TransactionReceipt(ctx, tx.Hash())
		if receipt != nil {
			log.With(
				log.String("tx_hash", receipt.TxHash.Hex()),
				log.String("block_hash", receipt.BlockHash.Hex()),
			).Info("Transaction included into block ")
			return nil
		}
		if innerErr != nil {
			log.Info("Transaction receipt not available yet, waiting...")
			time.Sleep(pollIntervalSecs * time.Second)
		}
	}
}

func createRegistrationStructure(
	dep *dittoentrypoint.DittoEntryPoint,
	privateKeyECDSA *ecdsa.PrivateKey,
) (iservicemanager.ISignatureUtilsSignatureWithSaltAndExpiry, error) {
	publicKey := privateKeyECDSA.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	address := crypto.PubkeyToAddress(*publicKeyECDSA)
	if !ok {
		return iservicemanager.ISignatureUtilsSignatureWithSaltAndExpiry{}, errors.New("error casting public key to ecdsa")
	}

	salt := [32]byte{}
	if _, err := rand.Read(salt[:]); err != nil {
		return iservicemanager.ISignatureUtilsSignatureWithSaltAndExpiry{}, err
	}

	expiry := big.NewInt(time.Now().Add(1 * time.Hour).Unix()) // 1 hour from now

	digestHash, err := dep.CalculateOperatorAVSRegistrationDigestHash(context.Background(), address, salt, expiry)
	if err != nil {
		return iservicemanager.ISignatureUtilsSignatureWithSaltAndExpiry{}, err
	}

	signature, err := crypto.Sign(digestHash[:], privateKeyECDSA)
	if err != nil {
		return iservicemanager.ISignatureUtilsSignatureWithSaltAndExpiry{}, err
	}

	// Adjusting 'v' for EIP-155
	// TODO: check if this value is the same for other chains
	const magicToAddToSign = 27
	signature[64] += byte(magicToAddToSign)

	return iservicemanager.ISignatureUtilsSignatureWithSaltAndExpiry{
		Signature: signature,
		Salt:      salt,
		Expiry:    expiry,
	}, nil
}
