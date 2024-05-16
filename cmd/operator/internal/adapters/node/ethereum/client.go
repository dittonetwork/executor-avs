package ethereum

import (
	"context"
	"fmt"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
)

type Client struct {
	client *ethclient.Client
}

func NewClient(client *ethclient.Client) *Client {
	return &Client{
		client: client,
	}
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

func (c *Client) EstimateGas(ctx context.Context, msg types.Transaction) (uint64, error) {
	return 0, nil
}
