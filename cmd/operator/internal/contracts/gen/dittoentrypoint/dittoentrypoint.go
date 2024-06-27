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
	ABI: "[{\"type\":\"function\",\"name\":\"activateExecutor\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"addStrategy\",\"inputs\":[{\"name\":\"_strategy\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"arrangeExecutors\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"calculateOperatorAVSRegistrationDigestHash\",\"inputs\":[{\"name\":\"operator\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"salt\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"expiry\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"canExecWorkflowCheck\",\"inputs\":[{\"name\":\"vaultAddress\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"workflowId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"gasPrice\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"cancelWorkflow\",\"inputs\":[{\"name\":\"workflowId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"deactivateExecutor\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"getActiveWorkflows\",\"inputs\":[{\"name\":\"from\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"to\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"pageWorkflows\",\"type\":\"tuple[]\",\"internalType\":\"structIDittoEntryPoint.Workflow[]\",\"components\":[{\"name\":\"vaultAddress\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"workflowId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getAllActiveWorkflows\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"tuple[]\",\"internalType\":\"structIDittoEntryPoint.Workflow[]\",\"components\":[{\"name\":\"vaultAddress\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"workflowId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getAmountActiveWorkflows\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getAmountExecutors\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getAssignedOperator\",\"inputs\":[{\"name\":\"delegatedSigner\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getSeedExternal\",\"inputs\":[{\"name\":\"blockNumber\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"seed\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"initialize\",\"inputs\":[{\"name\":\"newOwner\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_protocolFees\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"isExecutor\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"isValidExecutor\",\"inputs\":[{\"name\":\"blockNumber\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"executorAddress\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"registerWorkflow\",\"inputs\":[{\"name\":\"workflowId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"runMultipleWorkflows\",\"inputs\":[{\"name\":\"workflows\",\"type\":\"tuple[]\",\"internalType\":\"structIDittoEntryPoint.Workflow[]\",\"components\":[{\"name\":\"vaultAddress\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"workflowId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}],\"outputs\":[{\"name\":\"\",\"type\":\"bool[]\",\"internalType\":\"bool[]\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"runWorkflowWithRevert\",\"inputs\":[{\"name\":\"vaultAddress\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"workflowId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"runWorkflowWithoutRevert\",\"inputs\":[{\"name\":\"vaultAddress\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"workflowId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setDelegatedSigner\",\"inputs\":[{\"name\":\"newDelegatedSigner\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setParams\",\"inputs\":[{\"name\":\"blockSlotSizeNext\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"epochSizeNext\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setStrategies\",\"inputs\":[{\"name\":\"_strategies\",\"type\":\"address[]\",\"internalType\":\"address[]\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"startParamsInit\",\"inputs\":[{\"name\":\"blockSlotSize\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"epochSize\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"strategies\",\"type\":\"address[]\",\"internalType\":\"address[]\"},{\"name\":\"delegateManager\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"strategyManager\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"event\",\"name\":\"DittoEntryPointCancelWorkflow\",\"inputs\":[{\"name\":\"vault\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"workflowId\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"DittoEntryPointExecuterRegister\",\"inputs\":[{\"name\":\"executor\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"DittoEntryPointExecuterUnregister\",\"inputs\":[{\"name\":\"executor\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"DittoEntryPointRegisterWorkflow\",\"inputs\":[{\"name\":\"vault\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"workflowId\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"DittoEntryPointWorkflowSuccess\",\"inputs\":[{\"name\":\"vault\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"workflowId\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"error\",\"name\":\"DittoEntryPoint_AlreadyInitialized\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"DittoEntryPoint_AlreadyStartParamsInitialized\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"DittoEntryPoint_WorkflowReverted\",\"inputs\":[]}]",
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

