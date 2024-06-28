// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package getter_artifacts

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

// LibGlobalVarStateStake is an auto generated low-level Go binding around an user-defined struct.
type LibGlobalVarStateStake struct {
	Owner         common.Address
	TimeStamp     *big.Int
	BudsAmount    *big.Int
	FarmerTokenId *big.Int
}

// ArtifactsMetaData contains all meta data concerning the Artifacts contract.
var ArtifactsMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"name\":\"closeContest\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getCurrentApr\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getGlobalStakedBuds\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getNoOfChains\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getNumberOfStakers\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getTotalStakedFarmers\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"}],\"name\":\"getUserStakes\",\"outputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"timeStamp\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"budsAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"farmerTokenId\",\"type\":\"uint256\"}],\"internalType\":\"structLibGlobalVarState.Stake[]\",\"name\":\"\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getlocalStakedBuds\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"liquidity\",\"type\":\"uint256\"}],\"name\":\"setGlobalStakedBuds\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"chains\",\"type\":\"uint256\"}],\"name\":\"setNoOfChains\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_raidFees\",\"type\":\"uint256\"}],\"name\":\"setRaidFees\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_address\",\"type\":\"address\"}],\"name\":\"setRaidHandler\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"addresspayable\",\"name\":\"newAddress\",\"type\":\"address\"}],\"name\":\"setTreasury\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"startContest\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
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

