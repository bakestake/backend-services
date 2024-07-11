// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package artifacts

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

// ArtifactsMetaData contains all meta data concerning the Artifacts contract.
var ArtifactsMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"name\":\"ContestNotOpen\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"FarmerStakedAlready\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InsufficientFees\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InsufficientRaidFees\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InsufficientStake\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidData\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidForeignChainID\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidRiskLevel\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"MaxBoostReached\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NoStakeFound\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NotANarc\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NotOwnerOfAsset\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"UnexpectedResultLength\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"UnexpectedResultMismatch\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ZeroAddress\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"raider\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"isSuccess\",\"type\":\"bool\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"isBoosted\",\"type\":\"bool\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"rewardTaken\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"boostsUsedInLastSevenDays\",\"type\":\"uint256\"}],\"name\":\"Raided\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"timeStamp\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"rewardsDisbursed\",\"type\":\"uint256\"}],\"name\":\"RewardsCalculated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"budsAmount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"timeStamp\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"localStakedBudsCount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"latestAPR\",\"type\":\"uint256\"}],\"name\":\"Staked\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"budsAmount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"timeStamp\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"localStakedBudsCount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"latestAPR\",\"type\":\"uint256\"}],\"name\":\"UnStaked\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"messageId\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"reason\",\"type\":\"bytes\"}],\"name\":\"crossChainStakeFailed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"messageId\",\"type\":\"bytes32\"}],\"name\":\"recoveredFailedStake\",\"type\":\"event\"}]",
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

