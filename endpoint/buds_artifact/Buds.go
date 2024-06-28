// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package buds_artifact

import (
	"errors"
	"math/big"
	"strings"

	"github.com/ethereum/go-ethereum"
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

// ArtifactsMetaData contains all meta data concerning the Artifacts contract.
var ArtifactsMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"burn\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"burnFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"mintTo\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// ArtifactsABI is the input ABI used to generate the binding from.
// Deprecated: Use ArtifactsMetaData.ABI instead.
var ArtifactsABI = ArtifactsMetaData.ABI

// Artifacts is an auto generated Go binding around an Ethereum contract.
type Artifacts struct {
	ArtifactsCaller     // Read-only binding to the contract
	ArtifactsTransactor // Write-only binding to the contract
	ArtifactsFilterer   // Log filterer for contract events
}

// ArtifactsCaller is an auto generated read-only Go binding around an Ethereum contract.
type ArtifactsCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ArtifactsTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ArtifactsTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ArtifactsFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ArtifactsFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ArtifactsSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ArtifactsSession struct {
	Contract     *Artifacts        // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ArtifactsCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ArtifactsCallerSession struct {
	Contract *ArtifactsCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts    // Call options to use throughout this session
}

// ArtifactsTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ArtifactsTransactorSession struct {
	Contract     *ArtifactsTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// ArtifactsRaw is an auto generated low-level Go binding around an Ethereum contract.
type ArtifactsRaw struct {
	Contract *Artifacts // Generic contract binding to access the raw methods on
}

// ArtifactsCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ArtifactsCallerRaw struct {
	Contract *ArtifactsCaller // Generic read-only contract binding to access the raw methods on
}

// ArtifactsTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ArtifactsTransactorRaw struct {
	Contract *ArtifactsTransactor // Generic write-only contract binding to access the raw methods on
}

