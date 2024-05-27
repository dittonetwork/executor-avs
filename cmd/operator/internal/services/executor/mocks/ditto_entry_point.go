// Code generated by mockery v2.43.1. DO NOT EDIT.

package mocks

import (
	context "context"
	big "math/big"

	common "github.com/ethereum/go-ethereum/common"

	mock "github.com/stretchr/testify/mock"

	models "github.com/dittonetwork/executor-avs/cmd/operator/internal/models"

	types "github.com/ethereum/go-ethereum/core/types"
)

// DittoEntryPoint is an autogenerated mock type for the dittoEntryPoint type
type DittoEntryPoint struct {
	mock.Mock
}

type DittoEntryPoint_Expecter struct {
	mock *mock.Mock
}

func (_m *DittoEntryPoint) EXPECT() *DittoEntryPoint_Expecter {
	return &DittoEntryPoint_Expecter{mock: &_m.Mock}
}

// GetAllActiveWorkflows provides a mock function with given fields: ctx
func (_m *DittoEntryPoint) GetAllActiveWorkflows(ctx context.Context) ([]models.Workflow, error) {
	ret := _m.Called(ctx)

	if len(ret) == 0 {
		panic("no return value specified for GetAllActiveWorkflows")
	}

	var r0 []models.Workflow
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context) ([]models.Workflow, error)); ok {
		return rf(ctx)
	}
	if rf, ok := ret.Get(0).(func(context.Context) []models.Workflow); ok {
		r0 = rf(ctx)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]models.Workflow)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context) error); ok {
		r1 = rf(ctx)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DittoEntryPoint_GetAllActiveWorkflows_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetAllActiveWorkflows'
type DittoEntryPoint_GetAllActiveWorkflows_Call struct {
	*mock.Call
}

// GetAllActiveWorkflows is a helper method to define mock.On call
//   - ctx context.Context
func (_e *DittoEntryPoint_Expecter) GetAllActiveWorkflows(ctx interface{}) *DittoEntryPoint_GetAllActiveWorkflows_Call {
	return &DittoEntryPoint_GetAllActiveWorkflows_Call{Call: _e.mock.On("GetAllActiveWorkflows", ctx)}
}

func (_c *DittoEntryPoint_GetAllActiveWorkflows_Call) Run(run func(ctx context.Context)) *DittoEntryPoint_GetAllActiveWorkflows_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context))
	})
	return _c
}

func (_c *DittoEntryPoint_GetAllActiveWorkflows_Call) Return(_a0 []models.Workflow, _a1 error) *DittoEntryPoint_GetAllActiveWorkflows_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *DittoEntryPoint_GetAllActiveWorkflows_Call) RunAndReturn(run func(context.Context) ([]models.Workflow, error)) *DittoEntryPoint_GetAllActiveWorkflows_Call {
	_c.Call.Return(run)
	return _c
}

// GetRunWorkflowTx provides a mock function with given fields: ctx, vaultAddr, workflowID
func (_m *DittoEntryPoint) GetRunWorkflowTx(ctx context.Context, vaultAddr common.Address, workflowID *big.Int) (*types.Transaction, error) {
	ret := _m.Called(ctx, vaultAddr, workflowID)

	if len(ret) == 0 {
		panic("no return value specified for GetRunWorkflowTx")
	}

	var r0 *types.Transaction
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, common.Address, *big.Int) (*types.Transaction, error)); ok {
		return rf(ctx, vaultAddr, workflowID)
	}
	if rf, ok := ret.Get(0).(func(context.Context, common.Address, *big.Int) *types.Transaction); ok {
		r0 = rf(ctx, vaultAddr, workflowID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*types.Transaction)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, common.Address, *big.Int) error); ok {
		r1 = rf(ctx, vaultAddr, workflowID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DittoEntryPoint_GetRunWorkflowTx_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetRunWorkflowTx'
type DittoEntryPoint_GetRunWorkflowTx_Call struct {
	*mock.Call
}

// GetRunWorkflowTx is a helper method to define mock.On call
//   - ctx context.Context
//   - vaultAddr common.Address
//   - workflowID *big.Int
func (_e *DittoEntryPoint_Expecter) GetRunWorkflowTx(ctx interface{}, vaultAddr interface{}, workflowID interface{}) *DittoEntryPoint_GetRunWorkflowTx_Call {
	return &DittoEntryPoint_GetRunWorkflowTx_Call{Call: _e.mock.On("GetRunWorkflowTx", ctx, vaultAddr, workflowID)}
}

func (_c *DittoEntryPoint_GetRunWorkflowTx_Call) Run(run func(ctx context.Context, vaultAddr common.Address, workflowID *big.Int)) *DittoEntryPoint_GetRunWorkflowTx_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(common.Address), args[2].(*big.Int))
	})
	return _c
}

func (_c *DittoEntryPoint_GetRunWorkflowTx_Call) Return(_a0 *types.Transaction, _a1 error) *DittoEntryPoint_GetRunWorkflowTx_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *DittoEntryPoint_GetRunWorkflowTx_Call) RunAndReturn(run func(context.Context, common.Address, *big.Int) (*types.Transaction, error)) *DittoEntryPoint_GetRunWorkflowTx_Call {
	_c.Call.Return(run)
	return _c
}

