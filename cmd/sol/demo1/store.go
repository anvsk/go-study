// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package store

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
)

// StoreMetaData contains all meta data concerning the Store contract.
var StoreMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"stateMutability\":\"payable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"InvalidState\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OnlyBuyer\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OnlySeller\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ValueNotEven\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[],\"name\":\"Aborted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[],\"name\":\"ItemReceived\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[],\"name\":\"PurchaseConfirmed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[],\"name\":\"SellerRefunded\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"abort\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"buyer\",\"outputs\":[{\"internalType\":\"addresspayable\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"confirmPurchase\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"confirmReceived\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"refundSeller\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"seller\",\"outputs\":[{\"internalType\":\"addresspayable\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"state\",\"outputs\":[{\"internalType\":\"enumPurchase.State\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"value\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Bin: "0x608060405233600160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550600234610052919061010c565b600081905550346000546002610068919061013d565b1461009f576040517fbe3e4c4200000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b610197565b6000819050919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601260045260246000fd5b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b6000610117826100a4565b9150610122836100a4565b925082610132576101316100ae565b5b828204905092915050565b6000610148826100a4565b9150610153836100a4565b9250817fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff048311821515161561018c5761018b6100dd565b5b828202905092915050565b610a1d806101a66000396000f3fe60806040526004361061007b5760003560e01c806373fac6f01161004e57806373fac6f014610118578063c19d93fb1461012f578063c7981b1b1461015a578063d6960697146101715761007b565b806308551a531461008057806335a063b4146100ab5780633fa4f245146100c25780637150d8ae146100ed575b600080fd5b34801561008c57600080fd5b5061009561017b565b6040516100a2919061087d565b60405180910390f35b3480156100b757600080fd5b506100c06101a1565b005b3480156100ce57600080fd5b506100d7610359565b6040516100e491906108b1565b60405180910390f35b3480156100f957600080fd5b5061010261035f565b60405161010f919061087d565b60405180910390f35b34801561012457600080fd5b5061012d610385565b005b34801561013b57600080fd5b5061014461053e565b6040516101519190610943565b60405180910390f35b34801561016657600080fd5b5061016f610551565b005b610179610717565b005b600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614610228576040517f85d1f72600000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b600080600381111561023d5761023c6108cc565b5b600260149054906101000a900460ff16600381111561025f5761025e6108cc565b5b14610296576040517fbaf3f0f700000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b7f72c874aeff0b183a56e2b79c71b46e1aed4dee5e09862134b8821ba2fddbf8bf60405160405180910390a16003600260146101000a81548160ff021916908360038111156102e8576102e76108cc565b5b0217905550600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff166108fc479081150290604051600060405180830381858888f19350505050158015610355573d6000803e3d6000fd5b5050565b60005481565b600260009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b600260009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff161461040c576040517f86efbb5500000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6001806003811115610421576104206108cc565b5b600260149054906101000a900460ff166003811115610443576104426108cc565b5b1461047a576040517fbaf3f0f700000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b7fe89152acd703c9d8c7d28829d443260b411454d45394e7995815140c8cbcbcf760405160405180910390a160028060146101000a81548160ff021916908360038111156104cb576104ca6108cc565b5b0217905550600260009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff166108fc6000549081150290604051600060405180830381858888f1935050505015801561053a573d6000803e3d6000fd5b5050565b600260149054906101000a900460ff1681565b600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff16146105d8576040517f85d1f72600000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b60028060038111156105ed576105ec6108cc565b5b600260149054906101000a900460ff16600381111561060f5761060e6108cc565b5b14610646576040517fbaf3f0f700000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b7ffda69c32bcfdba840a167777906b173b607eb8b4d8853b97a80d26e613d858db60405160405180910390a16003600260146101000a81548160ff02191690836003811115610698576106976108cc565b5b0217905550600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff166108fc60005460036106e8919061098d565b9081150290604051600060405180830381858888f19350505050158015610713573d6000803e3d6000fd5b5050565b600080600381111561072c5761072b6108cc565b5b600260149054906101000a900460ff16600381111561074e5761074d6108cc565b5b14610785576040517fbaf3f0f700000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6000546002610794919061098d565b3414806107a057600080fd5b7fd5d55c8a68912e9a110618df8d5e2e83b8d83211c57a8ddd1203df92885dc88160405160405180910390a133600260006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055506001600260146101000a81548160ff02191690836003811115610833576108326108cc565b5b02179055505050565b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b60006108678261083c565b9050919050565b6108778161085c565b82525050565b6000602082019050610892600083018461086e565b92915050565b6000819050919050565b6108ab81610898565b82525050565b60006020820190506108c660008301846108a2565b92915050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052602160045260246000fd5b6004811061090c5761090b6108cc565b5b50565b600081905061091d826108fb565b919050565b600061092d8261090f565b9050919050565b61093d81610922565b82525050565b60006020820190506109586000830184610934565b92915050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b600061099882610898565b91506109a383610898565b9250817fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff04831182151516156109dc576109db61095e565b5b82820290509291505056fea2646970667358221220ebb0be08ed8e07ebc25fbe689a70099c8c66680ed377c5669010a4359db9a5d364736f6c634300080b0033",
}

