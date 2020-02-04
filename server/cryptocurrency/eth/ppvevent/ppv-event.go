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
const PpveventABI = "[{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_eventStartTime\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_eventEndTime\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_price\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"subscriptor\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"price\",\"type\":\"uint256\"}],\"name\":\"NewSubscription\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"subscriptor\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"price\",\"type\":\"uint256\"}],\"name\":\"NewUnsubscription\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[],\"name\":\"PPVEventEnded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[],\"name\":\"PPVEventStarted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"ppvEventPrice\",\"type\":\"uint256\"}],\"name\":\"PriceChanged\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_price\",\"type\":\"uint256\"}],\"name\":\"changePrice\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"ended\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"addresspayable\",\"name\":\"dest\",\"type\":\"address\"}],\"name\":\"eventEnd\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"ppvEventEnd\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"ppvEventPrice\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"ppvEventStart\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"invoiceId\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"signature\",\"type\":\"bytes\"}],\"name\":\"subscribe\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"invoiceId\",\"type\":\"bytes32\"}],\"name\":\"unsubscribe\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// PpveventBin is the compiled bytecode used for deploying new contracts.
var PpveventBin = "0x608060405234801561001057600080fd5b50604051610c2c380380610c2c8339818101604052606081101561003357600080fd5b5080516020820151604090920151600080546001600160a01b03191633179055600291909155600391909155600155610bbb806100716000396000f3fe60806040526004361061007b5760003560e01c8063ad9061e41161004e578063ad9061e4146101a9578063d1428312146101be578063ed34ce31146101e4578063f8c60146146101f95761007b565b806312fa6feb146100805780631b5a2b9f146100a95780632fd3308c14610158578063a2b40d191461017f575b600080fd5b34801561008c57600080fd5b50610095610223565b604080519115158252519081900360200190f35b610156600480360360408110156100bf57600080fd5b813591908101906040810160208201356401000000008111156100e157600080fd5b8201836020820111156100f357600080fd5b8035906020019184600183028401116401000000008311171561011557600080fd5b91908080601f01602080910402602001604051908101604052809392919081815260200183838082843760009201919091525092955061022c945050505050565b005b34801561016457600080fd5b5061016d610477565b60408051918252519081900360200190f35b34801561018b57600080fd5b50610156600480360360208110156101a257600080fd5b503561047d565b3480156101b557600080fd5b5061016d610551565b610156600480360360208110156101d457600080fd5b50356001600160a01b0316610557565b3480156101f057600080fd5b5061016d61071b565b34801561020557600080fd5b506101566004803603602081101561021c57600080fd5b5035610721565b60055460ff1681565b60055460ff16801561027b576040805162461bcd60e51b81526020600482015260136024820152721c1c1d88195d995b9d081a185cc8195b991959606a1b604482015290519081900360640190fd5b6001548034146102d2576040805162461bcd60e51b815260206004820181905260248201527f76616c756520646f6573206e6f74206d61746368206576656e74207072696365604482015290519081900360640190fd5b836102db610af9565b50600081815260046020908152604091829020825180840190935280548084526001909101546001600160a01b0316918301919091521561034d5760405162461bcd60e51b8152600401808060200182810382526022815260200180610b646022913960400191505060405180910390fd5b6001546040805160208082019390935280820189905230606090811b908201528151605481830301815260749091019091528051910120600090610390906109aa565b6000549091506001600160a01b03166103a982886109fb565b6001600160a01b0316146103ee5760405162461bcd60e51b815260040180806020018281038252602a815260200180610b3a602a913960400191505060405180910390fd5b6040805180820182523480825233602080840182815260008d8152600483528690209451855551600190940180546001600160a01b0319166001600160a01b039095169490941790935583519081529182015281517f1e05df24f73db39faa0c2d5d26727d08632debce09833123a69214ba943e07c2929181900390910190a150505050505050565b60035481565b6000546001600160a01b031633146104c65760405162461bcd60e51b8152600401808060200182810382526029815260200180610b116029913960400191505060405180910390fd5b60055460ff168015610515576040805162461bcd60e51b81526020600482015260136024820152721c1c1d88195d995b9d081a185cc8195b991959606a1b604482015290519081900360640190fd5b60018290556040805183815290517fa6dc15bdb68da224c66db4b3838d9a2b205138e8cff6774e57d0af91e196d6229181900360200190a15050565b60015481565b6000546001600160a01b031633146105a05760405162461bcd60e51b8152600401808060200182810382526029815260200180610b116029913960400191505060405180910390fd5b60055460ff1680156105ef576040805162461bcd60e51b81526020600482015260136024820152721c1c1d88195d995b9d081a185cc8195b991959606a1b604482015290519081900360640190fd5b60025480421015610647576040805162461bcd60e51b815260206004820152601960248201527f707076206576656e7420686173206e6f74207374617274656400000000000000604482015290519081900360640190fd5b6005805460ff191660011790556040516000906001600160a01b0385169047908381818185875af1925050503d806000811461069f576040519150601f19603f3d011682016040523d82523d6000602084013e6106a4565b606091505b50509050806106ec576040805162461bcd60e51b815260206004820152600f60248201526e1d1c985b9cd9995c8819985a5b1959608a1b604482015290519081900360640190fd5b6040517f545c09a6805d0d5edc24946e122843a3b3cfb83cef050826f143983353509e2f90600090a150505050565b60025481565b60055460ff168015610770576040805162461bcd60e51b81526020600482015260136024820152721c1c1d88195d995b9d081a185cc8195b991959606a1b604482015290519081900360640190fd5b6002548042106107c7576040805162461bcd60e51b815260206004820152601d60248201527f707076206576656e742068617320616c72656164792073746172746564000000604482015290519081900360640190fd5b6107cf610af9565b50600083815260046020908152604091829020825180840190935280548084526001909101546001600160a01b031691830191909152610856576040805162461bcd60e51b815260206004820152601960248201527f6e6f20737562736372697074696f6e20617661696c61626c6500000000000000604482015290519081900360640190fd5b60208101516001600160a01b031633146108b7576040805162461bcd60e51b815260206004820181905260248201527f7375627363726962657220646f6573206e6f74206d617463682073656e646572604482015290519081900360640190fd5b60008481526004602052604080822082815560010180546001600160a01b0319169055825190513391908381818185875af1925050503d8060008114610919576040519150601f19603f3d011682016040523d82523d6000602084013e61091e565b606091505b5050905080610966576040805162461bcd60e51b815260206004820152600f60248201526e1d1c985b9cd9995c8819985a5b1959608a1b604482015290519081900360640190fd5b815160408051338152602081019290925280517f79d9d14aed97c97b8dc663528ec3b03eecb8d467956fc03f37179b188a04efa49281900390910190a15050505050565b604080517f19457468657265756d205369676e6564204d6573736167653a0a333200000000602080830191909152603c8083019490945282518083039094018452605c909101909152815191012090565b600080600080610a0a85610a82565b92509250925060018684848460405160008152602001604052604051808581526020018460ff1660ff1681526020018381526020018281526020019450505050506020604051602081039080840390855afa158015610a6d573d6000803e3d6000fd5b5050604051601f190151979650505050505050565b60008060008351604114610add576040805162461bcd60e51b815260206004820152601860248201527f696e76616c6964207369676e6174757265206c656e6774680000000000000000604482015290519081900360640190fd5b5050506020810151604082015160609092015160001a92909190565b60408051808201909152600080825260208201529056fe6f6e6c792074686520636f6e7472616374206f776e65722063616e2063616c6c2066756e6374696f6e6d65737361676520776173206e6f74207369676e656420627920617574686f72697a656420706172747974686520696e766f6963652068617320616c7265616479206265656e207061796564a26469706673582212203a48ca30c4216a7375d2af3e5d02e04d2f0e55c8cfb617b04a8acf6b692ddf2464736f6c63430006020033"

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

