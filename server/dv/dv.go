// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package dv

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

// DaoVinciABI is the input ABI used to generate the binding from.
const DaoVinciABI = "[{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"address\"}],\"name\":\"balances\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"isOwner\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"payee\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"RewardDistributed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"payee\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"BalanceWithdrawn\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"constant\":false,\"inputs\":[],\"name\":\"initialize\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"initialize\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_payees\",\"type\":\"address[]\"},{\"name\":\"_amounts\",\"type\":\"uint256[]\"}],\"name\":\"distributeRewards\",\"outputs\":[],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"withdraw\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"usersAddress\",\"type\":\"address\"}],\"name\":\"withdrawOnBehalf\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// DaoVinci is an auto generated Go binding around an Ethereum contract.
type DaoVinci struct {
	DaoVinciCaller     // Read-only binding to the contract
	DaoVinciTransactor // Write-only binding to the contract
	DaoVinciFilterer   // Log filterer for contract events
}

// DaoVinciCaller is an auto generated read-only Go binding around an Ethereum contract.
type DaoVinciCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DaoVinciTransactor is an auto generated write-only Go binding around an Ethereum contract.
type DaoVinciTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DaoVinciFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type DaoVinciFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DaoVinciSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type DaoVinciSession struct {
	Contract     *DaoVinci         // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// DaoVinciCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type DaoVinciCallerSession struct {
	Contract *DaoVinciCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts   // Call options to use throughout this session
}

// DaoVinciTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type DaoVinciTransactorSession struct {
	Contract     *DaoVinciTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// DaoVinciRaw is an auto generated low-level Go binding around an Ethereum contract.
type DaoVinciRaw struct {
	Contract *DaoVinci // Generic contract binding to access the raw methods on
}

// DaoVinciCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type DaoVinciCallerRaw struct {
	Contract *DaoVinciCaller // Generic read-only contract binding to access the raw methods on
}

// DaoVinciTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type DaoVinciTransactorRaw struct {
	Contract *DaoVinciTransactor // Generic write-only contract binding to access the raw methods on
}

