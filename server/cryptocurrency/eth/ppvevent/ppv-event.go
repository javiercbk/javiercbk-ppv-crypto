// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package ppvevent

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

// PpveventABI is the input ABI used to generate the binding from.
const PpveventABI = "[{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_eventStartTime\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_eventEndTime\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_price\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"subscriptor\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"price\",\"type\":\"uint256\"}],\"name\":\"NewSubscription\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"subscriptor\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"price\",\"type\":\"uint256\"}],\"name\":\"NewUnsubscription\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[],\"name\":\"PPVEventEnded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[],\"name\":\"PPVEventStarted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"ppvEventPrice\",\"type\":\"uint256\"}],\"name\":\"PriceChanged\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_price\",\"type\":\"uint256\"}],\"name\":\"changePrice\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"ended\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"addresspayable\",\"name\":\"dest\",\"type\":\"address\"}],\"name\":\"eventEnd\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"ppvEventEnd\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"ppvEventPrice\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"ppvEventStart\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"subscribe\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"unsubscribe\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// PpveventBin is the compiled bytecode used for deploying new contracts.
var PpveventBin = "0x608060405234801561001057600080fd5b506040516108213803806108218339818101604052606081101561003357600080fd5b5080516020820151604090920151600080546001600160a01b031916331790556002919091556003919091556001556107b0806100716000396000f3fe60806040526004361061007b5760003560e01c8063ad9061e41161004e578063ad9061e414610104578063d142831214610119578063ed34ce311461013f578063fcae4484146101545761007b565b806312fa6feb146100805780632fd3308c146100a95780638f449a05146100d0578063a2b40d19146100da575b600080fd5b34801561008c57600080fd5b50610095610169565b604080519115158252519081900360200190f35b3480156100b557600080fd5b506100be610172565b60408051918252519081900360200190f35b6100d8610178565b005b3480156100e657600080fd5b506100d8600480360360208110156100fd57600080fd5b50356102c6565b34801561011057600080fd5b506100be61039a565b6100d86004803603602081101561012f57600080fd5b50356001600160a01b03166103a0565b34801561014b57600080fd5b506100be610564565b34801561016057600080fd5b506100d861056a565b60055460ff1681565b60035481565b60055460ff1680156101c7576040805162461bcd60e51b81526020600482015260136024820152721c1c1d88195d995b9d081a185cc8195b991959606a1b604482015290519081900360640190fd5b60015480341461021e576040805162461bcd60e51b815260206004820181905260248201527f76616c756520646f6573206e6f74206d61746368206576656e74207072696365604482015290519081900360640190fd5b3360009081526004602052604090205415610275576040805162461bcd60e51b8152602060048201526012602482015271185b1c9958591e481cdd589cd8dc9a58995960721b604482015290519081900360640190fd5b33600081815260046020908152604091829020349081905582519384529083015280517f1e05df24f73db39faa0c2d5d26727d08632debce09833123a69214ba943e07c29281900390910190a15050565b6000546001600160a01b0316331461030f5760405162461bcd60e51b81526004018080602001828103825260298152602001806107526029913960400191505060405180910390fd5b60055460ff16801561035e576040805162461bcd60e51b81526020600482015260136024820152721c1c1d88195d995b9d081a185cc8195b991959606a1b604482015290519081900360640190fd5b60018290556040805183815290517fa6dc15bdb68da224c66db4b3838d9a2b205138e8cff6774e57d0af91e196d6229181900360200190a15050565b60015481565b6000546001600160a01b031633146103e95760405162461bcd60e51b81526004018080602001828103825260298152602001806107526029913960400191505060405180910390fd5b60055460ff168015610438576040805162461bcd60e51b81526020600482015260136024820152721c1c1d88195d995b9d081a185cc8195b991959606a1b604482015290519081900360640190fd5b60025480421015610490576040805162461bcd60e51b815260206004820152601960248201527f707076206576656e7420686173206e6f74207374617274656400000000000000604482015290519081900360640190fd5b6005805460ff191660011790556040516000906001600160a01b0385169047908381818185875af1925050503d80600081146104e8576040519150601f19603f3d011682016040523d82523d6000602084013e6104ed565b606091505b5050905080610535576040805162461bcd60e51b815260206004820152600f60248201526e1d1c985b9cd9995c8819985a5b1959608a1b604482015290519081900360640190fd5b6040517f545c09a6805d0d5edc24946e122843a3b3cfb83cef050826f143983353509e2f90600090a150505050565b60025481565b60055460ff1680156105b9576040805162461bcd60e51b81526020600482015260136024820152721c1c1d88195d995b9d081a185cc8195b991959606a1b604482015290519081900360640190fd5b600254804210610610576040805162461bcd60e51b815260206004820152601d60248201527f707076206576656e742068617320616c72656164792073746172746564000000604482015290519081900360640190fd5b3360009081526004602052604090205480610672576040805162461bcd60e51b815260206004820152601960248201527f6e6f20737562736372697074696f6e20617661696c61626c6500000000000000604482015290519081900360640190fd5b336000818152600460205260408082208290555190919083908381818185875af1925050503d80600081146106c3576040519150601f19603f3d011682016040523d82523d6000602084013e6106c8565b606091505b5050905080610710576040805162461bcd60e51b815260206004820152600f60248201526e1d1c985b9cd9995c8819985a5b1959608a1b604482015290519081900360640190fd5b604080513381526020810184905281517f79d9d14aed97c97b8dc663528ec3b03eecb8d467956fc03f37179b188a04efa4929181900390910190a15050505056fe6f6e6c792074686520636f6e7472616374206f776e65722063616e2063616c6c2066756e6374696f6ea2646970667358221220c80f631dbd94801a925cbe163e9fcde9646f377866faab17275dbeb996bcd25764736f6c63430006000033"