// CalculateOperatorAVSRegistrationDigestHash is a free data retrieval call binding the contract method 0x8b82bf49.
//
// Solidity: function calculateOperatorAVSRegistrationDigestHash(address operator, bytes32 salt, uint256 expiry) view returns(bytes32)
func (_Dittoentrypoint *DittoentrypointCaller) CalculateOperatorAVSRegistrationDigestHash(opts *bind.CallOpts, operator common.Address, salt [32]byte, expiry *big.Int) ([32]byte, error) {
	var out []interface{}
	err := _Dittoentrypoint.contract.Call(opts, &out, "calculateOperatorAVSRegistrationDigestHash", operator, salt, expiry)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// CalculateOperatorAVSRegistrationDigestHash is a free data retrieval call binding the contract method 0x8b82bf49.
//
// Solidity: function calculateOperatorAVSRegistrationDigestHash(address operator, bytes32 salt, uint256 expiry) view returns(bytes32)
func (_Dittoentrypoint *DittoentrypointSession) CalculateOperatorAVSRegistrationDigestHash(operator common.Address, salt [32]byte, expiry *big.Int) ([32]byte, error) {
	return _Dittoentrypoint.Contract.CalculateOperatorAVSRegistrationDigestHash(&_Dittoentrypoint.CallOpts, operator, salt, expiry)
}

// CalculateOperatorAVSRegistrationDigestHash is a free data retrieval call binding the contract method 0x8b82bf49.
//
// Solidity: function calculateOperatorAVSRegistrationDigestHash(address operator, bytes32 salt, uint256 expiry) view returns(bytes32)
func (_Dittoentrypoint *DittoentrypointCallerSession) CalculateOperatorAVSRegistrationDigestHash(operator common.Address, salt [32]byte, expiry *big.Int) ([32]byte, error) {
	return _Dittoentrypoint.Contract.CalculateOperatorAVSRegistrationDigestHash(&_Dittoentrypoint.CallOpts, operator, salt, expiry)
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

// GetAssignedOperator is a free data retrieval call binding the contract method 0xe1a3fc74.
//
// Solidity: function getAssignedOperator(address delegatedSigner) view returns(address)
func (_Dittoentrypoint *DittoentrypointCaller) GetAssignedOperator(opts *bind.CallOpts, delegatedSigner common.Address) (common.Address, error) {
	var out []interface{}
	err := _Dittoentrypoint.contract.Call(opts, &out, "getAssignedOperator", delegatedSigner)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetAssignedOperator is a free data retrieval call binding the contract method 0xe1a3fc74.
//
// Solidity: function getAssignedOperator(address delegatedSigner) view returns(address)
func (_Dittoentrypoint *DittoentrypointSession) GetAssignedOperator(delegatedSigner common.Address) (common.Address, error) {
	return _Dittoentrypoint.Contract.GetAssignedOperator(&_Dittoentrypoint.CallOpts, delegatedSigner)
}

// GetAssignedOperator is a free data retrieval call binding the contract method 0xe1a3fc74.
//
// Solidity: function getAssignedOperator(address delegatedSigner) view returns(address)
func (_Dittoentrypoint *DittoentrypointCallerSession) GetAssignedOperator(delegatedSigner common.Address) (common.Address, error) {
	return _Dittoentrypoint.Contract.GetAssignedOperator(&_Dittoentrypoint.CallOpts, delegatedSigner)
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

// ActivateExecutor is a paid mutator transaction binding the contract method 0x4a1258ec.
//
// Solidity: function activateExecutor() returns()
func (_Dittoentrypoint *DittoentrypointTransactor) ActivateExecutor(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Dittoentrypoint.contract.Transact(opts, "activateExecutor")
}

// ActivateExecutor is a paid mutator transaction binding the contract method 0x4a1258ec.
//
// Solidity: function activateExecutor() returns()
func (_Dittoentrypoint *DittoentrypointSession) ActivateExecutor() (*types.Transaction, error) {
	return _Dittoentrypoint.Contract.ActivateExecutor(&_Dittoentrypoint.TransactOpts)
}

// ActivateExecutor is a paid mutator transaction binding the contract method 0x4a1258ec.
//
// Solidity: function activateExecutor() returns()
func (_Dittoentrypoint *DittoentrypointTransactorSession) ActivateExecutor() (*types.Transaction, error) {
	return _Dittoentrypoint.Contract.ActivateExecutor(&_Dittoentrypoint.TransactOpts)
}

// AddStrategy is a paid mutator transaction binding the contract method 0x223e5479.
//
// Solidity: function addStrategy(address _strategy) returns()
func (_Dittoentrypoint *DittoentrypointTransactor) AddStrategy(opts *bind.TransactOpts, _strategy common.Address) (*types.Transaction, error) {
	return _Dittoentrypoint.contract.Transact(opts, "addStrategy", _strategy)
}

// AddStrategy is a paid mutator transaction binding the contract method 0x223e5479.
//
// Solidity: function addStrategy(address _strategy) returns()
func (_Dittoentrypoint *DittoentrypointSession) AddStrategy(_strategy common.Address) (*types.Transaction, error) {
	return _Dittoentrypoint.Contract.AddStrategy(&_Dittoentrypoint.TransactOpts, _strategy)
}

// AddStrategy is a paid mutator transaction binding the contract method 0x223e5479.
//
// Solidity: function addStrategy(address _strategy) returns()
func (_Dittoentrypoint *DittoentrypointTransactorSession) AddStrategy(_strategy common.Address) (*types.Transaction, error) {
	return _Dittoentrypoint.Contract.AddStrategy(&_Dittoentrypoint.TransactOpts, _strategy)
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

// DeactivateExecutor is a paid mutator transaction binding the contract method 0x8b7b985c.
//
// Solidity: function deactivateExecutor() returns()
func (_Dittoentrypoint *DittoentrypointTransactor) DeactivateExecutor(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Dittoentrypoint.contract.Transact(opts, "deactivateExecutor")
}

// DeactivateExecutor is a paid mutator transaction binding the contract method 0x8b7b985c.
//
// Solidity: function deactivateExecutor() returns()
func (_Dittoentrypoint *DittoentrypointSession) DeactivateExecutor() (*types.Transaction, error) {
	return _Dittoentrypoint.Contract.DeactivateExecutor(&_Dittoentrypoint.TransactOpts)
}

// DeactivateExecutor is a paid mutator transaction binding the contract method 0x8b7b985c.
//
// Solidity: function deactivateExecutor() returns()
func (_Dittoentrypoint *DittoentrypointTransactorSession) DeactivateExecutor() (*types.Transaction, error) {
	return _Dittoentrypoint.Contract.DeactivateExecutor(&_Dittoentrypoint.TransactOpts)
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

// RunWorkflowWithoutRevert is a paid mutator transaction binding the contract method 0x46e2a68a.
//
// Solidity: function runWorkflowWithoutRevert(address vaultAddress, uint256 workflowId) returns(bool, bytes)
func (_Dittoentrypoint *DittoentrypointTransactor) RunWorkflowWithoutRevert(opts *bind.TransactOpts, vaultAddress common.Address, workflowId *big.Int) (*types.Transaction, error) {
	return _Dittoentrypoint.contract.Transact(opts, "runWorkflowWithoutRevert", vaultAddress, workflowId)
}

// RunWorkflowWithoutRevert is a paid mutator transaction binding the contract method 0x46e2a68a.
//
// Solidity: function runWorkflowWithoutRevert(address vaultAddress, uint256 workflowId) returns(bool, bytes)
func (_Dittoentrypoint *DittoentrypointSession) RunWorkflowWithoutRevert(vaultAddress common.Address, workflowId *big.Int) (*types.Transaction, error) {
	return _Dittoentrypoint.Contract.RunWorkflowWithoutRevert(&_Dittoentrypoint.TransactOpts, vaultAddress, workflowId)
}

// RunWorkflowWithoutRevert is a paid mutator transaction binding the contract method 0x46e2a68a.
//
// Solidity: function runWorkflowWithoutRevert(address vaultAddress, uint256 workflowId) returns(bool, bytes)
func (_Dittoentrypoint *DittoentrypointTransactorSession) RunWorkflowWithoutRevert(vaultAddress common.Address, workflowId *big.Int) (*types.Transaction, error) {
	return _Dittoentrypoint.Contract.RunWorkflowWithoutRevert(&_Dittoentrypoint.TransactOpts, vaultAddress, workflowId)
}

// SetDelegatedSigner is a paid mutator transaction binding the contract method 0x53863613.
//
// Solidity: function setDelegatedSigner(address newDelegatedSigner) returns()
func (_Dittoentrypoint *DittoentrypointTransactor) SetDelegatedSigner(opts *bind.TransactOpts, newDelegatedSigner common.Address) (*types.Transaction, error) {
	return _Dittoentrypoint.contract.Transact(opts, "setDelegatedSigner", newDelegatedSigner)
}

// SetDelegatedSigner is a paid mutator transaction binding the contract method 0x53863613.
//
// Solidity: function setDelegatedSigner(address newDelegatedSigner) returns()
func (_Dittoentrypoint *DittoentrypointSession) SetDelegatedSigner(newDelegatedSigner common.Address) (*types.Transaction, error) {
	return _Dittoentrypoint.Contract.SetDelegatedSigner(&_Dittoentrypoint.TransactOpts, newDelegatedSigner)
}

// SetDelegatedSigner is a paid mutator transaction binding the contract method 0x53863613.
//
// Solidity: function setDelegatedSigner(address newDelegatedSigner) returns()
func (_Dittoentrypoint *DittoentrypointTransactorSession) SetDelegatedSigner(newDelegatedSigner common.Address) (*types.Transaction, error) {
	return _Dittoentrypoint.Contract.SetDelegatedSigner(&_Dittoentrypoint.TransactOpts, newDelegatedSigner)
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

// SetStrategies is a paid mutator transaction binding the contract method 0x9a4620b7.
//
// Solidity: function setStrategies(address[] _strategies) returns()
func (_Dittoentrypoint *DittoentrypointTransactor) SetStrategies(opts *bind.TransactOpts, _strategies []common.Address) (*types.Transaction, error) {
	return _Dittoentrypoint.contract.Transact(opts, "setStrategies", _strategies)
}

// SetStrategies is a paid mutator transaction binding the contract method 0x9a4620b7.
//
// Solidity: function setStrategies(address[] _strategies) returns()
func (_Dittoentrypoint *DittoentrypointSession) SetStrategies(_strategies []common.Address) (*types.Transaction, error) {
	return _Dittoentrypoint.Contract.SetStrategies(&_Dittoentrypoint.TransactOpts, _strategies)
}

// SetStrategies is a paid mutator transaction binding the contract method 0x9a4620b7.
//
// Solidity: function setStrategies(address[] _strategies) returns()
func (_Dittoentrypoint *DittoentrypointTransactorSession) SetStrategies(_strategies []common.Address) (*types.Transaction, error) {
	return _Dittoentrypoint.Contract.SetStrategies(&_Dittoentrypoint.TransactOpts, _strategies)
}

// StartParamsInit is a paid mutator transaction binding the contract method 0x6f9a0cda.
//
// Solidity: function startParamsInit(uint256 blockSlotSize, uint256 epochSize, address[] strategies, address delegateManager, address strategyManager) returns()
func (_Dittoentrypoint *DittoentrypointTransactor) StartParamsInit(opts *bind.TransactOpts, blockSlotSize *big.Int, epochSize *big.Int, strategies []common.Address, delegateManager common.Address, strategyManager common.Address) (*types.Transaction, error) {
	return _Dittoentrypoint.contract.Transact(opts, "startParamsInit", blockSlotSize, epochSize, strategies, delegateManager, strategyManager)
}

// StartParamsInit is a paid mutator transaction binding the contract method 0x6f9a0cda.
//
// Solidity: function startParamsInit(uint256 blockSlotSize, uint256 epochSize, address[] strategies, address delegateManager, address strategyManager) returns()
func (_Dittoentrypoint *DittoentrypointSession) StartParamsInit(blockSlotSize *big.Int, epochSize *big.Int, strategies []common.Address, delegateManager common.Address, strategyManager common.Address) (*types.Transaction, error) {
	return _Dittoentrypoint.Contract.StartParamsInit(&_Dittoentrypoint.TransactOpts, blockSlotSize, epochSize, strategies, delegateManager, strategyManager)
}

// StartParamsInit is a paid mutator transaction binding the contract method 0x6f9a0cda.
//
// Solidity: function startParamsInit(uint256 blockSlotSize, uint256 epochSize, address[] strategies, address delegateManager, address strategyManager) returns()
func (_Dittoentrypoint *DittoentrypointTransactorSession) StartParamsInit(blockSlotSize *big.Int, epochSize *big.Int, strategies []common.Address, delegateManager common.Address, strategyManager common.Address) (*types.Transaction, error) {
	return _Dittoentrypoint.Contract.StartParamsInit(&_Dittoentrypoint.TransactOpts, blockSlotSize, epochSize, strategies, delegateManager, strategyManager)
}

// DittoentrypointDittoEntryPointCancelWorkflowIterator is returned from FilterDittoEntryPointCancelWorkflow and is used to iterate over the raw logs and unpacked data for DittoEntryPointCancelWorkflow events raised by the Dittoentrypoint contract.
type DittoentrypointDittoEntryPointCancelWorkflowIterator struct {
	Event *DittoentrypointDittoEntryPointCancelWorkflow // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *DittoentrypointDittoEntryPointCancelWorkflowIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DittoentrypointDittoEntryPointCancelWorkflow)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(DittoentrypointDittoEntryPointCancelWorkflow)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *DittoentrypointDittoEntryPointCancelWorkflowIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DittoentrypointDittoEntryPointCancelWorkflowIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DittoentrypointDittoEntryPointCancelWorkflow represents a DittoEntryPointCancelWorkflow event raised by the Dittoentrypoint contract.
type DittoentrypointDittoEntryPointCancelWorkflow struct {
	Vault      common.Address
	WorkflowId *big.Int
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterDittoEntryPointCancelWorkflow is a free log retrieval operation binding the contract event 0x7e0aef7b6397c5da713e203111cfb90ea0c9291d8ef058f35dae92313ad7a707.
//
// Solidity: event DittoEntryPointCancelWorkflow(address indexed vault, uint256 workflowId)
func (_Dittoentrypoint *DittoentrypointFilterer) FilterDittoEntryPointCancelWorkflow(opts *bind.FilterOpts, vault []common.Address) (*DittoentrypointDittoEntryPointCancelWorkflowIterator, error) {

	var vaultRule []interface{}
	for _, vaultItem := range vault {
		vaultRule = append(vaultRule, vaultItem)
	}

	logs, sub, err := _Dittoentrypoint.contract.FilterLogs(opts, "DittoEntryPointCancelWorkflow", vaultRule)
	if err != nil {
		return nil, err
	}
	return &DittoentrypointDittoEntryPointCancelWorkflowIterator{contract: _Dittoentrypoint.contract, event: "DittoEntryPointCancelWorkflow", logs: logs, sub: sub}, nil
}

// WatchDittoEntryPointCancelWorkflow is a free log subscription operation binding the contract event 0x7e0aef7b6397c5da713e203111cfb90ea0c9291d8ef058f35dae92313ad7a707.
//
// Solidity: event DittoEntryPointCancelWorkflow(address indexed vault, uint256 workflowId)
func (_Dittoentrypoint *DittoentrypointFilterer) WatchDittoEntryPointCancelWorkflow(opts *bind.WatchOpts, sink chan<- *DittoentrypointDittoEntryPointCancelWorkflow, vault []common.Address) (event.Subscription, error) {

	var vaultRule []interface{}
	for _, vaultItem := range vault {
		vaultRule = append(vaultRule, vaultItem)
	}

	logs, sub, err := _Dittoentrypoint.contract.WatchLogs(opts, "DittoEntryPointCancelWorkflow", vaultRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DittoentrypointDittoEntryPointCancelWorkflow)
				if err := _Dittoentrypoint.contract.UnpackLog(event, "DittoEntryPointCancelWorkflow", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseDittoEntryPointCancelWorkflow is a log parse operation binding the contract event 0x7e0aef7b6397c5da713e203111cfb90ea0c9291d8ef058f35dae92313ad7a707.
//
// Solidity: event DittoEntryPointCancelWorkflow(address indexed vault, uint256 workflowId)
func (_Dittoentrypoint *DittoentrypointFilterer) ParseDittoEntryPointCancelWorkflow(log types.Log) (*DittoentrypointDittoEntryPointCancelWorkflow, error) {
	event := new(DittoentrypointDittoEntryPointCancelWorkflow)
	if err := _Dittoentrypoint.contract.UnpackLog(event, "DittoEntryPointCancelWorkflow", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// DittoentrypointDittoEntryPointExecuterRegisterIterator is returned from FilterDittoEntryPointExecuterRegister and is used to iterate over the raw logs and unpacked data for DittoEntryPointExecuterRegister events raised by the Dittoentrypoint contract.
type DittoentrypointDittoEntryPointExecuterRegisterIterator struct {
	Event *DittoentrypointDittoEntryPointExecuterRegister // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *DittoentrypointDittoEntryPointExecuterRegisterIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DittoentrypointDittoEntryPointExecuterRegister)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(DittoentrypointDittoEntryPointExecuterRegister)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *DittoentrypointDittoEntryPointExecuterRegisterIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DittoentrypointDittoEntryPointExecuterRegisterIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DittoentrypointDittoEntryPointExecuterRegister represents a DittoEntryPointExecuterRegister event raised by the Dittoentrypoint contract.
type DittoentrypointDittoEntryPointExecuterRegister struct {
	Executor common.Address
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterDittoEntryPointExecuterRegister is a free log retrieval operation binding the contract event 0xe16d378d46c890b88b378d2ef260e841dc24ec60fc189be6e991fc540dca12c9.
//
// Solidity: event DittoEntryPointExecuterRegister(address indexed executor)
func (_Dittoentrypoint *DittoentrypointFilterer) FilterDittoEntryPointExecuterRegister(opts *bind.FilterOpts, executor []common.Address) (*DittoentrypointDittoEntryPointExecuterRegisterIterator, error) {

	var executorRule []interface{}
	for _, executorItem := range executor {
		executorRule = append(executorRule, executorItem)
	}

	logs, sub, err := _Dittoentrypoint.contract.FilterLogs(opts, "DittoEntryPointExecuterRegister", executorRule)
	if err != nil {
		return nil, err
	}
	return &DittoentrypointDittoEntryPointExecuterRegisterIterator{contract: _Dittoentrypoint.contract, event: "DittoEntryPointExecuterRegister", logs: logs, sub: sub}, nil
}

// WatchDittoEntryPointExecuterRegister is a free log subscription operation binding the contract event 0xe16d378d46c890b88b378d2ef260e841dc24ec60fc189be6e991fc540dca12c9.
//
// Solidity: event DittoEntryPointExecuterRegister(address indexed executor)
func (_Dittoentrypoint *DittoentrypointFilterer) WatchDittoEntryPointExecuterRegister(opts *bind.WatchOpts, sink chan<- *DittoentrypointDittoEntryPointExecuterRegister, executor []common.Address) (event.Subscription, error) {

	var executorRule []interface{}
	for _, executorItem := range executor {
		executorRule = append(executorRule, executorItem)
	}

	logs, sub, err := _Dittoentrypoint.contract.WatchLogs(opts, "DittoEntryPointExecuterRegister", executorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DittoentrypointDittoEntryPointExecuterRegister)
				if err := _Dittoentrypoint.contract.UnpackLog(event, "DittoEntryPointExecuterRegister", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseDittoEntryPointExecuterRegister is a log parse operation binding the contract event 0xe16d378d46c890b88b378d2ef260e841dc24ec60fc189be6e991fc540dca12c9.
//
// Solidity: event DittoEntryPointExecuterRegister(address indexed executor)
func (_Dittoentrypoint *DittoentrypointFilterer) ParseDittoEntryPointExecuterRegister(log types.Log) (*DittoentrypointDittoEntryPointExecuterRegister, error) {
	event := new(DittoentrypointDittoEntryPointExecuterRegister)
	if err := _Dittoentrypoint.contract.UnpackLog(event, "DittoEntryPointExecuterRegister", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// DittoentrypointDittoEntryPointExecuterUnregisterIterator is returned from FilterDittoEntryPointExecuterUnregister and is used to iterate over the raw logs and unpacked data for DittoEntryPointExecuterUnregister events raised by the Dittoentrypoint contract.
type DittoentrypointDittoEntryPointExecuterUnregisterIterator struct {
	Event *DittoentrypointDittoEntryPointExecuterUnregister // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *DittoentrypointDittoEntryPointExecuterUnregisterIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DittoentrypointDittoEntryPointExecuterUnregister)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(DittoentrypointDittoEntryPointExecuterUnregister)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *DittoentrypointDittoEntryPointExecuterUnregisterIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DittoentrypointDittoEntryPointExecuterUnregisterIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DittoentrypointDittoEntryPointExecuterUnregister represents a DittoEntryPointExecuterUnregister event raised by the Dittoentrypoint contract.
type DittoentrypointDittoEntryPointExecuterUnregister struct {
	Executor common.Address
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterDittoEntryPointExecuterUnregister is a free log retrieval operation binding the contract event 0xca999231fc65bf68e10375fc21ab441f4f58421c91854effcb9707f479efffb0.
//
// Solidity: event DittoEntryPointExecuterUnregister(address indexed executor)
func (_Dittoentrypoint *DittoentrypointFilterer) FilterDittoEntryPointExecuterUnregister(opts *bind.FilterOpts, executor []common.Address) (*DittoentrypointDittoEntryPointExecuterUnregisterIterator, error) {

	var executorRule []interface{}
	for _, executorItem := range executor {
		executorRule = append(executorRule, executorItem)
	}

	logs, sub, err := _Dittoentrypoint.contract.FilterLogs(opts, "DittoEntryPointExecuterUnregister", executorRule)
	if err != nil {
		return nil, err
	}
	return &DittoentrypointDittoEntryPointExecuterUnregisterIterator{contract: _Dittoentrypoint.contract, event: "DittoEntryPointExecuterUnregister", logs: logs, sub: sub}, nil
}

// WatchDittoEntryPointExecuterUnregister is a free log subscription operation binding the contract event 0xca999231fc65bf68e10375fc21ab441f4f58421c91854effcb9707f479efffb0.
//
// Solidity: event DittoEntryPointExecuterUnregister(address indexed executor)
func (_Dittoentrypoint *DittoentrypointFilterer) WatchDittoEntryPointExecuterUnregister(opts *bind.WatchOpts, sink chan<- *DittoentrypointDittoEntryPointExecuterUnregister, executor []common.Address) (event.Subscription, error) {

	var executorRule []interface{}
	for _, executorItem := range executor {
		executorRule = append(executorRule, executorItem)
	}

	logs, sub, err := _Dittoentrypoint.contract.WatchLogs(opts, "DittoEntryPointExecuterUnregister", executorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DittoentrypointDittoEntryPointExecuterUnregister)
				if err := _Dittoentrypoint.contract.UnpackLog(event, "DittoEntryPointExecuterUnregister", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseDittoEntryPointExecuterUnregister is a log parse operation binding the contract event 0xca999231fc65bf68e10375fc21ab441f4f58421c91854effcb9707f479efffb0.
//
// Solidity: event DittoEntryPointExecuterUnregister(address indexed executor)
func (_Dittoentrypoint *DittoentrypointFilterer) ParseDittoEntryPointExecuterUnregister(log types.Log) (*DittoentrypointDittoEntryPointExecuterUnregister, error) {
	event := new(DittoentrypointDittoEntryPointExecuterUnregister)
	if err := _Dittoentrypoint.contract.UnpackLog(event, "DittoEntryPointExecuterUnregister", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// DittoentrypointDittoEntryPointRegisterWorkflowIterator is returned from FilterDittoEntryPointRegisterWorkflow and is used to iterate over the raw logs and unpacked data for DittoEntryPointRegisterWorkflow events raised by the Dittoentrypoint contract.
type DittoentrypointDittoEntryPointRegisterWorkflowIterator struct {
	Event *DittoentrypointDittoEntryPointRegisterWorkflow // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *DittoentrypointDittoEntryPointRegisterWorkflowIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DittoentrypointDittoEntryPointRegisterWorkflow)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(DittoentrypointDittoEntryPointRegisterWorkflow)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *DittoentrypointDittoEntryPointRegisterWorkflowIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DittoentrypointDittoEntryPointRegisterWorkflowIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DittoentrypointDittoEntryPointRegisterWorkflow represents a DittoEntryPointRegisterWorkflow event raised by the Dittoentrypoint contract.
type DittoentrypointDittoEntryPointRegisterWorkflow struct {
	Vault      common.Address
	WorkflowId *big.Int
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterDittoEntryPointRegisterWorkflow is a free log retrieval operation binding the contract event 0x7d89995e8cf5bdb0963851260411a4a799c0a5a754120fab6b126845ce4bdf60.
//
// Solidity: event DittoEntryPointRegisterWorkflow(address indexed vault, uint256 workflowId)
func (_Dittoentrypoint *DittoentrypointFilterer) FilterDittoEntryPointRegisterWorkflow(opts *bind.FilterOpts, vault []common.Address) (*DittoentrypointDittoEntryPointRegisterWorkflowIterator, error) {

	var vaultRule []interface{}
	for _, vaultItem := range vault {
		vaultRule = append(vaultRule, vaultItem)
	}

	logs, sub, err := _Dittoentrypoint.contract.FilterLogs(opts, "DittoEntryPointRegisterWorkflow", vaultRule)
	if err != nil {
		return nil, err
	}
	return &DittoentrypointDittoEntryPointRegisterWorkflowIterator{contract: _Dittoentrypoint.contract, event: "DittoEntryPointRegisterWorkflow", logs: logs, sub: sub}, nil
}

// WatchDittoEntryPointRegisterWorkflow is a free log subscription operation binding the contract event 0x7d89995e8cf5bdb0963851260411a4a799c0a5a754120fab6b126845ce4bdf60.
//
// Solidity: event DittoEntryPointRegisterWorkflow(address indexed vault, uint256 workflowId)
func (_Dittoentrypoint *DittoentrypointFilterer) WatchDittoEntryPointRegisterWorkflow(opts *bind.WatchOpts, sink chan<- *DittoentrypointDittoEntryPointRegisterWorkflow, vault []common.Address) (event.Subscription, error) {

	var vaultRule []interface{}
	for _, vaultItem := range vault {
		vaultRule = append(vaultRule, vaultItem)
	}

	logs, sub, err := _Dittoentrypoint.contract.WatchLogs(opts, "DittoEntryPointRegisterWorkflow", vaultRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DittoentrypointDittoEntryPointRegisterWorkflow)
				if err := _Dittoentrypoint.contract.UnpackLog(event, "DittoEntryPointRegisterWorkflow", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseDittoEntryPointRegisterWorkflow is a log parse operation binding the contract event 0x7d89995e8cf5bdb0963851260411a4a799c0a5a754120fab6b126845ce4bdf60.
//
// Solidity: event DittoEntryPointRegisterWorkflow(address indexed vault, uint256 workflowId)
func (_Dittoentrypoint *DittoentrypointFilterer) ParseDittoEntryPointRegisterWorkflow(log types.Log) (*DittoentrypointDittoEntryPointRegisterWorkflow, error) {
	event := new(DittoentrypointDittoEntryPointRegisterWorkflow)
	if err := _Dittoentrypoint.contract.UnpackLog(event, "DittoEntryPointRegisterWorkflow", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// DittoentrypointDittoEntryPointWorkflowSuccessIterator is returned from FilterDittoEntryPointWorkflowSuccess and is used to iterate over the raw logs and unpacked data for DittoEntryPointWorkflowSuccess events raised by the Dittoentrypoint contract.
type DittoentrypointDittoEntryPointWorkflowSuccessIterator struct {
	Event *DittoentrypointDittoEntryPointWorkflowSuccess // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *DittoentrypointDittoEntryPointWorkflowSuccessIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DittoentrypointDittoEntryPointWorkflowSuccess)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(DittoentrypointDittoEntryPointWorkflowSuccess)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *DittoentrypointDittoEntryPointWorkflowSuccessIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DittoentrypointDittoEntryPointWorkflowSuccessIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DittoentrypointDittoEntryPointWorkflowSuccess represents a DittoEntryPointWorkflowSuccess event raised by the Dittoentrypoint contract.
type DittoentrypointDittoEntryPointWorkflowSuccess struct {
	Vault      common.Address
	WorkflowId *big.Int
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterDittoEntryPointWorkflowSuccess is a free log retrieval operation binding the contract event 0x7c32983556f0c282d44b1d191034a5141ca7e7d4739e98b160b854baa9b19005.
//
// Solidity: event DittoEntryPointWorkflowSuccess(address indexed vault, uint256 workflowId)
func (_Dittoentrypoint *DittoentrypointFilterer) FilterDittoEntryPointWorkflowSuccess(opts *bind.FilterOpts, vault []common.Address) (*DittoentrypointDittoEntryPointWorkflowSuccessIterator, error) {

	var vaultRule []interface{}
	for _, vaultItem := range vault {
		vaultRule = append(vaultRule, vaultItem)
	}

	logs, sub, err := _Dittoentrypoint.contract.FilterLogs(opts, "DittoEntryPointWorkflowSuccess", vaultRule)
	if err != nil {
		return nil, err
	}
	return &DittoentrypointDittoEntryPointWorkflowSuccessIterator{contract: _Dittoentrypoint.contract, event: "DittoEntryPointWorkflowSuccess", logs: logs, sub: sub}, nil
}

// WatchDittoEntryPointWorkflowSuccess is a free log subscription operation binding the contract event 0x7c32983556f0c282d44b1d191034a5141ca7e7d4739e98b160b854baa9b19005.
//
// Solidity: event DittoEntryPointWorkflowSuccess(address indexed vault, uint256 workflowId)
func (_Dittoentrypoint *DittoentrypointFilterer) WatchDittoEntryPointWorkflowSuccess(opts *bind.WatchOpts, sink chan<- *DittoentrypointDittoEntryPointWorkflowSuccess, vault []common.Address) (event.Subscription, error) {

	var vaultRule []interface{}
	for _, vaultItem := range vault {
		vaultRule = append(vaultRule, vaultItem)
	}

	logs, sub, err := _Dittoentrypoint.contract.WatchLogs(opts, "DittoEntryPointWorkflowSuccess", vaultRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DittoentrypointDittoEntryPointWorkflowSuccess)
				if err := _Dittoentrypoint.contract.UnpackLog(event, "DittoEntryPointWorkflowSuccess", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseDittoEntryPointWorkflowSuccess is a log parse operation binding the contract event 0x7c32983556f0c282d44b1d191034a5141ca7e7d4739e98b160b854baa9b19005.
//
// Solidity: event DittoEntryPointWorkflowSuccess(address indexed vault, uint256 workflowId)
func (_Dittoentrypoint *DittoentrypointFilterer) ParseDittoEntryPointWorkflowSuccess(log types.Log) (*DittoentrypointDittoEntryPointWorkflowSuccess, error) {
	event := new(DittoentrypointDittoEntryPointWorkflowSuccess)
	if err := _Dittoentrypoint.contract.UnpackLog(event, "DittoEntryPointWorkflowSuccess", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