// IsExecutor provides a mock function with given fields: ctx
func (_m *DittoEntryPoint) IsExecutor(ctx context.Context) (bool, error) {
	ret := _m.Called(ctx)

	if len(ret) == 0 {
		panic("no return value specified for IsExecutor")
	}

	var r0 bool
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context) (bool, error)); ok {
		return rf(ctx)
	}
	if rf, ok := ret.Get(0).(func(context.Context) bool); ok {
		r0 = rf(ctx)
	} else {
		r0 = ret.Get(0).(bool)
	}

	if rf, ok := ret.Get(1).(func(context.Context) error); ok {
		r1 = rf(ctx)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DittoEntryPoint_IsExecutor_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'IsExecutor'
type DittoEntryPoint_IsExecutor_Call struct {
	*mock.Call
}

// IsExecutor is a helper method to define mock.On call
//   - ctx context.Context
func (_e *DittoEntryPoint_Expecter) IsExecutor(ctx interface{}) *DittoEntryPoint_IsExecutor_Call {
	return &DittoEntryPoint_IsExecutor_Call{Call: _e.mock.On("IsExecutor", ctx)}
}

func (_c *DittoEntryPoint_IsExecutor_Call) Run(run func(ctx context.Context)) *DittoEntryPoint_IsExecutor_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context))
	})
	return _c
}

func (_c *DittoEntryPoint_IsExecutor_Call) Return(_a0 bool, _a1 error) *DittoEntryPoint_IsExecutor_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *DittoEntryPoint_IsExecutor_Call) RunAndReturn(run func(context.Context) (bool, error)) *DittoEntryPoint_IsExecutor_Call {
	_c.Call.Return(run)
	return _c
}

// IsValidExecutor provides a mock function with given fields: ctx, blockNumber
func (_m *DittoEntryPoint) IsValidExecutor(ctx context.Context, blockNumber *big.Int) (bool, error) {
	ret := _m.Called(ctx, blockNumber)

	if len(ret) == 0 {
		panic("no return value specified for IsValidExecutor")
	}

	var r0 bool
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *big.Int) (bool, error)); ok {
		return rf(ctx, blockNumber)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *big.Int) bool); ok {
		r0 = rf(ctx, blockNumber)
	} else {
		r0 = ret.Get(0).(bool)
	}

	if rf, ok := ret.Get(1).(func(context.Context, *big.Int) error); ok {
		r1 = rf(ctx, blockNumber)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DittoEntryPoint_IsValidExecutor_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'IsValidExecutor'
type DittoEntryPoint_IsValidExecutor_Call struct {
	*mock.Call
}

// IsValidExecutor is a helper method to define mock.On call
//   - ctx context.Context
//   - blockNumber *big.Int
func (_e *DittoEntryPoint_Expecter) IsValidExecutor(ctx interface{}, blockNumber interface{}) *DittoEntryPoint_IsValidExecutor_Call {
	return &DittoEntryPoint_IsValidExecutor_Call{Call: _e.mock.On("IsValidExecutor", ctx, blockNumber)}
}

func (_c *DittoEntryPoint_IsValidExecutor_Call) Run(run func(ctx context.Context, blockNumber *big.Int)) *DittoEntryPoint_IsValidExecutor_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(*big.Int))
	})
	return _c
}

func (_c *DittoEntryPoint_IsValidExecutor_Call) Return(_a0 bool, _a1 error) *DittoEntryPoint_IsValidExecutor_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *DittoEntryPoint_IsValidExecutor_Call) RunAndReturn(run func(context.Context, *big.Int) (bool, error)) *DittoEntryPoint_IsValidExecutor_Call {
	_c.Call.Return(run)
	return _c
}

// RunMultipleWorkflows provides a mock function with given fields: ctx, workflows
func (_m *DittoEntryPoint) RunMultipleWorkflows(ctx context.Context, workflows []models.Workflow) (*types.Transaction, error) {
	ret := _m.Called(ctx, workflows)

	if len(ret) == 0 {
		panic("no return value specified for RunMultipleWorkflows")
	}

	var r0 *types.Transaction
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, []models.Workflow) (*types.Transaction, error)); ok {
		return rf(ctx, workflows)
	}
	if rf, ok := ret.Get(0).(func(context.Context, []models.Workflow) *types.Transaction); ok {
		r0 = rf(ctx, workflows)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*types.Transaction)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, []models.Workflow) error); ok {
		r1 = rf(ctx, workflows)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DittoEntryPoint_RunMultipleWorkflows_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'RunMultipleWorkflows'
type DittoEntryPoint_RunMultipleWorkflows_Call struct {
	*mock.Call
}

// RunMultipleWorkflows is a helper method to define mock.On call
//   - ctx context.Context
//   - workflows []models.Workflow
func (_e *DittoEntryPoint_Expecter) RunMultipleWorkflows(ctx interface{}, workflows interface{}) *DittoEntryPoint_RunMultipleWorkflows_Call {
	return &DittoEntryPoint_RunMultipleWorkflows_Call{Call: _e.mock.On("RunMultipleWorkflows", ctx, workflows)}
}

func (_c *DittoEntryPoint_RunMultipleWorkflows_Call) Run(run func(ctx context.Context, workflows []models.Workflow)) *DittoEntryPoint_RunMultipleWorkflows_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].([]models.Workflow))
	})
	return _c
}

func (_c *DittoEntryPoint_RunMultipleWorkflows_Call) Return(_a0 *types.Transaction, _a1 error) *DittoEntryPoint_RunMultipleWorkflows_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *DittoEntryPoint_RunMultipleWorkflows_Call) RunAndReturn(run func(context.Context, []models.Workflow) (*types.Transaction, error)) *DittoEntryPoint_RunMultipleWorkflows_Call {
	_c.Call.Return(run)
	return _c
}

// NewDittoEntryPoint creates a new instance of DittoEntryPoint. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewDittoEntryPoint(t interface {
	mock.TestingT
	Cleanup(func())
}) *DittoEntryPoint {
	mock := &DittoEntryPoint{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
