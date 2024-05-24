package executor

import (
	"context"
	"fmt"
	"math/big"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"

	api "github.com/dittonetwork/executor-avs/api/operator"
	"github.com/dittonetwork/executor-avs/cmd/operator/internal/models"
	"github.com/dittonetwork/executor-avs/pkg/log"
	"github.com/dittonetwork/executor-avs/pkg/service"
)

//go:generate mockery --name EthereumClient --output ./mocks --outpkg mocks
type EthereumClient interface {
	SubscribeNewHead(ctx context.Context) (chan *types.Header, ethereum.Subscription, error)
	BlockByHash(ctx context.Context, hash common.Hash) (*types.Block, error)
	SimulateTransaction(ctx context.Context, tx *types.Transaction, blockNum *big.Int, result interface{}) error
	SendTransaction(ctx context.Context, tx *types.Transaction) error
}

//go:generate mockery --name DittoEntryPoint --output ./mocks --outpkg mocks
type DittoEntryPoint interface {
	GetAllActiveWorkflows(ctx context.Context) ([]models.Workflow, error)

	IsExecutor(ctx context.Context) (bool, error)
	IsValidExecutor(ctx context.Context, blockNumber *big.Int) (bool, error)
	GetRunWorkflowTx(ctx context.Context, vaultAddr common.Address, workflowID *big.Int) (*types.Transaction, error)
	RunMultipleWorkflows(ctx context.Context, workflows []models.Workflow) (*types.Transaction, error)
}

type Service struct {
	executor   *Executor
	client     EthereumClient
	entryPoint DittoEntryPoint
	status     api.ServiceStatusType

	isShuttingDown bool
	done           chan struct{}
}

func NewService(client EthereumClient, entryPoint DittoEntryPoint) *Service {
	executor := NewExecutor(client, entryPoint)

	return &Service{
		executor:       executor,
		client:         client,
		entryPoint:     entryPoint,
		status:         api.ServiceStatusTypeDown,
		isShuttingDown: false,
		done:           make(chan struct{}),
	}
}

func (s *Service) Start() {
	go s.start()
}
func (s *Service) start() {
	ctx := context.Background()

	log.Info("starting executor")
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

			if err = s.executor.Handle(ctx, block.Number()); err != nil {
				log.With(log.Err(err)).Error("handle block")
			}
		}
	}
}

func (s *Service) HandleBlock(ctx context.Context, block *types.Block) error {
	log.With(log.Int64("block_number", block.Number().Int64())).Info("checking if it is executor...")

	isExecutor, err := s.executor.CheckIsExecutor(ctx)
	if err != nil {
		return fmt.Errorf("check if is executor: %w", err)
	}

	if !isExecutor {
		log.Info("not executor ❌")

		if s.isShuttingDown {
			s.done <- struct{}{}
		}

		return nil
	}

	if err = s.executor.Handle(ctx, block.Number()); err != nil {
		return fmt.Errorf("executor handle: %w", err)
	}

	return nil
}

func (s *Service) Stop() {
	log.Info("stopping the executor service...")

	s.isShuttingDown = true
	s.status = api.ServiceStatusTypeDown

	<-s.done

	log.Info("executor is stopped 🫡")
}

var _ service.StartStopper = (*Service)(nil)
