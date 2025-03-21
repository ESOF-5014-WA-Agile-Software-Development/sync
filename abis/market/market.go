// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package market

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

// MarketMetaData contains all meta data concerning the Market contract.
var MarketMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"offerId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"seller\",\"type\":\"address\"}],\"name\":\"OfferCancelled\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"offerId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"seller\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"pricePerUnit\",\"type\":\"uint256\"}],\"name\":\"OfferCreated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"offerId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"seller\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"buyer\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"totalPrice\",\"type\":\"uint256\"}],\"name\":\"Purchased\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"offerId\",\"type\":\"uint256\"}],\"name\":\"cancelOffer\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_pricePerUnit\",\"type\":\"uint256\"}],\"name\":\"createOffer\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"offerCounter\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"offers\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"seller\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"pricePerUnit\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"isAvailable\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"offerId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"purchaseAmount\",\"type\":\"uint256\"}],\"name\":\"purchase\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"stateMutability\":\"payable\",\"type\":\"receive\"}]",
}

// MarketABI is the input ABI used to generate the binding from.
// Deprecated: Use MarketMetaData.ABI instead.
var MarketABI = MarketMetaData.ABI

// Market is an auto generated Go binding around an Ethereum contract.
type Market struct {
	MarketCaller     // Read-only binding to the contract
	MarketTransactor // Write-only binding to the contract
	MarketFilterer   // Log filterer for contract events
}

// MarketCaller is an auto generated read-only Go binding around an Ethereum contract.
type MarketCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MarketTransactor is an auto generated write-only Go binding around an Ethereum contract.
type MarketTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MarketFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type MarketFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MarketSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type MarketSession struct {
	Contract     *Market           // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// MarketCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type MarketCallerSession struct {
	Contract *MarketCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// MarketTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type MarketTransactorSession struct {
	Contract     *MarketTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// MarketRaw is an auto generated low-level Go binding around an Ethereum contract.
type MarketRaw struct {
	Contract *Market // Generic contract binding to access the raw methods on
}

// MarketCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type MarketCallerRaw struct {
	Contract *MarketCaller // Generic read-only contract binding to access the raw methods on
}

// MarketTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type MarketTransactorRaw struct {
	Contract *MarketTransactor // Generic write-only contract binding to access the raw methods on
}

