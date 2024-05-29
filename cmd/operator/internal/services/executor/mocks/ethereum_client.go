// Code generated by mockery v2.43.1. DO NOT EDIT.

package mocks

import (
	context "context"
	big "math/big"

	common "github.com/ethereum/go-ethereum/common"

	ethereum "github.com/ethereum/go-ethereum"

	mock "github.com/stretchr/testify/mock"

	types "github.com/ethereum/go-ethereum/core/types"
)

// EthereumClient is an autogenerated mock type for the ethereumClient type
type EthereumClient struct {
	mock.Mock
}

type EthereumClient_Expecter struct {
	mock *mock.Mock
}

func (_m *EthereumClient) EXPECT() *EthereumClient_Expecter {
	return &EthereumClient_Expecter{mock: &_m.Mock}
}

// BlockByHash provides a mock function with given fields: ctx, hash
func (_m *EthereumClient) BlockByHash(ctx context.Context, hash common.Hash) (*types.Block, error) {
	ret := _m.Called(ctx, hash)

	if len(ret) == 0 {
		panic("no return value specified for BlockByHash")
	}

	var r0 *types.Block
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, common.Hash) (*types.Block, error)); ok {
		return rf(ctx, hash)
	}
	if rf, ok := ret.Get(0).(func(context.Context, common.Hash) *types.Block); ok {
		r0 = rf(ctx, hash)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*types.Block)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, common.Hash) error); ok {
		r1 = rf(ctx, hash)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// EthereumClient_BlockByHash_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'BlockByHash'
type EthereumClient_BlockByHash_Call struct {
	*mock.Call
}

// BlockByHash is a helper method to define mock.On call
//   - ctx context.Context
//   - hash common.Hash
func (_e *EthereumClient_Expecter) BlockByHash(ctx interface{}, hash interface{}) *EthereumClient_BlockByHash_Call {
	return &EthereumClient_BlockByHash_Call{Call: _e.mock.On("BlockByHash", ctx, hash)}
}

func (_c *EthereumClient_BlockByHash_Call) Run(run func(ctx context.Context, hash common.Hash)) *EthereumClient_BlockByHash_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(common.Hash))
	})
	return _c
}

func (_c *EthereumClient_BlockByHash_Call) Return(_a0 *types.Block, _a1 error) *EthereumClient_BlockByHash_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *EthereumClient_BlockByHash_Call) RunAndReturn(run func(context.Context, common.Hash) (*types.Block, error)) *EthereumClient_BlockByHash_Call {
	_c.Call.Return(run)
	return _c
}

// GetBalance provides a mock function with given fields: ctx
func (_m *EthereumClient) GetBalance(ctx context.Context) (*big.Int, error) {
	ret := _m.Called(ctx)

	if len(ret) == 0 {
		panic("no return value specified for GetBalance")
	}

	var r0 *big.Int
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context) (*big.Int, error)); ok {
		return rf(ctx)
	}
	if rf, ok := ret.Get(0).(func(context.Context) *big.Int); ok {
		r0 = rf(ctx)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*big.Int)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context) error); ok {
		r1 = rf(ctx)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// EthereumClient_GetBalance_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetBalance'
type EthereumClient_GetBalance_Call struct {
	*mock.Call
}

// GetBalance is a helper method to define mock.On call
//   - ctx context.Context
func (_e *EthereumClient_Expecter) GetBalance(ctx interface{}) *EthereumClient_GetBalance_Call {
	return &EthereumClient_GetBalance_Call{Call: _e.mock.On("GetBalance", ctx)}
}

func (_c *EthereumClient_GetBalance_Call) Run(run func(ctx context.Context)) *EthereumClient_GetBalance_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context))
	})
	return _c
}

func (_c *EthereumClient_GetBalance_Call) Return(_a0 *big.Int, _a1 error) *EthereumClient_GetBalance_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *EthereumClient_GetBalance_Call) RunAndReturn(run func(context.Context) (*big.Int, error)) *EthereumClient_GetBalance_Call {
	_c.Call.Return(run)
	return _c
}

// SendTransaction provides a mock function with given fields: ctx, tx
func (_m *EthereumClient) SendTransaction(ctx context.Context, tx *types.Transaction) error {
	ret := _m.Called(ctx, tx)

	if len(ret) == 0 {
		panic("no return value specified for SendTransaction")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, *types.Transaction) error); ok {
		r0 = rf(ctx, tx)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// EthereumClient_SendTransaction_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'SendTransaction'
type EthereumClient_SendTransaction_Call struct {
	*mock.Call
}

// SendTransaction is a helper method to define mock.On call
//   - ctx context.Context
//   - tx *types.Transaction
func (_e *EthereumClient_Expecter) SendTransaction(ctx interface{}, tx interface{}) *EthereumClient_SendTransaction_Call {
	return &EthereumClient_SendTransaction_Call{Call: _e.mock.On("SendTransaction", ctx, tx)}
}

func (_c *EthereumClient_SendTransaction_Call) Run(run func(ctx context.Context, tx *types.Transaction)) *EthereumClient_SendTransaction_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(*types.Transaction))
	})
	return _c
}