// StoreABI is the input ABI used to generate the binding from.
// Deprecated: Use StoreMetaData.ABI instead.
var StoreABI = StoreMetaData.ABI

// StoreBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use StoreMetaData.Bin instead.
var StoreBin = StoreMetaData.Bin

// DeployStore deploys a new Ethereum contract, binding an instance of Store to it.
func DeployStore(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Store, error) {
	parsed, err := StoreMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(StoreBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Store{StoreCaller: StoreCaller{contract: contract}, StoreTransactor: StoreTransactor{contract: contract}, StoreFilterer: StoreFilterer{contract: contract}}, nil
}

// Store is an auto generated Go binding around an Ethereum contract.
type Store struct {
	StoreCaller     // Read-only binding to the contract
	StoreTransactor // Write-only binding to the contract
	StoreFilterer   // Log filterer for contract events
}

// StoreCaller is an auto generated read-only Go binding around an Ethereum contract.
type StoreCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StoreTransactor is an auto generated write-only Go binding around an Ethereum contract.
type StoreTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StoreFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type StoreFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StoreSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type StoreSession struct {
	Contract     *Store            // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// StoreCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type StoreCallerSession struct {
	Contract *StoreCaller  // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// StoreTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type StoreTransactorSession struct {
	Contract     *StoreTransactor  // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// StoreRaw is an auto generated low-level Go binding around an Ethereum contract.
type StoreRaw struct {
	Contract *Store // Generic contract binding to access the raw methods on
}

// StoreCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type StoreCallerRaw struct {
	Contract *StoreCaller // Generic read-only contract binding to access the raw methods on
}

// StoreTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type StoreTransactorRaw struct {
	Contract *StoreTransactor // Generic write-only contract binding to access the raw methods on
}

