package executor

import (
	"context"
	"errors"

	"github.com/ethereum/go-ethereum/core/types"

	"github.com/dittonetwork/executor-avs/pkg/log"
)

var (
	ErrBlockIsNil = errors.New("block is nil")
)

type Runner struct {
	client     EthereumClient
	entryPoint DittoEntryPoint
}

func NewRunner(client EthereumClient, entryPoint DittoEntryPoint) *Runner {
	return &Runner{
		client:     client,
		entryPoint: entryPoint,
	}
}

func (r *Runner) Handle(ctx context.Context, block *types.Block) error {
	if block == nil {
		return ErrBlockIsNil
	}

	log.Info("block", log.String("number", block.Number().String()))

	return nil
}
