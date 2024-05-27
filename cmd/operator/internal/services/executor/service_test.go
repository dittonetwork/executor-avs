package executor

import (
	"context"
	"testing"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/stretchr/testify/require"

	api "github.com/dittonetwork/executor-avs/api/operator"
	"github.com/dittonetwork/executor-avs/cmd/operator/internal/services/executor/mocks"
)

func TestService_GracefulShutdown(t *testing.T) {
	ctx := context.Background()

	blockHash := common.HexToHash("0x1234")

	executorHandler := mocks.NewExecutor(t)
	executorHandler.EXPECT().Handle(ctx, blockHash).Return(ErrUnregisteredExecutor)

	service := NewService(executorHandler)
	service.status = api.ServiceStatusTypeActive

	err := service.HandleBlock(ctx, blockHash)
	require.NoError(t, err)
	require.Equal(t, api.ServiceStatusTypeActive, service.GetStatus())

	go service.Stop()
	time.Sleep(100 * time.Millisecond)
	err = service.HandleBlock(ctx, blockHash)
	require.NoError(t, err)
	time.Sleep(100 * time.Millisecond)
	require.Equal(t, api.ServiceStatusTypeDown, service.GetStatus())
}