// ArtifactsRaidedIterator is returned from FilterRaided and is used to iterate over the raw logs and unpacked data for Raided events raised by the Artifacts contract.
type ArtifactsRaidedIterator struct {
	Event *ArtifactsRaided // Event containing the contract specifics and raw log

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
func (it *ArtifactsRaidedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ArtifactsRaided)
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
		it.Event = new(ArtifactsRaided)
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
func (it *ArtifactsRaidedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ArtifactsRaidedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ArtifactsRaided represents a Raided event raised by the Artifacts contract.
type ArtifactsRaided struct {
	Raider                    common.Address
	IsSuccess                 bool
	IsBoosted                 bool
	RewardTaken               *big.Int
	BoostsUsedInLastSevenDays *big.Int
	Raw                       types.Log // Blockchain specific contextual infos
}

// FilterRaided is a free log retrieval operation binding the contract event 0x7479100003832248ae43d829d4f7831a350510800f7475820666db5df32dc75c.
//
// Solidity: event Raided(address indexed raider, bool isSuccess, bool isBoosted, uint256 rewardTaken, uint256 boostsUsedInLastSevenDays)
func (_Artifacts *ArtifactsFilterer) FilterRaided(opts *bind.FilterOpts, raider []common.Address) (*ArtifactsRaidedIterator, error) {

	var raiderRule []interface{}
	for _, raiderItem := range raider {
		raiderRule = append(raiderRule, raiderItem)
	}

	logs, sub, err := _Artifacts.contract.FilterLogs(opts, "Raided", raiderRule)
	if err != nil {
		return nil, err
	}
	return &ArtifactsRaidedIterator{contract: _Artifacts.contract, event: "Raided", logs: logs, sub: sub}, nil
}

// WatchRaided is a free log subscription operation binding the contract event 0x7479100003832248ae43d829d4f7831a350510800f7475820666db5df32dc75c.
//
// Solidity: event Raided(address indexed raider, bool isSuccess, bool isBoosted, uint256 rewardTaken, uint256 boostsUsedInLastSevenDays)
func (_Artifacts *ArtifactsFilterer) WatchRaided(opts *bind.WatchOpts, sink chan<- *ArtifactsRaided, raider []common.Address) (event.Subscription, error) {

	var raiderRule []interface{}
	for _, raiderItem := range raider {
		raiderRule = append(raiderRule, raiderItem)
	}

	logs, sub, err := _Artifacts.contract.WatchLogs(opts, "Raided", raiderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ArtifactsRaided)
				if err := _Artifacts.contract.UnpackLog(event, "Raided", log); err != nil {
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

// ParseRaided is a log parse operation binding the contract event 0x7479100003832248ae43d829d4f7831a350510800f7475820666db5df32dc75c.
//
// Solidity: event Raided(address indexed raider, bool isSuccess, bool isBoosted, uint256 rewardTaken, uint256 boostsUsedInLastSevenDays)
func (_Artifacts *ArtifactsFilterer) ParseRaided(log types.Log) (*ArtifactsRaided, error) {
	event := new(ArtifactsRaided)
	if err := _Artifacts.contract.UnpackLog(event, "Raided", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ArtifactsRewardsCalculatedIterator is returned from FilterRewardsCalculated and is used to iterate over the raw logs and unpacked data for RewardsCalculated events raised by the Artifacts contract.
type ArtifactsRewardsCalculatedIterator struct {
	Event *ArtifactsRewardsCalculated // Event containing the contract specifics and raw log

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
func (it *ArtifactsRewardsCalculatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ArtifactsRewardsCalculated)
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
		it.Event = new(ArtifactsRewardsCalculated)
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
func (it *ArtifactsRewardsCalculatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ArtifactsRewardsCalculatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ArtifactsRewardsCalculated represents a RewardsCalculated event raised by the Artifacts contract.
type ArtifactsRewardsCalculated struct {
	TimeStamp        *big.Int
	RewardsDisbursed *big.Int
	Raw              types.Log // Blockchain specific contextual infos
}

// FilterRewardsCalculated is a free log retrieval operation binding the contract event 0x75b8bc4c9972548a1b952e44da1064fbdf4ea4b8ac8d9c27feebbb0946deb947.
//
// Solidity: event RewardsCalculated(uint256 timeStamp, uint256 rewardsDisbursed)
func (_Artifacts *ArtifactsFilterer) FilterRewardsCalculated(opts *bind.FilterOpts) (*ArtifactsRewardsCalculatedIterator, error) {

	logs, sub, err := _Artifacts.contract.FilterLogs(opts, "RewardsCalculated")
	if err != nil {
		return nil, err
	}
	return &ArtifactsRewardsCalculatedIterator{contract: _Artifacts.contract, event: "RewardsCalculated", logs: logs, sub: sub}, nil
}

// WatchRewardsCalculated is a free log subscription operation binding the contract event 0x75b8bc4c9972548a1b952e44da1064fbdf4ea4b8ac8d9c27feebbb0946deb947.
//
// Solidity: event RewardsCalculated(uint256 timeStamp, uint256 rewardsDisbursed)
func (_Artifacts *ArtifactsFilterer) WatchRewardsCalculated(opts *bind.WatchOpts, sink chan<- *ArtifactsRewardsCalculated) (event.Subscription, error) {

	logs, sub, err := _Artifacts.contract.WatchLogs(opts, "RewardsCalculated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ArtifactsRewardsCalculated)
				if err := _Artifacts.contract.UnpackLog(event, "RewardsCalculated", log); err != nil {
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

// ParseRewardsCalculated is a log parse operation binding the contract event 0x75b8bc4c9972548a1b952e44da1064fbdf4ea4b8ac8d9c27feebbb0946deb947.
//
// Solidity: event RewardsCalculated(uint256 timeStamp, uint256 rewardsDisbursed)
func (_Artifacts *ArtifactsFilterer) ParseRewardsCalculated(log types.Log) (*ArtifactsRewardsCalculated, error) {
	event := new(ArtifactsRewardsCalculated)
	if err := _Artifacts.contract.UnpackLog(event, "RewardsCalculated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ArtifactsStakedIterator is returned from FilterStaked and is used to iterate over the raw logs and unpacked data for Staked events raised by the Artifacts contract.
type ArtifactsStakedIterator struct {
	Event *ArtifactsStaked // Event containing the contract specifics and raw log

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
func (it *ArtifactsStakedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ArtifactsStaked)
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
		it.Event = new(ArtifactsStaked)
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
func (it *ArtifactsStakedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ArtifactsStakedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ArtifactsStaked represents a Staked event raised by the Artifacts contract.
type ArtifactsStaked struct {
	Owner                common.Address
	TokenId              *big.Int
	BudsAmount           *big.Int
	TimeStamp            *big.Int
	LocalStakedBudsCount *big.Int
	LatestAPR            *big.Int
	Raw                  types.Log // Blockchain specific contextual infos
}

// FilterStaked is a free log retrieval operation binding the contract event 0x6381ea17a5324d29cc015352644672ead5185c1c61a0d3a521eda97e35cec97e.
//
// Solidity: event Staked(address indexed owner, uint256 tokenId, uint256 budsAmount, uint256 timeStamp, uint256 localStakedBudsCount, uint256 latestAPR)
func (_Artifacts *ArtifactsFilterer) FilterStaked(opts *bind.FilterOpts, owner []common.Address) (*ArtifactsStakedIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}

	logs, sub, err := _Artifacts.contract.FilterLogs(opts, "Staked", ownerRule)
	if err != nil {
		return nil, err
	}
	return &ArtifactsStakedIterator{contract: _Artifacts.contract, event: "Staked", logs: logs, sub: sub}, nil
}

// WatchStaked is a free log subscription operation binding the contract event 0x6381ea17a5324d29cc015352644672ead5185c1c61a0d3a521eda97e35cec97e.
//
// Solidity: event Staked(address indexed owner, uint256 tokenId, uint256 budsAmount, uint256 timeStamp, uint256 localStakedBudsCount, uint256 latestAPR)
func (_Artifacts *ArtifactsFilterer) WatchStaked(opts *bind.WatchOpts, sink chan<- *ArtifactsStaked, owner []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}

	logs, sub, err := _Artifacts.contract.WatchLogs(opts, "Staked", ownerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ArtifactsStaked)
				if err := _Artifacts.contract.UnpackLog(event, "Staked", log); err != nil {
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

// ParseStaked is a log parse operation binding the contract event 0x6381ea17a5324d29cc015352644672ead5185c1c61a0d3a521eda97e35cec97e.
//
// Solidity: event Staked(address indexed owner, uint256 tokenId, uint256 budsAmount, uint256 timeStamp, uint256 localStakedBudsCount, uint256 latestAPR)
func (_Artifacts *ArtifactsFilterer) ParseStaked(log types.Log) (*ArtifactsStaked, error) {
	event := new(ArtifactsStaked)
	if err := _Artifacts.contract.UnpackLog(event, "Staked", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ArtifactsUnStakedIterator is returned from FilterUnStaked and is used to iterate over the raw logs and unpacked data for UnStaked events raised by the Artifacts contract.
type ArtifactsUnStakedIterator struct {
	Event *ArtifactsUnStaked // Event containing the contract specifics and raw log

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
func (it *ArtifactsUnStakedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ArtifactsUnStaked)
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
		it.Event = new(ArtifactsUnStaked)
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
func (it *ArtifactsUnStakedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ArtifactsUnStakedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ArtifactsUnStaked represents a UnStaked event raised by the Artifacts contract.
type ArtifactsUnStaked struct {
	Owner                common.Address
	TokenId              *big.Int
	BudsAmount           *big.Int
	TimeStamp            *big.Int
	LocalStakedBudsCount *big.Int
	LatestAPR            *big.Int
	Raw                  types.Log // Blockchain specific contextual infos
}

// FilterUnStaked is a free log retrieval operation binding the contract event 0xa4e1548e00706ebb653db94280851ebbad17444a4786dd50a9540a9c7eb7eeb5.
//
// Solidity: event UnStaked(address owner, uint256 tokenId, uint256 budsAmount, uint256 timeStamp, uint256 localStakedBudsCount, uint256 latestAPR)
func (_Artifacts *ArtifactsFilterer) FilterUnStaked(opts *bind.FilterOpts) (*ArtifactsUnStakedIterator, error) {

	logs, sub, err := _Artifacts.contract.FilterLogs(opts, "UnStaked")
	if err != nil {
		return nil, err
	}
	return &ArtifactsUnStakedIterator{contract: _Artifacts.contract, event: "UnStaked", logs: logs, sub: sub}, nil
}

// WatchUnStaked is a free log subscription operation binding the contract event 0xa4e1548e00706ebb653db94280851ebbad17444a4786dd50a9540a9c7eb7eeb5.
//
// Solidity: event UnStaked(address owner, uint256 tokenId, uint256 budsAmount, uint256 timeStamp, uint256 localStakedBudsCount, uint256 latestAPR)
func (_Artifacts *ArtifactsFilterer) WatchUnStaked(opts *bind.WatchOpts, sink chan<- *ArtifactsUnStaked) (event.Subscription, error) {

	logs, sub, err := _Artifacts.contract.WatchLogs(opts, "UnStaked")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ArtifactsUnStaked)
				if err := _Artifacts.contract.UnpackLog(event, "UnStaked", log); err != nil {
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

// ParseUnStaked is a log parse operation binding the contract event 0xa4e1548e00706ebb653db94280851ebbad17444a4786dd50a9540a9c7eb7eeb5.
//
// Solidity: event UnStaked(address owner, uint256 tokenId, uint256 budsAmount, uint256 timeStamp, uint256 localStakedBudsCount, uint256 latestAPR)
func (_Artifacts *ArtifactsFilterer) ParseUnStaked(log types.Log) (*ArtifactsUnStaked, error) {
	event := new(ArtifactsUnStaked)
	if err := _Artifacts.contract.UnpackLog(event, "UnStaked", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ArtifactsCrossChainStakeFailedIterator is returned from FilterCrossChainStakeFailed and is used to iterate over the raw logs and unpacked data for CrossChainStakeFailed events raised by the Artifacts contract.
type ArtifactsCrossChainStakeFailedIterator struct {
	Event *ArtifactsCrossChainStakeFailed // Event containing the contract specifics and raw log

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
func (it *ArtifactsCrossChainStakeFailedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ArtifactsCrossChainStakeFailed)
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
		it.Event = new(ArtifactsCrossChainStakeFailed)
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
func (it *ArtifactsCrossChainStakeFailedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ArtifactsCrossChainStakeFailedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ArtifactsCrossChainStakeFailed represents a CrossChainStakeFailed event raised by the Artifacts contract.
type ArtifactsCrossChainStakeFailed struct {
	MessageId [32]byte
	Reason    []byte
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterCrossChainStakeFailed is a free log retrieval operation binding the contract event 0x8fb2be4ad054a2859f094704bd579532f1339f92ca5867e73f25482e6db0c444.
//
// Solidity: event crossChainStakeFailed(bytes32 indexed messageId, bytes reason)
func (_Artifacts *ArtifactsFilterer) FilterCrossChainStakeFailed(opts *bind.FilterOpts, messageId [][32]byte) (*ArtifactsCrossChainStakeFailedIterator, error) {

	var messageIdRule []interface{}
	for _, messageIdItem := range messageId {
		messageIdRule = append(messageIdRule, messageIdItem)
	}

	logs, sub, err := _Artifacts.contract.FilterLogs(opts, "crossChainStakeFailed", messageIdRule)
	if err != nil {
		return nil, err
	}
	return &ArtifactsCrossChainStakeFailedIterator{contract: _Artifacts.contract, event: "crossChainStakeFailed", logs: logs, sub: sub}, nil
}

// WatchCrossChainStakeFailed is a free log subscription operation binding the contract event 0x8fb2be4ad054a2859f094704bd579532f1339f92ca5867e73f25482e6db0c444.
//
// Solidity: event crossChainStakeFailed(bytes32 indexed messageId, bytes reason)
func (_Artifacts *ArtifactsFilterer) WatchCrossChainStakeFailed(opts *bind.WatchOpts, sink chan<- *ArtifactsCrossChainStakeFailed, messageId [][32]byte) (event.Subscription, error) {

	var messageIdRule []interface{}
	for _, messageIdItem := range messageId {
		messageIdRule = append(messageIdRule, messageIdItem)
	}

	logs, sub, err := _Artifacts.contract.WatchLogs(opts, "crossChainStakeFailed", messageIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ArtifactsCrossChainStakeFailed)
				if err := _Artifacts.contract.UnpackLog(event, "crossChainStakeFailed", log); err != nil {
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

// ParseCrossChainStakeFailed is a log parse operation binding the contract event 0x8fb2be4ad054a2859f094704bd579532f1339f92ca5867e73f25482e6db0c444.
//
// Solidity: event crossChainStakeFailed(bytes32 indexed messageId, bytes reason)
func (_Artifacts *ArtifactsFilterer) ParseCrossChainStakeFailed(log types.Log) (*ArtifactsCrossChainStakeFailed, error) {
	event := new(ArtifactsCrossChainStakeFailed)
	if err := _Artifacts.contract.UnpackLog(event, "crossChainStakeFailed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ArtifactsRecoveredFailedStakeIterator is returned from FilterRecoveredFailedStake and is used to iterate over the raw logs and unpacked data for RecoveredFailedStake events raised by the Artifacts contract.
type ArtifactsRecoveredFailedStakeIterator struct {
	Event *ArtifactsRecoveredFailedStake // Event containing the contract specifics and raw log

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
func (it *ArtifactsRecoveredFailedStakeIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ArtifactsRecoveredFailedStake)
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
		it.Event = new(ArtifactsRecoveredFailedStake)
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
func (it *ArtifactsRecoveredFailedStakeIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ArtifactsRecoveredFailedStakeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ArtifactsRecoveredFailedStake represents a RecoveredFailedStake event raised by the Artifacts contract.
type ArtifactsRecoveredFailedStake struct {
	MessageId [32]byte
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterRecoveredFailedStake is a free log retrieval operation binding the contract event 0x26bb584128de32720d367ea13750ac19db93e0fff3e7be0e6ff0d3c32bfac94d.
//
// Solidity: event recoveredFailedStake(bytes32 indexed messageId)
func (_Artifacts *ArtifactsFilterer) FilterRecoveredFailedStake(opts *bind.FilterOpts, messageId [][32]byte) (*ArtifactsRecoveredFailedStakeIterator, error) {

	var messageIdRule []interface{}
	for _, messageIdItem := range messageId {
		messageIdRule = append(messageIdRule, messageIdItem)
	}

	logs, sub, err := _Artifacts.contract.FilterLogs(opts, "recoveredFailedStake", messageIdRule)
	if err != nil {
		return nil, err
	}
	return &ArtifactsRecoveredFailedStakeIterator{contract: _Artifacts.contract, event: "recoveredFailedStake", logs: logs, sub: sub}, nil
}

// WatchRecoveredFailedStake is a free log subscription operation binding the contract event 0x26bb584128de32720d367ea13750ac19db93e0fff3e7be0e6ff0d3c32bfac94d.
//
// Solidity: event recoveredFailedStake(bytes32 indexed messageId)
func (_Artifacts *ArtifactsFilterer) WatchRecoveredFailedStake(opts *bind.WatchOpts, sink chan<- *ArtifactsRecoveredFailedStake, messageId [][32]byte) (event.Subscription, error) {

	var messageIdRule []interface{}
	for _, messageIdItem := range messageId {
		messageIdRule = append(messageIdRule, messageIdItem)
	}

	logs, sub, err := _Artifacts.contract.WatchLogs(opts, "recoveredFailedStake", messageIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ArtifactsRecoveredFailedStake)
				if err := _Artifacts.contract.UnpackLog(event, "recoveredFailedStake", log); err != nil {
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

// ParseRecoveredFailedStake is a log parse operation binding the contract event 0x26bb584128de32720d367ea13750ac19db93e0fff3e7be0e6ff0d3c32bfac94d.
//
// Solidity: event recoveredFailedStake(bytes32 indexed messageId)
func (_Artifacts *ArtifactsFilterer) ParseRecoveredFailedStake(log types.Log) (*ArtifactsRecoveredFailedStake, error) {
	event := new(ArtifactsRecoveredFailedStake)
	if err := _Artifacts.contract.UnpackLog(event, "recoveredFailedStake", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
