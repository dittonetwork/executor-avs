// Code generated by mockery v2.43.2. DO NOT EDIT.

package mocks

import (
	context "context"
	big "math/big"

	common "github.com/ethereum/go-ethereum/common"

	mock "github.com/stretchr/testify/mock"

	types "github.com/ethereum/go-ethereum/core/types"
)

// EthClient is an autogenerated mock type for the EthClient type
type EthClient struct {
	mock.Mock
}

type EthClient_Expecter struct {
	mock *mock.Mock
}

func (_m *EthClient) EXPECT() *EthClient_Expecter {
	return &EthClient_Expecter{mock: &_m.Mock}
}

// BlockByHash provides a mock function with given fields: ctx, hash
func (_m *EthClient) BlockByHash(ctx context.Context, hash common.Hash) (*types.Block, error) {
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

// EthClient_BlockByHash_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'BlockByHash'
type EthClient_BlockByHash_Call struct {
	*mock.Call
}

// BlockByHash is a helper method to define mock.On call
//   - ctx context.Context
//   - hash common.Hash
func (_e *EthClient_Expecter) BlockByHash(ctx interface{}, hash interface{}) *EthClient_BlockByHash_Call {
	return &EthClient_BlockByHash_Call{Call: _e.mock.On("BlockByHash", ctx, hash)}
}

func (_c *EthClient_BlockByHash_Call) Run(run func(ctx context.Context, hash common.Hash)) *EthClient_BlockByHash_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(common.Hash))
	})
	return _c
}

func (_c *EthClient_BlockByHash_Call) Return(_a0 *types.Block, _a1 error) *EthClient_BlockByHash_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *EthClient_BlockByHash_Call) RunAndReturn(run func(context.Context, common.Hash) (*types.Block, error)) *EthClient_BlockByHash_Call {
	_c.Call.Return(run)
	return _c
}

// CodeAt provides a mock function with given fields: ctx, account, blockNumber
func (_m *EthClient) CodeAt(ctx context.Context, account common.Address, blockNumber *big.Int) ([]byte, error) {
	ret := _m.Called(ctx, account, blockNumber)

	if len(ret) == 0 {
		panic("no return value specified for CodeAt")
	}

	var r0 []byte
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, common.Address, *big.Int) ([]byte, error)); ok {
		return rf(ctx, account, blockNumber)
	}
	if rf, ok := ret.Get(0).(func(context.Context, common.Address, *big.Int) []byte); ok {
		r0 = rf(ctx, account, blockNumber)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]byte)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, common.Address, *big.Int) error); ok {
		r1 = rf(ctx, account, blockNumber)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// EthClient_CodeAt_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'CodeAt'
type EthClient_CodeAt_Call struct {
	*mock.Call
}

// CodeAt is a helper method to define mock.On call
//   - ctx context.Context
//   - account common.Address
//   - blockNumber *big.Int
func (_e *EthClient_Expecter) CodeAt(ctx interface{}, account interface{}, blockNumber interface{}) *EthClient_CodeAt_Call {
	return &EthClient_CodeAt_Call{Call: _e.mock.On("CodeAt", ctx, account, blockNumber)}
}

func (_c *EthClient_CodeAt_Call) Run(run func(ctx context.Context, account common.Address, blockNumber *big.Int)) *EthClient_CodeAt_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(common.Address), args[2].(*big.Int))
	})
	return _c
}

func (_c *EthClient_CodeAt_Call) Return(_a0 []byte, _a1 error) *EthClient_CodeAt_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *EthClient_CodeAt_Call) RunAndReturn(run func(context.Context, common.Address, *big.Int) ([]byte, error)) *EthClient_CodeAt_Call {
	_c.Call.Return(run)
	return _c
}

// HeaderByNumber provides a mock function with given fields: ctx, number
func (_m *EthClient) HeaderByNumber(ctx context.Context, number *big.Int) (*types.Header, error) {
	ret := _m.Called(ctx, number)

	if len(ret) == 0 {
		panic("no return value specified for HeaderByNumber")
	}

	var r0 *types.Header
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *big.Int) (*types.Header, error)); ok {
		return rf(ctx, number)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *big.Int) *types.Header); ok {
		r0 = rf(ctx, number)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*types.Header)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *big.Int) error); ok {
		r1 = rf(ctx, number)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// EthClient_HeaderByNumber_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'HeaderByNumber'
type EthClient_HeaderByNumber_Call struct {
	*mock.Call
}

// HeaderByNumber is a helper method to define mock.On call
//   - ctx context.Context
//   - number *big.Int
func (_e *EthClient_Expecter) HeaderByNumber(ctx interface{}, number interface{}) *EthClient_HeaderByNumber_Call {
	return &EthClient_HeaderByNumber_Call{Call: _e.mock.On("HeaderByNumber", ctx, number)}
}

func (_c *EthClient_HeaderByNumber_Call) Run(run func(ctx context.Context, number *big.Int)) *EthClient_HeaderByNumber_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(*big.Int))
	})
	return _c
}

func (_c *EthClient_HeaderByNumber_Call) Return(_a0 *types.Header, _a1 error) *EthClient_HeaderByNumber_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *EthClient_HeaderByNumber_Call) RunAndReturn(run func(context.Context, *big.Int) (*types.Header, error)) *EthClient_HeaderByNumber_Call {
	_c.Call.Return(run)
	return _c
}

// TransactionReceipt provides a mock function with given fields: ctx, txHash
func (_m *EthClient) TransactionReceipt(ctx context.Context, txHash common.Hash) (*types.Receipt, error) {
	ret := _m.Called(ctx, txHash)

	if len(ret) == 0 {
		panic("no return value specified for TransactionReceipt")
	}

	var r0 *types.Receipt
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, common.Hash) (*types.Receipt, error)); ok {
		return rf(ctx, txHash)
	}
	if rf, ok := ret.Get(0).(func(context.Context, common.Hash) *types.Receipt); ok {
		r0 = rf(ctx, txHash)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*types.Receipt)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, common.Hash) error); ok {
		r1 = rf(ctx, txHash)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// EthClient_TransactionReceipt_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'TransactionReceipt'
type EthClient_TransactionReceipt_Call struct {
	*mock.Call
}

// TransactionReceipt is a helper method to define mock.On call
//   - ctx context.Context
//   - txHash common.Hash
func (_e *EthClient_Expecter) TransactionReceipt(ctx interface{}, txHash interface{}) *EthClient_TransactionReceipt_Call {
	return &EthClient_TransactionReceipt_Call{Call: _e.mock.On("TransactionReceipt", ctx, txHash)}
}

func (_c *EthClient_TransactionReceipt_Call) Run(run func(ctx context.Context, txHash common.Hash)) *EthClient_TransactionReceipt_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(common.Hash))
	})
	return _c
}

func (_c *EthClient_TransactionReceipt_Call) Return(_a0 *types.Receipt, _a1 error) *EthClient_TransactionReceipt_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *EthClient_TransactionReceipt_Call) RunAndReturn(run func(context.Context, common.Hash) (*types.Receipt, error)) *EthClient_TransactionReceipt_Call {
	_c.Call.Return(run)
	return _c
}

// NewEthClient creates a new instance of EthClient. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewEthClient(t interface {
	mock.TestingT
	Cleanup(func())
}) *EthClient {
	mock := &EthClient{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
