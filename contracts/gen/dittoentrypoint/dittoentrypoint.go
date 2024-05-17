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
	ABI: "[{\"type\":\"function\",\"name\":\"activeWorkflows\",\"inputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"vaultAddress\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"workflowId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"arrangeExecutors\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"canExecWorkflowCheck\",\"inputs\":[{\"name\":\"vaultAddress\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"workflowId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"cancelWorkflow\",\"inputs\":[{\"name\":\"workflowId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"currentEpoch\",\"inputs\":[],\"outputs\":[{\"name\":\"startBlock\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"totalStake\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"blockSlotSize\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"epochSize\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"blockHash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"delegateManager\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"epochExecutors\",\"inputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"cumulativeStakeBefore\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"cumulativeStakeAfter\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"epochIdx\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"executors\",\"inputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getActiveWorkflows\",\"inputs\":[{\"name\":\"from\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"to\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"pageWorkflows\",\"type\":\"tuple[]\",\"internalType\":\"structIDittoEntryPoint.Workflow[]\",\"components\":[{\"name\":\"vaultAddress\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"workflowId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getAllActiveWorkflows\",\"inputs\":[],\"outputs\":[{\"name\":\"activeWorkflows_\",\"type\":\"tuple[]\",\"internalType\":\"structIDittoEntryPoint.Workflow[]\",\"components\":[{\"name\":\"vaultAddress\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"workflowId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getAmountActiveWorkflows\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getAmountExecutors\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getSeedExternal\",\"inputs\":[{\"name\":\"blockNumber\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"seed\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"initialize\",\"inputs\":[{\"name\":\"_newOwner\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"isExecutor\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"isValidExecutor\",\"inputs\":[{\"name\":\"blockNumber\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"executorAddress\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"nextEpoch\",\"inputs\":[],\"outputs\":[{\"name\":\"startBlock\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"totalStake\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"blockSlotSize\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"epochSize\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"blockHash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"operators\",\"inputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"currentStake\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"executorIndex\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"isInitialized\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"executorUnregister\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"owner\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"registerExecutor\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"registerWorkflow\",\"inputs\":[{\"name\":\"workflowId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"renounceOwnership\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"runMultipleWorkflows\",\"inputs\":[{\"name\":\"workflows\",\"type\":\"tuple[]\",\"internalType\":\"structIDittoEntryPoint.Workflow[]\",\"components\":[{\"name\":\"vaultAddress\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"workflowId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"runWorkflow\",\"inputs\":[{\"name\":\"vaultAddress\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"workflowId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setParams\",\"inputs\":[{\"name\":\"_blockSlotSizeNext\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"_epochSizeNext\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"startParamsInit\",\"inputs\":[{\"name\":\"_blockSlotSize\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"_epochSize\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"_strategy\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_delegateManager\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"strategy\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"transferOwnership\",\"inputs\":[{\"name\":\"newOwner\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"unregisterExecutor\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"event\",\"name\":\"OwnershipTransferred\",\"inputs\":[{\"name\":\"previousOwner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"newOwner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"error\",\"name\":\"DittoEntryPoint_AlreadyInitialized\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"DittoEntryPoint_AlreadyStartParamsInitialized\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"Ownable_NewOwnerCannotBeAddressZero\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"Ownable_SenderIsNotOwner\",\"inputs\":[{\"name\":\"sender\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"TransferHelper_SafeTransferNativeError\",\"inputs\":[]}]",
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

