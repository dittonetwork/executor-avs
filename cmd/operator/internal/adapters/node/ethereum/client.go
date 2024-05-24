package ethereum

import (
	"context"
	"crypto/ecdsa"
	"fmt"
	"math/big"

	"github.com/dittonetwork/executor-avs/pkg/hex"
	"github.com/dittonetwork/executor-avs/pkg/log"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
)

type Client struct {
	contractAddr string
	client       *ethclient.Client
	privateKey   *ecdsa.PrivateKey
}

func NewClient(client *ethclient.Client, contractAddr, privateKey string) (*Client, error) {
	pk, err := crypto.HexToECDSA(hex.ConvertTo16Bit(privateKey))
	if err != nil {
		return nil, fmt.Errorf("hex to ecdsa: %w", err)
	}

	return &Client{
		client:       client,
		contractAddr: contractAddr,
		privateKey:   pk,
	}, nil
}

func (c *Client) SubscribeNewHead(ctx context.Context) (chan *types.Header, ethereum.Subscription, error) {
	headers := make(chan *types.Header)
	subscription, err := c.client.SubscribeNewHead(ctx, headers)
	if err != nil {
		return nil, nil, fmt.Errorf("subscribe to new headers: %w", err)
	}

	return headers, subscription, nil
}

func (c *Client) BlockByHash(ctx context.Context, hash common.Hash) (*types.Block, error) {
	return c.client.BlockByHash(ctx, hash)
}

func (c *Client) SimulateTransaction(
	ctx context.Context, tx *types.Transaction, blockNum *big.Int, result interface{},
) error {
	err := c.client.Client().CallContext(ctx, &result, "eth_call", map[string]interface{}{
		"from":  crypto.PubkeyToAddress(c.privateKey.PublicKey),
		"to":    common.HexToAddress(c.contractAddr),
		"data":  hexutil.Encode(tx.Data()),
		"block": blockNum.Int64(),
	})
	if err != nil {
		return fmt.Errorf("call eth_call: %w", err)
	}

	log.With(log.Any("result", result)).Debug("simulate transaction done")

	return nil
}

func (c *Client) SendTransaction(ctx context.Context, tx *types.Transaction) error {
	// signedTx, err := types.SignTx(tx, types.NewEIP155Signer(tx.ChainId()), c.privateKey)
	// if err != nil {
	// 	return fmt.Errorf("sign tx: %w", err)
	// }

	if err := c.client.SendTransaction(ctx, tx); err != nil {
		return fmt.Errorf("send tx: %w", err)
	}

	return nil
}
