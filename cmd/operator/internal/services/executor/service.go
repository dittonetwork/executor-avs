package executor

import (
	"context"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"

	api "github.com/dittonetwork/executor-avs/api/operator"
	"github.com/dittonetwork/executor-avs/cmd/operator/internal/models"
	"github.com/dittonetwork/executor-avs/pkg/log"
)

type EthereumClient interface {
	SubscribeNewHead(ctx context.Context) (chan *types.Header, ethereum.Subscription, error)
	EstimateGas(ctx context.Context, msg types.Transaction) (uint64, error)
	BlockByHash(ctx context.Context, hash common.Hash) (*types.Block, error)
}

type DittoEntryPoint interface {
	GetAllActiveWorkflows(ctx context.Context) ([]models.Workflow, error)
	RegisterExecutor(ctx context.Context) error
	IsExecutor(ctx context.Context, executorAddr string) (bool, error)
	IsValidExecutor(ctx context.Context, blockNumber int64, executorAddr string) (bool, error)
	RunWorkflow(ctx context.Context, vaultAddr string, workflowID uint64) error
}

type Service struct {
	runner *Runner
	client EthereumClient
	status api.ServiceStatusType
	done   chan struct{}
}

func NewService(client EthereumClient, entryPoint DittoEntryPoint) *Service {
	runner := NewRunner(client, entryPoint)

	return &Service{
		runner: runner,
		client: client,
		status: api.ServiceStatusTypeDown,
		done:   make(chan struct{}),
	}
}

func (s *Service) Start() {
	ctx := context.Background()

	s.status = api.ServiceStatusTypeActive

	headers, sub, err := s.client.SubscribeNewHead(ctx)
	if err != nil {
		log.With(log.Err(err)).Fatal("subscribe new head")
	}

	for {
		select {
		case err = <-sub.Err():
			log.With(log.Err(err)).Error("subscription error")
		case header := <-headers:
			var block *types.Block
			block, err = s.client.BlockByHash(ctx, header.Hash())
			if err != nil {
				log.With(log.Err(err)).Error("get block by hash")
			}

			if err = s.runner.Handle(ctx, block); err != nil {
				log.With(log.Err(err)).Error("handle block")
			}
		case <-s.done:
			return
		}
	}
}

func (s *Service) Stop() {
	s.status = api.ServiceStatusTypeDown

	s.done <- struct{}{}
}
