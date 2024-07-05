// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package bindings

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

// DeveloperListMetaData contains all meta data concerning the DeveloperList contract.
var DeveloperListMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"function\",\"name\":\"addDeveloper\",\"inputs\":[{\"name\":\"addr\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"admin\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"commitChangeAdmin\",\"inputs\":[{\"name\":\"newAdmin\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"confirmChangeAdmin\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"disableDevVerify\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"enableDevVerify\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"enabled\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"initialize\",\"inputs\":[{\"name\":\"_admin\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"initialized\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"isDeveloper\",\"inputs\":[{\"name\":\"addr\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"pendingAdmin\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"removeDeveloper\",\"inputs\":[{\"name\":\"addr\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"event\",\"name\":\"AdminChanged\",\"inputs\":[{\"name\":\"newAdmin\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"AdminChanging\",\"inputs\":[{\"name\":\"newAdmin\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"DeveloperAdded\",\"inputs\":[{\"name\":\"addr\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"DeveloperRemoved\",\"inputs\":[{\"name\":\"addr\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"EnableStateChanged\",\"inputs\":[{\"name\":\"newState\",\"type\":\"bool\",\"indexed\":true,\"internalType\":\"bool\"}],\"anonymous\":false}]",
	Bin: "0x608060405234801561001057600080fd5b50610b16806100206000396000f3fe608060405234801561001057600080fd5b50600436106100d45760003560e01c80635eca4a7011610081578063db6619b01161005b578063db6619b0146101e1578063f851a440146101e9578063fb48270c1461020f57600080fd5b80635eca4a70146101825780639e23c209146101bb578063c4d66de8146101ce57600080fd5b806326782247116100b2578063267822471461012257806343e0c73a146101675780634fb9e9b71461016f57600080fd5b8063158ef93e146100d957806322fbf1e8146100fb578063238dafe014610110575b600080fd5b6000546100e69060ff1681565b60405190151581526020015b60405180910390f35b61010e610109366004610acc565b610217565b005b6000546100e690610100900460ff1681565b6001546101429073ffffffffffffffffffffffffffffffffffffffff1681565b60405173ffffffffffffffffffffffffffffffffffffffff90911681526020016100f2565b61010e6103aa565b61010e61017d366004610acc565b6104f4565b6100e6610190366004610acc565b73ffffffffffffffffffffffffffffffffffffffff1660009081526002602052604090205460ff1690565b61010e6101c9366004610acc565b6105ea565b61010e6101dc366004610acc565b610774565b61010e610856565b6000546101429062010000900473ffffffffffffffffffffffffffffffffffffffff1681565b61010e6109a6565b60005462010000900473ffffffffffffffffffffffffffffffffffffffff1633146102a3576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152600a60248201527f41646d696e206f6e6c790000000000000000000000000000000000000000000060448201526064015b60405180910390fd5b73ffffffffffffffffffffffffffffffffffffffff811660009081526002602052604090205460ff1615610333576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152600d60248201527f416c726561647920616464656400000000000000000000000000000000000000604482015260640161029a565b73ffffffffffffffffffffffffffffffffffffffff811660008181526002602052604080822080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00166001179055517f058fdae480ed8e99b762bceb2d39835a68ee3a4789cd84e5c90cd59722ba02099190a250565b60005462010000900473ffffffffffffffffffffffffffffffffffffffff163314610431576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152600a60248201527f41646d696e206f6e6c7900000000000000000000000000000000000000000000604482015260640161029a565b600054610100900460ff166104a2576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601060248201527f416c72656164792064697361626c656400000000000000000000000000000000604482015260640161029a565b600080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00ff1681556040517f733a7f99819dc7466bff56e7c0b6753b43b750a692f2a5bb4fe373815a0c7845908290a2565b60005462010000900473ffffffffffffffffffffffffffffffffffffffff16331461057b576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152600a60248201527f41646d696e206f6e6c7900000000000000000000000000000000000000000000604482015260640161029a565b600180547fffffffffffffffffffffffff00000000000000000000000000000000000000001673ffffffffffffffffffffffffffffffffffffffff83169081179091556040517faefcaa6215f99fe8c2f605dd268ee4d23a5b596bbca026e25ce8446187f4f1ba90600090a250565b60005462010000900473ffffffffffffffffffffffffffffffffffffffff163314610671576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152600a60248201527f41646d696e206f6e6c7900000000000000000000000000000000000000000000604482015260640161029a565b73ffffffffffffffffffffffffffffffffffffffff811660009081526002602052604090205460ff16610700576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152600f60248201527f4e6f74206120646576656c6f7065720000000000000000000000000000000000604482015260640161029a565b73ffffffffffffffffffffffffffffffffffffffff811660008181526002602052604080822080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00169055517f110a48e3e347ae018d4d40446e4e917b416f912dec489da19b4507bb9bb18cd49190a250565b60005460ff16156107e1576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601360248201527f416c726561647920696e697469616c697a656400000000000000000000000000604482015260640161029a565b600080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff0073ffffffffffffffffffffffffffffffffffffffff9093166201000002929092167fffffffffffffffffffff0000000000000000000000000000000000000000ff00909216919091176001179055565b60005462010000900473ffffffffffffffffffffffffffffffffffffffff1633146108dd576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152600a60248201527f41646d696e206f6e6c7900000000000000000000000000000000000000000000604482015260640161029a565b600054610100900460ff161561094f576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152600f60248201527f416c726561647920656e61626c65640000000000000000000000000000000000604482015260640161029a565b600080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00ff166101001781556040516001917f733a7f99819dc7466bff56e7c0b6753b43b750a692f2a5bb4fe373815a0c784591a2565b60015473ffffffffffffffffffffffffffffffffffffffff163314610a27576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152600e60248201527f4e65772061646d696e206f6e6c79000000000000000000000000000000000000604482015260640161029a565b60018054600080547fffffffffffffffffffff0000000000000000000000000000000000000000ffff1673ffffffffffffffffffffffffffffffffffffffff80841662010000908102929092178084557fffffffffffffffffffffffff00000000000000000000000000000000000000009094169094556040519204909216917f7ce7ec0b50378fb6c0186ffb5f48325f6593fcb4ca4386f21861af3129188f5c91a2565b600060208284031215610ade57600080fd5b813573ffffffffffffffffffffffffffffffffffffffff81168114610b0257600080fd5b939250505056fea164736f6c634300080f000a",
}

