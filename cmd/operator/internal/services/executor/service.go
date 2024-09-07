package executor

import (
	"context"
	"sync"

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
	IsAutoDeactivate() bool
}

type Service struct {
	executor executor

	status api.ServiceStatusType
	ctx    context.Context
	cancel context.CancelFunc
}

func NewService(executorHandler executor) *Service {
	return &Service{
		executor: executorHandler,
		status:   api.ServiceStatusTypeDown,
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
	go s.run()
}

func (s *Service) run() {
	s.ctx, s.cancel = context.WithCancel(context.Background())

	log.Info("starting executor")
	s.status = api.ServiceStatusTypeActive

	go func() {
		err := s.executor.Activate(s.ctx)
		if err != nil {
			log.With(log.Err(err)).Fatal("activate executor")
		}
	}()

	blocks, sub, err := s.executor.SubscribeToNewBlocks(s.ctx)
	if err != nil {
		log.With(log.Err(err)).Fatal("subscribe to new blocks")
	}
	defer sub.Unsubscribe()

	var wg sync.WaitGroup
	func() {
		for {
			select {
			case err = <-sub.Err():
				log.With(log.Err(err)).Fatal("subscription error")
			case block := <-blocks:
				wg.Add(1)
				go func(b *types.Header) {
					defer wg.Done()
					if err = s.executor.Handle(s.ctx, b.Hash()); err != nil {
						log.With(log.Err(err)).Error("handle block")
					}
				}(block)
			case <-s.ctx.Done():
				return
			}
		}
	}()

	wg.Wait()
}

func (s *Service) Stop() {
	log.Info("stopping the executor service...")

	if s.executor.IsAutoDeactivate() {
		err := s.executor.Deactivate(context.Background())
		if err != nil {
			log.With(log.Err(err)).Error("deactivate executor")
		}
	}

	s.cancel()
	s.status = api.ServiceStatusTypeDown

	log.Info("executor is stopped 🫡")
}

var _ service.StartStopper = (*Service)(nil)