// DeployPpvevent deploys a new Ethereum contract, binding an instance of Ppvevent to it.
func DeployPpvevent(auth *bind.TransactOpts, backend bind.ContractBackend, _eventStartTime *big.Int, _eventEndTime *big.Int, _price *big.Int) (common.Address, *types.Transaction, *Ppvevent, error) {
	parsed, err := abi.JSON(strings.NewReader(PpveventABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(PpveventBin), backend, _eventStartTime, _eventEndTime, _price)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Ppvevent{PpveventCaller: PpveventCaller{contract: contract}, PpveventTransactor: PpveventTransactor{contract: contract}, PpveventFilterer: PpveventFilterer{contract: contract}}, nil
}

// Ppvevent is an auto generated Go binding around an Ethereum contract.
type Ppvevent struct {
	PpveventCaller     // Read-only binding to the contract
	PpveventTransactor // Write-only binding to the contract
	PpveventFilterer   // Log filterer for contract events
}

// PpveventCaller is an auto generated read-only Go binding around an Ethereum contract.
type PpveventCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PpveventTransactor is an auto generated write-only Go binding around an Ethereum contract.
type PpveventTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PpveventFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type PpveventFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PpveventSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type PpveventSession struct {
	Contract     *Ppvevent         // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// PpveventCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type PpveventCallerSession struct {
	Contract *PpveventCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts   // Call options to use throughout this session
}

// PpveventTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type PpveventTransactorSession struct {
	Contract     *PpveventTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// PpveventRaw is an auto generated low-level Go binding around an Ethereum contract.
type PpveventRaw struct {
	Contract *Ppvevent // Generic contract binding to access the raw methods on
}

// PpveventCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type PpveventCallerRaw struct {
	Contract *PpveventCaller // Generic read-only contract binding to access the raw methods on
}

// PpveventTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type PpveventTransactorRaw struct {
	Contract *PpveventTransactor // Generic write-only contract binding to access the raw methods on
}