// GetCurrentApr is a free data retrieval call binding the contract method 0x5ec6bd9f.
//
// Solidity: function getCurrentApr() view returns(uint256)
func (_Artifacts *ArtifactsCaller) GetCurrentApr(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Artifacts.contract.Call(opts, &out, "getCurrentApr")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetCurrentApr is a free data retrieval call binding the contract method 0x5ec6bd9f.
//
// Solidity: function getCurrentApr() view returns(uint256)
func (_Artifacts *ArtifactsSession) GetCurrentApr() (*big.Int, error) {
	return _Artifacts.Contract.GetCurrentApr(&_Artifacts.CallOpts)
}

// GetCurrentApr is a free data retrieval call binding the contract method 0x5ec6bd9f.
//
// Solidity: function getCurrentApr() view returns(uint256)
func (_Artifacts *ArtifactsCallerSession) GetCurrentApr() (*big.Int, error) {
	return _Artifacts.Contract.GetCurrentApr(&_Artifacts.CallOpts)
}

// GetGlobalStakedBuds is a free data retrieval call binding the contract method 0x0903b335.
//
// Solidity: function getGlobalStakedBuds() view returns(uint256)
func (_Artifacts *ArtifactsCaller) GetGlobalStakedBuds(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Artifacts.contract.Call(opts, &out, "getGlobalStakedBuds")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetGlobalStakedBuds is a free data retrieval call binding the contract method 0x0903b335.
//
// Solidity: function getGlobalStakedBuds() view returns(uint256)
func (_Artifacts *ArtifactsSession) GetGlobalStakedBuds() (*big.Int, error) {
	return _Artifacts.Contract.GetGlobalStakedBuds(&_Artifacts.CallOpts)
}

// GetGlobalStakedBuds is a free data retrieval call binding the contract method 0x0903b335.
//
// Solidity: function getGlobalStakedBuds() view returns(uint256)
func (_Artifacts *ArtifactsCallerSession) GetGlobalStakedBuds() (*big.Int, error) {
	return _Artifacts.Contract.GetGlobalStakedBuds(&_Artifacts.CallOpts)
}

// GetNoOfChains is a free data retrieval call binding the contract method 0xd8b84491.
//
// Solidity: function getNoOfChains() view returns(uint256)
func (_Artifacts *ArtifactsCaller) GetNoOfChains(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Artifacts.contract.Call(opts, &out, "getNoOfChains")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetNoOfChains is a free data retrieval call binding the contract method 0xd8b84491.
//
// Solidity: function getNoOfChains() view returns(uint256)
func (_Artifacts *ArtifactsSession) GetNoOfChains() (*big.Int, error) {
	return _Artifacts.Contract.GetNoOfChains(&_Artifacts.CallOpts)
}

// GetNoOfChains is a free data retrieval call binding the contract method 0xd8b84491.
//
// Solidity: function getNoOfChains() view returns(uint256)
func (_Artifacts *ArtifactsCallerSession) GetNoOfChains() (*big.Int, error) {
	return _Artifacts.Contract.GetNoOfChains(&_Artifacts.CallOpts)
}

// GetNumberOfStakers is a free data retrieval call binding the contract method 0x658b6729.
//
// Solidity: function getNumberOfStakers() view returns(uint256)
func (_Artifacts *ArtifactsCaller) GetNumberOfStakers(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Artifacts.contract.Call(opts, &out, "getNumberOfStakers")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetNumberOfStakers is a free data retrieval call binding the contract method 0x658b6729.
//
// Solidity: function getNumberOfStakers() view returns(uint256)
func (_Artifacts *ArtifactsSession) GetNumberOfStakers() (*big.Int, error) {
	return _Artifacts.Contract.GetNumberOfStakers(&_Artifacts.CallOpts)
}

// GetNumberOfStakers is a free data retrieval call binding the contract method 0x658b6729.
//
// Solidity: function getNumberOfStakers() view returns(uint256)
func (_Artifacts *ArtifactsCallerSession) GetNumberOfStakers() (*big.Int, error) {
	return _Artifacts.Contract.GetNumberOfStakers(&_Artifacts.CallOpts)
}

// GetTotalStakedFarmers is a free data retrieval call binding the contract method 0xc9990c2a.
//
// Solidity: function getTotalStakedFarmers() view returns(uint256)
func (_Artifacts *ArtifactsCaller) GetTotalStakedFarmers(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Artifacts.contract.Call(opts, &out, "getTotalStakedFarmers")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetTotalStakedFarmers is a free data retrieval call binding the contract method 0xc9990c2a.
//
// Solidity: function getTotalStakedFarmers() view returns(uint256)
func (_Artifacts *ArtifactsSession) GetTotalStakedFarmers() (*big.Int, error) {
	return _Artifacts.Contract.GetTotalStakedFarmers(&_Artifacts.CallOpts)
}

// GetTotalStakedFarmers is a free data retrieval call binding the contract method 0xc9990c2a.
//
// Solidity: function getTotalStakedFarmers() view returns(uint256)
func (_Artifacts *ArtifactsCallerSession) GetTotalStakedFarmers() (*big.Int, error) {
	return _Artifacts.Contract.GetTotalStakedFarmers(&_Artifacts.CallOpts)
}

// GetUserStakes is a free data retrieval call binding the contract method 0x842e2981.
//
// Solidity: function getUserStakes(address user) view returns((address,uint256,uint256,uint256)[])
func (_Artifacts *ArtifactsCaller) GetUserStakes(opts *bind.CallOpts, user common.Address) ([]LibGlobalVarStateStake, error) {
	var out []interface{}
	err := _Artifacts.contract.Call(opts, &out, "getUserStakes", user)

	if err != nil {
		return *new([]LibGlobalVarStateStake), err
	}

	out0 := *abi.ConvertType(out[0], new([]LibGlobalVarStateStake)).(*[]LibGlobalVarStateStake)

	return out0, err

}

// GetUserStakes is a free data retrieval call binding the contract method 0x842e2981.
//
// Solidity: function getUserStakes(address user) view returns((address,uint256,uint256,uint256)[])
func (_Artifacts *ArtifactsSession) GetUserStakes(user common.Address) ([]LibGlobalVarStateStake, error) {
	return _Artifacts.Contract.GetUserStakes(&_Artifacts.CallOpts, user)
}

// GetUserStakes is a free data retrieval call binding the contract method 0x842e2981.
//
// Solidity: function getUserStakes(address user) view returns((address,uint256,uint256,uint256)[])
func (_Artifacts *ArtifactsCallerSession) GetUserStakes(user common.Address) ([]LibGlobalVarStateStake, error) {
	return _Artifacts.Contract.GetUserStakes(&_Artifacts.CallOpts, user)
}

// GetlocalStakedBuds is a free data retrieval call binding the contract method 0x4269e94c.
//
// Solidity: function getlocalStakedBuds() view returns(uint256)
func (_Artifacts *ArtifactsCaller) GetlocalStakedBuds(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Artifacts.contract.Call(opts, &out, "getlocalStakedBuds")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetlocalStakedBuds is a free data retrieval call binding the contract method 0x4269e94c.
//
// Solidity: function getlocalStakedBuds() view returns(uint256)
func (_Artifacts *ArtifactsSession) GetlocalStakedBuds() (*big.Int, error) {
	return _Artifacts.Contract.GetlocalStakedBuds(&_Artifacts.CallOpts)
}

// GetlocalStakedBuds is a free data retrieval call binding the contract method 0x4269e94c.
//
// Solidity: function getlocalStakedBuds() view returns(uint256)
func (_Artifacts *ArtifactsCallerSession) GetlocalStakedBuds() (*big.Int, error) {
	return _Artifacts.Contract.GetlocalStakedBuds(&_Artifacts.CallOpts)
}

// CloseContest is a paid mutator transaction binding the contract method 0xa094c632.
//
// Solidity: function closeContest() returns()
func (_Artifacts *ArtifactsTransactor) CloseContest(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Artifacts.contract.Transact(opts, "closeContest")
}

// CloseContest is a paid mutator transaction binding the contract method 0xa094c632.
//
// Solidity: function closeContest() returns()
func (_Artifacts *ArtifactsSession) CloseContest() (*types.Transaction, error) {
	return _Artifacts.Contract.CloseContest(&_Artifacts.TransactOpts)
}

// CloseContest is a paid mutator transaction binding the contract method 0xa094c632.
//
// Solidity: function closeContest() returns()
func (_Artifacts *ArtifactsTransactorSession) CloseContest() (*types.Transaction, error) {
	return _Artifacts.Contract.CloseContest(&_Artifacts.TransactOpts)
}

// SetGlobalStakedBuds is a paid mutator transaction binding the contract method 0x2d6fe6e1.
//
// Solidity: function setGlobalStakedBuds(uint256 liquidity) returns()
func (_Artifacts *ArtifactsTransactor) SetGlobalStakedBuds(opts *bind.TransactOpts, liquidity *big.Int) (*types.Transaction, error) {
	return _Artifacts.contract.Transact(opts, "setGlobalStakedBuds", liquidity)
}

// SetGlobalStakedBuds is a paid mutator transaction binding the contract method 0x2d6fe6e1.
//
// Solidity: function setGlobalStakedBuds(uint256 liquidity) returns()
func (_Artifacts *ArtifactsSession) SetGlobalStakedBuds(liquidity *big.Int) (*types.Transaction, error) {
	return _Artifacts.Contract.SetGlobalStakedBuds(&_Artifacts.TransactOpts, liquidity)
}

// SetGlobalStakedBuds is a paid mutator transaction binding the contract method 0x2d6fe6e1.
//
// Solidity: function setGlobalStakedBuds(uint256 liquidity) returns()
func (_Artifacts *ArtifactsTransactorSession) SetGlobalStakedBuds(liquidity *big.Int) (*types.Transaction, error) {
	return _Artifacts.Contract.SetGlobalStakedBuds(&_Artifacts.TransactOpts, liquidity)
}

// SetNoOfChains is a paid mutator transaction binding the contract method 0x25970f6d.
//
// Solidity: function setNoOfChains(uint256 chains) returns()
func (_Artifacts *ArtifactsTransactor) SetNoOfChains(opts *bind.TransactOpts, chains *big.Int) (*types.Transaction, error) {
	return _Artifacts.contract.Transact(opts, "setNoOfChains", chains)
}

// SetNoOfChains is a paid mutator transaction binding the contract method 0x25970f6d.
//
// Solidity: function setNoOfChains(uint256 chains) returns()
func (_Artifacts *ArtifactsSession) SetNoOfChains(chains *big.Int) (*types.Transaction, error) {
	return _Artifacts.Contract.SetNoOfChains(&_Artifacts.TransactOpts, chains)
}

// SetNoOfChains is a paid mutator transaction binding the contract method 0x25970f6d.
//
// Solidity: function setNoOfChains(uint256 chains) returns()
func (_Artifacts *ArtifactsTransactorSession) SetNoOfChains(chains *big.Int) (*types.Transaction, error) {
	return _Artifacts.Contract.SetNoOfChains(&_Artifacts.TransactOpts, chains)
}

// SetRaidFees is a paid mutator transaction binding the contract method 0xadb73624.
//
// Solidity: function setRaidFees(uint256 _raidFees) returns()
func (_Artifacts *ArtifactsTransactor) SetRaidFees(opts *bind.TransactOpts, _raidFees *big.Int) (*types.Transaction, error) {
	return _Artifacts.contract.Transact(opts, "setRaidFees", _raidFees)
}

// SetRaidFees is a paid mutator transaction binding the contract method 0xadb73624.
//
// Solidity: function setRaidFees(uint256 _raidFees) returns()
func (_Artifacts *ArtifactsSession) SetRaidFees(_raidFees *big.Int) (*types.Transaction, error) {
	return _Artifacts.Contract.SetRaidFees(&_Artifacts.TransactOpts, _raidFees)
}

// SetRaidFees is a paid mutator transaction binding the contract method 0xadb73624.
//
// Solidity: function setRaidFees(uint256 _raidFees) returns()
func (_Artifacts *ArtifactsTransactorSession) SetRaidFees(_raidFees *big.Int) (*types.Transaction, error) {
	return _Artifacts.Contract.SetRaidFees(&_Artifacts.TransactOpts, _raidFees)
}

// SetRaidHandler is a paid mutator transaction binding the contract method 0xe2197110.
//
// Solidity: function setRaidHandler(address _address) returns()
func (_Artifacts *ArtifactsTransactor) SetRaidHandler(opts *bind.TransactOpts, _address common.Address) (*types.Transaction, error) {
	return _Artifacts.contract.Transact(opts, "setRaidHandler", _address)
}

// SetRaidHandler is a paid mutator transaction binding the contract method 0xe2197110.
//
// Solidity: function setRaidHandler(address _address) returns()
func (_Artifacts *ArtifactsSession) SetRaidHandler(_address common.Address) (*types.Transaction, error) {
	return _Artifacts.Contract.SetRaidHandler(&_Artifacts.TransactOpts, _address)
}

// SetRaidHandler is a paid mutator transaction binding the contract method 0xe2197110.
//
// Solidity: function setRaidHandler(address _address) returns()
func (_Artifacts *ArtifactsTransactorSession) SetRaidHandler(_address common.Address) (*types.Transaction, error) {
	return _Artifacts.Contract.SetRaidHandler(&_Artifacts.TransactOpts, _address)
}

// SetTreasury is a paid mutator transaction binding the contract method 0xf0f44260.
//
// Solidity: function setTreasury(address newAddress) returns()
func (_Artifacts *ArtifactsTransactor) SetTreasury(opts *bind.TransactOpts, newAddress common.Address) (*types.Transaction, error) {
	return _Artifacts.contract.Transact(opts, "setTreasury", newAddress)
}

// SetTreasury is a paid mutator transaction binding the contract method 0xf0f44260.
//
// Solidity: function setTreasury(address newAddress) returns()
func (_Artifacts *ArtifactsSession) SetTreasury(newAddress common.Address) (*types.Transaction, error) {
	return _Artifacts.Contract.SetTreasury(&_Artifacts.TransactOpts, newAddress)
}

// SetTreasury is a paid mutator transaction binding the contract method 0xf0f44260.
//
// Solidity: function setTreasury(address newAddress) returns()
func (_Artifacts *ArtifactsTransactorSession) SetTreasury(newAddress common.Address) (*types.Transaction, error) {
	return _Artifacts.Contract.SetTreasury(&_Artifacts.TransactOpts, newAddress)
}

// StartContest is a paid mutator transaction binding the contract method 0x1765634b.
//
// Solidity: function startContest() returns()
func (_Artifacts *ArtifactsTransactor) StartContest(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Artifacts.contract.Transact(opts, "startContest")
}

// StartContest is a paid mutator transaction binding the contract method 0x1765634b.
//
// Solidity: function startContest() returns()
func (_Artifacts *ArtifactsSession) StartContest() (*types.Transaction, error) {
	return _Artifacts.Contract.StartContest(&_Artifacts.TransactOpts)
}

// StartContest is a paid mutator transaction binding the contract method 0x1765634b.
//
// Solidity: function startContest() returns()
func (_Artifacts *ArtifactsTransactorSession) StartContest() (*types.Transaction, error) {
	return _Artifacts.Contract.StartContest(&_Artifacts.TransactOpts)
}
