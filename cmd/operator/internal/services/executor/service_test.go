//nolint:testpackage // it's necessary to use the unexported fields for testing
package executor

import (
	"math/big"
	"testing"
	"time"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"

	api "github.com/dittonetwork/executor-avs/api/operator"
	"github.com/dittonetwork/executor-avs/cmd/operator/internal/services/executor/mocks"
)

type mockSubscription struct {
	errChan      chan error
	unsubscribed bool
}

// NewMockSubscription creates and returns a new mock Subscription.
func NewMockSubscription() ethereum.Subscription {
	return &mockSubscription{
		errChan: make(chan error, 1), // Buffered channel to allow non-blocking sends.
	}
}

// Unsubscribe stops the subscription. It closes the error channel,
// following the contract specified by the Subscription interface.
func (m *mockSubscription) Unsubscribe() {
	// Use a flag to ensure idempotency, i.e., multiple calls to Unsubscribe do not panic.
	if !m.unsubscribed {
		close(m.errChan)
		m.unsubscribed = true
	}
}

// Err returns the error channel. The channel is closed by Unsubscribe.
func (m *mockSubscription) Err() <-chan error {
	return m.errChan
}

func TestService_GracefulShutdown(t *testing.T) {
	header := &types.Header{Number: big.NewInt(100)}
	headersChan := make(chan *types.Header, 1)

	executorHandler := mocks.NewExecutor(t)
	executorHandler.EXPECT().Activate(mock.Anything).Return(nil)
	executorHandler.EXPECT().Handle(mock.Anything, header.Hash()).Return(ErrUnregisteredExecutor)
	executorHandler.EXPECT().IsAutoDeactivate().Return(true)
	executorHandler.EXPECT().Deactivate(mock.Anything).Return(nil)
	executorHandler.EXPECT().SubscribeToNewBlocks(mock.Anything).Return(headersChan, NewMockSubscription(), nil)

	service := NewService(executorHandler)
	service.status = api.ServiceStatusTypeActive

	service.Start()

	time.Sleep(100 * time.Millisecond)
	require.Equal(t, api.ServiceStatusTypeActive, service.GetStatus())

	headersChan <- header
	time.Sleep(time.Second)
	service.Stop()

	headersChan <- header
	time.Sleep(time.Second)

	require.Equal(t, api.ServiceStatusTypeDown, service.GetStatus())
}
