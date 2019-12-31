// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package smartcontract

import (
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
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = abi.U256
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
)

// SmartcontractABI is the input ABI used to generate the binding from.
const SmartcontractABI = "[{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_eventStartTime\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_eventEndTime\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_price\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"subscriptor\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"price\",\"type\":\"uint256\"}],\"name\":\"NewSubscription\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"subscriptor\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"price\",\"type\":\"uint256\"}],\"name\":\"NewUnsubscription\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[],\"name\":\"PPVEventEnded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[],\"name\":\"PPVEventStarted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"ppvEventPrice\",\"type\":\"uint256\"}],\"name\":\"PriceChanged\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_price\",\"type\":\"uint256\"}],\"name\":\"changePrice\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"ended\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"addresspayable\",\"name\":\"dest\",\"type\":\"address\"}],\"name\":\"eventEnd\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"ppvEventEnd\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"ppvEventPrice\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"ppvEventStart\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"subscribe\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"unsubscribe\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// Smartcontract is an auto generated Go binding around an Ethereum contract.
type Smartcontract struct {
	SmartcontractCaller     // Read-only binding to the contract
	SmartcontractTransactor // Write-only binding to the contract
	SmartcontractFilterer   // Log filterer for contract events
}

// SmartcontractCaller is an auto generated read-only Go binding around an Ethereum contract.
type SmartcontractCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SmartcontractTransactor is an auto generated write-only Go binding around an Ethereum contract.
type SmartcontractTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SmartcontractFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type SmartcontractFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SmartcontractSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type SmartcontractSession struct {
	Contract     *Smartcontract    // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// SmartcontractCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type SmartcontractCallerSession struct {
	Contract *SmartcontractCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts        // Call options to use throughout this session
}

// SmartcontractTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type SmartcontractTransactorSession struct {
	Contract     *SmartcontractTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts        // Transaction auth options to use throughout this session
}

// SmartcontractRaw is an auto generated low-level Go binding around an Ethereum contract.
type SmartcontractRaw struct {
	Contract *Smartcontract // Generic contract binding to access the raw methods on
}

// SmartcontractCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type SmartcontractCallerRaw struct {
	Contract *SmartcontractCaller // Generic read-only contract binding to access the raw methods on
}

// SmartcontractTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type SmartcontractTransactorRaw struct {
	Contract *SmartcontractTransactor // Generic write-only contract binding to access the raw methods on
}