// DeveloperListABI is the input ABI used to generate the binding from.
// Deprecated: Use DeveloperListMetaData.ABI instead.
var DeveloperListABI = DeveloperListMetaData.ABI

// DeveloperListBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use DeveloperListMetaData.Bin instead.
var DeveloperListBin = DeveloperListMetaData.Bin

// DeployDeveloperList deploys a new Ethereum contract, binding an instance of DeveloperList to it.
func DeployDeveloperList(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *DeveloperList, error) {
	parsed, err := DeveloperListMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(DeveloperListBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &DeveloperList{DeveloperListCaller: DeveloperListCaller{contract: contract}, DeveloperListTransactor: DeveloperListTransactor{contract: contract}, DeveloperListFilterer: DeveloperListFilterer{contract: contract}}, nil
}

// DeveloperList is an auto generated Go binding around an Ethereum contract.
type DeveloperList struct {
	DeveloperListCaller     // Read-only binding to the contract
	DeveloperListTransactor // Write-only binding to the contract
	DeveloperListFilterer   // Log filterer for contract events
}

// DeveloperListCaller is an auto generated read-only Go binding around an Ethereum contract.
type DeveloperListCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DeveloperListTransactor is an auto generated write-only Go binding around an Ethereum contract.
type DeveloperListTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DeveloperListFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type DeveloperListFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DeveloperListSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type DeveloperListSession struct {
	Contract     *DeveloperList    // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// DeveloperListCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type DeveloperListCallerSession struct {
	Contract *DeveloperListCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts        // Call options to use throughout this session
}

// DeveloperListTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type DeveloperListTransactorSession struct {
	Contract     *DeveloperListTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts        // Transaction auth options to use throughout this session
}

// DeveloperListRaw is an auto generated low-level Go binding around an Ethereum contract.
type DeveloperListRaw struct {
	Contract *DeveloperList // Generic contract binding to access the raw methods on
}

// DeveloperListCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type DeveloperListCallerRaw struct {
	Contract *DeveloperListCaller // Generic read-only contract binding to access the raw methods on
}

// DeveloperListTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type DeveloperListTransactorRaw struct {
	Contract *DeveloperListTransactor // Generic write-only contract binding to access the raw methods on
}