// NewStore creates a new instance of Store, bound to a specific deployed contract.
func NewStore(address common.Address, backend bind.ContractBackend) (*Store, error) {
	contract, err := bindStore(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Store{StoreCaller: StoreCaller{contract: contract}, StoreTransactor: StoreTransactor{contract: contract}, StoreFilterer: StoreFilterer{contract: contract}}, nil
}

// NewStoreCaller creates a new read-only instance of Store, bound to a specific deployed contract.
func NewStoreCaller(address common.Address, caller bind.ContractCaller) (*StoreCaller, error) {
	contract, err := bindStore(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &StoreCaller{contract: contract}, nil
}

// NewStoreTransactor creates a new write-only instance of Store, bound to a specific deployed contract.
func NewStoreTransactor(address common.Address, transactor bind.ContractTransactor) (*StoreTransactor, error) {
	contract, err := bindStore(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &StoreTransactor{contract: contract}, nil
}

// NewStoreFilterer creates a new log filterer instance of Store, bound to a specific deployed contract.
func NewStoreFilterer(address common.Address, filterer bind.ContractFilterer) (*StoreFilterer, error) {
	contract, err := bindStore(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &StoreFilterer{contract: contract}, nil
}

// bindStore binds a generic wrapper to an already deployed contract.
func bindStore(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(StoreABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Store *StoreRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Store.Contract.StoreCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Store *StoreRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Store.Contract.StoreTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Store *StoreRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Store.Contract.StoreTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Store *StoreCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Store.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Store *StoreTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Store.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Store *StoreTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Store.Contract.contract.Transact(opts, method, params...)
}

// Buyer is a free data retrieval call binding the contract method 0x7150d8ae.
//
// Solidity: function buyer() view returns(address)
func (_Store *StoreCaller) Buyer(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Store.contract.Call(opts, &out, "buyer")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Buyer is a free data retrieval call binding the contract method 0x7150d8ae.
//
// Solidity: function buyer() view returns(address)
func (_Store *StoreSession) Buyer() (common.Address, error) {
	return _Store.Contract.Buyer(&_Store.CallOpts)
}

// Buyer is a free data retrieval call binding the contract method 0x7150d8ae.
//
// Solidity: function buyer() view returns(address)
func (_Store *StoreCallerSession) Buyer() (common.Address, error) {
	return _Store.Contract.Buyer(&_Store.CallOpts)
}

// Seller is a free data retrieval call binding the contract method 0x08551a53.
//
// Solidity: function seller() view returns(address)
func (_Store *StoreCaller) Seller(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Store.contract.Call(opts, &out, "seller")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Seller is a free data retrieval call binding the contract method 0x08551a53.
//
// Solidity: function seller() view returns(address)
func (_Store *StoreSession) Seller() (common.Address, error) {
	return _Store.Contract.Seller(&_Store.CallOpts)
}

// Seller is a free data retrieval call binding the contract method 0x08551a53.
//
// Solidity: function seller() view returns(address)
func (_Store *StoreCallerSession) Seller() (common.Address, error) {
	return _Store.Contract.Seller(&_Store.CallOpts)
}

// State is a free data retrieval call binding the contract method 0xc19d93fb.
//
// Solidity: function state() view returns(uint8)
func (_Store *StoreCaller) State(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _Store.contract.Call(opts, &out, "state")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// State is a free data retrieval call binding the contract method 0xc19d93fb.
//
// Solidity: function state() view returns(uint8)
func (_Store *StoreSession) State() (uint8, error) {
	return _Store.Contract.State(&_Store.CallOpts)
}

// State is a free data retrieval call binding the contract method 0xc19d93fb.
//
// Solidity: function state() view returns(uint8)
func (_Store *StoreCallerSession) State() (uint8, error) {
	return _Store.Contract.State(&_Store.CallOpts)
}

// Value is a free data retrieval call binding the contract method 0x3fa4f245.
//
// Solidity: function value() view returns(uint256)
func (_Store *StoreCaller) Value(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Store.contract.Call(opts, &out, "value")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Value is a free data retrieval call binding the contract method 0x3fa4f245.
//
// Solidity: function value() view returns(uint256)
func (_Store *StoreSession) Value() (*big.Int, error) {
	return _Store.Contract.Value(&_Store.CallOpts)
}

// Value is a free data retrieval call binding the contract method 0x3fa4f245.
//
// Solidity: function value() view returns(uint256)
func (_Store *StoreCallerSession) Value() (*big.Int, error) {
	return _Store.Contract.Value(&_Store.CallOpts)
}

// Abort is a paid mutator transaction binding the contract method 0x35a063b4.
//
// Solidity: function abort() returns()
func (_Store *StoreTransactor) Abort(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Store.contract.Transact(opts, "abort")
}

// Abort is a paid mutator transaction binding the contract method 0x35a063b4.
//
// Solidity: function abort() returns()
func (_Store *StoreSession) Abort() (*types.Transaction, error) {
	return _Store.Contract.Abort(&_Store.TransactOpts)
}

// Abort is a paid mutator transaction binding the contract method 0x35a063b4.
//
// Solidity: function abort() returns()
func (_Store *StoreTransactorSession) Abort() (*types.Transaction, error) {
	return _Store.Contract.Abort(&_Store.TransactOpts)
}

// ConfirmPurchase is a paid mutator transaction binding the contract method 0xd6960697.
//
// Solidity: function confirmPurchase() payable returns()
func (_Store *StoreTransactor) ConfirmPurchase(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Store.contract.Transact(opts, "confirmPurchase")
}

// ConfirmPurchase is a paid mutator transaction binding the contract method 0xd6960697.
//
// Solidity: function confirmPurchase() payable returns()
func (_Store *StoreSession) ConfirmPurchase() (*types.Transaction, error) {
	return _Store.Contract.ConfirmPurchase(&_Store.TransactOpts)
}

// ConfirmPurchase is a paid mutator transaction binding the contract method 0xd6960697.
//
// Solidity: function confirmPurchase() payable returns()
func (_Store *StoreTransactorSession) ConfirmPurchase() (*types.Transaction, error) {
	return _Store.Contract.ConfirmPurchase(&_Store.TransactOpts)
}

// ConfirmReceived is a paid mutator transaction binding the contract method 0x73fac6f0.
//
// Solidity: function confirmReceived() returns()
func (_Store *StoreTransactor) ConfirmReceived(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Store.contract.Transact(opts, "confirmReceived")
}

// ConfirmReceived is a paid mutator transaction binding the contract method 0x73fac6f0.
//
// Solidity: function confirmReceived() returns()
func (_Store *StoreSession) ConfirmReceived() (*types.Transaction, error) {
	return _Store.Contract.ConfirmReceived(&_Store.TransactOpts)
}

// ConfirmReceived is a paid mutator transaction binding the contract method 0x73fac6f0.
//
// Solidity: function confirmReceived() returns()
func (_Store *StoreTransactorSession) ConfirmReceived() (*types.Transaction, error) {
	return _Store.Contract.ConfirmReceived(&_Store.TransactOpts)
}

// RefundSeller is a paid mutator transaction binding the contract method 0xc7981b1b.
//
// Solidity: function refundSeller() returns()
func (_Store *StoreTransactor) RefundSeller(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Store.contract.Transact(opts, "refundSeller")
}

// RefundSeller is a paid mutator transaction binding the contract method 0xc7981b1b.
//
// Solidity: function refundSeller() returns()
func (_Store *StoreSession) RefundSeller() (*types.Transaction, error) {
	return _Store.Contract.RefundSeller(&_Store.TransactOpts)
}

// RefundSeller is a paid mutator transaction binding the contract method 0xc7981b1b.
//
// Solidity: function refundSeller() returns()
func (_Store *StoreTransactorSession) RefundSeller() (*types.Transaction, error) {
	return _Store.Contract.RefundSeller(&_Store.TransactOpts)
}

// StoreAbortedIterator is returned from FilterAborted and is used to iterate over the raw logs and unpacked data for Aborted events raised by the Store contract.
type StoreAbortedIterator struct {
	Event *StoreAborted // Event containing the contract specifics and raw log

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
func (it *StoreAbortedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StoreAborted)
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
		it.Event = new(StoreAborted)
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
func (it *StoreAbortedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StoreAbortedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StoreAborted represents a Aborted event raised by the Store contract.
type StoreAborted struct {
	Raw types.Log // Blockchain specific contextual infos
}

// FilterAborted is a free log retrieval operation binding the contract event 0x72c874aeff0b183a56e2b79c71b46e1aed4dee5e09862134b8821ba2fddbf8bf.
//
// Solidity: event Aborted()
func (_Store *StoreFilterer) FilterAborted(opts *bind.FilterOpts) (*StoreAbortedIterator, error) {

	logs, sub, err := _Store.contract.FilterLogs(opts, "Aborted")
	if err != nil {
		return nil, err
	}
	return &StoreAbortedIterator{contract: _Store.contract, event: "Aborted", logs: logs, sub: sub}, nil
}

// WatchAborted is a free log subscription operation binding the contract event 0x72c874aeff0b183a56e2b79c71b46e1aed4dee5e09862134b8821ba2fddbf8bf.
//
// Solidity: event Aborted()
func (_Store *StoreFilterer) WatchAborted(opts *bind.WatchOpts, sink chan<- *StoreAborted) (event.Subscription, error) {

	logs, sub, err := _Store.contract.WatchLogs(opts, "Aborted")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StoreAborted)
				if err := _Store.contract.UnpackLog(event, "Aborted", log); err != nil {
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

// ParseAborted is a log parse operation binding the contract event 0x72c874aeff0b183a56e2b79c71b46e1aed4dee5e09862134b8821ba2fddbf8bf.
//
// Solidity: event Aborted()
func (_Store *StoreFilterer) ParseAborted(log types.Log) (*StoreAborted, error) {
	event := new(StoreAborted)
	if err := _Store.contract.UnpackLog(event, "Aborted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StoreItemReceivedIterator is returned from FilterItemReceived and is used to iterate over the raw logs and unpacked data for ItemReceived events raised by the Store contract.
type StoreItemReceivedIterator struct {
	Event *StoreItemReceived // Event containing the contract specifics and raw log

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
func (it *StoreItemReceivedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StoreItemReceived)
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
		it.Event = new(StoreItemReceived)
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
func (it *StoreItemReceivedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StoreItemReceivedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StoreItemReceived represents a ItemReceived event raised by the Store contract.
type StoreItemReceived struct {
	Raw types.Log // Blockchain specific contextual infos
}

// FilterItemReceived is a free log retrieval operation binding the contract event 0xe89152acd703c9d8c7d28829d443260b411454d45394e7995815140c8cbcbcf7.
//
// Solidity: event ItemReceived()
func (_Store *StoreFilterer) FilterItemReceived(opts *bind.FilterOpts) (*StoreItemReceivedIterator, error) {

	logs, sub, err := _Store.contract.FilterLogs(opts, "ItemReceived")
	if err != nil {
		return nil, err
	}
	return &StoreItemReceivedIterator{contract: _Store.contract, event: "ItemReceived", logs: logs, sub: sub}, nil
}

// WatchItemReceived is a free log subscription operation binding the contract event 0xe89152acd703c9d8c7d28829d443260b411454d45394e7995815140c8cbcbcf7.
//
// Solidity: event ItemReceived()
func (_Store *StoreFilterer) WatchItemReceived(opts *bind.WatchOpts, sink chan<- *StoreItemReceived) (event.Subscription, error) {

	logs, sub, err := _Store.contract.WatchLogs(opts, "ItemReceived")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StoreItemReceived)
				if err := _Store.contract.UnpackLog(event, "ItemReceived", log); err != nil {
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

// ParseItemReceived is a log parse operation binding the contract event 0xe89152acd703c9d8c7d28829d443260b411454d45394e7995815140c8cbcbcf7.
//
// Solidity: event ItemReceived()
func (_Store *StoreFilterer) ParseItemReceived(log types.Log) (*StoreItemReceived, error) {
	event := new(StoreItemReceived)
	if err := _Store.contract.UnpackLog(event, "ItemReceived", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StorePurchaseConfirmedIterator is returned from FilterPurchaseConfirmed and is used to iterate over the raw logs and unpacked data for PurchaseConfirmed events raised by the Store contract.
type StorePurchaseConfirmedIterator struct {
	Event *StorePurchaseConfirmed // Event containing the contract specifics and raw log

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
func (it *StorePurchaseConfirmedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StorePurchaseConfirmed)
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
		it.Event = new(StorePurchaseConfirmed)
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
func (it *StorePurchaseConfirmedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StorePurchaseConfirmedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StorePurchaseConfirmed represents a PurchaseConfirmed event raised by the Store contract.
type StorePurchaseConfirmed struct {
	Raw types.Log // Blockchain specific contextual infos
}

// FilterPurchaseConfirmed is a free log retrieval operation binding the contract event 0xd5d55c8a68912e9a110618df8d5e2e83b8d83211c57a8ddd1203df92885dc881.
//
// Solidity: event PurchaseConfirmed()
func (_Store *StoreFilterer) FilterPurchaseConfirmed(opts *bind.FilterOpts) (*StorePurchaseConfirmedIterator, error) {

	logs, sub, err := _Store.contract.FilterLogs(opts, "PurchaseConfirmed")
	if err != nil {
		return nil, err
	}
	return &StorePurchaseConfirmedIterator{contract: _Store.contract, event: "PurchaseConfirmed", logs: logs, sub: sub}, nil
}

// WatchPurchaseConfirmed is a free log subscription operation binding the contract event 0xd5d55c8a68912e9a110618df8d5e2e83b8d83211c57a8ddd1203df92885dc881.
//
// Solidity: event PurchaseConfirmed()
func (_Store *StoreFilterer) WatchPurchaseConfirmed(opts *bind.WatchOpts, sink chan<- *StorePurchaseConfirmed) (event.Subscription, error) {

	logs, sub, err := _Store.contract.WatchLogs(opts, "PurchaseConfirmed")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StorePurchaseConfirmed)
				if err := _Store.contract.UnpackLog(event, "PurchaseConfirmed", log); err != nil {
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

// ParsePurchaseConfirmed is a log parse operation binding the contract event 0xd5d55c8a68912e9a110618df8d5e2e83b8d83211c57a8ddd1203df92885dc881.
//
// Solidity: event PurchaseConfirmed()
func (_Store *StoreFilterer) ParsePurchaseConfirmed(log types.Log) (*StorePurchaseConfirmed, error) {
	event := new(StorePurchaseConfirmed)
	if err := _Store.contract.UnpackLog(event, "PurchaseConfirmed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StoreSellerRefundedIterator is returned from FilterSellerRefunded and is used to iterate over the raw logs and unpacked data for SellerRefunded events raised by the Store contract.
type StoreSellerRefundedIterator struct {
	Event *StoreSellerRefunded // Event containing the contract specifics and raw log

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
func (it *StoreSellerRefundedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StoreSellerRefunded)
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
		it.Event = new(StoreSellerRefunded)
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
func (it *StoreSellerRefundedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StoreSellerRefundedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StoreSellerRefunded represents a SellerRefunded event raised by the Store contract.
type StoreSellerRefunded struct {
	Raw types.Log // Blockchain specific contextual infos
}

// FilterSellerRefunded is a free log retrieval operation binding the contract event 0xfda69c32bcfdba840a167777906b173b607eb8b4d8853b97a80d26e613d858db.
//
// Solidity: event SellerRefunded()
func (_Store *StoreFilterer) FilterSellerRefunded(opts *bind.FilterOpts) (*StoreSellerRefundedIterator, error) {

	logs, sub, err := _Store.contract.FilterLogs(opts, "SellerRefunded")
	if err != nil {
		return nil, err
	}
	return &StoreSellerRefundedIterator{contract: _Store.contract, event: "SellerRefunded", logs: logs, sub: sub}, nil
}

// WatchSellerRefunded is a free log subscription operation binding the contract event 0xfda69c32bcfdba840a167777906b173b607eb8b4d8853b97a80d26e613d858db.
//
// Solidity: event SellerRefunded()
func (_Store *StoreFilterer) WatchSellerRefunded(opts *bind.WatchOpts, sink chan<- *StoreSellerRefunded) (event.Subscription, error) {

	logs, sub, err := _Store.contract.WatchLogs(opts, "SellerRefunded")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StoreSellerRefunded)
				if err := _Store.contract.UnpackLog(event, "SellerRefunded", log); err != nil {
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

// ParseSellerRefunded is a log parse operation binding the contract event 0xfda69c32bcfdba840a167777906b173b607eb8b4d8853b97a80d26e613d858db.
//
// Solidity: event SellerRefunded()
func (_Store *StoreFilterer) ParseSellerRefunded(log types.Log) (*StoreSellerRefunded, error) {
	event := new(StoreSellerRefunded)
	if err := _Store.contract.UnpackLog(event, "SellerRefunded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
