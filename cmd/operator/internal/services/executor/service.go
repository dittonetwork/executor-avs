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
	UnregisterExecutor(ctx context.Context) error
	IsExecutor(ctx context.Context, executorAddr string) (bool, error)
	IsValidExecutor(ctx context.Context, blockNumber int64, executorAddr string) (bool, error)
	RunWorkflow(ctx context.Context, vaultAddr string, workflowID uint64) error
}

type Service struct {
	executor        *Executor
	client          EthereumClient
	entryPoint      DittoEntryPoint
	contractAddress string
	status          api.ServiceStatusType

	isShuttingDown bool
	done           chan struct{}
}

func NewService(client EthereumClient, entryPoint DittoEntryPoint, contractAddress string) *Service {
	executor := NewExecutor(client, entryPoint, contractAddress)

	return &Service{
		executor:        executor,
		client:          client,
		entryPoint:      entryPoint,
		contractAddress: contractAddress,
		status:          api.ServiceStatusTypeDown,
		isShuttingDown:  false,
		done:            make(chan struct{}),
	}
}

func (s *Service) Start() {
	go s.start()
}
func (s *Service) start() {
	log.Info("starting executor")

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

				continue
			}

			log.With(log.Int64("block_number", block.Number().Int64())).Info("checking if it is executor...")

			var isExecutor bool
			isExecutor, err = s.executor.CheckIsExecutor(ctx)
			if err != nil {
				log.With(log.Err(err)).Error("checking is executor error")

				continue
			}

			if !isExecutor {
				log.Info("not executor ❌")

				if s.isShuttingDown {
					s.done <- struct{}{}

					return
				}

				continue
			}

			if err = s.executor.Handle(ctx, block); err != nil {
				log.With(log.Err(err)).Error("handle block")
			}
		}
	}
}

func (s *Service) Stop() {
	ctx := context.Background()

	log.Info("stopping the executor service...")

	if err := s.entryPoint.UnregisterExecutor(ctx); err != nil {
		log.With(log.Err(err)).Error("failed to unregister executor")
	}

	log.Info("unregistering the executor...")

	s.isShuttingDown = true
	s.status = api.ServiceStatusTypeDown

	<-s.done

	log.Info("executor is stopped 🫡")
}