// Subscribe is a paid mutator transaction binding the contract method 0x1b5a2b9f.
//
// Solidity: function subscribe(bytes32 invoiceId, bytes signature) returns()
func (_Ppvevent *PpveventTransactor) Subscribe(opts *bind.TransactOpts, invoiceId [32]byte, signature []byte) (*types.Transaction, error) {
	return _Ppvevent.contract.Transact(opts, "subscribe", invoiceId, signature)
}

// Subscribe is a paid mutator transaction binding the contract method 0x1b5a2b9f.
//
// Solidity: function subscribe(bytes32 invoiceId, bytes signature) returns()
func (_Ppvevent *PpveventSession) Subscribe(invoiceId [32]byte, signature []byte) (*types.Transaction, error) {
	return _Ppvevent.Contract.Subscribe(&_Ppvevent.TransactOpts, invoiceId, signature)
}

// Subscribe is a paid mutator transaction binding the contract method 0x1b5a2b9f.
//
// Solidity: function subscribe(bytes32 invoiceId, bytes signature) returns()
func (_Ppvevent *PpveventTransactorSession) Subscribe(invoiceId [32]byte, signature []byte) (*types.Transaction, error) {
	return _Ppvevent.Contract.Subscribe(&_Ppvevent.TransactOpts, invoiceId, signature)
}

// Unsubscribe is a paid mutator transaction binding the contract method 0xf8c60146.
//
// Solidity: function unsubscribe(bytes32 invoiceId) returns()
func (_Ppvevent *PpveventTransactor) Unsubscribe(opts *bind.TransactOpts, invoiceId [32]byte) (*types.Transaction, error) {
	return _Ppvevent.contract.Transact(opts, "unsubscribe", invoiceId)
}

// Unsubscribe is a paid mutator transaction binding the contract method 0xf8c60146.
//
// Solidity: function unsubscribe(bytes32 invoiceId) returns()
func (_Ppvevent *PpveventSession) Unsubscribe(invoiceId [32]byte) (*types.Transaction, error) {
	return _Ppvevent.Contract.Unsubscribe(&_Ppvevent.TransactOpts, invoiceId)
}

// Unsubscribe is a paid mutator transaction binding the contract method 0xf8c60146.
//
// Solidity: function unsubscribe(bytes32 invoiceId) returns()
func (_Ppvevent *PpveventTransactorSession) Unsubscribe(invoiceId [32]byte) (*types.Transaction, error) {
	return _Ppvevent.Contract.Unsubscribe(&_Ppvevent.TransactOpts, invoiceId)
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