// NewSmartcontract creates a new instance of Smartcontract, bound to a specific deployed contract.
func NewSmartcontract(address common.Address, backend bind.ContractBackend) (*Smartcontract, error) {
	contract, err := bindSmartcontract(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Smartcontract{SmartcontractCaller: SmartcontractCaller{contract: contract}, SmartcontractTransactor: SmartcontractTransactor{contract: contract}, SmartcontractFilterer: SmartcontractFilterer{contract: contract}}, nil
}

// NewSmartcontractCaller creates a new read-only instance of Smartcontract, bound to a specific deployed contract.
func NewSmartcontractCaller(address common.Address, caller bind.ContractCaller) (*SmartcontractCaller, error) {
	contract, err := bindSmartcontract(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &SmartcontractCaller{contract: contract}, nil
}

// NewSmartcontractTransactor creates a new write-only instance of Smartcontract, bound to a specific deployed contract.
func NewSmartcontractTransactor(address common.Address, transactor bind.ContractTransactor) (*SmartcontractTransactor, error) {
	contract, err := bindSmartcontract(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &SmartcontractTransactor{contract: contract}, nil
}

// NewSmartcontractFilterer creates a new log filterer instance of Smartcontract, bound to a specific deployed contract.
func NewSmartcontractFilterer(address common.Address, filterer bind.ContractFilterer) (*SmartcontractFilterer, error) {
	contract, err := bindSmartcontract(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &SmartcontractFilterer{contract: contract}, nil
}

// bindSmartcontract binds a generic wrapper to an already deployed contract.
func bindSmartcontract(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(SmartcontractABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Smartcontract *SmartcontractRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Smartcontract.Contract.SmartcontractCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Smartcontract *SmartcontractRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Smartcontract.Contract.SmartcontractTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Smartcontract *SmartcontractRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Smartcontract.Contract.SmartcontractTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Smartcontract *SmartcontractCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Smartcontract.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Smartcontract *SmartcontractTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Smartcontract.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Smartcontract *SmartcontractTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Smartcontract.Contract.contract.Transact(opts, method, params...)
}

// ChangePrice is a paid mutator transaction binding the contract method 0xa2b40d19.
//
// Solidity: function changePrice(uint256 _price) returns()
func (_Smartcontract *SmartcontractTransactor) ChangePrice(opts *bind.TransactOpts, _price *big.Int) (*types.Transaction, error) {
	return _Smartcontract.contract.Transact(opts, "changePrice", _price)
}

// ChangePrice is a paid mutator transaction binding the contract method 0xa2b40d19.
//
// Solidity: function changePrice(uint256 _price) returns()
func (_Smartcontract *SmartcontractSession) ChangePrice(_price *big.Int) (*types.Transaction, error) {
	return _Smartcontract.Contract.ChangePrice(&_Smartcontract.TransactOpts, _price)
}

// ChangePrice is a paid mutator transaction binding the contract method 0xa2b40d19.
//
// Solidity: function changePrice(uint256 _price) returns()
func (_Smartcontract *SmartcontractTransactorSession) ChangePrice(_price *big.Int) (*types.Transaction, error) {
	return _Smartcontract.Contract.ChangePrice(&_Smartcontract.TransactOpts, _price)
}

// Ended is a paid mutator transaction binding the contract method 0x12fa6feb.
//
// Solidity: function ended() returns(bool)
func (_Smartcontract *SmartcontractTransactor) Ended(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Smartcontract.contract.Transact(opts, "ended")
}

// Ended is a paid mutator transaction binding the contract method 0x12fa6feb.
//
// Solidity: function ended() returns(bool)
func (_Smartcontract *SmartcontractSession) Ended() (*types.Transaction, error) {
	return _Smartcontract.Contract.Ended(&_Smartcontract.TransactOpts)
}

// Ended is a paid mutator transaction binding the contract method 0x12fa6feb.
//
// Solidity: function ended() returns(bool)
func (_Smartcontract *SmartcontractTransactorSession) Ended() (*types.Transaction, error) {
	return _Smartcontract.Contract.Ended(&_Smartcontract.TransactOpts)
}

// EventEnd is a paid mutator transaction binding the contract method 0xd1428312.
//
// Solidity: function eventEnd(address dest) returns()
func (_Smartcontract *SmartcontractTransactor) EventEnd(opts *bind.TransactOpts, dest common.Address) (*types.Transaction, error) {
	return _Smartcontract.contract.Transact(opts, "eventEnd", dest)
}

// EventEnd is a paid mutator transaction binding the contract method 0xd1428312.
//
// Solidity: function eventEnd(address dest) returns()
func (_Smartcontract *SmartcontractSession) EventEnd(dest common.Address) (*types.Transaction, error) {
	return _Smartcontract.Contract.EventEnd(&_Smartcontract.TransactOpts, dest)
}

// EventEnd is a paid mutator transaction binding the contract method 0xd1428312.
//
// Solidity: function eventEnd(address dest) returns()
func (_Smartcontract *SmartcontractTransactorSession) EventEnd(dest common.Address) (*types.Transaction, error) {
	return _Smartcontract.Contract.EventEnd(&_Smartcontract.TransactOpts, dest)
}

// PpvEventEnd is a paid mutator transaction binding the contract method 0x2fd3308c.
//
// Solidity: function ppvEventEnd() returns(uint256)
func (_Smartcontract *SmartcontractTransactor) PpvEventEnd(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Smartcontract.contract.Transact(opts, "ppvEventEnd")
}

// PpvEventEnd is a paid mutator transaction binding the contract method 0x2fd3308c.
//
// Solidity: function ppvEventEnd() returns(uint256)
func (_Smartcontract *SmartcontractSession) PpvEventEnd() (*types.Transaction, error) {
	return _Smartcontract.Contract.PpvEventEnd(&_Smartcontract.TransactOpts)
}

// PpvEventEnd is a paid mutator transaction binding the contract method 0x2fd3308c.
//
// Solidity: function ppvEventEnd() returns(uint256)
func (_Smartcontract *SmartcontractTransactorSession) PpvEventEnd() (*types.Transaction, error) {
	return _Smartcontract.Contract.PpvEventEnd(&_Smartcontract.TransactOpts)
}

// PpvEventPrice is a paid mutator transaction binding the contract method 0xad9061e4.
//
// Solidity: function ppvEventPrice() returns(uint256)
func (_Smartcontract *SmartcontractTransactor) PpvEventPrice(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Smartcontract.contract.Transact(opts, "ppvEventPrice")
}

// PpvEventPrice is a paid mutator transaction binding the contract method 0xad9061e4.
//
// Solidity: function ppvEventPrice() returns(uint256)
func (_Smartcontract *SmartcontractSession) PpvEventPrice() (*types.Transaction, error) {
	return _Smartcontract.Contract.PpvEventPrice(&_Smartcontract.TransactOpts)
}

// PpvEventPrice is a paid mutator transaction binding the contract method 0xad9061e4.
//
// Solidity: function ppvEventPrice() returns(uint256)
func (_Smartcontract *SmartcontractTransactorSession) PpvEventPrice() (*types.Transaction, error) {
	return _Smartcontract.Contract.PpvEventPrice(&_Smartcontract.TransactOpts)
}

// PpvEventStart is a paid mutator transaction binding the contract method 0xed34ce31.
//
// Solidity: function ppvEventStart() returns(uint256)
func (_Smartcontract *SmartcontractTransactor) PpvEventStart(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Smartcontract.contract.Transact(opts, "ppvEventStart")
}

// PpvEventStart is a paid mutator transaction binding the contract method 0xed34ce31.
//
// Solidity: function ppvEventStart() returns(uint256)
func (_Smartcontract *SmartcontractSession) PpvEventStart() (*types.Transaction, error) {
	return _Smartcontract.Contract.PpvEventStart(&_Smartcontract.TransactOpts)
}

// PpvEventStart is a paid mutator transaction binding the contract method 0xed34ce31.
//
// Solidity: function ppvEventStart() returns(uint256)
func (_Smartcontract *SmartcontractTransactorSession) PpvEventStart() (*types.Transaction, error) {
	return _Smartcontract.Contract.PpvEventStart(&_Smartcontract.TransactOpts)
}

// Subscribe is a paid mutator transaction binding the contract method 0x8f449a05.
//
// Solidity: function subscribe() returns()
func (_Smartcontract *SmartcontractTransactor) Subscribe(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Smartcontract.contract.Transact(opts, "subscribe")
}

// Subscribe is a paid mutator transaction binding the contract method 0x8f449a05.
//
// Solidity: function subscribe() returns()
func (_Smartcontract *SmartcontractSession) Subscribe() (*types.Transaction, error) {
	return _Smartcontract.Contract.Subscribe(&_Smartcontract.TransactOpts)
}

// Subscribe is a paid mutator transaction binding the contract method 0x8f449a05.
//
// Solidity: function subscribe() returns()
func (_Smartcontract *SmartcontractTransactorSession) Subscribe() (*types.Transaction, error) {
	return _Smartcontract.Contract.Subscribe(&_Smartcontract.TransactOpts)
}

// Unsubscribe is a paid mutator transaction binding the contract method 0xfcae4484.
//
// Solidity: function unsubscribe() returns()
func (_Smartcontract *SmartcontractTransactor) Unsubscribe(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Smartcontract.contract.Transact(opts, "unsubscribe")
}

// Unsubscribe is a paid mutator transaction binding the contract method 0xfcae4484.
//
// Solidity: function unsubscribe() returns()
func (_Smartcontract *SmartcontractSession) Unsubscribe() (*types.Transaction, error) {
	return _Smartcontract.Contract.Unsubscribe(&_Smartcontract.TransactOpts)
}

// Unsubscribe is a paid mutator transaction binding the contract method 0xfcae4484.
//
// Solidity: function unsubscribe() returns()
func (_Smartcontract *SmartcontractTransactorSession) Unsubscribe() (*types.Transaction, error) {
	return _Smartcontract.Contract.Unsubscribe(&_Smartcontract.TransactOpts)
}

// SmartcontractNewSubscriptionIterator is returned from FilterNewSubscription and is used to iterate over the raw logs and unpacked data for NewSubscription events raised by the Smartcontract contract.
type SmartcontractNewSubscriptionIterator struct {
	Event *SmartcontractNewSubscription // Event containing the contract specifics and raw log

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
func (it *SmartcontractNewSubscriptionIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SmartcontractNewSubscription)
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
		it.Event = new(SmartcontractNewSubscription)
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
func (it *SmartcontractNewSubscriptionIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SmartcontractNewSubscriptionIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SmartcontractNewSubscription represents a NewSubscription event raised by the Smartcontract contract.
type SmartcontractNewSubscription struct {
	Subscriptor common.Address
	Price       *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterNewSubscription is a free log retrieval operation binding the contract event 0x1e05df24f73db39faa0c2d5d26727d08632debce09833123a69214ba943e07c2.
//
// Solidity: event NewSubscription(address subscriptor, uint256 price)
func (_Smartcontract *SmartcontractFilterer) FilterNewSubscription(opts *bind.FilterOpts) (*SmartcontractNewSubscriptionIterator, error) {

	logs, sub, err := _Smartcontract.contract.FilterLogs(opts, "NewSubscription")
	if err != nil {
		return nil, err
	}
	return &SmartcontractNewSubscriptionIterator{contract: _Smartcontract.contract, event: "NewSubscription", logs: logs, sub: sub}, nil
}

// WatchNewSubscription is a free log subscription operation binding the contract event 0x1e05df24f73db39faa0c2d5d26727d08632debce09833123a69214ba943e07c2.
//
// Solidity: event NewSubscription(address subscriptor, uint256 price)
func (_Smartcontract *SmartcontractFilterer) WatchNewSubscription(opts *bind.WatchOpts, sink chan<- *SmartcontractNewSubscription) (event.Subscription, error) {

	logs, sub, err := _Smartcontract.contract.WatchLogs(opts, "NewSubscription")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SmartcontractNewSubscription)
				if err := _Smartcontract.contract.UnpackLog(event, "NewSubscription", log); err != nil {
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

// ParseNewSubscription is a log parse operation binding the contract event 0x1e05df24f73db39faa0c2d5d26727d08632debce09833123a69214ba943e07c2.
//
// Solidity: event NewSubscription(address subscriptor, uint256 price)
func (_Smartcontract *SmartcontractFilterer) ParseNewSubscription(log types.Log) (*SmartcontractNewSubscription, error) {
	event := new(SmartcontractNewSubscription)
	if err := _Smartcontract.contract.UnpackLog(event, "NewSubscription", log); err != nil {
		return nil, err
	}
	return event, nil
}

// SmartcontractNewUnsubscriptionIterator is returned from FilterNewUnsubscription and is used to iterate over the raw logs and unpacked data for NewUnsubscription events raised by the Smartcontract contract.
type SmartcontractNewUnsubscriptionIterator struct {
	Event *SmartcontractNewUnsubscription // Event containing the contract specifics and raw log

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
func (it *SmartcontractNewUnsubscriptionIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SmartcontractNewUnsubscription)
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
		it.Event = new(SmartcontractNewUnsubscription)
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
func (it *SmartcontractNewUnsubscriptionIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SmartcontractNewUnsubscriptionIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SmartcontractNewUnsubscription represents a NewUnsubscription event raised by the Smartcontract contract.
type SmartcontractNewUnsubscription struct {
	Subscriptor common.Address
	Price       *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterNewUnsubscription is a free log retrieval operation binding the contract event 0x79d9d14aed97c97b8dc663528ec3b03eecb8d467956fc03f37179b188a04efa4.
//
// Solidity: event NewUnsubscription(address subscriptor, uint256 price)
func (_Smartcontract *SmartcontractFilterer) FilterNewUnsubscription(opts *bind.FilterOpts) (*SmartcontractNewUnsubscriptionIterator, error) {

	logs, sub, err := _Smartcontract.contract.FilterLogs(opts, "NewUnsubscription")
	if err != nil {
		return nil, err
	}
	return &SmartcontractNewUnsubscriptionIterator{contract: _Smartcontract.contract, event: "NewUnsubscription", logs: logs, sub: sub}, nil
}

// WatchNewUnsubscription is a free log subscription operation binding the contract event 0x79d9d14aed97c97b8dc663528ec3b03eecb8d467956fc03f37179b188a04efa4.
//
// Solidity: event NewUnsubscription(address subscriptor, uint256 price)
func (_Smartcontract *SmartcontractFilterer) WatchNewUnsubscription(opts *bind.WatchOpts, sink chan<- *SmartcontractNewUnsubscription) (event.Subscription, error) {

	logs, sub, err := _Smartcontract.contract.WatchLogs(opts, "NewUnsubscription")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SmartcontractNewUnsubscription)
				if err := _Smartcontract.contract.UnpackLog(event, "NewUnsubscription", log); err != nil {
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

// ParseNewUnsubscription is a log parse operation binding the contract event 0x79d9d14aed97c97b8dc663528ec3b03eecb8d467956fc03f37179b188a04efa4.
//
// Solidity: event NewUnsubscription(address subscriptor, uint256 price)
func (_Smartcontract *SmartcontractFilterer) ParseNewUnsubscription(log types.Log) (*SmartcontractNewUnsubscription, error) {
	event := new(SmartcontractNewUnsubscription)
	if err := _Smartcontract.contract.UnpackLog(event, "NewUnsubscription", log); err != nil {
		return nil, err
	}
	return event, nil
}

// SmartcontractPPVEventEndedIterator is returned from FilterPPVEventEnded and is used to iterate over the raw logs and unpacked data for PPVEventEnded events raised by the Smartcontract contract.
type SmartcontractPPVEventEndedIterator struct {
	Event *SmartcontractPPVEventEnded // Event containing the contract specifics and raw log

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
func (it *SmartcontractPPVEventEndedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SmartcontractPPVEventEnded)
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
		it.Event = new(SmartcontractPPVEventEnded)
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
func (it *SmartcontractPPVEventEndedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SmartcontractPPVEventEndedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SmartcontractPPVEventEnded represents a PPVEventEnded event raised by the Smartcontract contract.
type SmartcontractPPVEventEnded struct {
	Raw types.Log // Blockchain specific contextual infos
}

// FilterPPVEventEnded is a free log retrieval operation binding the contract event 0x545c09a6805d0d5edc24946e122843a3b3cfb83cef050826f143983353509e2f.
//
// Solidity: event PPVEventEnded()
func (_Smartcontract *SmartcontractFilterer) FilterPPVEventEnded(opts *bind.FilterOpts) (*SmartcontractPPVEventEndedIterator, error) {

	logs, sub, err := _Smartcontract.contract.FilterLogs(opts, "PPVEventEnded")
	if err != nil {
		return nil, err
	}
	return &SmartcontractPPVEventEndedIterator{contract: _Smartcontract.contract, event: "PPVEventEnded", logs: logs, sub: sub}, nil
}

// WatchPPVEventEnded is a free log subscription operation binding the contract event 0x545c09a6805d0d5edc24946e122843a3b3cfb83cef050826f143983353509e2f.
//
// Solidity: event PPVEventEnded()
func (_Smartcontract *SmartcontractFilterer) WatchPPVEventEnded(opts *bind.WatchOpts, sink chan<- *SmartcontractPPVEventEnded) (event.Subscription, error) {

	logs, sub, err := _Smartcontract.contract.WatchLogs(opts, "PPVEventEnded")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SmartcontractPPVEventEnded)
				if err := _Smartcontract.contract.UnpackLog(event, "PPVEventEnded", log); err != nil {
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

// ParsePPVEventEnded is a log parse operation binding the contract event 0x545c09a6805d0d5edc24946e122843a3b3cfb83cef050826f143983353509e2f.
//
// Solidity: event PPVEventEnded()
func (_Smartcontract *SmartcontractFilterer) ParsePPVEventEnded(log types.Log) (*SmartcontractPPVEventEnded, error) {
	event := new(SmartcontractPPVEventEnded)
	if err := _Smartcontract.contract.UnpackLog(event, "PPVEventEnded", log); err != nil {
		return nil, err
	}
	return event, nil
}

// SmartcontractPPVEventStartedIterator is returned from FilterPPVEventStarted and is used to iterate over the raw logs and unpacked data for PPVEventStarted events raised by the Smartcontract contract.
type SmartcontractPPVEventStartedIterator struct {
	Event *SmartcontractPPVEventStarted // Event containing the contract specifics and raw log

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
func (it *SmartcontractPPVEventStartedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SmartcontractPPVEventStarted)
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
		it.Event = new(SmartcontractPPVEventStarted)
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
func (it *SmartcontractPPVEventStartedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SmartcontractPPVEventStartedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SmartcontractPPVEventStarted represents a PPVEventStarted event raised by the Smartcontract contract.
type SmartcontractPPVEventStarted struct {
	Raw types.Log // Blockchain specific contextual infos
}

// FilterPPVEventStarted is a free log retrieval operation binding the contract event 0x0e10146934a460c3f27a8ca4c63fdad156eed0ae2051edc0c1839268ade6c9ad.
//
// Solidity: event PPVEventStarted()
func (_Smartcontract *SmartcontractFilterer) FilterPPVEventStarted(opts *bind.FilterOpts) (*SmartcontractPPVEventStartedIterator, error) {

	logs, sub, err := _Smartcontract.contract.FilterLogs(opts, "PPVEventStarted")
	if err != nil {
		return nil, err
	}
	return &SmartcontractPPVEventStartedIterator{contract: _Smartcontract.contract, event: "PPVEventStarted", logs: logs, sub: sub}, nil
}

// WatchPPVEventStarted is a free log subscription operation binding the contract event 0x0e10146934a460c3f27a8ca4c63fdad156eed0ae2051edc0c1839268ade6c9ad.
//
// Solidity: event PPVEventStarted()
func (_Smartcontract *SmartcontractFilterer) WatchPPVEventStarted(opts *bind.WatchOpts, sink chan<- *SmartcontractPPVEventStarted) (event.Subscription, error) {

	logs, sub, err := _Smartcontract.contract.WatchLogs(opts, "PPVEventStarted")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SmartcontractPPVEventStarted)
				if err := _Smartcontract.contract.UnpackLog(event, "PPVEventStarted", log); err != nil {
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

// ParsePPVEventStarted is a log parse operation binding the contract event 0x0e10146934a460c3f27a8ca4c63fdad156eed0ae2051edc0c1839268ade6c9ad.
//
// Solidity: event PPVEventStarted()
func (_Smartcontract *SmartcontractFilterer) ParsePPVEventStarted(log types.Log) (*SmartcontractPPVEventStarted, error) {
	event := new(SmartcontractPPVEventStarted)
	if err := _Smartcontract.contract.UnpackLog(event, "PPVEventStarted", log); err != nil {
		return nil, err
	}
	return event, nil
}

// SmartcontractPriceChangedIterator is returned from FilterPriceChanged and is used to iterate over the raw logs and unpacked data for PriceChanged events raised by the Smartcontract contract.
type SmartcontractPriceChangedIterator struct {
	Event *SmartcontractPriceChanged // Event containing the contract specifics and raw log

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
func (it *SmartcontractPriceChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SmartcontractPriceChanged)
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
		it.Event = new(SmartcontractPriceChanged)
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
func (it *SmartcontractPriceChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SmartcontractPriceChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SmartcontractPriceChanged represents a PriceChanged event raised by the Smartcontract contract.
type SmartcontractPriceChanged struct {
	PpvEventPrice *big.Int
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterPriceChanged is a free log retrieval operation binding the contract event 0xa6dc15bdb68da224c66db4b3838d9a2b205138e8cff6774e57d0af91e196d622.
//
// Solidity: event PriceChanged(uint256 ppvEventPrice)
func (_Smartcontract *SmartcontractFilterer) FilterPriceChanged(opts *bind.FilterOpts) (*SmartcontractPriceChangedIterator, error) {

	logs, sub, err := _Smartcontract.contract.FilterLogs(opts, "PriceChanged")
	if err != nil {
		return nil, err
	}
	return &SmartcontractPriceChangedIterator{contract: _Smartcontract.contract, event: "PriceChanged", logs: logs, sub: sub}, nil
}

// WatchPriceChanged is a free log subscription operation binding the contract event 0xa6dc15bdb68da224c66db4b3838d9a2b205138e8cff6774e57d0af91e196d622.
//
// Solidity: event PriceChanged(uint256 ppvEventPrice)
func (_Smartcontract *SmartcontractFilterer) WatchPriceChanged(opts *bind.WatchOpts, sink chan<- *SmartcontractPriceChanged) (event.Subscription, error) {

	logs, sub, err := _Smartcontract.contract.WatchLogs(opts, "PriceChanged")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SmartcontractPriceChanged)
				if err := _Smartcontract.contract.UnpackLog(event, "PriceChanged", log); err != nil {
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

// ParsePriceChanged is a log parse operation binding the contract event 0xa6dc15bdb68da224c66db4b3838d9a2b205138e8cff6774e57d0af91e196d622.
//
// Solidity: event PriceChanged(uint256 ppvEventPrice)
func (_Smartcontract *SmartcontractFilterer) ParsePriceChanged(log types.Log) (*SmartcontractPriceChanged, error) {
	event := new(SmartcontractPriceChanged)
	if err := _Smartcontract.contract.UnpackLog(event, "PriceChanged", log); err != nil {
		return nil, err
	}
	return event, nil
}
