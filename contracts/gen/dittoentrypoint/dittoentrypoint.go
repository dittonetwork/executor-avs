// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package dittoentrypoint

import (
	"errors"
	"math/big"
	"strings"

	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/event"
)

// Reference imports to suppress errors if they are not otherwise used.
var (
	_ = errors.New
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
	_ = abi.ConvertType
)

// IDittoEntryPointWorkflow is an auto generated low-level Go binding around an user-defined struct.
type IDittoEntryPointWorkflow struct {
	VaultAddress common.Address
	WorkflowId   *big.Int
}

// DittoentrypointMetaData contains all meta data concerning the Dittoentrypoint contract.
var DittoentrypointMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"function\",\"name\":\"arrangeExecutors\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"canExecWorkflowCheck\",\"inputs\":[{\"name\":\"vaultAddress\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"workflowId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"gasPrice\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"cancelWorkflow\",\"inputs\":[{\"name\":\"workflowId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"getActiveWorkflows\",\"inputs\":[{\"name\":\"from\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"to\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"pageWorkflows\",\"type\":\"tuple[]\",\"internalType\":\"structIDittoEntryPoint.Workflow[]\",\"components\":[{\"name\":\"vaultAddress\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"workflowId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getAllActiveWorkflows\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"tuple[]\",\"internalType\":\"structIDittoEntryPoint.Workflow[]\",\"components\":[{\"name\":\"vaultAddress\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"workflowId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getAmountActiveWorkflows\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getAmountExecutors\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getSeedExternal\",\"inputs\":[{\"name\":\"blockNumber\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"seed\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"initialize\",\"inputs\":[{\"name\":\"newOwner\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_protocolFees\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"isExecutor\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"isValidExecutor\",\"inputs\":[{\"name\":\"blockNumber\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"executorAddress\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"registerExecutor\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"registerWorkflow\",\"inputs\":[{\"name\":\"workflowId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"runMultipleWorkflows\",\"inputs\":[{\"name\":\"workflows\",\"type\":\"tuple[]\",\"internalType\":\"structIDittoEntryPoint.Workflow[]\",\"components\":[{\"name\":\"vaultAddress\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"workflowId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}],\"outputs\":[{\"name\":\"\",\"type\":\"bool[]\",\"internalType\":\"bool[]\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"runWorkflow\",\"inputs\":[{\"name\":\"vaultAddress\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"workflowId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"runWorkflowWithRevert\",\"inputs\":[{\"name\":\"vaultAddress\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"workflowId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setParams\",\"inputs\":[{\"name\":\"blockSlotSizeNext\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"epochSizeNext\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"startParamsInit\",\"inputs\":[{\"name\":\"blockSlotSize\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"epochSize\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"strategy\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"delegateManager\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"unregisterExecutor\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"error\",\"name\":\"DittoEntryPoint_AlreadyInitialized\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"DittoEntryPoint_AlreadyStartParamsInitialized\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"DittoEntryPoint_WorkflowReverted\",\"inputs\":[]}]",
}

// DittoentrypointABI is the input ABI used to generate the binding from.
// Deprecated: Use DittoentrypointMetaData.ABI instead.
var DittoentrypointABI = DittoentrypointMetaData.ABI

// Dittoentrypoint is an auto generated Go binding around an Ethereum contract.
type Dittoentrypoint struct {
	DittoentrypointCaller     // Read-only binding to the contract
	DittoentrypointTransactor // Write-only binding to the contract
	DittoentrypointFilterer   // Log filterer for contract events
}

// DittoentrypointCaller is an auto generated read-only Go binding around an Ethereum contract.
type DittoentrypointCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DittoentrypointTransactor is an auto generated write-only Go binding around an Ethereum contract.
type DittoentrypointTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DittoentrypointFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type DittoentrypointFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DittoentrypointSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type DittoentrypointSession struct {
	Contract     *Dittoentrypoint  // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// DittoentrypointCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type DittoentrypointCallerSession struct {
	Contract *DittoentrypointCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts          // Call options to use throughout this session
}

// DittoentrypointTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type DittoentrypointTransactorSession struct {
	Contract     *DittoentrypointTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts          // Transaction auth options to use throughout this session
}

