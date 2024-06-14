// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package iservicemanager

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

// ISignatureUtilsSignatureWithSaltAndExpiry is an auto generated low-level Go binding around an user-defined struct.
type ISignatureUtilsSignatureWithSaltAndExpiry struct {
	Signature []byte
	Salt      [32]byte
	Expiry    *big.Int
}

// IservicemanagerMetaData contains all meta data concerning the Iservicemanager contract.
var IservicemanagerMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"function\",\"name\":\"deregisterOperatorFromAVS\",\"inputs\":[{\"name\":\"operator\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"getOperatorRestakedStrategies\",\"inputs\":[{\"name\":\"operator\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"address[]\",\"internalType\":\"address[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getRestakeableStrategies\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address[]\",\"internalType\":\"address[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"registerOperatorToAVS\",\"inputs\":[{\"name\":\"operator\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"operatorSignature\",\"type\":\"tuple\",\"internalType\":\"structISignatureUtils.SignatureWithSaltAndExpiry\",\"components\":[{\"name\":\"signature\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"salt\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"expiry\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"updateAVSMetadataURI\",\"inputs\":[{\"name\":\"_metadataURI\",\"type\":\"string\",\"internalType\":\"string\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"}]",
}

// IservicemanagerABI is the input ABI used to generate the binding from.
// Deprecated: Use IservicemanagerMetaData.ABI instead.
var IservicemanagerABI = IservicemanagerMetaData.ABI

// Iservicemanager is an auto generated Go binding around an Ethereum contract.
type Iservicemanager struct {
	IservicemanagerCaller     // Read-only binding to the contract
	IservicemanagerTransactor // Write-only binding to the contract
	IservicemanagerFilterer   // Log filterer for contract events
}

// IservicemanagerCaller is an auto generated read-only Go binding around an Ethereum contract.
type IservicemanagerCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IservicemanagerTransactor is an auto generated write-only Go binding around an Ethereum contract.
type IservicemanagerTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IservicemanagerFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IservicemanagerFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IservicemanagerSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IservicemanagerSession struct {
	Contract     *Iservicemanager  // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// IservicemanagerCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IservicemanagerCallerSession struct {
	Contract *IservicemanagerCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts          // Call options to use throughout this session
}

// IservicemanagerTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IservicemanagerTransactorSession struct {
	Contract     *IservicemanagerTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts          // Transaction auth options to use throughout this session
}

// IservicemanagerRaw is an auto generated low-level Go binding around an Ethereum contract.
type IservicemanagerRaw struct {
	Contract *Iservicemanager // Generic contract binding to access the raw methods on
}

// IservicemanagerCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IservicemanagerCallerRaw struct {
	Contract *IservicemanagerCaller // Generic read-only contract binding to access the raw methods on
}

// IservicemanagerTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IservicemanagerTransactorRaw struct {
	Contract *IservicemanagerTransactor // Generic write-only contract binding to access the raw methods on
}

