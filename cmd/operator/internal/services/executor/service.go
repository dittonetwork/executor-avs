package executor

import (
	"context"

	api "github.com/dittonetwork/executor-avs/api/operator"
	"github.com/dittonetwork/executor-avs/pkg/log"
	"github.com/dittonetwork/executor-avs/pkg/service"
)

type Service struct {
	executor Executor

	status api.ServiceStatusType
	ctx    context.Context
	cancel context.CancelFunc
}

func NewService(executorHandler Executor) *Service {
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

	// could throw only if cache creation failed, not during execution
	err := s.executor.LongPollingLoop(s.ctx)
	if err != nil {
		log.With(log.Err(err)).Error("error starting long poll lopo")
		s.Stop()
	}
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
