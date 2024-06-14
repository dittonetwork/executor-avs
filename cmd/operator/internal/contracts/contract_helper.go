package contracts

import (
	"context"
	"crypto/ecdsa"
	"fmt"
	"math/big"
	"time"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
)

type ContractHelper struct {
	Client          *ethclient.Client
	ChainID         *big.Int
	PrivateKey      *ecdsa.PrivateKey
	KeyedTransactor *bind.TransactOpts
}

func NewContractHelper(ethClientURL string, privateKey string) (*ContractHelper, error) {
	const dialTimeoutSecs = 5
	ctx, cancel := context.WithTimeout(context.Background(), dialTimeoutSecs*time.Second)
	defer cancel()

	ethClient, err := ethclient.DialContext(ctx, ethClientURL)
	if err != nil {
		return nil, fmt.Errorf("connection to RPC node: %w", err)
	}

	var privateKeyECDSA *ecdsa.PrivateKey
	privateKeyECDSA, err = crypto.HexToECDSA(privateKey)
	if err != nil {
		return nil, fmt.Errorf("private key hex to ECDSA: %w", err)
	}

	chainID, err := ethClient.NetworkID(context.Background())
	if err != nil {
		return nil, fmt.Errorf("get network id: %w", err)
	}

	keyedTransactor, err := bind.NewKeyedTransactorWithChainID(privateKeyECDSA, chainID)
	if err != nil {
		return nil, fmt.Errorf("new keyd transactor with chain id: %w", err)
	}

	return &ContractHelper{
		Client:          ethClient,
		ChainID:         chainID,
		PrivateKey:      privateKeyECDSA,
		KeyedTransactor: keyedTransactor,
	}, nil
}