func (_c *EthereumClient_SendTransaction_Call) Return(_a0 error) *EthereumClient_SendTransaction_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *EthereumClient_SendTransaction_Call) RunAndReturn(run func(context.Context, *types.Transaction) error) *EthereumClient_SendTransaction_Call {
	_c.Call.Return(run)
	return _c
}

// SimulateTransaction provides a mock function with given fields: ctx, tx, blockNum
func (_m *EthereumClient) SimulateTransaction(ctx context.Context, tx *types.Transaction, blockNum *big.Int) (string, error) {
	ret := _m.Called(ctx, tx, blockNum)

	if len(ret) == 0 {
		panic("no return value specified for SimulateTransaction")
	}

	var r0 string
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *types.Transaction, *big.Int) (string, error)); ok {
		return rf(ctx, tx, blockNum)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *types.Transaction, *big.Int) string); ok {
		r0 = rf(ctx, tx, blockNum)
	} else {
		r0 = ret.Get(0).(string)
	}

	if rf, ok := ret.Get(1).(func(context.Context, *types.Transaction, *big.Int) error); ok {
		r1 = rf(ctx, tx, blockNum)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// EthereumClient_SimulateTransaction_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'SimulateTransaction'
type EthereumClient_SimulateTransaction_Call struct {
	*mock.Call
}

// SimulateTransaction is a helper method to define mock.On call
//   - ctx context.Context
//   - tx *types.Transaction
//   - blockNum *big.Int
func (_e *EthereumClient_Expecter) SimulateTransaction(ctx interface{}, tx interface{}, blockNum interface{}) *EthereumClient_SimulateTransaction_Call {
	return &EthereumClient_SimulateTransaction_Call{Call: _e.mock.On("SimulateTransaction", ctx, tx, blockNum)}
}

func (_c *EthereumClient_SimulateTransaction_Call) Run(run func(ctx context.Context, tx *types.Transaction, blockNum *big.Int)) *EthereumClient_SimulateTransaction_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(*types.Transaction), args[2].(*big.Int))
	})
	return _c
}

func (_c *EthereumClient_SimulateTransaction_Call) Return(_a0 string, _a1 error) *EthereumClient_SimulateTransaction_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *EthereumClient_SimulateTransaction_Call) RunAndReturn(run func(context.Context, *types.Transaction, *big.Int) (string, error)) *EthereumClient_SimulateTransaction_Call {
	_c.Call.Return(run)
	return _c
}

// SubscribeNewHead provides a mock function with given fields: ctx
func (_m *EthereumClient) SubscribeNewHead(ctx context.Context) (chan *types.Header, ethereum.Subscription, error) {
	ret := _m.Called(ctx)

	if len(ret) == 0 {
		panic("no return value specified for SubscribeNewHead")
	}

	var r0 chan *types.Header
	var r1 ethereum.Subscription
	var r2 error
	if rf, ok := ret.Get(0).(func(context.Context) (chan *types.Header, ethereum.Subscription, error)); ok {
		return rf(ctx)
	}
	if rf, ok := ret.Get(0).(func(context.Context) chan *types.Header); ok {
		r0 = rf(ctx)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(chan *types.Header)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context) ethereum.Subscription); ok {
		r1 = rf(ctx)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(ethereum.Subscription)
		}
	}

	if rf, ok := ret.Get(2).(func(context.Context) error); ok {
		r2 = rf(ctx)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// EthereumClient_SubscribeNewHead_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'SubscribeNewHead'
type EthereumClient_SubscribeNewHead_Call struct {
	*mock.Call
}

// SubscribeNewHead is a helper method to define mock.On call
//   - ctx context.Context
func (_e *EthereumClient_Expecter) SubscribeNewHead(ctx interface{}) *EthereumClient_SubscribeNewHead_Call {
	return &EthereumClient_SubscribeNewHead_Call{Call: _e.mock.On("SubscribeNewHead", ctx)}
}

func (_c *EthereumClient_SubscribeNewHead_Call) Run(run func(ctx context.Context)) *EthereumClient_SubscribeNewHead_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context))
	})
	return _c
}

func (_c *EthereumClient_SubscribeNewHead_Call) Return(_a0 chan *types.Header, _a1 ethereum.Subscription, _a2 error) *EthereumClient_SubscribeNewHead_Call {
	_c.Call.Return(_a0, _a1, _a2)
	return _c
}

func (_c *EthereumClient_SubscribeNewHead_Call) RunAndReturn(run func(context.Context) (chan *types.Header, ethereum.Subscription, error)) *EthereumClient_SubscribeNewHead_Call {
	_c.Call.Return(run)
	return _c
}

// NewEthereumClient creates a new instance of EthereumClient. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewEthereumClient(t interface {
	mock.TestingT
	Cleanup(func())
}) *EthereumClient {
	mock := &EthereumClient{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
