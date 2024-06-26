package executor

import (
	"context"
	"errors"
	"fmt"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"

	api "github.com/dittonetwork/executor-avs/api/operator"
	"github.com/dittonetwork/executor-avs/pkg/log"
	"github.com/dittonetwork/executor-avs/pkg/service"
)

//go:generate mockery --name executor --output ./mocks --outpkg mocks
type executor interface {
	SubscribeToNewBlocks(ctx context.Context) (chan *types.Header, ethereum.Subscription, error)
	Handle(ctx context.Context, blockHash common.Hash) error
	Deactivate(ctx context.Context) error
	Activate(ctx context.Context) error
}

type Service struct {
	executor executor

	status         api.ServiceStatusType
	isShuttingDown bool
	done           chan struct{}
}

func NewService(executorHandler executor) *Service {
	return &Service{
		executor:       executorHandler,
		status:         api.ServiceStatusTypeDown,
		isShuttingDown: false,
		done:           make(chan struct{}),
	}
}

func (s *Service) GetName() string {
	return ""
}

func (s *Service) GetID() string {
	return ""
}

func (s *Service) GetDescription() string {
	return ""
}

func (s *Service) GetStatus() api.ServiceStatusType {
	return s.status
}

func (s *Service) Start() {
	go s.start()
}
func (s *Service) start() {
	ctx := context.Background()

	log.Info("starting executor")
	s.status = api.ServiceStatusTypeActive

	go func() {
		err := s.executor.Activate(ctx)
		if err != nil {
			log.With(log.Err(err)).Fatal("activate executor")
		}
	}()

	blocks, sub, err := s.executor.SubscribeToNewBlocks(ctx)
	if err != nil {
		log.With(log.Err(err)).Fatal("subscribe to new blocks")
	}

	for {
		select {
		case err = <-sub.Err():
			log.With(log.Err(err)).Error("subscription error")
		case block := <-blocks:
			if err = s.executor.Handle(ctx, block.Hash()); err != nil {
				log.With(log.Err(err)).Error("handle block")
			}
		}
	}
}

func (s *Service) HandleBlock(ctx context.Context, blockHash common.Hash) error {
	if err := s.executor.Handle(ctx, blockHash); err != nil {
		if errors.Is(err, ErrUnregisteredExecutor) {
			if s.isShuttingDown {
				s.done <- struct{}{}
			}
		} else {
			return fmt.Errorf("executor handle: %w", err)
		}
	}

	return nil
}

func (s *Service) Stop() {
	log.Info("stopping the executor service...")

	err := s.executor.Deactivate(context.Background())
	if err != nil {
		log.With(log.Err(err)).Error("deactivate executor")
	}

	s.isShuttingDown = true
	<-s.done
	s.status = api.ServiceStatusTypeDown

	log.Info("executor is stopped 🫡")
}

var _ service.StartStopper = (*Service)(nil)