// NewArtifacts creates a new instance of Artifacts, bound to a specific deployed contract.
func NewArtifacts(address common.Address, backend bind.ContractBackend) (*Artifacts, error) {
	contract, err := bindArtifacts(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Artifacts{ArtifactsCaller: ArtifactsCaller{contract: contract}, ArtifactsTransactor: ArtifactsTransactor{contract: contract}, ArtifactsFilterer: ArtifactsFilterer{contract: contract}}, nil
}

// NewArtifactsCaller creates a new read-only instance of Artifacts, bound to a specific deployed contract.
func NewArtifactsCaller(address common.Address, caller bind.ContractCaller) (*ArtifactsCaller, error) {
	contract, err := bindArtifacts(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ArtifactsCaller{contract: contract}, nil
}

// NewArtifactsTransactor creates a new write-only instance of Artifacts, bound to a specific deployed contract.
func NewArtifactsTransactor(address common.Address, transactor bind.ContractTransactor) (*ArtifactsTransactor, error) {
	contract, err := bindArtifacts(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ArtifactsTransactor{contract: contract}, nil
}

// NewArtifactsFilterer creates a new log filterer instance of Artifacts, bound to a specific deployed contract.
func NewArtifactsFilterer(address common.Address, filterer bind.ContractFilterer) (*ArtifactsFilterer, error) {
	contract, err := bindArtifacts(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ArtifactsFilterer{contract: contract}, nil
}

// bindArtifacts binds a generic wrapper to an already deployed contract.
func bindArtifacts(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := ArtifactsMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Artifacts *ArtifactsRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Artifacts.Contract.ArtifactsCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Artifacts *ArtifactsRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Artifacts.Contract.ArtifactsTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Artifacts *ArtifactsRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Artifacts.Contract.ArtifactsTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Artifacts *ArtifactsCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Artifacts.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Artifacts *ArtifactsTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Artifacts.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Artifacts *ArtifactsTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Artifacts.Contract.contract.Transact(opts, method, params...)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_Artifacts *ArtifactsCaller) BalanceOf(opts *bind.CallOpts, account common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Artifacts.contract.Call(opts, &out, "balanceOf", account)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_Artifacts *ArtifactsSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _Artifacts.Contract.BalanceOf(&_Artifacts.CallOpts, account)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_Artifacts *ArtifactsCallerSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _Artifacts.Contract.BalanceOf(&_Artifacts.CallOpts, account)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_Artifacts *ArtifactsCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Artifacts.contract.Call(opts, &out, "totalSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_Artifacts *ArtifactsSession) TotalSupply() (*big.Int, error) {
	return _Artifacts.Contract.TotalSupply(&_Artifacts.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_Artifacts *ArtifactsCallerSession) TotalSupply() (*big.Int, error) {
	return _Artifacts.Contract.TotalSupply(&_Artifacts.CallOpts)
}

// Burn is a paid mutator transaction binding the contract method 0x42966c68.
//
// Solidity: function burn(uint256 amount) returns()
func (_Artifacts *ArtifactsTransactor) Burn(opts *bind.TransactOpts, amount *big.Int) (*types.Transaction, error) {
	return _Artifacts.contract.Transact(opts, "burn", amount)
}

// Burn is a paid mutator transaction binding the contract method 0x42966c68.
//
// Solidity: function burn(uint256 amount) returns()
func (_Artifacts *ArtifactsSession) Burn(amount *big.Int) (*types.Transaction, error) {
	return _Artifacts.Contract.Burn(&_Artifacts.TransactOpts, amount)
}

// Burn is a paid mutator transaction binding the contract method 0x42966c68.
//
// Solidity: function burn(uint256 amount) returns()
func (_Artifacts *ArtifactsTransactorSession) Burn(amount *big.Int) (*types.Transaction, error) {
	return _Artifacts.Contract.Burn(&_Artifacts.TransactOpts, amount)
}

// BurnFrom is a paid mutator transaction binding the contract method 0x79cc6790.
//
// Solidity: function burnFrom(address from, uint256 amount) returns()
func (_Artifacts *ArtifactsTransactor) BurnFrom(opts *bind.TransactOpts, from common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Artifacts.contract.Transact(opts, "burnFrom", from, amount)
}

// BurnFrom is a paid mutator transaction binding the contract method 0x79cc6790.
//
// Solidity: function burnFrom(address from, uint256 amount) returns()
func (_Artifacts *ArtifactsSession) BurnFrom(from common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Artifacts.Contract.BurnFrom(&_Artifacts.TransactOpts, from, amount)
}

// BurnFrom is a paid mutator transaction binding the contract method 0x79cc6790.
//
// Solidity: function burnFrom(address from, uint256 amount) returns()
func (_Artifacts *ArtifactsTransactorSession) BurnFrom(from common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Artifacts.Contract.BurnFrom(&_Artifacts.TransactOpts, from, amount)
}

// MintTo is a paid mutator transaction binding the contract method 0x449a52f8.
//
// Solidity: function mintTo(address to, uint256 amount) returns()
func (_Artifacts *ArtifactsTransactor) MintTo(opts *bind.TransactOpts, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Artifacts.contract.Transact(opts, "mintTo", to, amount)
}

// MintTo is a paid mutator transaction binding the contract method 0x449a52f8.
//
// Solidity: function mintTo(address to, uint256 amount) returns()
func (_Artifacts *ArtifactsSession) MintTo(to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Artifacts.Contract.MintTo(&_Artifacts.TransactOpts, to, amount)
}

// MintTo is a paid mutator transaction binding the contract method 0x449a52f8.
//
// Solidity: function mintTo(address to, uint256 amount) returns()
func (_Artifacts *ArtifactsTransactorSession) MintTo(to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Artifacts.Contract.MintTo(&_Artifacts.TransactOpts, to, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 amount) returns(bool)
func (_Artifacts *ArtifactsTransactor) Transfer(opts *bind.TransactOpts, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Artifacts.contract.Transact(opts, "transfer", to, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 amount) returns(bool)
func (_Artifacts *ArtifactsSession) Transfer(to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Artifacts.Contract.Transfer(&_Artifacts.TransactOpts, to, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 amount) returns(bool)
func (_Artifacts *ArtifactsTransactorSession) Transfer(to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Artifacts.Contract.Transfer(&_Artifacts.TransactOpts, to, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 amount) returns(bool)
func (_Artifacts *ArtifactsTransactor) TransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Artifacts.contract.Transact(opts, "transferFrom", from, to, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 amount) returns(bool)
func (_Artifacts *ArtifactsSession) TransferFrom(from common.Address, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Artifacts.Contract.TransferFrom(&_Artifacts.TransactOpts, from, to, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 amount) returns(bool)
func (_Artifacts *ArtifactsTransactorSession) TransferFrom(from common.Address, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Artifacts.Contract.TransferFrom(&_Artifacts.TransactOpts, from, to, amount)
}
