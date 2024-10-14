//nolint:testpackage // it's necessary to use the unexported fields for testing
package executor

import (
	"testing"
	"time"

	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"

	api "github.com/dittonetwork/executor-avs/api/operator"
	"github.com/dittonetwork/executor-avs/cmd/operator/internal/services/executor/mocks"
)

func TestService_GracefulShutdown(t *testing.T) {
	executorHandler := mocks.NewExecutor(t)
	executorHandler.EXPECT().Activate(mock.Anything).Return(nil)
	executorHandler.EXPECT().IsAutoDeactivate().Return(true)
	executorHandler.EXPECT().Deactivate(mock.Anything).Return(nil)
	executorHandler.EXPECT().LongPollingLoop(mock.Anything).Return(nil)

	service := NewService(executorHandler)
	service.status = api.ServiceStatusTypeActive

	service.Start()
	time.Sleep(100 * time.Millisecond)
	require.Equal(t, api.ServiceStatusTypeActive, service.GetStatus())

	service.Stop()
	time.Sleep(100 * time.Millisecond)
	require.Equal(t, api.ServiceStatusTypeDown, service.GetStatus())
}