// NewPpvevent creates a new instance of Ppvevent, bound to a specific deployed contract.
func NewPpvevent(address common.Address, backend bind.ContractBackend) (*Ppvevent, error) {
	contract, err := bindPpvevent(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Ppvevent{PpveventCaller: PpveventCaller{contract: contract}, PpveventTransactor: PpveventTransactor{contract: contract}, PpveventFilterer: PpveventFilterer{contract: contract}}, nil
}

// NewPpveventCaller creates a new read-only instance of Ppvevent, bound to a specific deployed contract.
func NewPpveventCaller(address common.Address, caller bind.ContractCaller) (*PpveventCaller, error) {
	contract, err := bindPpvevent(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &PpveventCaller{contract: contract}, nil
}

// NewPpveventTransactor creates a new write-only instance of Ppvevent, bound to a specific deployed contract.
func NewPpveventTransactor(address common.Address, transactor bind.ContractTransactor) (*PpveventTransactor, error) {
	contract, err := bindPpvevent(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &PpveventTransactor{contract: contract}, nil
}

// NewPpveventFilterer creates a new log filterer instance of Ppvevent, bound to a specific deployed contract.
func NewPpveventFilterer(address common.Address, filterer bind.ContractFilterer) (*PpveventFilterer, error) {
	contract, err := bindPpvevent(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &PpveventFilterer{contract: contract}, nil
}

// bindPpvevent binds a generic wrapper to an already deployed contract.
func bindPpvevent(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(PpveventABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Ppvevent *PpveventRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Ppvevent.Contract.PpveventCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Ppvevent *PpveventRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Ppvevent.Contract.PpveventTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Ppvevent *PpveventRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Ppvevent.Contract.PpveventTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Ppvevent *PpveventCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Ppvevent.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Ppvevent *PpveventTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Ppvevent.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Ppvevent *PpveventTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Ppvevent.Contract.contract.Transact(opts, method, params...)
}

// ChangePrice is a paid mutator transaction binding the contract method 0xa2b40d19.
//
// Solidity: function changePrice(uint256 _price) returns()
func (_Ppvevent *PpveventTransactor) ChangePrice(opts *bind.TransactOpts, _price *big.Int) (*types.Transaction, error) {
	return _Ppvevent.contract.Transact(opts, "changePrice", _price)
}

// ChangePrice is a paid mutator transaction binding the contract method 0xa2b40d19.
//
// Solidity: function changePrice(uint256 _price) returns()
func (_Ppvevent *PpveventSession) ChangePrice(_price *big.Int) (*types.Transaction, error) {
	return _Ppvevent.Contract.ChangePrice(&_Ppvevent.TransactOpts, _price)
}

// ChangePrice is a paid mutator transaction binding the contract method 0xa2b40d19.
//
// Solidity: function changePrice(uint256 _price) returns()
func (_Ppvevent *PpveventTransactorSession) ChangePrice(_price *big.Int) (*types.Transaction, error) {
	return _Ppvevent.Contract.ChangePrice(&_Ppvevent.TransactOpts, _price)
}

// Ended is a paid mutator transaction binding the contract method 0x12fa6feb.
//
// Solidity: function ended() returns(bool)
func (_Ppvevent *PpveventTransactor) Ended(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Ppvevent.contract.Transact(opts, "ended")
}

// Ended is a paid mutator transaction binding the contract method 0x12fa6feb.
//
// Solidity: function ended() returns(bool)
func (_Ppvevent *PpveventSession) Ended() (*types.Transaction, error) {
	return _Ppvevent.Contract.Ended(&_Ppvevent.TransactOpts)
}

// Ended is a paid mutator transaction binding the contract method 0x12fa6feb.
//
// Solidity: function ended() returns(bool)
func (_Ppvevent *PpveventTransactorSession) Ended() (*types.Transaction, error) {
	return _Ppvevent.Contract.Ended(&_Ppvevent.TransactOpts)
}

// EventEnd is a paid mutator transaction binding the contract method 0xd1428312.
//
// Solidity: function eventEnd(address dest) returns()
func (_Ppvevent *PpveventTransactor) EventEnd(opts *bind.TransactOpts, dest common.Address) (*types.Transaction, error) {
	return _Ppvevent.contract.Transact(opts, "eventEnd", dest)
}

// EventEnd is a paid mutator transaction binding the contract method 0xd1428312.
//
// Solidity: function eventEnd(address dest) returns()
func (_Ppvevent *PpveventSession) EventEnd(dest common.Address) (*types.Transaction, error) {
	return _Ppvevent.Contract.EventEnd(&_Ppvevent.TransactOpts, dest)
}

// EventEnd is a paid mutator transaction binding the contract method 0xd1428312.
//
// Solidity: function eventEnd(address dest) returns()
func (_Ppvevent *PpveventTransactorSession) EventEnd(dest common.Address) (*types.Transaction, error) {
	return _Ppvevent.Contract.EventEnd(&_Ppvevent.TransactOpts, dest)
}

// PpvEventEnd is a paid mutator transaction binding the contract method 0x2fd3308c.
//
// Solidity: function ppvEventEnd() returns(uint256)
func (_Ppvevent *PpveventTransactor) PpvEventEnd(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Ppvevent.contract.Transact(opts, "ppvEventEnd")
}

// PpvEventEnd is a paid mutator transaction binding the contract method 0x2fd3308c.
//
// Solidity: function ppvEventEnd() returns(uint256)
func (_Ppvevent *PpveventSession) PpvEventEnd() (*types.Transaction, error) {
	return _Ppvevent.Contract.PpvEventEnd(&_Ppvevent.TransactOpts)
}

// PpvEventEnd is a paid mutator transaction binding the contract method 0x2fd3308c.
//
// Solidity: function ppvEventEnd() returns(uint256)
func (_Ppvevent *PpveventTransactorSession) PpvEventEnd() (*types.Transaction, error) {
	return _Ppvevent.Contract.PpvEventEnd(&_Ppvevent.TransactOpts)
}

// PpvEventPrice is a paid mutator transaction binding the contract method 0xad9061e4.
//
// Solidity: function ppvEventPrice() returns(uint256)
func (_Ppvevent *PpveventTransactor) PpvEventPrice(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Ppvevent.contract.Transact(opts, "ppvEventPrice")
}

// PpvEventPrice is a paid mutator transaction binding the contract method 0xad9061e4.
//
// Solidity: function ppvEventPrice() returns(uint256)
func (_Ppvevent *PpveventSession) PpvEventPrice() (*types.Transaction, error) {
	return _Ppvevent.Contract.PpvEventPrice(&_Ppvevent.TransactOpts)
}

// PpvEventPrice is a paid mutator transaction binding the contract method 0xad9061e4.
//
// Solidity: function ppvEventPrice() returns(uint256)
func (_Ppvevent *PpveventTransactorSession) PpvEventPrice() (*types.Transaction, error) {
	return _Ppvevent.Contract.PpvEventPrice(&_Ppvevent.TransactOpts)
}

// PpvEventStart is a paid mutator transaction binding the contract method 0xed34ce31.
//
// Solidity: function ppvEventStart() returns(uint256)
func (_Ppvevent *PpveventTransactor) PpvEventStart(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Ppvevent.contract.Transact(opts, "ppvEventStart")
}

// PpvEventStart is a paid mutator transaction binding the contract method 0xed34ce31.
//
// Solidity: function ppvEventStart() returns(uint256)
func (_Ppvevent *PpveventSession) PpvEventStart() (*types.Transaction, error) {
	return _Ppvevent.Contract.PpvEventStart(&_Ppvevent.TransactOpts)
}

// PpvEventStart is a paid mutator transaction binding the contract method 0xed34ce31.
//
// Solidity: function ppvEventStart() returns(uint256)
func (_Ppvevent *PpveventTransactorSession) PpvEventStart() (*types.Transaction, error) {
	return _Ppvevent.Contract.PpvEventStart(&_Ppvevent.TransactOpts)
}

// Subscribe is a paid mutator transaction binding the contract method 0x8f449a05.
//
// Solidity: function subscribe() returns()
func (_Ppvevent *PpveventTransactor) Subscribe(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Ppvevent.contract.Transact(opts, "subscribe")
}

// Subscribe is a paid mutator transaction binding the contract method 0x8f449a05.
//
// Solidity: function subscribe() returns()
func (_Ppvevent *PpveventSession) Subscribe() (*types.Transaction, error) {
	return _Ppvevent.Contract.Subscribe(&_Ppvevent.TransactOpts)
}

// Subscribe is a paid mutator transaction binding the contract method 0x8f449a05.
//
// Solidity: function subscribe() returns()
func (_Ppvevent *PpveventTransactorSession) Subscribe() (*types.Transaction, error) {
	return _Ppvevent.Contract.Subscribe(&_Ppvevent.TransactOpts)
}

// Unsubscribe is a paid mutator transaction binding the contract method 0xfcae4484.
//
// Solidity: function unsubscribe() returns()
func (_Ppvevent *PpveventTransactor) Unsubscribe(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Ppvevent.contract.Transact(opts, "unsubscribe")
}

// Unsubscribe is a paid mutator transaction binding the contract method 0xfcae4484.
//
// Solidity: function unsubscribe() returns()
func (_Ppvevent *PpveventSession) Unsubscribe() (*types.Transaction, error) {
	return _Ppvevent.Contract.Unsubscribe(&_Ppvevent.TransactOpts)
}

// Unsubscribe is a paid mutator transaction binding the contract method 0xfcae4484.
//
// Solidity: function unsubscribe() returns()
func (_Ppvevent *PpveventTransactorSession) Unsubscribe() (*types.Transaction, error) {
	return _Ppvevent.Contract.Unsubscribe(&_Ppvevent.TransactOpts)
}

// PpveventNewSubscriptionIterator is returned from FilterNewSubscription and is used to iterate over the raw logs and unpacked data for NewSubscription events raised by the Ppvevent contract.
type PpveventNewSubscriptionIterator struct {
	Event *PpveventNewSubscription // Event containing the contract specifics and raw log

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
func (it *PpveventNewSubscriptionIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PpveventNewSubscription)
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
		it.Event = new(PpveventNewSubscription)
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
func (it *PpveventNewSubscriptionIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PpveventNewSubscriptionIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PpveventNewSubscription represents a NewSubscription event raised by the Ppvevent contract.
type PpveventNewSubscription struct {
	Subscriptor common.Address
	Price       *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterNewSubscription is a free log retrieval operation binding the contract event 0x1e05df24f73db39faa0c2d5d26727d08632debce09833123a69214ba943e07c2.
//
// Solidity: event NewSubscription(address subscriptor, uint256 price)
func (_Ppvevent *PpveventFilterer) FilterNewSubscription(opts *bind.FilterOpts) (*PpveventNewSubscriptionIterator, error) {

	logs, sub, err := _Ppvevent.contract.FilterLogs(opts, "NewSubscription")
	if err != nil {
		return nil, err
	}
	return &PpveventNewSubscriptionIterator{contract: _Ppvevent.contract, event: "NewSubscription", logs: logs, sub: sub}, nil
}

// WatchNewSubscription is a free log subscription operation binding the contract event 0x1e05df24f73db39faa0c2d5d26727d08632debce09833123a69214ba943e07c2.
//
// Solidity: event NewSubscription(address subscriptor, uint256 price)
func (_Ppvevent *PpveventFilterer) WatchNewSubscription(opts *bind.WatchOpts, sink chan<- *PpveventNewSubscription) (event.Subscription, error) {

	logs, sub, err := _Ppvevent.contract.WatchLogs(opts, "NewSubscription")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PpveventNewSubscription)
				if err := _Ppvevent.contract.UnpackLog(event, "NewSubscription", log); err != nil {
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
func (_Ppvevent *PpveventFilterer) ParseNewSubscription(log types.Log) (*PpveventNewSubscription, error) {
	event := new(PpveventNewSubscription)
	if err := _Ppvevent.contract.UnpackLog(event, "NewSubscription", log); err != nil {
		return nil, err
	}
	return event, nil
}

// PpveventNewUnsubscriptionIterator is returned from FilterNewUnsubscription and is used to iterate over the raw logs and unpacked data for NewUnsubscription events raised by the Ppvevent contract.
type PpveventNewUnsubscriptionIterator struct {
	Event *PpveventNewUnsubscription // Event containing the contract specifics and raw log

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
func (it *PpveventNewUnsubscriptionIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PpveventNewUnsubscription)
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
		it.Event = new(PpveventNewUnsubscription)
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
func (it *PpveventNewUnsubscriptionIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PpveventNewUnsubscriptionIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PpveventNewUnsubscription represents a NewUnsubscription event raised by the Ppvevent contract.
type PpveventNewUnsubscription struct {
	Subscriptor common.Address
	Price       *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterNewUnsubscription is a free log retrieval operation binding the contract event 0x79d9d14aed97c97b8dc663528ec3b03eecb8d467956fc03f37179b188a04efa4.
//
// Solidity: event NewUnsubscription(address subscriptor, uint256 price)
func (_Ppvevent *PpveventFilterer) FilterNewUnsubscription(opts *bind.FilterOpts) (*PpveventNewUnsubscriptionIterator, error) {

	logs, sub, err := _Ppvevent.contract.FilterLogs(opts, "NewUnsubscription")
	if err != nil {
		return nil, err
	}
	return &PpveventNewUnsubscriptionIterator{contract: _Ppvevent.contract, event: "NewUnsubscription", logs: logs, sub: sub}, nil
}

// WatchNewUnsubscription is a free log subscription operation binding the contract event 0x79d9d14aed97c97b8dc663528ec3b03eecb8d467956fc03f37179b188a04efa4.
//
// Solidity: event NewUnsubscription(address subscriptor, uint256 price)
func (_Ppvevent *PpveventFilterer) WatchNewUnsubscription(opts *bind.WatchOpts, sink chan<- *PpveventNewUnsubscription) (event.Subscription, error) {

	logs, sub, err := _Ppvevent.contract.WatchLogs(opts, "NewUnsubscription")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PpveventNewUnsubscription)
				if err := _Ppvevent.contract.UnpackLog(event, "NewUnsubscription", log); err != nil {
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
func (_Ppvevent *PpveventFilterer) ParseNewUnsubscription(log types.Log) (*PpveventNewUnsubscription, error) {
	event := new(PpveventNewUnsubscription)
	if err := _Ppvevent.contract.UnpackLog(event, "NewUnsubscription", log); err != nil {
		return nil, err
	}
	return event, nil
}

// PpveventPPVEventEndedIterator is returned from FilterPPVEventEnded and is used to iterate over the raw logs and unpacked data for PPVEventEnded events raised by the Ppvevent contract.
type PpveventPPVEventEndedIterator struct {
	Event *PpveventPPVEventEnded // Event containing the contract specifics and raw log

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
func (it *PpveventPPVEventEndedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PpveventPPVEventEnded)
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
		it.Event = new(PpveventPPVEventEnded)
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
func (it *PpveventPPVEventEndedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PpveventPPVEventEndedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PpveventPPVEventEnded represents a PPVEventEnded event raised by the Ppvevent contract.
type PpveventPPVEventEnded struct {
	Raw types.Log // Blockchain specific contextual infos
}

// FilterPPVEventEnded is a free log retrieval operation binding the contract event 0x545c09a6805d0d5edc24946e122843a3b3cfb83cef050826f143983353509e2f.
//
// Solidity: event PPVEventEnded()
func (_Ppvevent *PpveventFilterer) FilterPPVEventEnded(opts *bind.FilterOpts) (*PpveventPPVEventEndedIterator, error) {

	logs, sub, err := _Ppvevent.contract.FilterLogs(opts, "PPVEventEnded")
	if err != nil {
		return nil, err
	}
	return &PpveventPPVEventEndedIterator{contract: _Ppvevent.contract, event: "PPVEventEnded", logs: logs, sub: sub}, nil
}

// WatchPPVEventEnded is a free log subscription operation binding the contract event 0x545c09a6805d0d5edc24946e122843a3b3cfb83cef050826f143983353509e2f.
//
// Solidity: event PPVEventEnded()
func (_Ppvevent *PpveventFilterer) WatchPPVEventEnded(opts *bind.WatchOpts, sink chan<- *PpveventPPVEventEnded) (event.Subscription, error) {

	logs, sub, err := _Ppvevent.contract.WatchLogs(opts, "PPVEventEnded")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PpveventPPVEventEnded)
				if err := _Ppvevent.contract.UnpackLog(event, "PPVEventEnded", log); err != nil {
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
func (_Ppvevent *PpveventFilterer) ParsePPVEventEnded(log types.Log) (*PpveventPPVEventEnded, error) {
	event := new(PpveventPPVEventEnded)
	if err := _Ppvevent.contract.UnpackLog(event, "PPVEventEnded", log); err != nil {
		return nil, err
	}
	return event, nil
}

// PpveventPPVEventStartedIterator is returned from FilterPPVEventStarted and is used to iterate over the raw logs and unpacked data for PPVEventStarted events raised by the Ppvevent contract.
type PpveventPPVEventStartedIterator struct {
	Event *PpveventPPVEventStarted // Event containing the contract specifics and raw log

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
func (it *PpveventPPVEventStartedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PpveventPPVEventStarted)
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
		it.Event = new(PpveventPPVEventStarted)
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
func (it *PpveventPPVEventStartedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PpveventPPVEventStartedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PpveventPPVEventStarted represents a PPVEventStarted event raised by the Ppvevent contract.
type PpveventPPVEventStarted struct {
	Raw types.Log // Blockchain specific contextual infos
}

// FilterPPVEventStarted is a free log retrieval operation binding the contract event 0x0e10146934a460c3f27a8ca4c63fdad156eed0ae2051edc0c1839268ade6c9ad.
//
// Solidity: event PPVEventStarted()
func (_Ppvevent *PpveventFilterer) FilterPPVEventStarted(opts *bind.FilterOpts) (*PpveventPPVEventStartedIterator, error) {

	logs, sub, err := _Ppvevent.contract.FilterLogs(opts, "PPVEventStarted")
	if err != nil {
		return nil, err
	}
	return &PpveventPPVEventStartedIterator{contract: _Ppvevent.contract, event: "PPVEventStarted", logs: logs, sub: sub}, nil
}

// WatchPPVEventStarted is a free log subscription operation binding the contract event 0x0e10146934a460c3f27a8ca4c63fdad156eed0ae2051edc0c1839268ade6c9ad.
//
// Solidity: event PPVEventStarted()
func (_Ppvevent *PpveventFilterer) WatchPPVEventStarted(opts *bind.WatchOpts, sink chan<- *PpveventPPVEventStarted) (event.Subscription, error) {

	logs, sub, err := _Ppvevent.contract.WatchLogs(opts, "PPVEventStarted")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PpveventPPVEventStarted)
				if err := _Ppvevent.contract.UnpackLog(event, "PPVEventStarted", log); err != nil {
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
func (_Ppvevent *PpveventFilterer) ParsePPVEventStarted(log types.Log) (*PpveventPPVEventStarted, error) {
	event := new(PpveventPPVEventStarted)
	if err := _Ppvevent.contract.UnpackLog(event, "PPVEventStarted", log); err != nil {
		return nil, err
	}
	return event, nil
}

// PpveventPriceChangedIterator is returned from FilterPriceChanged and is used to iterate over the raw logs and unpacked data for PriceChanged events raised by the Ppvevent contract.
type PpveventPriceChangedIterator struct {
	Event *PpveventPriceChanged // Event containing the contract specifics and raw log

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
func (it *PpveventPriceChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PpveventPriceChanged)
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
		it.Event = new(PpveventPriceChanged)
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
func (it *PpveventPriceChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PpveventPriceChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PpveventPriceChanged represents a PriceChanged event raised by the Ppvevent contract.
type PpveventPriceChanged struct {
	PpvEventPrice *big.Int
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterPriceChanged is a free log retrieval operation binding the contract event 0xa6dc15bdb68da224c66db4b3838d9a2b205138e8cff6774e57d0af91e196d622.
//
// Solidity: event PriceChanged(uint256 ppvEventPrice)
func (_Ppvevent *PpveventFilterer) FilterPriceChanged(opts *bind.FilterOpts) (*PpveventPriceChangedIterator, error) {

	logs, sub, err := _Ppvevent.contract.FilterLogs(opts, "PriceChanged")
	if err != nil {
		return nil, err
	}
	return &PpveventPriceChangedIterator{contract: _Ppvevent.contract, event: "PriceChanged", logs: logs, sub: sub}, nil
}

// WatchPriceChanged is a free log subscription operation binding the contract event 0xa6dc15bdb68da224c66db4b3838d9a2b205138e8cff6774e57d0af91e196d622.
//
// Solidity: event PriceChanged(uint256 ppvEventPrice)
func (_Ppvevent *PpveventFilterer) WatchPriceChanged(opts *bind.WatchOpts, sink chan<- *PpveventPriceChanged) (event.Subscription, error) {

	logs, sub, err := _Ppvevent.contract.WatchLogs(opts, "PriceChanged")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PpveventPriceChanged)
				if err := _Ppvevent.contract.UnpackLog(event, "PriceChanged", log); err != nil {
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
func (_Ppvevent *PpveventFilterer) ParsePriceChanged(log types.Log) (*PpveventPriceChanged, error) {
	event := new(PpveventPriceChanged)
	if err := _Ppvevent.contract.UnpackLog(event, "PriceChanged", log); err != nil {
		return nil, err
	}
	return event, nil
}