// NewDaoVinci creates a new instance of DaoVinci, bound to a specific deployed contract.
func NewDaoVinci(address common.Address, backend bind.ContractBackend) (*DaoVinci, error) {
	contract, err := bindDaoVinci(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &DaoVinci{DaoVinciCaller: DaoVinciCaller{contract: contract}, DaoVinciTransactor: DaoVinciTransactor{contract: contract}, DaoVinciFilterer: DaoVinciFilterer{contract: contract}}, nil
}

// NewDaoVinciCaller creates a new read-only instance of DaoVinci, bound to a specific deployed contract.
func NewDaoVinciCaller(address common.Address, caller bind.ContractCaller) (*DaoVinciCaller, error) {
	contract, err := bindDaoVinci(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &DaoVinciCaller{contract: contract}, nil
}

// NewDaoVinciTransactor creates a new write-only instance of DaoVinci, bound to a specific deployed contract.
func NewDaoVinciTransactor(address common.Address, transactor bind.ContractTransactor) (*DaoVinciTransactor, error) {
	contract, err := bindDaoVinci(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &DaoVinciTransactor{contract: contract}, nil
}

// NewDaoVinciFilterer creates a new log filterer instance of DaoVinci, bound to a specific deployed contract.
func NewDaoVinciFilterer(address common.Address, filterer bind.ContractFilterer) (*DaoVinciFilterer, error) {
	contract, err := bindDaoVinci(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &DaoVinciFilterer{contract: contract}, nil
}

// bindDaoVinci binds a generic wrapper to an already deployed contract.
func bindDaoVinci(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(DaoVinciABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_DaoVinci *DaoVinciRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _DaoVinci.Contract.DaoVinciCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_DaoVinci *DaoVinciRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _DaoVinci.Contract.DaoVinciTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_DaoVinci *DaoVinciRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _DaoVinci.Contract.DaoVinciTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_DaoVinci *DaoVinciCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _DaoVinci.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_DaoVinci *DaoVinciTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _DaoVinci.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_DaoVinci *DaoVinciTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _DaoVinci.Contract.contract.Transact(opts, method, params...)
}

// Balances is a free data retrieval call binding the contract method 0x27e235e3.
//
// Solidity: function balances(address ) constant returns(uint256)
func (_DaoVinci *DaoVinciCaller) Balances(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _DaoVinci.contract.Call(opts, out, "balances", arg0)
	return *ret0, err
}

// Balances is a free data retrieval call binding the contract method 0x27e235e3.
//
// Solidity: function balances(address ) constant returns(uint256)
func (_DaoVinci *DaoVinciSession) Balances(arg0 common.Address) (*big.Int, error) {
	return _DaoVinci.Contract.Balances(&_DaoVinci.CallOpts, arg0)
}

// Balances is a free data retrieval call binding the contract method 0x27e235e3.
//
// Solidity: function balances(address ) constant returns(uint256)
func (_DaoVinci *DaoVinciCallerSession) Balances(arg0 common.Address) (*big.Int, error) {
	return _DaoVinci.Contract.Balances(&_DaoVinci.CallOpts, arg0)
}

// IsOwner is a free data retrieval call binding the contract method 0x8f32d59b.
//
// Solidity: function isOwner() constant returns(bool)
func (_DaoVinci *DaoVinciCaller) IsOwner(opts *bind.CallOpts) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _DaoVinci.contract.Call(opts, out, "isOwner")
	return *ret0, err
}

// IsOwner is a free data retrieval call binding the contract method 0x8f32d59b.
//
// Solidity: function isOwner() constant returns(bool)
func (_DaoVinci *DaoVinciSession) IsOwner() (bool, error) {
	return _DaoVinci.Contract.IsOwner(&_DaoVinci.CallOpts)
}

// IsOwner is a free data retrieval call binding the contract method 0x8f32d59b.
//
// Solidity: function isOwner() constant returns(bool)
func (_DaoVinci *DaoVinciCallerSession) IsOwner() (bool, error) {
	return _DaoVinci.Contract.IsOwner(&_DaoVinci.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_DaoVinci *DaoVinciCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _DaoVinci.contract.Call(opts, out, "owner")
	return *ret0, err
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_DaoVinci *DaoVinciSession) Owner() (common.Address, error) {
	return _DaoVinci.Contract.Owner(&_DaoVinci.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_DaoVinci *DaoVinciCallerSession) Owner() (common.Address, error) {
	return _DaoVinci.Contract.Owner(&_DaoVinci.CallOpts)
}

// DistributeRewards is a paid mutator transaction binding the contract method 0x143ba4f3.
//
// Solidity: function distributeRewards(address[] _payees, uint256[] _amounts) returns()
func (_DaoVinci *DaoVinciTransactor) DistributeRewards(opts *bind.TransactOpts, _payees []common.Address, _amounts []*big.Int) (*types.Transaction, error) {
	return _DaoVinci.contract.Transact(opts, "distributeRewards", _payees, _amounts)
}

// DistributeRewards is a paid mutator transaction binding the contract method 0x143ba4f3.
//
// Solidity: function distributeRewards(address[] _payees, uint256[] _amounts) returns()
func (_DaoVinci *DaoVinciSession) DistributeRewards(_payees []common.Address, _amounts []*big.Int) (*types.Transaction, error) {
	return _DaoVinci.Contract.DistributeRewards(&_DaoVinci.TransactOpts, _payees, _amounts)
}

// DistributeRewards is a paid mutator transaction binding the contract method 0x143ba4f3.
//
// Solidity: function distributeRewards(address[] _payees, uint256[] _amounts) returns()
func (_DaoVinci *DaoVinciTransactorSession) DistributeRewards(_payees []common.Address, _amounts []*big.Int) (*types.Transaction, error) {
	return _DaoVinci.Contract.DistributeRewards(&_DaoVinci.TransactOpts, _payees, _amounts)
}

// Initialize is a paid mutator transaction binding the contract method 0xc4d66de8.
//
// Solidity: function initialize(address sender) returns()
func (_DaoVinci *DaoVinciTransactor) Initialize(opts *bind.TransactOpts, sender common.Address) (*types.Transaction, error) {
	return _DaoVinci.contract.Transact(opts, "initialize", sender)
}

// Initialize is a paid mutator transaction binding the contract method 0xc4d66de8.
//
// Solidity: function initialize(address sender) returns()
func (_DaoVinci *DaoVinciSession) Initialize(sender common.Address) (*types.Transaction, error) {
	return _DaoVinci.Contract.Initialize(&_DaoVinci.TransactOpts, sender)
}

// Initialize is a paid mutator transaction binding the contract method 0xc4d66de8.
//
// Solidity: function initialize(address sender) returns()
func (_DaoVinci *DaoVinciTransactorSession) Initialize(sender common.Address) (*types.Transaction, error) {
	return _DaoVinci.Contract.Initialize(&_DaoVinci.TransactOpts, sender)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_DaoVinci *DaoVinciTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _DaoVinci.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_DaoVinci *DaoVinciSession) RenounceOwnership() (*types.Transaction, error) {
	return _DaoVinci.Contract.RenounceOwnership(&_DaoVinci.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_DaoVinci *DaoVinciTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _DaoVinci.Contract.RenounceOwnership(&_DaoVinci.TransactOpts)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_DaoVinci *DaoVinciTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _DaoVinci.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_DaoVinci *DaoVinciSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _DaoVinci.Contract.TransferOwnership(&_DaoVinci.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_DaoVinci *DaoVinciTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _DaoVinci.Contract.TransferOwnership(&_DaoVinci.TransactOpts, newOwner)
}

// Withdraw is a paid mutator transaction binding the contract method 0x3ccfd60b.
//
// Solidity: function withdraw() returns()
func (_DaoVinci *DaoVinciTransactor) Withdraw(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _DaoVinci.contract.Transact(opts, "withdraw")
}

// Withdraw is a paid mutator transaction binding the contract method 0x3ccfd60b.
//
// Solidity: function withdraw() returns()
func (_DaoVinci *DaoVinciSession) Withdraw() (*types.Transaction, error) {
	return _DaoVinci.Contract.Withdraw(&_DaoVinci.TransactOpts)
}

// Withdraw is a paid mutator transaction binding the contract method 0x3ccfd60b.
//
// Solidity: function withdraw() returns()
func (_DaoVinci *DaoVinciTransactorSession) Withdraw() (*types.Transaction, error) {
	return _DaoVinci.Contract.Withdraw(&_DaoVinci.TransactOpts)
}

// WithdrawOnBehalf is a paid mutator transaction binding the contract method 0xa22b97f5.
//
// Solidity: function withdrawOnBehalf(address usersAddress) returns()
func (_DaoVinci *DaoVinciTransactor) WithdrawOnBehalf(opts *bind.TransactOpts, usersAddress common.Address) (*types.Transaction, error) {
	return _DaoVinci.contract.Transact(opts, "withdrawOnBehalf", usersAddress)
}

// WithdrawOnBehalf is a paid mutator transaction binding the contract method 0xa22b97f5.
//
// Solidity: function withdrawOnBehalf(address usersAddress) returns()
func (_DaoVinci *DaoVinciSession) WithdrawOnBehalf(usersAddress common.Address) (*types.Transaction, error) {
	return _DaoVinci.Contract.WithdrawOnBehalf(&_DaoVinci.TransactOpts, usersAddress)
}

// WithdrawOnBehalf is a paid mutator transaction binding the contract method 0xa22b97f5.
//
// Solidity: function withdrawOnBehalf(address usersAddress) returns()
func (_DaoVinci *DaoVinciTransactorSession) WithdrawOnBehalf(usersAddress common.Address) (*types.Transaction, error) {
	return _DaoVinci.Contract.WithdrawOnBehalf(&_DaoVinci.TransactOpts, usersAddress)
}

// DaoVinciBalanceWithdrawnIterator is returned from FilterBalanceWithdrawn and is used to iterate over the raw logs and unpacked data for BalanceWithdrawn events raised by the DaoVinci contract.
type DaoVinciBalanceWithdrawnIterator struct {
	Event *DaoVinciBalanceWithdrawn // Event containing the contract specifics and raw log

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
func (it *DaoVinciBalanceWithdrawnIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DaoVinciBalanceWithdrawn)
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
		it.Event = new(DaoVinciBalanceWithdrawn)
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
func (it *DaoVinciBalanceWithdrawnIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DaoVinciBalanceWithdrawnIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DaoVinciBalanceWithdrawn represents a BalanceWithdrawn event raised by the DaoVinci contract.
type DaoVinciBalanceWithdrawn struct {
	Payee  common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterBalanceWithdrawn is a free log retrieval operation binding the contract event 0xddc398b321237a8d40ac914388309c2f52a08c134e4dc4ce61e32f57cb7d80f1.
//
// Solidity: event BalanceWithdrawn(address payee, uint256 amount)
func (_DaoVinci *DaoVinciFilterer) FilterBalanceWithdrawn(opts *bind.FilterOpts) (*DaoVinciBalanceWithdrawnIterator, error) {

	logs, sub, err := _DaoVinci.contract.FilterLogs(opts, "BalanceWithdrawn")
	if err != nil {
		return nil, err
	}
	return &DaoVinciBalanceWithdrawnIterator{contract: _DaoVinci.contract, event: "BalanceWithdrawn", logs: logs, sub: sub}, nil
}

// WatchBalanceWithdrawn is a free log subscription operation binding the contract event 0xddc398b321237a8d40ac914388309c2f52a08c134e4dc4ce61e32f57cb7d80f1.
//
// Solidity: event BalanceWithdrawn(address payee, uint256 amount)
func (_DaoVinci *DaoVinciFilterer) WatchBalanceWithdrawn(opts *bind.WatchOpts, sink chan<- *DaoVinciBalanceWithdrawn) (event.Subscription, error) {

	logs, sub, err := _DaoVinci.contract.WatchLogs(opts, "BalanceWithdrawn")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DaoVinciBalanceWithdrawn)
				if err := _DaoVinci.contract.UnpackLog(event, "BalanceWithdrawn", log); err != nil {
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

// DaoVinciOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the DaoVinci contract.
type DaoVinciOwnershipTransferredIterator struct {
	Event *DaoVinciOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *DaoVinciOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DaoVinciOwnershipTransferred)
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
		it.Event = new(DaoVinciOwnershipTransferred)
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
func (it *DaoVinciOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DaoVinciOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DaoVinciOwnershipTransferred represents a OwnershipTransferred event raised by the DaoVinci contract.
type DaoVinciOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_DaoVinci *DaoVinciFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*DaoVinciOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _DaoVinci.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &DaoVinciOwnershipTransferredIterator{contract: _DaoVinci.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_DaoVinci *DaoVinciFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *DaoVinciOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _DaoVinci.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DaoVinciOwnershipTransferred)
				if err := _DaoVinci.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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

// DaoVinciRewardDistributedIterator is returned from FilterRewardDistributed and is used to iterate over the raw logs and unpacked data for RewardDistributed events raised by the DaoVinci contract.
type DaoVinciRewardDistributedIterator struct {
	Event *DaoVinciRewardDistributed // Event containing the contract specifics and raw log

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
func (it *DaoVinciRewardDistributedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DaoVinciRewardDistributed)
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
		it.Event = new(DaoVinciRewardDistributed)
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
func (it *DaoVinciRewardDistributedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DaoVinciRewardDistributedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DaoVinciRewardDistributed represents a RewardDistributed event raised by the DaoVinci contract.
type DaoVinciRewardDistributed struct {
	Payee  common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterRewardDistributed is a free log retrieval operation binding the contract event 0xe34918ff1c7084970068b53fd71ad6d8b04e9f15d3886cbf006443e6cdc52ea6.
//
// Solidity: event RewardDistributed(address payee, uint256 amount)
func (_DaoVinci *DaoVinciFilterer) FilterRewardDistributed(opts *bind.FilterOpts) (*DaoVinciRewardDistributedIterator, error) {

	logs, sub, err := _DaoVinci.contract.FilterLogs(opts, "RewardDistributed")
	if err != nil {
		return nil, err
	}
	return &DaoVinciRewardDistributedIterator{contract: _DaoVinci.contract, event: "RewardDistributed", logs: logs, sub: sub}, nil
}

// WatchRewardDistributed is a free log subscription operation binding the contract event 0xe34918ff1c7084970068b53fd71ad6d8b04e9f15d3886cbf006443e6cdc52ea6.
//
// Solidity: event RewardDistributed(address payee, uint256 amount)
func (_DaoVinci *DaoVinciFilterer) WatchRewardDistributed(opts *bind.WatchOpts, sink chan<- *DaoVinciRewardDistributed) (event.Subscription, error) {

	logs, sub, err := _DaoVinci.contract.WatchLogs(opts, "RewardDistributed")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DaoVinciRewardDistributed)
				if err := _DaoVinci.contract.UnpackLog(event, "RewardDistributed", log); err != nil {
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