// NewIservicemanager creates a new instance of Iservicemanager, bound to a specific deployed contract.
func NewIservicemanager(address common.Address, backend bind.ContractBackend) (*Iservicemanager, error) {
	contract, err := bindIservicemanager(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Iservicemanager{IservicemanagerCaller: IservicemanagerCaller{contract: contract}, IservicemanagerTransactor: IservicemanagerTransactor{contract: contract}, IservicemanagerFilterer: IservicemanagerFilterer{contract: contract}}, nil
}

// NewIservicemanagerCaller creates a new read-only instance of Iservicemanager, bound to a specific deployed contract.
func NewIservicemanagerCaller(address common.Address, caller bind.ContractCaller) (*IservicemanagerCaller, error) {
	contract, err := bindIservicemanager(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IservicemanagerCaller{contract: contract}, nil
}

// NewIservicemanagerTransactor creates a new write-only instance of Iservicemanager, bound to a specific deployed contract.
func NewIservicemanagerTransactor(address common.Address, transactor bind.ContractTransactor) (*IservicemanagerTransactor, error) {
	contract, err := bindIservicemanager(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IservicemanagerTransactor{contract: contract}, nil
}

// NewIservicemanagerFilterer creates a new log filterer instance of Iservicemanager, bound to a specific deployed contract.
func NewIservicemanagerFilterer(address common.Address, filterer bind.ContractFilterer) (*IservicemanagerFilterer, error) {
	contract, err := bindIservicemanager(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IservicemanagerFilterer{contract: contract}, nil
}

// bindIservicemanager binds a generic wrapper to an already deployed contract.
func bindIservicemanager(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := IservicemanagerMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Iservicemanager *IservicemanagerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Iservicemanager.Contract.IservicemanagerCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Iservicemanager *IservicemanagerRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Iservicemanager.Contract.IservicemanagerTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Iservicemanager *IservicemanagerRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Iservicemanager.Contract.IservicemanagerTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Iservicemanager *IservicemanagerCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Iservicemanager.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Iservicemanager *IservicemanagerTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Iservicemanager.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Iservicemanager *IservicemanagerTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Iservicemanager.Contract.contract.Transact(opts, method, params...)
}

// GetOperatorRestakedStrategies is a free data retrieval call binding the contract method 0x33cfb7b7.
//
// Solidity: function getOperatorRestakedStrategies(address operator) view returns(address[])
func (_Iservicemanager *IservicemanagerCaller) GetOperatorRestakedStrategies(opts *bind.CallOpts, operator common.Address) ([]common.Address, error) {
	var out []interface{}
	err := _Iservicemanager.contract.Call(opts, &out, "getOperatorRestakedStrategies", operator)

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// GetOperatorRestakedStrategies is a free data retrieval call binding the contract method 0x33cfb7b7.
//
// Solidity: function getOperatorRestakedStrategies(address operator) view returns(address[])
func (_Iservicemanager *IservicemanagerSession) GetOperatorRestakedStrategies(operator common.Address) ([]common.Address, error) {
	return _Iservicemanager.Contract.GetOperatorRestakedStrategies(&_Iservicemanager.CallOpts, operator)
}

// GetOperatorRestakedStrategies is a free data retrieval call binding the contract method 0x33cfb7b7.
//
// Solidity: function getOperatorRestakedStrategies(address operator) view returns(address[])
func (_Iservicemanager *IservicemanagerCallerSession) GetOperatorRestakedStrategies(operator common.Address) ([]common.Address, error) {
	return _Iservicemanager.Contract.GetOperatorRestakedStrategies(&_Iservicemanager.CallOpts, operator)
}

// GetRestakeableStrategies is a free data retrieval call binding the contract method 0xe481af9d.
//
// Solidity: function getRestakeableStrategies() view returns(address[])
func (_Iservicemanager *IservicemanagerCaller) GetRestakeableStrategies(opts *bind.CallOpts) ([]common.Address, error) {
	var out []interface{}
	err := _Iservicemanager.contract.Call(opts, &out, "getRestakeableStrategies")

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// GetRestakeableStrategies is a free data retrieval call binding the contract method 0xe481af9d.
//
// Solidity: function getRestakeableStrategies() view returns(address[])
func (_Iservicemanager *IservicemanagerSession) GetRestakeableStrategies() ([]common.Address, error) {
	return _Iservicemanager.Contract.GetRestakeableStrategies(&_Iservicemanager.CallOpts)
}

// GetRestakeableStrategies is a free data retrieval call binding the contract method 0xe481af9d.
//
// Solidity: function getRestakeableStrategies() view returns(address[])
func (_Iservicemanager *IservicemanagerCallerSession) GetRestakeableStrategies() ([]common.Address, error) {
	return _Iservicemanager.Contract.GetRestakeableStrategies(&_Iservicemanager.CallOpts)
}

// DeregisterOperatorFromAVS is a paid mutator transaction binding the contract method 0xa364f4da.
//
// Solidity: function deregisterOperatorFromAVS(address operator) returns()
func (_Iservicemanager *IservicemanagerTransactor) DeregisterOperatorFromAVS(opts *bind.TransactOpts, operator common.Address) (*types.Transaction, error) {
	return _Iservicemanager.contract.Transact(opts, "deregisterOperatorFromAVS", operator)
}

// DeregisterOperatorFromAVS is a paid mutator transaction binding the contract method 0xa364f4da.
//
// Solidity: function deregisterOperatorFromAVS(address operator) returns()
func (_Iservicemanager *IservicemanagerSession) DeregisterOperatorFromAVS(operator common.Address) (*types.Transaction, error) {
	return _Iservicemanager.Contract.DeregisterOperatorFromAVS(&_Iservicemanager.TransactOpts, operator)
}

// DeregisterOperatorFromAVS is a paid mutator transaction binding the contract method 0xa364f4da.
//
// Solidity: function deregisterOperatorFromAVS(address operator) returns()
func (_Iservicemanager *IservicemanagerTransactorSession) DeregisterOperatorFromAVS(operator common.Address) (*types.Transaction, error) {
	return _Iservicemanager.Contract.DeregisterOperatorFromAVS(&_Iservicemanager.TransactOpts, operator)
}

// RegisterOperatorToAVS is a paid mutator transaction binding the contract method 0x9926ee7d.
//
// Solidity: function registerOperatorToAVS(address operator, (bytes,bytes32,uint256) operatorSignature) returns()
func (_Iservicemanager *IservicemanagerTransactor) RegisterOperatorToAVS(opts *bind.TransactOpts, operator common.Address, operatorSignature ISignatureUtilsSignatureWithSaltAndExpiry) (*types.Transaction, error) {
	return _Iservicemanager.contract.Transact(opts, "registerOperatorToAVS", operator, operatorSignature)
}

// RegisterOperatorToAVS is a paid mutator transaction binding the contract method 0x9926ee7d.
//
// Solidity: function registerOperatorToAVS(address operator, (bytes,bytes32,uint256) operatorSignature) returns()
func (_Iservicemanager *IservicemanagerSession) RegisterOperatorToAVS(operator common.Address, operatorSignature ISignatureUtilsSignatureWithSaltAndExpiry) (*types.Transaction, error) {
	return _Iservicemanager.Contract.RegisterOperatorToAVS(&_Iservicemanager.TransactOpts, operator, operatorSignature)
}

// RegisterOperatorToAVS is a paid mutator transaction binding the contract method 0x9926ee7d.
//
// Solidity: function registerOperatorToAVS(address operator, (bytes,bytes32,uint256) operatorSignature) returns()
func (_Iservicemanager *IservicemanagerTransactorSession) RegisterOperatorToAVS(operator common.Address, operatorSignature ISignatureUtilsSignatureWithSaltAndExpiry) (*types.Transaction, error) {
	return _Iservicemanager.Contract.RegisterOperatorToAVS(&_Iservicemanager.TransactOpts, operator, operatorSignature)
}

// UpdateAVSMetadataURI is a paid mutator transaction binding the contract method 0xa98fb355.
//
// Solidity: function updateAVSMetadataURI(string _metadataURI) returns()
func (_Iservicemanager *IservicemanagerTransactor) UpdateAVSMetadataURI(opts *bind.TransactOpts, _metadataURI string) (*types.Transaction, error) {
	return _Iservicemanager.contract.Transact(opts, "updateAVSMetadataURI", _metadataURI)
}

// UpdateAVSMetadataURI is a paid mutator transaction binding the contract method 0xa98fb355.
//
// Solidity: function updateAVSMetadataURI(string _metadataURI) returns()
func (_Iservicemanager *IservicemanagerSession) UpdateAVSMetadataURI(_metadataURI string) (*types.Transaction, error) {
	return _Iservicemanager.Contract.UpdateAVSMetadataURI(&_Iservicemanager.TransactOpts, _metadataURI)
}

// UpdateAVSMetadataURI is a paid mutator transaction binding the contract method 0xa98fb355.
//
// Solidity: function updateAVSMetadataURI(string _metadataURI) returns()
func (_Iservicemanager *IservicemanagerTransactorSession) UpdateAVSMetadataURI(_metadataURI string) (*types.Transaction, error) {
	return _Iservicemanager.Contract.UpdateAVSMetadataURI(&_Iservicemanager.TransactOpts, _metadataURI)
}