// NewMarket creates a new instance of Market, bound to a specific deployed contract.
func NewMarket(address common.Address, backend bind.ContractBackend) (*Market, error) {
	contract, err := bindMarket(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Market{MarketCaller: MarketCaller{contract: contract}, MarketTransactor: MarketTransactor{contract: contract}, MarketFilterer: MarketFilterer{contract: contract}}, nil
}

// NewMarketCaller creates a new read-only instance of Market, bound to a specific deployed contract.
func NewMarketCaller(address common.Address, caller bind.ContractCaller) (*MarketCaller, error) {
	contract, err := bindMarket(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &MarketCaller{contract: contract}, nil
}

// NewMarketTransactor creates a new write-only instance of Market, bound to a specific deployed contract.
func NewMarketTransactor(address common.Address, transactor bind.ContractTransactor) (*MarketTransactor, error) {
	contract, err := bindMarket(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &MarketTransactor{contract: contract}, nil
}

// NewMarketFilterer creates a new log filterer instance of Market, bound to a specific deployed contract.
func NewMarketFilterer(address common.Address, filterer bind.ContractFilterer) (*MarketFilterer, error) {
	contract, err := bindMarket(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &MarketFilterer{contract: contract}, nil
}

// bindMarket binds a generic wrapper to an already deployed contract.
func bindMarket(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := MarketMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Market *MarketRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Market.Contract.MarketCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Market *MarketRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Market.Contract.MarketTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Market *MarketRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Market.Contract.MarketTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Market *MarketCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Market.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Market *MarketTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Market.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Market *MarketTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Market.Contract.contract.Transact(opts, method, params...)
}

// OfferCounter is a free data retrieval call binding the contract method 0xb7a3c1da.
//
// Solidity: function offerCounter() view returns(uint256)
func (_Market *MarketCaller) OfferCounter(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Market.contract.Call(opts, &out, "offerCounter")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// OfferCounter is a free data retrieval call binding the contract method 0xb7a3c1da.
//
// Solidity: function offerCounter() view returns(uint256)
func (_Market *MarketSession) OfferCounter() (*big.Int, error) {
	return _Market.Contract.OfferCounter(&_Market.CallOpts)
}

// OfferCounter is a free data retrieval call binding the contract method 0xb7a3c1da.
//
// Solidity: function offerCounter() view returns(uint256)
func (_Market *MarketCallerSession) OfferCounter() (*big.Int, error) {
	return _Market.Contract.OfferCounter(&_Market.CallOpts)
}

// Offers is a free data retrieval call binding the contract method 0x8a72ea6a.
//
// Solidity: function offers(uint256 ) view returns(address seller, uint256 amount, uint256 pricePerUnit, bool isAvailable)
func (_Market *MarketCaller) Offers(opts *bind.CallOpts, arg0 *big.Int) (struct {
	Seller       common.Address
	Amount       *big.Int
	PricePerUnit *big.Int
	IsAvailable  bool
}, error) {
	var out []interface{}
	err := _Market.contract.Call(opts, &out, "offers", arg0)

	outstruct := new(struct {
		Seller       common.Address
		Amount       *big.Int
		PricePerUnit *big.Int
		IsAvailable  bool
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Seller = *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	outstruct.Amount = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.PricePerUnit = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	outstruct.IsAvailable = *abi.ConvertType(out[3], new(bool)).(*bool)

	return *outstruct, err

}

// Offers is a free data retrieval call binding the contract method 0x8a72ea6a.
//
// Solidity: function offers(uint256 ) view returns(address seller, uint256 amount, uint256 pricePerUnit, bool isAvailable)
func (_Market *MarketSession) Offers(arg0 *big.Int) (struct {
	Seller       common.Address
	Amount       *big.Int
	PricePerUnit *big.Int
	IsAvailable  bool
}, error) {
	return _Market.Contract.Offers(&_Market.CallOpts, arg0)
}

// Offers is a free data retrieval call binding the contract method 0x8a72ea6a.
//
// Solidity: function offers(uint256 ) view returns(address seller, uint256 amount, uint256 pricePerUnit, bool isAvailable)
func (_Market *MarketCallerSession) Offers(arg0 *big.Int) (struct {
	Seller       common.Address
	Amount       *big.Int
	PricePerUnit *big.Int
	IsAvailable  bool
}, error) {
	return _Market.Contract.Offers(&_Market.CallOpts, arg0)
}

// CancelOffer is a paid mutator transaction binding the contract method 0xef706adf.
//
// Solidity: function cancelOffer(uint256 offerId) returns()
func (_Market *MarketTransactor) CancelOffer(opts *bind.TransactOpts, offerId *big.Int) (*types.Transaction, error) {
	return _Market.contract.Transact(opts, "cancelOffer", offerId)
}

// CancelOffer is a paid mutator transaction binding the contract method 0xef706adf.
//
// Solidity: function cancelOffer(uint256 offerId) returns()
func (_Market *MarketSession) CancelOffer(offerId *big.Int) (*types.Transaction, error) {
	return _Market.Contract.CancelOffer(&_Market.TransactOpts, offerId)
}

// CancelOffer is a paid mutator transaction binding the contract method 0xef706adf.
//
// Solidity: function cancelOffer(uint256 offerId) returns()
func (_Market *MarketTransactorSession) CancelOffer(offerId *big.Int) (*types.Transaction, error) {
	return _Market.Contract.CancelOffer(&_Market.TransactOpts, offerId)
}

// CreateOffer is a paid mutator transaction binding the contract method 0x1afed875.
//
// Solidity: function createOffer(uint256 _amount, uint256 _pricePerUnit) returns()
func (_Market *MarketTransactor) CreateOffer(opts *bind.TransactOpts, _amount *big.Int, _pricePerUnit *big.Int) (*types.Transaction, error) {
	return _Market.contract.Transact(opts, "createOffer", _amount, _pricePerUnit)
}

// CreateOffer is a paid mutator transaction binding the contract method 0x1afed875.
//
// Solidity: function createOffer(uint256 _amount, uint256 _pricePerUnit) returns()
func (_Market *MarketSession) CreateOffer(_amount *big.Int, _pricePerUnit *big.Int) (*types.Transaction, error) {
	return _Market.Contract.CreateOffer(&_Market.TransactOpts, _amount, _pricePerUnit)
}

// CreateOffer is a paid mutator transaction binding the contract method 0x1afed875.
//
// Solidity: function createOffer(uint256 _amount, uint256 _pricePerUnit) returns()
func (_Market *MarketTransactorSession) CreateOffer(_amount *big.Int, _pricePerUnit *big.Int) (*types.Transaction, error) {
	return _Market.Contract.CreateOffer(&_Market.TransactOpts, _amount, _pricePerUnit)
}

// Purchase is a paid mutator transaction binding the contract method 0x70876c98.
//
// Solidity: function purchase(uint256 offerId, uint256 purchaseAmount) payable returns()
func (_Market *MarketTransactor) Purchase(opts *bind.TransactOpts, offerId *big.Int, purchaseAmount *big.Int) (*types.Transaction, error) {
	return _Market.contract.Transact(opts, "purchase", offerId, purchaseAmount)
}

// Purchase is a paid mutator transaction binding the contract method 0x70876c98.
//
// Solidity: function purchase(uint256 offerId, uint256 purchaseAmount) payable returns()
func (_Market *MarketSession) Purchase(offerId *big.Int, purchaseAmount *big.Int) (*types.Transaction, error) {
	return _Market.Contract.Purchase(&_Market.TransactOpts, offerId, purchaseAmount)
}

// Purchase is a paid mutator transaction binding the contract method 0x70876c98.
//
// Solidity: function purchase(uint256 offerId, uint256 purchaseAmount) payable returns()
func (_Market *MarketTransactorSession) Purchase(offerId *big.Int, purchaseAmount *big.Int) (*types.Transaction, error) {
	return _Market.Contract.Purchase(&_Market.TransactOpts, offerId, purchaseAmount)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_Market *MarketTransactor) Receive(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Market.contract.RawTransact(opts, nil) // calldata is disallowed for receive function
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_Market *MarketSession) Receive() (*types.Transaction, error) {
	return _Market.Contract.Receive(&_Market.TransactOpts)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_Market *MarketTransactorSession) Receive() (*types.Transaction, error) {
	return _Market.Contract.Receive(&_Market.TransactOpts)
}

// MarketOfferCancelledIterator is returned from FilterOfferCancelled and is used to iterate over the raw logs and unpacked data for OfferCancelled events raised by the Market contract.
type MarketOfferCancelledIterator struct {
	Event *MarketOfferCancelled // Event containing the contract specifics and raw log

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
func (it *MarketOfferCancelledIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MarketOfferCancelled)
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
		it.Event = new(MarketOfferCancelled)
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
func (it *MarketOfferCancelledIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MarketOfferCancelledIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MarketOfferCancelled represents a OfferCancelled event raised by the Market contract.
type MarketOfferCancelled struct {
	OfferId *big.Int
	Seller  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterOfferCancelled is a free log retrieval operation binding the contract event 0x1f51377b3e685a0e2419f9bb4ba7c07ec54936353ba3d0fb3c6538dab6766222.
//
// Solidity: event OfferCancelled(uint256 indexed offerId, address indexed seller)
func (_Market *MarketFilterer) FilterOfferCancelled(opts *bind.FilterOpts, offerId []*big.Int, seller []common.Address) (*MarketOfferCancelledIterator, error) {

	var offerIdRule []interface{}
	for _, offerIdItem := range offerId {
		offerIdRule = append(offerIdRule, offerIdItem)
	}
	var sellerRule []interface{}
	for _, sellerItem := range seller {
		sellerRule = append(sellerRule, sellerItem)
	}

	logs, sub, err := _Market.contract.FilterLogs(opts, "OfferCancelled", offerIdRule, sellerRule)
	if err != nil {
		return nil, err
	}
	return &MarketOfferCancelledIterator{contract: _Market.contract, event: "OfferCancelled", logs: logs, sub: sub}, nil
}

// WatchOfferCancelled is a free log subscription operation binding the contract event 0x1f51377b3e685a0e2419f9bb4ba7c07ec54936353ba3d0fb3c6538dab6766222.
//
// Solidity: event OfferCancelled(uint256 indexed offerId, address indexed seller)
func (_Market *MarketFilterer) WatchOfferCancelled(opts *bind.WatchOpts, sink chan<- *MarketOfferCancelled, offerId []*big.Int, seller []common.Address) (event.Subscription, error) {

	var offerIdRule []interface{}
	for _, offerIdItem := range offerId {
		offerIdRule = append(offerIdRule, offerIdItem)
	}
	var sellerRule []interface{}
	for _, sellerItem := range seller {
		sellerRule = append(sellerRule, sellerItem)
	}

	logs, sub, err := _Market.contract.WatchLogs(opts, "OfferCancelled", offerIdRule, sellerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MarketOfferCancelled)
				if err := _Market.contract.UnpackLog(event, "OfferCancelled", log); err != nil {
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

// ParseOfferCancelled is a log parse operation binding the contract event 0x1f51377b3e685a0e2419f9bb4ba7c07ec54936353ba3d0fb3c6538dab6766222.
//
// Solidity: event OfferCancelled(uint256 indexed offerId, address indexed seller)
func (_Market *MarketFilterer) ParseOfferCancelled(log types.Log) (*MarketOfferCancelled, error) {
	event := new(MarketOfferCancelled)
	if err := _Market.contract.UnpackLog(event, "OfferCancelled", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MarketOfferCreatedIterator is returned from FilterOfferCreated and is used to iterate over the raw logs and unpacked data for OfferCreated events raised by the Market contract.
type MarketOfferCreatedIterator struct {
	Event *MarketOfferCreated // Event containing the contract specifics and raw log

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
func (it *MarketOfferCreatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MarketOfferCreated)
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
		it.Event = new(MarketOfferCreated)
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
func (it *MarketOfferCreatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MarketOfferCreatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MarketOfferCreated represents a OfferCreated event raised by the Market contract.
type MarketOfferCreated struct {
	OfferId      *big.Int
	Seller       common.Address
	Amount       *big.Int
	PricePerUnit *big.Int
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterOfferCreated is a free log retrieval operation binding the contract event 0xbf26a47e2418fff23bee5882ae59627d33e81fa75ab2b1fa08de6cf5f2fa3761.
//
// Solidity: event OfferCreated(uint256 indexed offerId, address indexed seller, uint256 amount, uint256 pricePerUnit)
func (_Market *MarketFilterer) FilterOfferCreated(opts *bind.FilterOpts, offerId []*big.Int, seller []common.Address) (*MarketOfferCreatedIterator, error) {

	var offerIdRule []interface{}
	for _, offerIdItem := range offerId {
		offerIdRule = append(offerIdRule, offerIdItem)
	}
	var sellerRule []interface{}
	for _, sellerItem := range seller {
		sellerRule = append(sellerRule, sellerItem)
	}

	logs, sub, err := _Market.contract.FilterLogs(opts, "OfferCreated", offerIdRule, sellerRule)
	if err != nil {
		return nil, err
	}
	return &MarketOfferCreatedIterator{contract: _Market.contract, event: "OfferCreated", logs: logs, sub: sub}, nil
}

// WatchOfferCreated is a free log subscription operation binding the contract event 0xbf26a47e2418fff23bee5882ae59627d33e81fa75ab2b1fa08de6cf5f2fa3761.
//
// Solidity: event OfferCreated(uint256 indexed offerId, address indexed seller, uint256 amount, uint256 pricePerUnit)
func (_Market *MarketFilterer) WatchOfferCreated(opts *bind.WatchOpts, sink chan<- *MarketOfferCreated, offerId []*big.Int, seller []common.Address) (event.Subscription, error) {

	var offerIdRule []interface{}
	for _, offerIdItem := range offerId {
		offerIdRule = append(offerIdRule, offerIdItem)
	}
	var sellerRule []interface{}
	for _, sellerItem := range seller {
		sellerRule = append(sellerRule, sellerItem)
	}

	logs, sub, err := _Market.contract.WatchLogs(opts, "OfferCreated", offerIdRule, sellerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MarketOfferCreated)
				if err := _Market.contract.UnpackLog(event, "OfferCreated", log); err != nil {
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

// ParseOfferCreated is a log parse operation binding the contract event 0xbf26a47e2418fff23bee5882ae59627d33e81fa75ab2b1fa08de6cf5f2fa3761.
//
// Solidity: event OfferCreated(uint256 indexed offerId, address indexed seller, uint256 amount, uint256 pricePerUnit)
func (_Market *MarketFilterer) ParseOfferCreated(log types.Log) (*MarketOfferCreated, error) {
	event := new(MarketOfferCreated)
	if err := _Market.contract.UnpackLog(event, "OfferCreated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MarketPurchasedIterator is returned from FilterPurchased and is used to iterate over the raw logs and unpacked data for Purchased events raised by the Market contract.
type MarketPurchasedIterator struct {
	Event *MarketPurchased // Event containing the contract specifics and raw log

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
func (it *MarketPurchasedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MarketPurchased)
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
		it.Event = new(MarketPurchased)
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
func (it *MarketPurchasedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MarketPurchasedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MarketPurchased represents a Purchased event raised by the Market contract.
type MarketPurchased struct {
	OfferId    *big.Int
	Seller     common.Address
	Buyer      common.Address
	Amount     *big.Int
	TotalPrice *big.Int
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterPurchased is a free log retrieval operation binding the contract event 0x5455e1f43c674b0e790485adf9e64c058a505bbff2e540eba3a4e612a3a85ec4.
//
// Solidity: event Purchased(uint256 indexed offerId, address indexed seller, address indexed buyer, uint256 amount, uint256 totalPrice)
func (_Market *MarketFilterer) FilterPurchased(opts *bind.FilterOpts, offerId []*big.Int, seller []common.Address, buyer []common.Address) (*MarketPurchasedIterator, error) {

	var offerIdRule []interface{}
	for _, offerIdItem := range offerId {
		offerIdRule = append(offerIdRule, offerIdItem)
	}
	var sellerRule []interface{}
	for _, sellerItem := range seller {
		sellerRule = append(sellerRule, sellerItem)
	}
	var buyerRule []interface{}
	for _, buyerItem := range buyer {
		buyerRule = append(buyerRule, buyerItem)
	}

	logs, sub, err := _Market.contract.FilterLogs(opts, "Purchased", offerIdRule, sellerRule, buyerRule)
	if err != nil {
		return nil, err
	}
	return &MarketPurchasedIterator{contract: _Market.contract, event: "Purchased", logs: logs, sub: sub}, nil
}

// WatchPurchased is a free log subscription operation binding the contract event 0x5455e1f43c674b0e790485adf9e64c058a505bbff2e540eba3a4e612a3a85ec4.
//
// Solidity: event Purchased(uint256 indexed offerId, address indexed seller, address indexed buyer, uint256 amount, uint256 totalPrice)
func (_Market *MarketFilterer) WatchPurchased(opts *bind.WatchOpts, sink chan<- *MarketPurchased, offerId []*big.Int, seller []common.Address, buyer []common.Address) (event.Subscription, error) {

	var offerIdRule []interface{}
	for _, offerIdItem := range offerId {
		offerIdRule = append(offerIdRule, offerIdItem)
	}
	var sellerRule []interface{}
	for _, sellerItem := range seller {
		sellerRule = append(sellerRule, sellerItem)
	}
	var buyerRule []interface{}
	for _, buyerItem := range buyer {
		buyerRule = append(buyerRule, buyerItem)
	}

	logs, sub, err := _Market.contract.WatchLogs(opts, "Purchased", offerIdRule, sellerRule, buyerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MarketPurchased)
				if err := _Market.contract.UnpackLog(event, "Purchased", log); err != nil {
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

// ParsePurchased is a log parse operation binding the contract event 0x5455e1f43c674b0e790485adf9e64c058a505bbff2e540eba3a4e612a3a85ec4.
//
// Solidity: event Purchased(uint256 indexed offerId, address indexed seller, address indexed buyer, uint256 amount, uint256 totalPrice)
func (_Market *MarketFilterer) ParsePurchased(log types.Log) (*MarketPurchased, error) {
	event := new(MarketPurchased)
	if err := _Market.contract.UnpackLog(event, "Purchased", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