// ActiveWorkflows is a free data retrieval call binding the contract method 0x54ba805b.
//
// Solidity: function activeWorkflows(uint256 ) view returns(address vaultAddress, uint256 workflowId)
func (_Dittoentrypoint *DittoentrypointCaller) ActiveWorkflows(opts *bind.CallOpts, arg0 *big.Int) (struct {
	VaultAddress common.Address
	WorkflowId   *big.Int
}, error) {
	var out []interface{}
	err := _Dittoentrypoint.contract.Call(opts, &out, "activeWorkflows", arg0)

	outstruct := new(struct {
		VaultAddress common.Address
		WorkflowId   *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.VaultAddress = *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	outstruct.WorkflowId = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// ActiveWorkflows is a free data retrieval call binding the contract method 0x54ba805b.
//
// Solidity: function activeWorkflows(uint256 ) view returns(address vaultAddress, uint256 workflowId)
func (_Dittoentrypoint *DittoentrypointSession) ActiveWorkflows(arg0 *big.Int) (struct {
	VaultAddress common.Address
	WorkflowId   *big.Int
}, error) {
	return _Dittoentrypoint.Contract.ActiveWorkflows(&_Dittoentrypoint.CallOpts, arg0)
}

// ActiveWorkflows is a free data retrieval call binding the contract method 0x54ba805b.
//
// Solidity: function activeWorkflows(uint256 ) view returns(address vaultAddress, uint256 workflowId)
func (_Dittoentrypoint *DittoentrypointCallerSession) ActiveWorkflows(arg0 *big.Int) (struct {
	VaultAddress common.Address
	WorkflowId   *big.Int
}, error) {
	return _Dittoentrypoint.Contract.ActiveWorkflows(&_Dittoentrypoint.CallOpts, arg0)
}

// CanExecWorkflowCheck is a free data retrieval call binding the contract method 0x416c3941.
//
// Solidity: function canExecWorkflowCheck(address vaultAddress, uint256 workflowId) view returns(bool)
func (_Dittoentrypoint *DittoentrypointCaller) CanExecWorkflowCheck(opts *bind.CallOpts, vaultAddress common.Address, workflowId *big.Int) (bool, error) {
	var out []interface{}
	err := _Dittoentrypoint.contract.Call(opts, &out, "canExecWorkflowCheck", vaultAddress, workflowId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// CanExecWorkflowCheck is a free data retrieval call binding the contract method 0x416c3941.
//
// Solidity: function canExecWorkflowCheck(address vaultAddress, uint256 workflowId) view returns(bool)
func (_Dittoentrypoint *DittoentrypointSession) CanExecWorkflowCheck(vaultAddress common.Address, workflowId *big.Int) (bool, error) {
	return _Dittoentrypoint.Contract.CanExecWorkflowCheck(&_Dittoentrypoint.CallOpts, vaultAddress, workflowId)
}

// CanExecWorkflowCheck is a free data retrieval call binding the contract method 0x416c3941.
//
// Solidity: function canExecWorkflowCheck(address vaultAddress, uint256 workflowId) view returns(bool)
func (_Dittoentrypoint *DittoentrypointCallerSession) CanExecWorkflowCheck(vaultAddress common.Address, workflowId *big.Int) (bool, error) {
	return _Dittoentrypoint.Contract.CanExecWorkflowCheck(&_Dittoentrypoint.CallOpts, vaultAddress, workflowId)
}

// CurrentEpoch is a free data retrieval call binding the contract method 0x76671808.
//
// Solidity: function currentEpoch() view returns(uint256 startBlock, uint256 totalStake, uint256 blockSlotSize, uint256 epochSize, bytes32 blockHash)
func (_Dittoentrypoint *DittoentrypointCaller) CurrentEpoch(opts *bind.CallOpts) (struct {
	StartBlock    *big.Int
	TotalStake    *big.Int
	BlockSlotSize *big.Int
	EpochSize     *big.Int
	BlockHash     [32]byte
}, error) {
	var out []interface{}
	err := _Dittoentrypoint.contract.Call(opts, &out, "currentEpoch")

	outstruct := new(struct {
		StartBlock    *big.Int
		TotalStake    *big.Int
		BlockSlotSize *big.Int
		EpochSize     *big.Int
		BlockHash     [32]byte
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.StartBlock = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.TotalStake = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.BlockSlotSize = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	outstruct.EpochSize = *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)
	outstruct.BlockHash = *abi.ConvertType(out[4], new([32]byte)).(*[32]byte)

	return *outstruct, err

}

// CurrentEpoch is a free data retrieval call binding the contract method 0x76671808.
//
// Solidity: function currentEpoch() view returns(uint256 startBlock, uint256 totalStake, uint256 blockSlotSize, uint256 epochSize, bytes32 blockHash)
func (_Dittoentrypoint *DittoentrypointSession) CurrentEpoch() (struct {
	StartBlock    *big.Int
	TotalStake    *big.Int
	BlockSlotSize *big.Int
	EpochSize     *big.Int
	BlockHash     [32]byte
}, error) {
	return _Dittoentrypoint.Contract.CurrentEpoch(&_Dittoentrypoint.CallOpts)
}

// CurrentEpoch is a free data retrieval call binding the contract method 0x76671808.
//
// Solidity: function currentEpoch() view returns(uint256 startBlock, uint256 totalStake, uint256 blockSlotSize, uint256 epochSize, bytes32 blockHash)
func (_Dittoentrypoint *DittoentrypointCallerSession) CurrentEpoch() (struct {
	StartBlock    *big.Int
	TotalStake    *big.Int
	BlockSlotSize *big.Int
	EpochSize     *big.Int
	BlockHash     [32]byte
}, error) {
	return _Dittoentrypoint.Contract.CurrentEpoch(&_Dittoentrypoint.CallOpts)
}

// DelegateManager is a free data retrieval call binding the contract method 0x388e9104.
//
// Solidity: function delegateManager() view returns(address)
func (_Dittoentrypoint *DittoentrypointCaller) DelegateManager(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Dittoentrypoint.contract.Call(opts, &out, "delegateManager")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// DelegateManager is a free data retrieval call binding the contract method 0x388e9104.
//
// Solidity: function delegateManager() view returns(address)
func (_Dittoentrypoint *DittoentrypointSession) DelegateManager() (common.Address, error) {
	return _Dittoentrypoint.Contract.DelegateManager(&_Dittoentrypoint.CallOpts)
}

// DelegateManager is a free data retrieval call binding the contract method 0x388e9104.
//
// Solidity: function delegateManager() view returns(address)
func (_Dittoentrypoint *DittoentrypointCallerSession) DelegateManager() (common.Address, error) {
	return _Dittoentrypoint.Contract.DelegateManager(&_Dittoentrypoint.CallOpts)
}

// EpochExecutors is a free data retrieval call binding the contract method 0x1b5b3ef2.
//
// Solidity: function epochExecutors(uint256 , address ) view returns(uint256 cumulativeStakeBefore, uint256 cumulativeStakeAfter)
func (_Dittoentrypoint *DittoentrypointCaller) EpochExecutors(opts *bind.CallOpts, arg0 *big.Int, arg1 common.Address) (struct {
	CumulativeStakeBefore *big.Int
	CumulativeStakeAfter  *big.Int
}, error) {
	var out []interface{}
	err := _Dittoentrypoint.contract.Call(opts, &out, "epochExecutors", arg0, arg1)

	outstruct := new(struct {
		CumulativeStakeBefore *big.Int
		CumulativeStakeAfter  *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.CumulativeStakeBefore = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.CumulativeStakeAfter = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// EpochExecutors is a free data retrieval call binding the contract method 0x1b5b3ef2.
//
// Solidity: function epochExecutors(uint256 , address ) view returns(uint256 cumulativeStakeBefore, uint256 cumulativeStakeAfter)
func (_Dittoentrypoint *DittoentrypointSession) EpochExecutors(arg0 *big.Int, arg1 common.Address) (struct {
	CumulativeStakeBefore *big.Int
	CumulativeStakeAfter  *big.Int
}, error) {
	return _Dittoentrypoint.Contract.EpochExecutors(&_Dittoentrypoint.CallOpts, arg0, arg1)
}

// EpochExecutors is a free data retrieval call binding the contract method 0x1b5b3ef2.
//
// Solidity: function epochExecutors(uint256 , address ) view returns(uint256 cumulativeStakeBefore, uint256 cumulativeStakeAfter)
func (_Dittoentrypoint *DittoentrypointCallerSession) EpochExecutors(arg0 *big.Int, arg1 common.Address) (struct {
	CumulativeStakeBefore *big.Int
	CumulativeStakeAfter  *big.Int
}, error) {
	return _Dittoentrypoint.Contract.EpochExecutors(&_Dittoentrypoint.CallOpts, arg0, arg1)
}

// EpochIdx is a free data retrieval call binding the contract method 0x664c3741.
//
// Solidity: function epochIdx() view returns(uint256)
func (_Dittoentrypoint *DittoentrypointCaller) EpochIdx(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Dittoentrypoint.contract.Call(opts, &out, "epochIdx")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// EpochIdx is a free data retrieval call binding the contract method 0x664c3741.
//
// Solidity: function epochIdx() view returns(uint256)
func (_Dittoentrypoint *DittoentrypointSession) EpochIdx() (*big.Int, error) {
	return _Dittoentrypoint.Contract.EpochIdx(&_Dittoentrypoint.CallOpts)
}

// EpochIdx is a free data retrieval call binding the contract method 0x664c3741.
//
// Solidity: function epochIdx() view returns(uint256)
func (_Dittoentrypoint *DittoentrypointCallerSession) EpochIdx() (*big.Int, error) {
	return _Dittoentrypoint.Contract.EpochIdx(&_Dittoentrypoint.CallOpts)
}

// Executors is a free data retrieval call binding the contract method 0xf97a05df.
//
// Solidity: function executors(uint256 ) view returns(address)
func (_Dittoentrypoint *DittoentrypointCaller) Executors(opts *bind.CallOpts, arg0 *big.Int) (common.Address, error) {
	var out []interface{}
	err := _Dittoentrypoint.contract.Call(opts, &out, "executors", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Executors is a free data retrieval call binding the contract method 0xf97a05df.
//
// Solidity: function executors(uint256 ) view returns(address)
func (_Dittoentrypoint *DittoentrypointSession) Executors(arg0 *big.Int) (common.Address, error) {
	return _Dittoentrypoint.Contract.Executors(&_Dittoentrypoint.CallOpts, arg0)
}

// Executors is a free data retrieval call binding the contract method 0xf97a05df.
//
// Solidity: function executors(uint256 ) view returns(address)
func (_Dittoentrypoint *DittoentrypointCallerSession) Executors(arg0 *big.Int) (common.Address, error) {
	return _Dittoentrypoint.Contract.Executors(&_Dittoentrypoint.CallOpts, arg0)
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
// Solidity: function getAllActiveWorkflows() view returns((address,uint256)[] activeWorkflows_)
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
// Solidity: function getAllActiveWorkflows() view returns((address,uint256)[] activeWorkflows_)
func (_Dittoentrypoint *DittoentrypointSession) GetAllActiveWorkflows() ([]IDittoEntryPointWorkflow, error) {
	return _Dittoentrypoint.Contract.GetAllActiveWorkflows(&_Dittoentrypoint.CallOpts)
}

// GetAllActiveWorkflows is a free data retrieval call binding the contract method 0xcd92d3f1.
//
// Solidity: function getAllActiveWorkflows() view returns((address,uint256)[] activeWorkflows_)
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

// NextEpoch is a free data retrieval call binding the contract method 0xaea0e78b.
//
// Solidity: function nextEpoch() view returns(uint256 startBlock, uint256 totalStake, uint256 blockSlotSize, uint256 epochSize, bytes32 blockHash)
func (_Dittoentrypoint *DittoentrypointCaller) NextEpoch(opts *bind.CallOpts) (struct {
	StartBlock    *big.Int
	TotalStake    *big.Int
	BlockSlotSize *big.Int
	EpochSize     *big.Int
	BlockHash     [32]byte
}, error) {
	var out []interface{}
	err := _Dittoentrypoint.contract.Call(opts, &out, "nextEpoch")

	outstruct := new(struct {
		StartBlock    *big.Int
		TotalStake    *big.Int
		BlockSlotSize *big.Int
		EpochSize     *big.Int
		BlockHash     [32]byte
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.StartBlock = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.TotalStake = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.BlockSlotSize = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	outstruct.EpochSize = *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)
	outstruct.BlockHash = *abi.ConvertType(out[4], new([32]byte)).(*[32]byte)

	return *outstruct, err

}

// NextEpoch is a free data retrieval call binding the contract method 0xaea0e78b.
//
// Solidity: function nextEpoch() view returns(uint256 startBlock, uint256 totalStake, uint256 blockSlotSize, uint256 epochSize, bytes32 blockHash)
func (_Dittoentrypoint *DittoentrypointSession) NextEpoch() (struct {
	StartBlock    *big.Int
	TotalStake    *big.Int
	BlockSlotSize *big.Int
	EpochSize     *big.Int
	BlockHash     [32]byte
}, error) {
	return _Dittoentrypoint.Contract.NextEpoch(&_Dittoentrypoint.CallOpts)
}

// NextEpoch is a free data retrieval call binding the contract method 0xaea0e78b.
//
// Solidity: function nextEpoch() view returns(uint256 startBlock, uint256 totalStake, uint256 blockSlotSize, uint256 epochSize, bytes32 blockHash)
func (_Dittoentrypoint *DittoentrypointCallerSession) NextEpoch() (struct {
	StartBlock    *big.Int
	TotalStake    *big.Int
	BlockSlotSize *big.Int
	EpochSize     *big.Int
	BlockHash     [32]byte
}, error) {
	return _Dittoentrypoint.Contract.NextEpoch(&_Dittoentrypoint.CallOpts)
}

// Operators is a free data retrieval call binding the contract method 0x13e7c9d8.
//
// Solidity: function operators(address ) view returns(uint256 currentStake, uint256 executorIndex, bool isInitialized, bool executorUnregister)
func (_Dittoentrypoint *DittoentrypointCaller) Operators(opts *bind.CallOpts, arg0 common.Address) (struct {
	CurrentStake       *big.Int
	ExecutorIndex      *big.Int
	IsInitialized      bool
	ExecutorUnregister bool
}, error) {
	var out []interface{}
	err := _Dittoentrypoint.contract.Call(opts, &out, "operators", arg0)

	outstruct := new(struct {
		CurrentStake       *big.Int
		ExecutorIndex      *big.Int
		IsInitialized      bool
		ExecutorUnregister bool
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.CurrentStake = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.ExecutorIndex = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.IsInitialized = *abi.ConvertType(out[2], new(bool)).(*bool)
	outstruct.ExecutorUnregister = *abi.ConvertType(out[3], new(bool)).(*bool)

	return *outstruct, err

}

// Operators is a free data retrieval call binding the contract method 0x13e7c9d8.
//
// Solidity: function operators(address ) view returns(uint256 currentStake, uint256 executorIndex, bool isInitialized, bool executorUnregister)
func (_Dittoentrypoint *DittoentrypointSession) Operators(arg0 common.Address) (struct {
	CurrentStake       *big.Int
	ExecutorIndex      *big.Int
	IsInitialized      bool
	ExecutorUnregister bool
}, error) {
	return _Dittoentrypoint.Contract.Operators(&_Dittoentrypoint.CallOpts, arg0)
}

// Operators is a free data retrieval call binding the contract method 0x13e7c9d8.
//
// Solidity: function operators(address ) view returns(uint256 currentStake, uint256 executorIndex, bool isInitialized, bool executorUnregister)
func (_Dittoentrypoint *DittoentrypointCallerSession) Operators(arg0 common.Address) (struct {
	CurrentStake       *big.Int
	ExecutorIndex      *big.Int
	IsInitialized      bool
	ExecutorUnregister bool
}, error) {
	return _Dittoentrypoint.Contract.Operators(&_Dittoentrypoint.CallOpts, arg0)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Dittoentrypoint *DittoentrypointCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Dittoentrypoint.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Dittoentrypoint *DittoentrypointSession) Owner() (common.Address, error) {
	return _Dittoentrypoint.Contract.Owner(&_Dittoentrypoint.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Dittoentrypoint *DittoentrypointCallerSession) Owner() (common.Address, error) {
	return _Dittoentrypoint.Contract.Owner(&_Dittoentrypoint.CallOpts)
}

// Strategy is a free data retrieval call binding the contract method 0xa8c62e76.
//
// Solidity: function strategy() view returns(address)
func (_Dittoentrypoint *DittoentrypointCaller) Strategy(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Dittoentrypoint.contract.Call(opts, &out, "strategy")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Strategy is a free data retrieval call binding the contract method 0xa8c62e76.
//
// Solidity: function strategy() view returns(address)
func (_Dittoentrypoint *DittoentrypointSession) Strategy() (common.Address, error) {
	return _Dittoentrypoint.Contract.Strategy(&_Dittoentrypoint.CallOpts)
}

// Strategy is a free data retrieval call binding the contract method 0xa8c62e76.
//
// Solidity: function strategy() view returns(address)
func (_Dittoentrypoint *DittoentrypointCallerSession) Strategy() (common.Address, error) {
	return _Dittoentrypoint.Contract.Strategy(&_Dittoentrypoint.CallOpts)
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

// Initialize is a paid mutator transaction binding the contract method 0xc4d66de8.
//
// Solidity: function initialize(address _newOwner) returns()
func (_Dittoentrypoint *DittoentrypointTransactor) Initialize(opts *bind.TransactOpts, _newOwner common.Address) (*types.Transaction, error) {
	return _Dittoentrypoint.contract.Transact(opts, "initialize", _newOwner)
}

// Initialize is a paid mutator transaction binding the contract method 0xc4d66de8.
//
// Solidity: function initialize(address _newOwner) returns()
func (_Dittoentrypoint *DittoentrypointSession) Initialize(_newOwner common.Address) (*types.Transaction, error) {
	return _Dittoentrypoint.Contract.Initialize(&_Dittoentrypoint.TransactOpts, _newOwner)
}

// Initialize is a paid mutator transaction binding the contract method 0xc4d66de8.
//
// Solidity: function initialize(address _newOwner) returns()
func (_Dittoentrypoint *DittoentrypointTransactorSession) Initialize(_newOwner common.Address) (*types.Transaction, error) {
	return _Dittoentrypoint.Contract.Initialize(&_Dittoentrypoint.TransactOpts, _newOwner)
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

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Dittoentrypoint *DittoentrypointTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Dittoentrypoint.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Dittoentrypoint *DittoentrypointSession) RenounceOwnership() (*types.Transaction, error) {
	return _Dittoentrypoint.Contract.RenounceOwnership(&_Dittoentrypoint.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Dittoentrypoint *DittoentrypointTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _Dittoentrypoint.Contract.RenounceOwnership(&_Dittoentrypoint.TransactOpts)
}

// RunMultipleWorkflows is a paid mutator transaction binding the contract method 0x1370159c.
//
// Solidity: function runMultipleWorkflows((address,uint256)[] workflows) returns()
func (_Dittoentrypoint *DittoentrypointTransactor) RunMultipleWorkflows(opts *bind.TransactOpts, workflows []IDittoEntryPointWorkflow) (*types.Transaction, error) {
	return _Dittoentrypoint.contract.Transact(opts, "runMultipleWorkflows", workflows)
}

// RunMultipleWorkflows is a paid mutator transaction binding the contract method 0x1370159c.
//
// Solidity: function runMultipleWorkflows((address,uint256)[] workflows) returns()
func (_Dittoentrypoint *DittoentrypointSession) RunMultipleWorkflows(workflows []IDittoEntryPointWorkflow) (*types.Transaction, error) {
	return _Dittoentrypoint.Contract.RunMultipleWorkflows(&_Dittoentrypoint.TransactOpts, workflows)
}

// RunMultipleWorkflows is a paid mutator transaction binding the contract method 0x1370159c.
//
// Solidity: function runMultipleWorkflows((address,uint256)[] workflows) returns()
func (_Dittoentrypoint *DittoentrypointTransactorSession) RunMultipleWorkflows(workflows []IDittoEntryPointWorkflow) (*types.Transaction, error) {
	return _Dittoentrypoint.Contract.RunMultipleWorkflows(&_Dittoentrypoint.TransactOpts, workflows)
}

// RunWorkflow is a paid mutator transaction binding the contract method 0x109bf98d.
//
// Solidity: function runWorkflow(address vaultAddress, uint256 workflowId) returns()
func (_Dittoentrypoint *DittoentrypointTransactor) RunWorkflow(opts *bind.TransactOpts, vaultAddress common.Address, workflowId *big.Int) (*types.Transaction, error) {
	return _Dittoentrypoint.contract.Transact(opts, "runWorkflow", vaultAddress, workflowId)
}

// RunWorkflow is a paid mutator transaction binding the contract method 0x109bf98d.
//
// Solidity: function runWorkflow(address vaultAddress, uint256 workflowId) returns()
func (_Dittoentrypoint *DittoentrypointSession) RunWorkflow(vaultAddress common.Address, workflowId *big.Int) (*types.Transaction, error) {
	return _Dittoentrypoint.Contract.RunWorkflow(&_Dittoentrypoint.TransactOpts, vaultAddress, workflowId)
}

// RunWorkflow is a paid mutator transaction binding the contract method 0x109bf98d.
//
// Solidity: function runWorkflow(address vaultAddress, uint256 workflowId) returns()
func (_Dittoentrypoint *DittoentrypointTransactorSession) RunWorkflow(vaultAddress common.Address, workflowId *big.Int) (*types.Transaction, error) {
	return _Dittoentrypoint.Contract.RunWorkflow(&_Dittoentrypoint.TransactOpts, vaultAddress, workflowId)
}

// SetParams is a paid mutator transaction binding the contract method 0xc0324c77.
//
// Solidity: function setParams(uint256 _blockSlotSizeNext, uint256 _epochSizeNext) returns()
func (_Dittoentrypoint *DittoentrypointTransactor) SetParams(opts *bind.TransactOpts, _blockSlotSizeNext *big.Int, _epochSizeNext *big.Int) (*types.Transaction, error) {
	return _Dittoentrypoint.contract.Transact(opts, "setParams", _blockSlotSizeNext, _epochSizeNext)
}

// SetParams is a paid mutator transaction binding the contract method 0xc0324c77.
//
// Solidity: function setParams(uint256 _blockSlotSizeNext, uint256 _epochSizeNext) returns()
func (_Dittoentrypoint *DittoentrypointSession) SetParams(_blockSlotSizeNext *big.Int, _epochSizeNext *big.Int) (*types.Transaction, error) {
	return _Dittoentrypoint.Contract.SetParams(&_Dittoentrypoint.TransactOpts, _blockSlotSizeNext, _epochSizeNext)
}

// SetParams is a paid mutator transaction binding the contract method 0xc0324c77.
//
// Solidity: function setParams(uint256 _blockSlotSizeNext, uint256 _epochSizeNext) returns()
func (_Dittoentrypoint *DittoentrypointTransactorSession) SetParams(_blockSlotSizeNext *big.Int, _epochSizeNext *big.Int) (*types.Transaction, error) {
	return _Dittoentrypoint.Contract.SetParams(&_Dittoentrypoint.TransactOpts, _blockSlotSizeNext, _epochSizeNext)
}

// StartParamsInit is a paid mutator transaction binding the contract method 0x05ce8c4d.
//
// Solidity: function startParamsInit(uint256 _blockSlotSize, uint256 _epochSize, address _strategy, address _delegateManager) returns()
func (_Dittoentrypoint *DittoentrypointTransactor) StartParamsInit(opts *bind.TransactOpts, _blockSlotSize *big.Int, _epochSize *big.Int, _strategy common.Address, _delegateManager common.Address) (*types.Transaction, error) {
	return _Dittoentrypoint.contract.Transact(opts, "startParamsInit", _blockSlotSize, _epochSize, _strategy, _delegateManager)
}

// StartParamsInit is a paid mutator transaction binding the contract method 0x05ce8c4d.
//
// Solidity: function startParamsInit(uint256 _blockSlotSize, uint256 _epochSize, address _strategy, address _delegateManager) returns()
func (_Dittoentrypoint *DittoentrypointSession) StartParamsInit(_blockSlotSize *big.Int, _epochSize *big.Int, _strategy common.Address, _delegateManager common.Address) (*types.Transaction, error) {
	return _Dittoentrypoint.Contract.StartParamsInit(&_Dittoentrypoint.TransactOpts, _blockSlotSize, _epochSize, _strategy, _delegateManager)
}

// StartParamsInit is a paid mutator transaction binding the contract method 0x05ce8c4d.
//
// Solidity: function startParamsInit(uint256 _blockSlotSize, uint256 _epochSize, address _strategy, address _delegateManager) returns()
func (_Dittoentrypoint *DittoentrypointTransactorSession) StartParamsInit(_blockSlotSize *big.Int, _epochSize *big.Int, _strategy common.Address, _delegateManager common.Address) (*types.Transaction, error) {
	return _Dittoentrypoint.Contract.StartParamsInit(&_Dittoentrypoint.TransactOpts, _blockSlotSize, _epochSize, _strategy, _delegateManager)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Dittoentrypoint *DittoentrypointTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _Dittoentrypoint.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Dittoentrypoint *DittoentrypointSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Dittoentrypoint.Contract.TransferOwnership(&_Dittoentrypoint.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Dittoentrypoint *DittoentrypointTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Dittoentrypoint.Contract.TransferOwnership(&_Dittoentrypoint.TransactOpts, newOwner)
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

// DittoentrypointOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the Dittoentrypoint contract.
type DittoentrypointOwnershipTransferredIterator struct {
	Event *DittoentrypointOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *DittoentrypointOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DittoentrypointOwnershipTransferred)
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
		it.Event = new(DittoentrypointOwnershipTransferred)
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
func (it *DittoentrypointOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DittoentrypointOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DittoentrypointOwnershipTransferred represents a OwnershipTransferred event raised by the Dittoentrypoint contract.
type DittoentrypointOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Dittoentrypoint *DittoentrypointFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*DittoentrypointOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Dittoentrypoint.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &DittoentrypointOwnershipTransferredIterator{contract: _Dittoentrypoint.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Dittoentrypoint *DittoentrypointFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *DittoentrypointOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Dittoentrypoint.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DittoentrypointOwnershipTransferred)
				if err := _Dittoentrypoint.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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

// ParseOwnershipTransferred is a log parse operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Dittoentrypoint *DittoentrypointFilterer) ParseOwnershipTransferred(log types.Log) (*DittoentrypointOwnershipTransferred, error) {
	event := new(DittoentrypointOwnershipTransferred)
	if err := _Dittoentrypoint.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