// DittoentrypointRaw is an auto generated low-level Go binding around an Ethereum contract.
type DittoentrypointRaw struct {
	Contract *Dittoentrypoint // Generic contract binding to access the raw methods on
}

// DittoentrypointCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type DittoentrypointCallerRaw struct {
	Contract *DittoentrypointCaller // Generic read-only contract binding to access the raw methods on
}

// DittoentrypointTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type DittoentrypointTransactorRaw struct {
	Contract *DittoentrypointTransactor // Generic write-only contract binding to access the raw methods on
}

// NewDittoentrypoint creates a new instance of Dittoentrypoint, bound to a specific deployed contract.
func NewDittoentrypoint(address common.Address, backend bind.ContractBackend) (*Dittoentrypoint, error) {
	contract, err := bindDittoentrypoint(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Dittoentrypoint{DittoentrypointCaller: DittoentrypointCaller{contract: contract}, DittoentrypointTransactor: DittoentrypointTransactor{contract: contract}, DittoentrypointFilterer: DittoentrypointFilterer{contract: contract}}, nil
}

// NewDittoentrypointCaller creates a new read-only instance of Dittoentrypoint, bound to a specific deployed contract.
func NewDittoentrypointCaller(address common.Address, caller bind.ContractCaller) (*DittoentrypointCaller, error) {
	contract, err := bindDittoentrypoint(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &DittoentrypointCaller{contract: contract}, nil
}

// NewDittoentrypointTransactor creates a new write-only instance of Dittoentrypoint, bound to a specific deployed contract.
func NewDittoentrypointTransactor(address common.Address, transactor bind.ContractTransactor) (*DittoentrypointTransactor, error) {
	contract, err := bindDittoentrypoint(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &DittoentrypointTransactor{contract: contract}, nil
}

// NewDittoentrypointFilterer creates a new log filterer instance of Dittoentrypoint, bound to a specific deployed contract.
func NewDittoentrypointFilterer(address common.Address, filterer bind.ContractFilterer) (*DittoentrypointFilterer, error) {
	contract, err := bindDittoentrypoint(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &DittoentrypointFilterer{contract: contract}, nil
}

// bindDittoentrypoint binds a generic wrapper to an already deployed contract.
func bindDittoentrypoint(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := DittoentrypointMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Dittoentrypoint *DittoentrypointRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Dittoentrypoint.Contract.DittoentrypointCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Dittoentrypoint *DittoentrypointRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Dittoentrypoint.Contract.DittoentrypointTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Dittoentrypoint *DittoentrypointRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Dittoentrypoint.Contract.DittoentrypointTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Dittoentrypoint *DittoentrypointCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Dittoentrypoint.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Dittoentrypoint *DittoentrypointTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Dittoentrypoint.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Dittoentrypoint *DittoentrypointTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Dittoentrypoint.Contract.contract.Transact(opts, method, params...)
}

// CanExecWorkflowCheck is a free data retrieval call binding the contract method 0x39b2afd4.
//
// Solidity: function canExecWorkflowCheck(address vaultAddress, uint256 workflowId, uint256 gasPrice) view returns(bool)
func (_Dittoentrypoint *DittoentrypointCaller) CanExecWorkflowCheck(opts *bind.CallOpts, vaultAddress common.Address, workflowId *big.Int, gasPrice *big.Int) (bool, error) {
	var out []interface{}
	err := _Dittoentrypoint.contract.Call(opts, &out, "canExecWorkflowCheck", vaultAddress, workflowId, gasPrice)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// CanExecWorkflowCheck is a free data retrieval call binding the contract method 0x39b2afd4.
//
// Solidity: function canExecWorkflowCheck(address vaultAddress, uint256 workflowId, uint256 gasPrice) view returns(bool)
func (_Dittoentrypoint *DittoentrypointSession) CanExecWorkflowCheck(vaultAddress common.Address, workflowId *big.Int, gasPrice *big.Int) (bool, error) {
	return _Dittoentrypoint.Contract.CanExecWorkflowCheck(&_Dittoentrypoint.CallOpts, vaultAddress, workflowId, gasPrice)
}

// CanExecWorkflowCheck is a free data retrieval call binding the contract method 0x39b2afd4.
//
// Solidity: function canExecWorkflowCheck(address vaultAddress, uint256 workflowId, uint256 gasPrice) view returns(bool)
func (_Dittoentrypoint *DittoentrypointCallerSession) CanExecWorkflowCheck(vaultAddress common.Address, workflowId *big.Int, gasPrice *big.Int) (bool, error) {
	return _Dittoentrypoint.Contract.CanExecWorkflowCheck(&_Dittoentrypoint.CallOpts, vaultAddress, workflowId, gasPrice)
}

// GetActiveWorkflows is a free data retrieval call binding the contract method 0x57d81f2a.
//
// Solidity: function getActiveWorkflows(uint256 from, uint256 to) view returns((address,uint256)[] pageWorkflows)
func (_Dittoentrypoint *DittoentrypointCaller) GetActiveWorkflows(opts *bind.CallOpts, from *big.Int, to *big.Int) ([]IDittoEntryPointWorkflow, error) {
	var out []interface{}
	err := _Dittoentrypoint.contract.Call(opts, &out, "getActiveWorkflows", from, to)

	if err != nil {
		return *new([]IDittoEntryPointWorkflow), err
	}

	out0 := *abi.ConvertType(out[0], new([]IDittoEntryPointWorkflow)).(*[]IDittoEntryPointWorkflow)

	return out0, err

}

// GetActiveWorkflows is a free data retrieval call binding the contract method 0x57d81f2a.
//
// Solidity: function getActiveWorkflows(uint256 from, uint256 to) view returns((address,uint256)[] pageWorkflows)
func (_Dittoentrypoint *DittoentrypointSession) GetActiveWorkflows(from *big.Int, to *big.Int) ([]IDittoEntryPointWorkflow, error) {
	return _Dittoentrypoint.Contract.GetActiveWorkflows(&_Dittoentrypoint.CallOpts, from, to)
}

// GetActiveWorkflows is a free data retrieval call binding the contract method 0x57d81f2a.
//
// Solidity: function getActiveWorkflows(uint256 from, uint256 to) view returns((address,uint256)[] pageWorkflows)
func (_Dittoentrypoint *DittoentrypointCallerSession) GetActiveWorkflows(from *big.Int, to *big.Int) ([]IDittoEntryPointWorkflow, error) {
	return _Dittoentrypoint.Contract.GetActiveWorkflows(&_Dittoentrypoint.CallOpts, from, to)
}

// GetAllActiveWorkflows is a free data retrieval call binding the contract method 0xcd92d3f1.
//
// Solidity: function getAllActiveWorkflows() view returns((address,uint256)[])
func (_Dittoentrypoint *DittoentrypointCaller) GetAllActiveWorkflows(opts *bind.CallOpts) ([]IDittoEntryPointWorkflow, error) {
	var out []interface{}
	err := _Dittoentrypoint.contract.Call(opts, &out, "getAllActiveWorkflows")

	if err != nil {
		return *new([]IDittoEntryPointWorkflow), err
	}

	out0 := *abi.ConvertType(out[0], new([]IDittoEntryPointWorkflow)).(*[]IDittoEntryPointWorkflow)

	return out0, err

}

// GetAllActiveWorkflows is a free data retrieval call binding the contract method 0xcd92d3f1.
//
// Solidity: function getAllActiveWorkflows() view returns((address,uint256)[])
func (_Dittoentrypoint *DittoentrypointSession) GetAllActiveWorkflows() ([]IDittoEntryPointWorkflow, error) {
	return _Dittoentrypoint.Contract.GetAllActiveWorkflows(&_Dittoentrypoint.CallOpts)
}

// GetAllActiveWorkflows is a free data retrieval call binding the contract method 0xcd92d3f1.
//
// Solidity: function getAllActiveWorkflows() view returns((address,uint256)[])
func (_Dittoentrypoint *DittoentrypointCallerSession) GetAllActiveWorkflows() ([]IDittoEntryPointWorkflow, error) {
	return _Dittoentrypoint.Contract.GetAllActiveWorkflows(&_Dittoentrypoint.CallOpts)
}

// GetAmountActiveWorkflows is a free data retrieval call binding the contract method 0x5bed4069.
//
// Solidity: function getAmountActiveWorkflows() view returns(uint256)
func (_Dittoentrypoint *DittoentrypointCaller) GetAmountActiveWorkflows(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Dittoentrypoint.contract.Call(opts, &out, "getAmountActiveWorkflows")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetAmountActiveWorkflows is a free data retrieval call binding the contract method 0x5bed4069.
//
// Solidity: function getAmountActiveWorkflows() view returns(uint256)
func (_Dittoentrypoint *DittoentrypointSession) GetAmountActiveWorkflows() (*big.Int, error) {
	return _Dittoentrypoint.Contract.GetAmountActiveWorkflows(&_Dittoentrypoint.CallOpts)
}

// GetAmountActiveWorkflows is a free data retrieval call binding the contract method 0x5bed4069.
//
// Solidity: function getAmountActiveWorkflows() view returns(uint256)
func (_Dittoentrypoint *DittoentrypointCallerSession) GetAmountActiveWorkflows() (*big.Int, error) {
	return _Dittoentrypoint.Contract.GetAmountActiveWorkflows(&_Dittoentrypoint.CallOpts)
}

// GetAmountExecutors is a free data retrieval call binding the contract method 0xcccbb947.
//
// Solidity: function getAmountExecutors() view returns(uint256)
func (_Dittoentrypoint *DittoentrypointCaller) GetAmountExecutors(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Dittoentrypoint.contract.Call(opts, &out, "getAmountExecutors")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetAmountExecutors is a free data retrieval call binding the contract method 0xcccbb947.
//
// Solidity: function getAmountExecutors() view returns(uint256)
func (_Dittoentrypoint *DittoentrypointSession) GetAmountExecutors() (*big.Int, error) {
	return _Dittoentrypoint.Contract.GetAmountExecutors(&_Dittoentrypoint.CallOpts)
}

// GetAmountExecutors is a free data retrieval call binding the contract method 0xcccbb947.
//
// Solidity: function getAmountExecutors() view returns(uint256)
func (_Dittoentrypoint *DittoentrypointCallerSession) GetAmountExecutors() (*big.Int, error) {
	return _Dittoentrypoint.Contract.GetAmountExecutors(&_Dittoentrypoint.CallOpts)
}

// GetSeedExternal is a free data retrieval call binding the contract method 0x8063f630.
//
// Solidity: function getSeedExternal(uint256 blockNumber) view returns(uint256 seed)
func (_Dittoentrypoint *DittoentrypointCaller) GetSeedExternal(opts *bind.CallOpts, blockNumber *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Dittoentrypoint.contract.Call(opts, &out, "getSeedExternal", blockNumber)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetSeedExternal is a free data retrieval call binding the contract method 0x8063f630.
//
// Solidity: function getSeedExternal(uint256 blockNumber) view returns(uint256 seed)
func (_Dittoentrypoint *DittoentrypointSession) GetSeedExternal(blockNumber *big.Int) (*big.Int, error) {
	return _Dittoentrypoint.Contract.GetSeedExternal(&_Dittoentrypoint.CallOpts, blockNumber)
}

// GetSeedExternal is a free data retrieval call binding the contract method 0x8063f630.
//
// Solidity: function getSeedExternal(uint256 blockNumber) view returns(uint256 seed)
func (_Dittoentrypoint *DittoentrypointCallerSession) GetSeedExternal(blockNumber *big.Int) (*big.Int, error) {
	return _Dittoentrypoint.Contract.GetSeedExternal(&_Dittoentrypoint.CallOpts, blockNumber)
}

// IsExecutor is a free data retrieval call binding the contract method 0xa429352f.
//
// Solidity: function isExecutor() view returns(bool)
func (_Dittoentrypoint *DittoentrypointCaller) IsExecutor(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _Dittoentrypoint.contract.Call(opts, &out, "isExecutor")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsExecutor is a free data retrieval call binding the contract method 0xa429352f.
//
// Solidity: function isExecutor() view returns(bool)
func (_Dittoentrypoint *DittoentrypointSession) IsExecutor() (bool, error) {
	return _Dittoentrypoint.Contract.IsExecutor(&_Dittoentrypoint.CallOpts)
}

// IsExecutor is a free data retrieval call binding the contract method 0xa429352f.
//
// Solidity: function isExecutor() view returns(bool)
func (_Dittoentrypoint *DittoentrypointCallerSession) IsExecutor() (bool, error) {
	return _Dittoentrypoint.Contract.IsExecutor(&_Dittoentrypoint.CallOpts)
}

// IsValidExecutor is a free data retrieval call binding the contract method 0x8f35dc01.
//
// Solidity: function isValidExecutor(uint256 blockNumber, address executorAddress) view returns(bool)
func (_Dittoentrypoint *DittoentrypointCaller) IsValidExecutor(opts *bind.CallOpts, blockNumber *big.Int, executorAddress common.Address) (bool, error) {
	var out []interface{}
	err := _Dittoentrypoint.contract.Call(opts, &out, "isValidExecutor", blockNumber, executorAddress)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsValidExecutor is a free data retrieval call binding the contract method 0x8f35dc01.
//
// Solidity: function isValidExecutor(uint256 blockNumber, address executorAddress) view returns(bool)
func (_Dittoentrypoint *DittoentrypointSession) IsValidExecutor(blockNumber *big.Int, executorAddress common.Address) (bool, error) {
	return _Dittoentrypoint.Contract.IsValidExecutor(&_Dittoentrypoint.CallOpts, blockNumber, executorAddress)
}

// IsValidExecutor is a free data retrieval call binding the contract method 0x8f35dc01.
//
// Solidity: function isValidExecutor(uint256 blockNumber, address executorAddress) view returns(bool)
func (_Dittoentrypoint *DittoentrypointCallerSession) IsValidExecutor(blockNumber *big.Int, executorAddress common.Address) (bool, error) {
	return _Dittoentrypoint.Contract.IsValidExecutor(&_Dittoentrypoint.CallOpts, blockNumber, executorAddress)
}

// ArrangeExecutors is a paid mutator transaction binding the contract method 0x92052498.
//
// Solidity: function arrangeExecutors() returns()
func (_Dittoentrypoint *DittoentrypointTransactor) ArrangeExecutors(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Dittoentrypoint.contract.Transact(opts, "arrangeExecutors")
}

// ArrangeExecutors is a paid mutator transaction binding the contract method 0x92052498.
//
// Solidity: function arrangeExecutors() returns()
func (_Dittoentrypoint *DittoentrypointSession) ArrangeExecutors() (*types.Transaction, error) {
	return _Dittoentrypoint.Contract.ArrangeExecutors(&_Dittoentrypoint.TransactOpts)
}

// ArrangeExecutors is a paid mutator transaction binding the contract method 0x92052498.
//
// Solidity: function arrangeExecutors() returns()
func (_Dittoentrypoint *DittoentrypointTransactorSession) ArrangeExecutors() (*types.Transaction, error) {
	return _Dittoentrypoint.Contract.ArrangeExecutors(&_Dittoentrypoint.TransactOpts)
}

// CancelWorkflow is a paid mutator transaction binding the contract method 0xd0c81e98.
//
// Solidity: function cancelWorkflow(uint256 workflowId) returns()
func (_Dittoentrypoint *DittoentrypointTransactor) CancelWorkflow(opts *bind.TransactOpts, workflowId *big.Int) (*types.Transaction, error) {
	return _Dittoentrypoint.contract.Transact(opts, "cancelWorkflow", workflowId)
}

// CancelWorkflow is a paid mutator transaction binding the contract method 0xd0c81e98.
//
// Solidity: function cancelWorkflow(uint256 workflowId) returns()
func (_Dittoentrypoint *DittoentrypointSession) CancelWorkflow(workflowId *big.Int) (*types.Transaction, error) {
	return _Dittoentrypoint.Contract.CancelWorkflow(&_Dittoentrypoint.TransactOpts, workflowId)
}

// CancelWorkflow is a paid mutator transaction binding the contract method 0xd0c81e98.
//
// Solidity: function cancelWorkflow(uint256 workflowId) returns()
func (_Dittoentrypoint *DittoentrypointTransactorSession) CancelWorkflow(workflowId *big.Int) (*types.Transaction, error) {
	return _Dittoentrypoint.Contract.CancelWorkflow(&_Dittoentrypoint.TransactOpts, workflowId)
}

// Initialize is a paid mutator transaction binding the contract method 0x485cc955.
//
// Solidity: function initialize(address newOwner, address _protocolFees) returns()
func (_Dittoentrypoint *DittoentrypointTransactor) Initialize(opts *bind.TransactOpts, newOwner common.Address, _protocolFees common.Address) (*types.Transaction, error) {
	return _Dittoentrypoint.contract.Transact(opts, "initialize", newOwner, _protocolFees)
}

// Initialize is a paid mutator transaction binding the contract method 0x485cc955.
//
// Solidity: function initialize(address newOwner, address _protocolFees) returns()
func (_Dittoentrypoint *DittoentrypointSession) Initialize(newOwner common.Address, _protocolFees common.Address) (*types.Transaction, error) {
	return _Dittoentrypoint.Contract.Initialize(&_Dittoentrypoint.TransactOpts, newOwner, _protocolFees)
}

// Initialize is a paid mutator transaction binding the contract method 0x485cc955.
//
// Solidity: function initialize(address newOwner, address _protocolFees) returns()
func (_Dittoentrypoint *DittoentrypointTransactorSession) Initialize(newOwner common.Address, _protocolFees common.Address) (*types.Transaction, error) {
	return _Dittoentrypoint.Contract.Initialize(&_Dittoentrypoint.TransactOpts, newOwner, _protocolFees)
}

// RegisterExecutor is a paid mutator transaction binding the contract method 0x2afabb98.
//
// Solidity: function registerExecutor() returns()
func (_Dittoentrypoint *DittoentrypointTransactor) RegisterExecutor(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Dittoentrypoint.contract.Transact(opts, "registerExecutor")
}

// RegisterExecutor is a paid mutator transaction binding the contract method 0x2afabb98.
//
// Solidity: function registerExecutor() returns()
func (_Dittoentrypoint *DittoentrypointSession) RegisterExecutor() (*types.Transaction, error) {
	return _Dittoentrypoint.Contract.RegisterExecutor(&_Dittoentrypoint.TransactOpts)
}

// RegisterExecutor is a paid mutator transaction binding the contract method 0x2afabb98.
//
// Solidity: function registerExecutor() returns()
func (_Dittoentrypoint *DittoentrypointTransactorSession) RegisterExecutor() (*types.Transaction, error) {
	return _Dittoentrypoint.Contract.RegisterExecutor(&_Dittoentrypoint.TransactOpts)
}

// RegisterWorkflow is a paid mutator transaction binding the contract method 0xa61f676b.
//
// Solidity: function registerWorkflow(uint256 workflowId) returns()
func (_Dittoentrypoint *DittoentrypointTransactor) RegisterWorkflow(opts *bind.TransactOpts, workflowId *big.Int) (*types.Transaction, error) {
	return _Dittoentrypoint.contract.Transact(opts, "registerWorkflow", workflowId)
}

// RegisterWorkflow is a paid mutator transaction binding the contract method 0xa61f676b.
//
// Solidity: function registerWorkflow(uint256 workflowId) returns()
func (_Dittoentrypoint *DittoentrypointSession) RegisterWorkflow(workflowId *big.Int) (*types.Transaction, error) {
	return _Dittoentrypoint.Contract.RegisterWorkflow(&_Dittoentrypoint.TransactOpts, workflowId)
}

// RegisterWorkflow is a paid mutator transaction binding the contract method 0xa61f676b.
//
// Solidity: function registerWorkflow(uint256 workflowId) returns()
func (_Dittoentrypoint *DittoentrypointTransactorSession) RegisterWorkflow(workflowId *big.Int) (*types.Transaction, error) {
	return _Dittoentrypoint.Contract.RegisterWorkflow(&_Dittoentrypoint.TransactOpts, workflowId)
}

// RunMultipleWorkflows is a paid mutator transaction binding the contract method 0x1370159c.
//
// Solidity: function runMultipleWorkflows((address,uint256)[] workflows) returns(bool[])
func (_Dittoentrypoint *DittoentrypointTransactor) RunMultipleWorkflows(opts *bind.TransactOpts, workflows []IDittoEntryPointWorkflow) (*types.Transaction, error) {
	return _Dittoentrypoint.contract.Transact(opts, "runMultipleWorkflows", workflows)
}

// RunMultipleWorkflows is a paid mutator transaction binding the contract method 0x1370159c.
//
// Solidity: function runMultipleWorkflows((address,uint256)[] workflows) returns(bool[])
func (_Dittoentrypoint *DittoentrypointSession) RunMultipleWorkflows(workflows []IDittoEntryPointWorkflow) (*types.Transaction, error) {
	return _Dittoentrypoint.Contract.RunMultipleWorkflows(&_Dittoentrypoint.TransactOpts, workflows)
}

// RunMultipleWorkflows is a paid mutator transaction binding the contract method 0x1370159c.
//
// Solidity: function runMultipleWorkflows((address,uint256)[] workflows) returns(bool[])
func (_Dittoentrypoint *DittoentrypointTransactorSession) RunMultipleWorkflows(workflows []IDittoEntryPointWorkflow) (*types.Transaction, error) {
	return _Dittoentrypoint.Contract.RunMultipleWorkflows(&_Dittoentrypoint.TransactOpts, workflows)
}

// RunWorkflow is a paid mutator transaction binding the contract method 0x109bf98d.
//
// Solidity: function runWorkflow(address vaultAddress, uint256 workflowId) returns(bool)
func (_Dittoentrypoint *DittoentrypointTransactor) RunWorkflow(opts *bind.TransactOpts, vaultAddress common.Address, workflowId *big.Int) (*types.Transaction, error) {
	return _Dittoentrypoint.contract.Transact(opts, "runWorkflow", vaultAddress, workflowId)
}

// RunWorkflow is a paid mutator transaction binding the contract method 0x109bf98d.
//
// Solidity: function runWorkflow(address vaultAddress, uint256 workflowId) returns(bool)
func (_Dittoentrypoint *DittoentrypointSession) RunWorkflow(vaultAddress common.Address, workflowId *big.Int) (*types.Transaction, error) {
	return _Dittoentrypoint.Contract.RunWorkflow(&_Dittoentrypoint.TransactOpts, vaultAddress, workflowId)
}

// RunWorkflow is a paid mutator transaction binding the contract method 0x109bf98d.
//
// Solidity: function runWorkflow(address vaultAddress, uint256 workflowId) returns(bool)
func (_Dittoentrypoint *DittoentrypointTransactorSession) RunWorkflow(vaultAddress common.Address, workflowId *big.Int) (*types.Transaction, error) {
	return _Dittoentrypoint.Contract.RunWorkflow(&_Dittoentrypoint.TransactOpts, vaultAddress, workflowId)
}

// RunWorkflowWithRevert is a paid mutator transaction binding the contract method 0x61c5557f.
//
// Solidity: function runWorkflowWithRevert(address vaultAddress, uint256 workflowId) returns()
func (_Dittoentrypoint *DittoentrypointTransactor) RunWorkflowWithRevert(opts *bind.TransactOpts, vaultAddress common.Address, workflowId *big.Int) (*types.Transaction, error) {
	return _Dittoentrypoint.contract.Transact(opts, "runWorkflowWithRevert", vaultAddress, workflowId)
}

// RunWorkflowWithRevert is a paid mutator transaction binding the contract method 0x61c5557f.
//
// Solidity: function runWorkflowWithRevert(address vaultAddress, uint256 workflowId) returns()
func (_Dittoentrypoint *DittoentrypointSession) RunWorkflowWithRevert(vaultAddress common.Address, workflowId *big.Int) (*types.Transaction, error) {
	return _Dittoentrypoint.Contract.RunWorkflowWithRevert(&_Dittoentrypoint.TransactOpts, vaultAddress, workflowId)
}

// RunWorkflowWithRevert is a paid mutator transaction binding the contract method 0x61c5557f.
//
// Solidity: function runWorkflowWithRevert(address vaultAddress, uint256 workflowId) returns()
func (_Dittoentrypoint *DittoentrypointTransactorSession) RunWorkflowWithRevert(vaultAddress common.Address, workflowId *big.Int) (*types.Transaction, error) {
	return _Dittoentrypoint.Contract.RunWorkflowWithRevert(&_Dittoentrypoint.TransactOpts, vaultAddress, workflowId)
}

// SetParams is a paid mutator transaction binding the contract method 0xc0324c77.
//
// Solidity: function setParams(uint256 blockSlotSizeNext, uint256 epochSizeNext) returns()
func (_Dittoentrypoint *DittoentrypointTransactor) SetParams(opts *bind.TransactOpts, blockSlotSizeNext *big.Int, epochSizeNext *big.Int) (*types.Transaction, error) {
	return _Dittoentrypoint.contract.Transact(opts, "setParams", blockSlotSizeNext, epochSizeNext)
}

// SetParams is a paid mutator transaction binding the contract method 0xc0324c77.
//
// Solidity: function setParams(uint256 blockSlotSizeNext, uint256 epochSizeNext) returns()
func (_Dittoentrypoint *DittoentrypointSession) SetParams(blockSlotSizeNext *big.Int, epochSizeNext *big.Int) (*types.Transaction, error) {
	return _Dittoentrypoint.Contract.SetParams(&_Dittoentrypoint.TransactOpts, blockSlotSizeNext, epochSizeNext)
}

// SetParams is a paid mutator transaction binding the contract method 0xc0324c77.
//
// Solidity: function setParams(uint256 blockSlotSizeNext, uint256 epochSizeNext) returns()
func (_Dittoentrypoint *DittoentrypointTransactorSession) SetParams(blockSlotSizeNext *big.Int, epochSizeNext *big.Int) (*types.Transaction, error) {
	return _Dittoentrypoint.Contract.SetParams(&_Dittoentrypoint.TransactOpts, blockSlotSizeNext, epochSizeNext)
}

// StartParamsInit is a paid mutator transaction binding the contract method 0x05ce8c4d.
//
// Solidity: function startParamsInit(uint256 blockSlotSize, uint256 epochSize, address strategy, address delegateManager) returns()
func (_Dittoentrypoint *DittoentrypointTransactor) StartParamsInit(opts *bind.TransactOpts, blockSlotSize *big.Int, epochSize *big.Int, strategy common.Address, delegateManager common.Address) (*types.Transaction, error) {
	return _Dittoentrypoint.contract.Transact(opts, "startParamsInit", blockSlotSize, epochSize, strategy, delegateManager)
}

// StartParamsInit is a paid mutator transaction binding the contract method 0x05ce8c4d.
//
// Solidity: function startParamsInit(uint256 blockSlotSize, uint256 epochSize, address strategy, address delegateManager) returns()
func (_Dittoentrypoint *DittoentrypointSession) StartParamsInit(blockSlotSize *big.Int, epochSize *big.Int, strategy common.Address, delegateManager common.Address) (*types.Transaction, error) {
	return _Dittoentrypoint.Contract.StartParamsInit(&_Dittoentrypoint.TransactOpts, blockSlotSize, epochSize, strategy, delegateManager)
}

// StartParamsInit is a paid mutator transaction binding the contract method 0x05ce8c4d.
//
// Solidity: function startParamsInit(uint256 blockSlotSize, uint256 epochSize, address strategy, address delegateManager) returns()
func (_Dittoentrypoint *DittoentrypointTransactorSession) StartParamsInit(blockSlotSize *big.Int, epochSize *big.Int, strategy common.Address, delegateManager common.Address) (*types.Transaction, error) {
	return _Dittoentrypoint.Contract.StartParamsInit(&_Dittoentrypoint.TransactOpts, blockSlotSize, epochSize, strategy, delegateManager)
}

// UnregisterExecutor is a paid mutator transaction binding the contract method 0x868a660f.
//
// Solidity: function unregisterExecutor() returns()
func (_Dittoentrypoint *DittoentrypointTransactor) UnregisterExecutor(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Dittoentrypoint.contract.Transact(opts, "unregisterExecutor")
}

// UnregisterExecutor is a paid mutator transaction binding the contract method 0x868a660f.
//
// Solidity: function unregisterExecutor() returns()
func (_Dittoentrypoint *DittoentrypointSession) UnregisterExecutor() (*types.Transaction, error) {
	return _Dittoentrypoint.Contract.UnregisterExecutor(&_Dittoentrypoint.TransactOpts)
}

// UnregisterExecutor is a paid mutator transaction binding the contract method 0x868a660f.
//
// Solidity: function unregisterExecutor() returns()
func (_Dittoentrypoint *DittoentrypointTransactorSession) UnregisterExecutor() (*types.Transaction, error) {
	return _Dittoentrypoint.Contract.UnregisterExecutor(&_Dittoentrypoint.TransactOpts)
}
