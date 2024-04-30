package executor

import api "github.com/dittonetwork/executor-avs/api/operator"

type Service struct {
	status api.ServiceStatusType
}

func NewService() *Service {
	return &Service{
		status: api.ServiceStatusTypeDown,
	}
}

func (s *Service) Start() {
	// Start ethereum block watcher
}

func (s *Service) Stop() {
	// Stop ethereum block watcher
}