// NewDeveloperList creates a new instance of DeveloperList, bound to a specific deployed contract.
func NewDeveloperList(address common.Address, backend bind.ContractBackend) (*DeveloperList, error) {
	contract, err := bindDeveloperList(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &DeveloperList{DeveloperListCaller: DeveloperListCaller{contract: contract}, DeveloperListTransactor: DeveloperListTransactor{contract: contract}, DeveloperListFilterer: DeveloperListFilterer{contract: contract}}, nil
}

// NewDeveloperListCaller creates a new read-only instance of DeveloperList, bound to a specific deployed contract.
func NewDeveloperListCaller(address common.Address, caller bind.ContractCaller) (*DeveloperListCaller, error) {
	contract, err := bindDeveloperList(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &DeveloperListCaller{contract: contract}, nil
}

// NewDeveloperListTransactor creates a new write-only instance of DeveloperList, bound to a specific deployed contract.
func NewDeveloperListTransactor(address common.Address, transactor bind.ContractTransactor) (*DeveloperListTransactor, error) {
	contract, err := bindDeveloperList(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &DeveloperListTransactor{contract: contract}, nil
}

// NewDeveloperListFilterer creates a new log filterer instance of DeveloperList, bound to a specific deployed contract.
func NewDeveloperListFilterer(address common.Address, filterer bind.ContractFilterer) (*DeveloperListFilterer, error) {
	contract, err := bindDeveloperList(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &DeveloperListFilterer{contract: contract}, nil
}

// bindDeveloperList binds a generic wrapper to an already deployed contract.
func bindDeveloperList(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := DeveloperListMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_DeveloperList *DeveloperListRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _DeveloperList.Contract.DeveloperListCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_DeveloperList *DeveloperListRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _DeveloperList.Contract.DeveloperListTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_DeveloperList *DeveloperListRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _DeveloperList.Contract.DeveloperListTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_DeveloperList *DeveloperListCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _DeveloperList.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_DeveloperList *DeveloperListTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _DeveloperList.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_DeveloperList *DeveloperListTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _DeveloperList.Contract.contract.Transact(opts, method, params...)
}

// Admin is a free data retrieval call binding the contract method 0xf851a440.
//
// Solidity: function admin() view returns(address)
func (_DeveloperList *DeveloperListCaller) Admin(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _DeveloperList.contract.Call(opts, &out, "admin")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Admin is a free data retrieval call binding the contract method 0xf851a440.
//
// Solidity: function admin() view returns(address)
func (_DeveloperList *DeveloperListSession) Admin() (common.Address, error) {
	return _DeveloperList.Contract.Admin(&_DeveloperList.CallOpts)
}

// Admin is a free data retrieval call binding the contract method 0xf851a440.
//
// Solidity: function admin() view returns(address)
func (_DeveloperList *DeveloperListCallerSession) Admin() (common.Address, error) {
	return _DeveloperList.Contract.Admin(&_DeveloperList.CallOpts)
}

// Enabled is a free data retrieval call binding the contract method 0x238dafe0.
//
// Solidity: function enabled() view returns(bool)
func (_DeveloperList *DeveloperListCaller) Enabled(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _DeveloperList.contract.Call(opts, &out, "enabled")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Enabled is a free data retrieval call binding the contract method 0x238dafe0.
//
// Solidity: function enabled() view returns(bool)
func (_DeveloperList *DeveloperListSession) Enabled() (bool, error) {
	return _DeveloperList.Contract.Enabled(&_DeveloperList.CallOpts)
}

// Enabled is a free data retrieval call binding the contract method 0x238dafe0.
//
// Solidity: function enabled() view returns(bool)
func (_DeveloperList *DeveloperListCallerSession) Enabled() (bool, error) {
	return _DeveloperList.Contract.Enabled(&_DeveloperList.CallOpts)
}

// Initialized is a free data retrieval call binding the contract method 0x158ef93e.
//
// Solidity: function initialized() view returns(bool)
func (_DeveloperList *DeveloperListCaller) Initialized(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _DeveloperList.contract.Call(opts, &out, "initialized")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Initialized is a free data retrieval call binding the contract method 0x158ef93e.
//
// Solidity: function initialized() view returns(bool)
func (_DeveloperList *DeveloperListSession) Initialized() (bool, error) {
	return _DeveloperList.Contract.Initialized(&_DeveloperList.CallOpts)
}

// Initialized is a free data retrieval call binding the contract method 0x158ef93e.
//
// Solidity: function initialized() view returns(bool)
func (_DeveloperList *DeveloperListCallerSession) Initialized() (bool, error) {
	return _DeveloperList.Contract.Initialized(&_DeveloperList.CallOpts)
}

// IsDeveloper is a free data retrieval call binding the contract method 0x5eca4a70.
//
// Solidity: function isDeveloper(address addr) view returns(bool)
func (_DeveloperList *DeveloperListCaller) IsDeveloper(opts *bind.CallOpts, addr common.Address) (bool, error) {
	var out []interface{}
	err := _DeveloperList.contract.Call(opts, &out, "isDeveloper", addr)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsDeveloper is a free data retrieval call binding the contract method 0x5eca4a70.
//
// Solidity: function isDeveloper(address addr) view returns(bool)
func (_DeveloperList *DeveloperListSession) IsDeveloper(addr common.Address) (bool, error) {
	return _DeveloperList.Contract.IsDeveloper(&_DeveloperList.CallOpts, addr)
}

// IsDeveloper is a free data retrieval call binding the contract method 0x5eca4a70.
//
// Solidity: function isDeveloper(address addr) view returns(bool)
func (_DeveloperList *DeveloperListCallerSession) IsDeveloper(addr common.Address) (bool, error) {
	return _DeveloperList.Contract.IsDeveloper(&_DeveloperList.CallOpts, addr)
}

// PendingAdmin is a free data retrieval call binding the contract method 0x26782247.
//
// Solidity: function pendingAdmin() view returns(address)
func (_DeveloperList *DeveloperListCaller) PendingAdmin(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _DeveloperList.contract.Call(opts, &out, "pendingAdmin")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// PendingAdmin is a free data retrieval call binding the contract method 0x26782247.
//
// Solidity: function pendingAdmin() view returns(address)
func (_DeveloperList *DeveloperListSession) PendingAdmin() (common.Address, error) {
	return _DeveloperList.Contract.PendingAdmin(&_DeveloperList.CallOpts)
}

// PendingAdmin is a free data retrieval call binding the contract method 0x26782247.
//
// Solidity: function pendingAdmin() view returns(address)
func (_DeveloperList *DeveloperListCallerSession) PendingAdmin() (common.Address, error) {
	return _DeveloperList.Contract.PendingAdmin(&_DeveloperList.CallOpts)
}

// AddDeveloper is a paid mutator transaction binding the contract method 0x22fbf1e8.
//
// Solidity: function addDeveloper(address addr) returns()
func (_DeveloperList *DeveloperListTransactor) AddDeveloper(opts *bind.TransactOpts, addr common.Address) (*types.Transaction, error) {
	return _DeveloperList.contract.Transact(opts, "addDeveloper", addr)
}

// AddDeveloper is a paid mutator transaction binding the contract method 0x22fbf1e8.
//
// Solidity: function addDeveloper(address addr) returns()
func (_DeveloperList *DeveloperListSession) AddDeveloper(addr common.Address) (*types.Transaction, error) {
	return _DeveloperList.Contract.AddDeveloper(&_DeveloperList.TransactOpts, addr)
}

// AddDeveloper is a paid mutator transaction binding the contract method 0x22fbf1e8.
//
// Solidity: function addDeveloper(address addr) returns()
func (_DeveloperList *DeveloperListTransactorSession) AddDeveloper(addr common.Address) (*types.Transaction, error) {
	return _DeveloperList.Contract.AddDeveloper(&_DeveloperList.TransactOpts, addr)
}

// CommitChangeAdmin is a paid mutator transaction binding the contract method 0x4fb9e9b7.
//
// Solidity: function commitChangeAdmin(address newAdmin) returns()
func (_DeveloperList *DeveloperListTransactor) CommitChangeAdmin(opts *bind.TransactOpts, newAdmin common.Address) (*types.Transaction, error) {
	return _DeveloperList.contract.Transact(opts, "commitChangeAdmin", newAdmin)
}

// CommitChangeAdmin is a paid mutator transaction binding the contract method 0x4fb9e9b7.
//
// Solidity: function commitChangeAdmin(address newAdmin) returns()
func (_DeveloperList *DeveloperListSession) CommitChangeAdmin(newAdmin common.Address) (*types.Transaction, error) {
	return _DeveloperList.Contract.CommitChangeAdmin(&_DeveloperList.TransactOpts, newAdmin)
}

// CommitChangeAdmin is a paid mutator transaction binding the contract method 0x4fb9e9b7.
//
// Solidity: function commitChangeAdmin(address newAdmin) returns()
func (_DeveloperList *DeveloperListTransactorSession) CommitChangeAdmin(newAdmin common.Address) (*types.Transaction, error) {
	return _DeveloperList.Contract.CommitChangeAdmin(&_DeveloperList.TransactOpts, newAdmin)
}

// ConfirmChangeAdmin is a paid mutator transaction binding the contract method 0xfb48270c.
//
// Solidity: function confirmChangeAdmin() returns()
func (_DeveloperList *DeveloperListTransactor) ConfirmChangeAdmin(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _DeveloperList.contract.Transact(opts, "confirmChangeAdmin")
}

// ConfirmChangeAdmin is a paid mutator transaction binding the contract method 0xfb48270c.
//
// Solidity: function confirmChangeAdmin() returns()
func (_DeveloperList *DeveloperListSession) ConfirmChangeAdmin() (*types.Transaction, error) {
	return _DeveloperList.Contract.ConfirmChangeAdmin(&_DeveloperList.TransactOpts)
}

// ConfirmChangeAdmin is a paid mutator transaction binding the contract method 0xfb48270c.
//
// Solidity: function confirmChangeAdmin() returns()
func (_DeveloperList *DeveloperListTransactorSession) ConfirmChangeAdmin() (*types.Transaction, error) {
	return _DeveloperList.Contract.ConfirmChangeAdmin(&_DeveloperList.TransactOpts)
}

// DisableDevVerify is a paid mutator transaction binding the contract method 0x43e0c73a.
//
// Solidity: function disableDevVerify() returns()
func (_DeveloperList *DeveloperListTransactor) DisableDevVerify(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _DeveloperList.contract.Transact(opts, "disableDevVerify")
}

// DisableDevVerify is a paid mutator transaction binding the contract method 0x43e0c73a.
//
// Solidity: function disableDevVerify() returns()
func (_DeveloperList *DeveloperListSession) DisableDevVerify() (*types.Transaction, error) {
	return _DeveloperList.Contract.DisableDevVerify(&_DeveloperList.TransactOpts)
}

// DisableDevVerify is a paid mutator transaction binding the contract method 0x43e0c73a.
//
// Solidity: function disableDevVerify() returns()
func (_DeveloperList *DeveloperListTransactorSession) DisableDevVerify() (*types.Transaction, error) {
	return _DeveloperList.Contract.DisableDevVerify(&_DeveloperList.TransactOpts)
}

// EnableDevVerify is a paid mutator transaction binding the contract method 0xdb6619b0.
//
// Solidity: function enableDevVerify() returns()
func (_DeveloperList *DeveloperListTransactor) EnableDevVerify(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _DeveloperList.contract.Transact(opts, "enableDevVerify")
}

// EnableDevVerify is a paid mutator transaction binding the contract method 0xdb6619b0.
//
// Solidity: function enableDevVerify() returns()
func (_DeveloperList *DeveloperListSession) EnableDevVerify() (*types.Transaction, error) {
	return _DeveloperList.Contract.EnableDevVerify(&_DeveloperList.TransactOpts)
}

// EnableDevVerify is a paid mutator transaction binding the contract method 0xdb6619b0.
//
// Solidity: function enableDevVerify() returns()
func (_DeveloperList *DeveloperListTransactorSession) EnableDevVerify() (*types.Transaction, error) {
	return _DeveloperList.Contract.EnableDevVerify(&_DeveloperList.TransactOpts)
}

// Initialize is a paid mutator transaction binding the contract method 0xc4d66de8.
//
// Solidity: function initialize(address _admin) returns()
func (_DeveloperList *DeveloperListTransactor) Initialize(opts *bind.TransactOpts, _admin common.Address) (*types.Transaction, error) {
	return _DeveloperList.contract.Transact(opts, "initialize", _admin)
}

// Initialize is a paid mutator transaction binding the contract method 0xc4d66de8.
//
// Solidity: function initialize(address _admin) returns()
func (_DeveloperList *DeveloperListSession) Initialize(_admin common.Address) (*types.Transaction, error) {
	return _DeveloperList.Contract.Initialize(&_DeveloperList.TransactOpts, _admin)
}

// Initialize is a paid mutator transaction binding the contract method 0xc4d66de8.
//
// Solidity: function initialize(address _admin) returns()
func (_DeveloperList *DeveloperListTransactorSession) Initialize(_admin common.Address) (*types.Transaction, error) {
	return _DeveloperList.Contract.Initialize(&_DeveloperList.TransactOpts, _admin)
}

// RemoveDeveloper is a paid mutator transaction binding the contract method 0x9e23c209.
//
// Solidity: function removeDeveloper(address addr) returns()
func (_DeveloperList *DeveloperListTransactor) RemoveDeveloper(opts *bind.TransactOpts, addr common.Address) (*types.Transaction, error) {
	return _DeveloperList.contract.Transact(opts, "removeDeveloper", addr)
}

// RemoveDeveloper is a paid mutator transaction binding the contract method 0x9e23c209.
//
// Solidity: function removeDeveloper(address addr) returns()
func (_DeveloperList *DeveloperListSession) RemoveDeveloper(addr common.Address) (*types.Transaction, error) {
	return _DeveloperList.Contract.RemoveDeveloper(&_DeveloperList.TransactOpts, addr)
}

// RemoveDeveloper is a paid mutator transaction binding the contract method 0x9e23c209.
//
// Solidity: function removeDeveloper(address addr) returns()
func (_DeveloperList *DeveloperListTransactorSession) RemoveDeveloper(addr common.Address) (*types.Transaction, error) {
	return _DeveloperList.Contract.RemoveDeveloper(&_DeveloperList.TransactOpts, addr)
}

// DeveloperListAdminChangedIterator is returned from FilterAdminChanged and is used to iterate over the raw logs and unpacked data for AdminChanged events raised by the DeveloperList contract.
type DeveloperListAdminChangedIterator struct {
	Event *DeveloperListAdminChanged // Event containing the contract specifics and raw log

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
func (it *DeveloperListAdminChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DeveloperListAdminChanged)
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
		it.Event = new(DeveloperListAdminChanged)
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
func (it *DeveloperListAdminChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DeveloperListAdminChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DeveloperListAdminChanged represents a AdminChanged event raised by the DeveloperList contract.
type DeveloperListAdminChanged struct {
	NewAdmin common.Address
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterAdminChanged is a free log retrieval operation binding the contract event 0x7ce7ec0b50378fb6c0186ffb5f48325f6593fcb4ca4386f21861af3129188f5c.
//
// Solidity: event AdminChanged(address indexed newAdmin)
func (_DeveloperList *DeveloperListFilterer) FilterAdminChanged(opts *bind.FilterOpts, newAdmin []common.Address) (*DeveloperListAdminChangedIterator, error) {

	var newAdminRule []interface{}
	for _, newAdminItem := range newAdmin {
		newAdminRule = append(newAdminRule, newAdminItem)
	}

	logs, sub, err := _DeveloperList.contract.FilterLogs(opts, "AdminChanged", newAdminRule)
	if err != nil {
		return nil, err
	}
	return &DeveloperListAdminChangedIterator{contract: _DeveloperList.contract, event: "AdminChanged", logs: logs, sub: sub}, nil
}

// WatchAdminChanged is a free log subscription operation binding the contract event 0x7ce7ec0b50378fb6c0186ffb5f48325f6593fcb4ca4386f21861af3129188f5c.
//
// Solidity: event AdminChanged(address indexed newAdmin)
func (_DeveloperList *DeveloperListFilterer) WatchAdminChanged(opts *bind.WatchOpts, sink chan<- *DeveloperListAdminChanged, newAdmin []common.Address) (event.Subscription, error) {

	var newAdminRule []interface{}
	for _, newAdminItem := range newAdmin {
		newAdminRule = append(newAdminRule, newAdminItem)
	}

	logs, sub, err := _DeveloperList.contract.WatchLogs(opts, "AdminChanged", newAdminRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DeveloperListAdminChanged)
				if err := _DeveloperList.contract.UnpackLog(event, "AdminChanged", log); err != nil {
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

// ParseAdminChanged is a log parse operation binding the contract event 0x7ce7ec0b50378fb6c0186ffb5f48325f6593fcb4ca4386f21861af3129188f5c.
//
// Solidity: event AdminChanged(address indexed newAdmin)
func (_DeveloperList *DeveloperListFilterer) ParseAdminChanged(log types.Log) (*DeveloperListAdminChanged, error) {
	event := new(DeveloperListAdminChanged)
	if err := _DeveloperList.contract.UnpackLog(event, "AdminChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// DeveloperListAdminChangingIterator is returned from FilterAdminChanging and is used to iterate over the raw logs and unpacked data for AdminChanging events raised by the DeveloperList contract.
type DeveloperListAdminChangingIterator struct {
	Event *DeveloperListAdminChanging // Event containing the contract specifics and raw log

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
func (it *DeveloperListAdminChangingIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DeveloperListAdminChanging)
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
		it.Event = new(DeveloperListAdminChanging)
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
func (it *DeveloperListAdminChangingIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DeveloperListAdminChangingIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DeveloperListAdminChanging represents a AdminChanging event raised by the DeveloperList contract.
type DeveloperListAdminChanging struct {
	NewAdmin common.Address
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterAdminChanging is a free log retrieval operation binding the contract event 0xaefcaa6215f99fe8c2f605dd268ee4d23a5b596bbca026e25ce8446187f4f1ba.
//
// Solidity: event AdminChanging(address indexed newAdmin)
func (_DeveloperList *DeveloperListFilterer) FilterAdminChanging(opts *bind.FilterOpts, newAdmin []common.Address) (*DeveloperListAdminChangingIterator, error) {

	var newAdminRule []interface{}
	for _, newAdminItem := range newAdmin {
		newAdminRule = append(newAdminRule, newAdminItem)
	}

	logs, sub, err := _DeveloperList.contract.FilterLogs(opts, "AdminChanging", newAdminRule)
	if err != nil {
		return nil, err
	}
	return &DeveloperListAdminChangingIterator{contract: _DeveloperList.contract, event: "AdminChanging", logs: logs, sub: sub}, nil
}

// WatchAdminChanging is a free log subscription operation binding the contract event 0xaefcaa6215f99fe8c2f605dd268ee4d23a5b596bbca026e25ce8446187f4f1ba.
//
// Solidity: event AdminChanging(address indexed newAdmin)
func (_DeveloperList *DeveloperListFilterer) WatchAdminChanging(opts *bind.WatchOpts, sink chan<- *DeveloperListAdminChanging, newAdmin []common.Address) (event.Subscription, error) {

	var newAdminRule []interface{}
	for _, newAdminItem := range newAdmin {
		newAdminRule = append(newAdminRule, newAdminItem)
	}

	logs, sub, err := _DeveloperList.contract.WatchLogs(opts, "AdminChanging", newAdminRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DeveloperListAdminChanging)
				if err := _DeveloperList.contract.UnpackLog(event, "AdminChanging", log); err != nil {
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

// ParseAdminChanging is a log parse operation binding the contract event 0xaefcaa6215f99fe8c2f605dd268ee4d23a5b596bbca026e25ce8446187f4f1ba.
//
// Solidity: event AdminChanging(address indexed newAdmin)
func (_DeveloperList *DeveloperListFilterer) ParseAdminChanging(log types.Log) (*DeveloperListAdminChanging, error) {
	event := new(DeveloperListAdminChanging)
	if err := _DeveloperList.contract.UnpackLog(event, "AdminChanging", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// DeveloperListDeveloperAddedIterator is returned from FilterDeveloperAdded and is used to iterate over the raw logs and unpacked data for DeveloperAdded events raised by the DeveloperList contract.
type DeveloperListDeveloperAddedIterator struct {
	Event *DeveloperListDeveloperAdded // Event containing the contract specifics and raw log

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
func (it *DeveloperListDeveloperAddedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DeveloperListDeveloperAdded)
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
		it.Event = new(DeveloperListDeveloperAdded)
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
func (it *DeveloperListDeveloperAddedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DeveloperListDeveloperAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DeveloperListDeveloperAdded represents a DeveloperAdded event raised by the DeveloperList contract.
type DeveloperListDeveloperAdded struct {
	Addr common.Address
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterDeveloperAdded is a free log retrieval operation binding the contract event 0x058fdae480ed8e99b762bceb2d39835a68ee3a4789cd84e5c90cd59722ba0209.
//
// Solidity: event DeveloperAdded(address indexed addr)
func (_DeveloperList *DeveloperListFilterer) FilterDeveloperAdded(opts *bind.FilterOpts, addr []common.Address) (*DeveloperListDeveloperAddedIterator, error) {

	var addrRule []interface{}
	for _, addrItem := range addr {
		addrRule = append(addrRule, addrItem)
	}

	logs, sub, err := _DeveloperList.contract.FilterLogs(opts, "DeveloperAdded", addrRule)
	if err != nil {
		return nil, err
	}
	return &DeveloperListDeveloperAddedIterator{contract: _DeveloperList.contract, event: "DeveloperAdded", logs: logs, sub: sub}, nil
}

// WatchDeveloperAdded is a free log subscription operation binding the contract event 0x058fdae480ed8e99b762bceb2d39835a68ee3a4789cd84e5c90cd59722ba0209.
//
// Solidity: event DeveloperAdded(address indexed addr)
func (_DeveloperList *DeveloperListFilterer) WatchDeveloperAdded(opts *bind.WatchOpts, sink chan<- *DeveloperListDeveloperAdded, addr []common.Address) (event.Subscription, error) {

	var addrRule []interface{}
	for _, addrItem := range addr {
		addrRule = append(addrRule, addrItem)
	}

	logs, sub, err := _DeveloperList.contract.WatchLogs(opts, "DeveloperAdded", addrRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DeveloperListDeveloperAdded)
				if err := _DeveloperList.contract.UnpackLog(event, "DeveloperAdded", log); err != nil {
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

// ParseDeveloperAdded is a log parse operation binding the contract event 0x058fdae480ed8e99b762bceb2d39835a68ee3a4789cd84e5c90cd59722ba0209.
//
// Solidity: event DeveloperAdded(address indexed addr)
func (_DeveloperList *DeveloperListFilterer) ParseDeveloperAdded(log types.Log) (*DeveloperListDeveloperAdded, error) {
	event := new(DeveloperListDeveloperAdded)
	if err := _DeveloperList.contract.UnpackLog(event, "DeveloperAdded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// DeveloperListDeveloperRemovedIterator is returned from FilterDeveloperRemoved and is used to iterate over the raw logs and unpacked data for DeveloperRemoved events raised by the DeveloperList contract.
type DeveloperListDeveloperRemovedIterator struct {
	Event *DeveloperListDeveloperRemoved // Event containing the contract specifics and raw log

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
func (it *DeveloperListDeveloperRemovedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DeveloperListDeveloperRemoved)
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
		it.Event = new(DeveloperListDeveloperRemoved)
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
func (it *DeveloperListDeveloperRemovedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DeveloperListDeveloperRemovedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DeveloperListDeveloperRemoved represents a DeveloperRemoved event raised by the DeveloperList contract.
type DeveloperListDeveloperRemoved struct {
	Addr common.Address
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterDeveloperRemoved is a free log retrieval operation binding the contract event 0x110a48e3e347ae018d4d40446e4e917b416f912dec489da19b4507bb9bb18cd4.
//
// Solidity: event DeveloperRemoved(address indexed addr)
func (_DeveloperList *DeveloperListFilterer) FilterDeveloperRemoved(opts *bind.FilterOpts, addr []common.Address) (*DeveloperListDeveloperRemovedIterator, error) {

	var addrRule []interface{}
	for _, addrItem := range addr {
		addrRule = append(addrRule, addrItem)
	}

	logs, sub, err := _DeveloperList.contract.FilterLogs(opts, "DeveloperRemoved", addrRule)
	if err != nil {
		return nil, err
	}
	return &DeveloperListDeveloperRemovedIterator{contract: _DeveloperList.contract, event: "DeveloperRemoved", logs: logs, sub: sub}, nil
}

// WatchDeveloperRemoved is a free log subscription operation binding the contract event 0x110a48e3e347ae018d4d40446e4e917b416f912dec489da19b4507bb9bb18cd4.
//
// Solidity: event DeveloperRemoved(address indexed addr)
func (_DeveloperList *DeveloperListFilterer) WatchDeveloperRemoved(opts *bind.WatchOpts, sink chan<- *DeveloperListDeveloperRemoved, addr []common.Address) (event.Subscription, error) {

	var addrRule []interface{}
	for _, addrItem := range addr {
		addrRule = append(addrRule, addrItem)
	}

	logs, sub, err := _DeveloperList.contract.WatchLogs(opts, "DeveloperRemoved", addrRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DeveloperListDeveloperRemoved)
				if err := _DeveloperList.contract.UnpackLog(event, "DeveloperRemoved", log); err != nil {
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

// ParseDeveloperRemoved is a log parse operation binding the contract event 0x110a48e3e347ae018d4d40446e4e917b416f912dec489da19b4507bb9bb18cd4.
//
// Solidity: event DeveloperRemoved(address indexed addr)
func (_DeveloperList *DeveloperListFilterer) ParseDeveloperRemoved(log types.Log) (*DeveloperListDeveloperRemoved, error) {
	event := new(DeveloperListDeveloperRemoved)
	if err := _DeveloperList.contract.UnpackLog(event, "DeveloperRemoved", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// DeveloperListEnableStateChangedIterator is returned from FilterEnableStateChanged and is used to iterate over the raw logs and unpacked data for EnableStateChanged events raised by the DeveloperList contract.
type DeveloperListEnableStateChangedIterator struct {
	Event *DeveloperListEnableStateChanged // Event containing the contract specifics and raw log

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
func (it *DeveloperListEnableStateChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DeveloperListEnableStateChanged)
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
		it.Event = new(DeveloperListEnableStateChanged)
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
func (it *DeveloperListEnableStateChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DeveloperListEnableStateChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DeveloperListEnableStateChanged represents a EnableStateChanged event raised by the DeveloperList contract.
type DeveloperListEnableStateChanged struct {
	NewState bool
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterEnableStateChanged is a free log retrieval operation binding the contract event 0x733a7f99819dc7466bff56e7c0b6753b43b750a692f2a5bb4fe373815a0c7845.
//
// Solidity: event EnableStateChanged(bool indexed newState)
func (_DeveloperList *DeveloperListFilterer) FilterEnableStateChanged(opts *bind.FilterOpts, newState []bool) (*DeveloperListEnableStateChangedIterator, error) {

	var newStateRule []interface{}
	for _, newStateItem := range newState {
		newStateRule = append(newStateRule, newStateItem)
	}

	logs, sub, err := _DeveloperList.contract.FilterLogs(opts, "EnableStateChanged", newStateRule)
	if err != nil {
		return nil, err
	}
	return &DeveloperListEnableStateChangedIterator{contract: _DeveloperList.contract, event: "EnableStateChanged", logs: logs, sub: sub}, nil
}

// WatchEnableStateChanged is a free log subscription operation binding the contract event 0x733a7f99819dc7466bff56e7c0b6753b43b750a692f2a5bb4fe373815a0c7845.
//
// Solidity: event EnableStateChanged(bool indexed newState)
func (_DeveloperList *DeveloperListFilterer) WatchEnableStateChanged(opts *bind.WatchOpts, sink chan<- *DeveloperListEnableStateChanged, newState []bool) (event.Subscription, error) {

	var newStateRule []interface{}
	for _, newStateItem := range newState {
		newStateRule = append(newStateRule, newStateItem)
	}

	logs, sub, err := _DeveloperList.contract.WatchLogs(opts, "EnableStateChanged", newStateRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DeveloperListEnableStateChanged)
				if err := _DeveloperList.contract.UnpackLog(event, "EnableStateChanged", log); err != nil {
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

// ParseEnableStateChanged is a log parse operation binding the contract event 0x733a7f99819dc7466bff56e7c0b6753b43b750a692f2a5bb4fe373815a0c7845.
//
// Solidity: event EnableStateChanged(bool indexed newState)
func (_DeveloperList *DeveloperListFilterer) ParseEnableStateChanged(log types.Log) (*DeveloperListEnableStateChanged, error) {
	event := new(DeveloperListEnableStateChanged)
	if err := _DeveloperList.contract.UnpackLog(event, "EnableStateChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
