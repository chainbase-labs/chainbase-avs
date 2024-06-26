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

// IAVSStrategyParam is an auto generated low-level Go binding around an user-defined struct.
type IAVSStrategyParam struct {
	Strategy   common.Address
	Multiplier *big.Int
}

// ISignatureUtilsSignatureWithSaltAndExpiry is an auto generated low-level Go binding around an user-defined struct.

// IAVSMetaData contains all meta data concerning the IAVS contract.
var IAVSMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"function\",\"name\":\"canRegister\",\"inputs\":[{\"name\":\"operator\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"operators\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address[]\",\"internalType\":\"address[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"registerOperator\",\"inputs\":[{\"name\":\"operatorSignature\",\"type\":\"tuple\",\"internalType\":\"structISignatureUtils.SignatureWithSaltAndExpiry\",\"components\":[{\"name\":\"signature\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"salt\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"expiry\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"strategyParams\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"tuple[]\",\"internalType\":\"structIAVS.StrategyParam[]\",\"components\":[{\"name\":\"strategy\",\"type\":\"address\",\"internalType\":\"contractIStrategy\"},{\"name\":\"multiplier\",\"type\":\"uint96\",\"internalType\":\"uint96\"}]}],\"stateMutability\":\"view\"},{\"type\":\"event\",\"name\":\"OperatorAdded\",\"inputs\":[{\"name\":\"operator\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OperatorRemoved\",\"inputs\":[{\"name\":\"operator\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false}]",
}

// IAVSABI is the input ABI used to generate the binding from.
// Deprecated: Use IAVSMetaData.ABI instead.
var IAVSABI = IAVSMetaData.ABI

// IAVS is an auto generated Go binding around an Ethereum contract.
type IAVS struct {
	IAVSCaller     // Read-only binding to the contract
	IAVSTransactor // Write-only binding to the contract
	IAVSFilterer   // Log filterer for contract events
}

// IAVSCaller is an auto generated read-only Go binding around an Ethereum contract.
type IAVSCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IAVSTransactor is an auto generated write-only Go binding around an Ethereum contract.
type IAVSTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IAVSFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IAVSFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IAVSSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IAVSSession struct {
	Contract     *IAVS             // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// IAVSCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IAVSCallerSession struct {
	Contract *IAVSCaller   // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// IAVSTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IAVSTransactorSession struct {
	Contract     *IAVSTransactor   // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// IAVSRaw is an auto generated low-level Go binding around an Ethereum contract.
type IAVSRaw struct {
	Contract *IAVS // Generic contract binding to access the raw methods on
}

// IAVSCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IAVSCallerRaw struct {
	Contract *IAVSCaller // Generic read-only contract binding to access the raw methods on
}

// IAVSTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IAVSTransactorRaw struct {
	Contract *IAVSTransactor // Generic write-only contract binding to access the raw methods on
}

// NewIAVS creates a new instance of IAVS, bound to a specific deployed contract.
func NewIAVS(address common.Address, backend bind.ContractBackend) (*IAVS, error) {
	contract, err := bindIAVS(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IAVS{IAVSCaller: IAVSCaller{contract: contract}, IAVSTransactor: IAVSTransactor{contract: contract}, IAVSFilterer: IAVSFilterer{contract: contract}}, nil
}

// NewIAVSCaller creates a new read-only instance of IAVS, bound to a specific deployed contract.
func NewIAVSCaller(address common.Address, caller bind.ContractCaller) (*IAVSCaller, error) {
	contract, err := bindIAVS(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IAVSCaller{contract: contract}, nil
}

// NewIAVSTransactor creates a new write-only instance of IAVS, bound to a specific deployed contract.
func NewIAVSTransactor(address common.Address, transactor bind.ContractTransactor) (*IAVSTransactor, error) {
	contract, err := bindIAVS(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IAVSTransactor{contract: contract}, nil
}

// NewIAVSFilterer creates a new log filterer instance of IAVS, bound to a specific deployed contract.
func NewIAVSFilterer(address common.Address, filterer bind.ContractFilterer) (*IAVSFilterer, error) {
	contract, err := bindIAVS(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IAVSFilterer{contract: contract}, nil
}

// bindIAVS binds a generic wrapper to an already deployed contract.
func bindIAVS(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := IAVSMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IAVS *IAVSRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IAVS.Contract.IAVSCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IAVS *IAVSRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IAVS.Contract.IAVSTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IAVS *IAVSRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IAVS.Contract.IAVSTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IAVS *IAVSCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IAVS.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IAVS *IAVSTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IAVS.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IAVS *IAVSTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IAVS.Contract.contract.Transact(opts, method, params...)
}

// CanRegister is a free data retrieval call binding the contract method 0x320d46d4.
//
// Solidity: function canRegister(address operator) view returns(bool, string)
func (_IAVS *IAVSCaller) CanRegister(opts *bind.CallOpts, operator common.Address) (bool, string, error) {
	var out []interface{}
	err := _IAVS.contract.Call(opts, &out, "canRegister", operator)

	if err != nil {
		return *new(bool), *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)
	out1 := *abi.ConvertType(out[1], new(string)).(*string)

	return out0, out1, err

}

// CanRegister is a free data retrieval call binding the contract method 0x320d46d4.
//
// Solidity: function canRegister(address operator) view returns(bool, string)
func (_IAVS *IAVSSession) CanRegister(operator common.Address) (bool, string, error) {
	return _IAVS.Contract.CanRegister(&_IAVS.CallOpts, operator)
}

// CanRegister is a free data retrieval call binding the contract method 0x320d46d4.
//
// Solidity: function canRegister(address operator) view returns(bool, string)
func (_IAVS *IAVSCallerSession) CanRegister(operator common.Address) (bool, string, error) {
	return _IAVS.Contract.CanRegister(&_IAVS.CallOpts, operator)
}

// Operators is a free data retrieval call binding the contract method 0xe673df8a.
//
// Solidity: function operators() view returns(address[])
func (_IAVS *IAVSCaller) Operators(opts *bind.CallOpts) ([]common.Address, error) {
	var out []interface{}
	err := _IAVS.contract.Call(opts, &out, "operators")

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// Operators is a free data retrieval call binding the contract method 0xe673df8a.
//
// Solidity: function operators() view returns(address[])
func (_IAVS *IAVSSession) Operators() ([]common.Address, error) {
	return _IAVS.Contract.Operators(&_IAVS.CallOpts)
}

// Operators is a free data retrieval call binding the contract method 0xe673df8a.
//
// Solidity: function operators() view returns(address[])
func (_IAVS *IAVSCallerSession) Operators() ([]common.Address, error) {
	return _IAVS.Contract.Operators(&_IAVS.CallOpts)
}

// StrategyParams is a free data retrieval call binding the contract method 0xf57f33d0.
//
// Solidity: function strategyParams() view returns((address,uint96)[])
func (_IAVS *IAVSCaller) StrategyParams(opts *bind.CallOpts) ([]IAVSStrategyParam, error) {
	var out []interface{}
	err := _IAVS.contract.Call(opts, &out, "strategyParams")

	if err != nil {
		return *new([]IAVSStrategyParam), err
	}

	out0 := *abi.ConvertType(out[0], new([]IAVSStrategyParam)).(*[]IAVSStrategyParam)

	return out0, err

}

// StrategyParams is a free data retrieval call binding the contract method 0xf57f33d0.
//
// Solidity: function strategyParams() view returns((address,uint96)[])
func (_IAVS *IAVSSession) StrategyParams() ([]IAVSStrategyParam, error) {
	return _IAVS.Contract.StrategyParams(&_IAVS.CallOpts)
}

// StrategyParams is a free data retrieval call binding the contract method 0xf57f33d0.
//
// Solidity: function strategyParams() view returns((address,uint96)[])
func (_IAVS *IAVSCallerSession) StrategyParams() ([]IAVSStrategyParam, error) {
	return _IAVS.Contract.StrategyParams(&_IAVS.CallOpts)
}

// RegisterOperator is a paid mutator transaction binding the contract method 0x8317781d.
//
// Solidity: function registerOperator((bytes,bytes32,uint256) operatorSignature) returns()
func (_IAVS *IAVSTransactor) RegisterOperator(opts *bind.TransactOpts, operatorSignature ISignatureUtilsSignatureWithSaltAndExpiry) (*types.Transaction, error) {
	return _IAVS.contract.Transact(opts, "registerOperator", operatorSignature)
}

// RegisterOperator is a paid mutator transaction binding the contract method 0x8317781d.
//
// Solidity: function registerOperator((bytes,bytes32,uint256) operatorSignature) returns()
func (_IAVS *IAVSSession) RegisterOperator(operatorSignature ISignatureUtilsSignatureWithSaltAndExpiry) (*types.Transaction, error) {
	return _IAVS.Contract.RegisterOperator(&_IAVS.TransactOpts, operatorSignature)
}

// RegisterOperator is a paid mutator transaction binding the contract method 0x8317781d.
//
// Solidity: function registerOperator((bytes,bytes32,uint256) operatorSignature) returns()
func (_IAVS *IAVSTransactorSession) RegisterOperator(operatorSignature ISignatureUtilsSignatureWithSaltAndExpiry) (*types.Transaction, error) {
	return _IAVS.Contract.RegisterOperator(&_IAVS.TransactOpts, operatorSignature)
}

// IAVSOperatorAddedIterator is returned from FilterOperatorAdded and is used to iterate over the raw logs and unpacked data for OperatorAdded events raised by the IAVS contract.
type IAVSOperatorAddedIterator struct {
	Event *IAVSOperatorAdded // Event containing the contract specifics and raw log

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
func (it *IAVSOperatorAddedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IAVSOperatorAdded)
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
		it.Event = new(IAVSOperatorAdded)
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
func (it *IAVSOperatorAddedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IAVSOperatorAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IAVSOperatorAdded represents a OperatorAdded event raised by the IAVS contract.
type IAVSOperatorAdded struct {
	Operator common.Address
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterOperatorAdded is a free log retrieval operation binding the contract event 0xac6fa858e9350a46cec16539926e0fde25b7629f84b5a72bffaae4df888ae86d.
//
// Solidity: event OperatorAdded(address indexed operator)
func (_IAVS *IAVSFilterer) FilterOperatorAdded(opts *bind.FilterOpts, operator []common.Address) (*IAVSOperatorAddedIterator, error) {

	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _IAVS.contract.FilterLogs(opts, "OperatorAdded", operatorRule)
	if err != nil {
		return nil, err
	}
	return &IAVSOperatorAddedIterator{contract: _IAVS.contract, event: "OperatorAdded", logs: logs, sub: sub}, nil
}

// WatchOperatorAdded is a free log subscription operation binding the contract event 0xac6fa858e9350a46cec16539926e0fde25b7629f84b5a72bffaae4df888ae86d.
//
// Solidity: event OperatorAdded(address indexed operator)
func (_IAVS *IAVSFilterer) WatchOperatorAdded(opts *bind.WatchOpts, sink chan<- *IAVSOperatorAdded, operator []common.Address) (event.Subscription, error) {

	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _IAVS.contract.WatchLogs(opts, "OperatorAdded", operatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IAVSOperatorAdded)
				if err := _IAVS.contract.UnpackLog(event, "OperatorAdded", log); err != nil {
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

// ParseOperatorAdded is a log parse operation binding the contract event 0xac6fa858e9350a46cec16539926e0fde25b7629f84b5a72bffaae4df888ae86d.
//
// Solidity: event OperatorAdded(address indexed operator)
func (_IAVS *IAVSFilterer) ParseOperatorAdded(log types.Log) (*IAVSOperatorAdded, error) {
	event := new(IAVSOperatorAdded)
	if err := _IAVS.contract.UnpackLog(event, "OperatorAdded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IAVSOperatorRemovedIterator is returned from FilterOperatorRemoved and is used to iterate over the raw logs and unpacked data for OperatorRemoved events raised by the IAVS contract.
type IAVSOperatorRemovedIterator struct {
	Event *IAVSOperatorRemoved // Event containing the contract specifics and raw log

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
func (it *IAVSOperatorRemovedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IAVSOperatorRemoved)
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
		it.Event = new(IAVSOperatorRemoved)
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
func (it *IAVSOperatorRemovedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IAVSOperatorRemovedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IAVSOperatorRemoved represents a OperatorRemoved event raised by the IAVS contract.
type IAVSOperatorRemoved struct {
	Operator common.Address
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterOperatorRemoved is a free log retrieval operation binding the contract event 0x80c0b871b97b595b16a7741c1b06fed0c6f6f558639f18ccbce50724325dc40d.
//
// Solidity: event OperatorRemoved(address indexed operator)
func (_IAVS *IAVSFilterer) FilterOperatorRemoved(opts *bind.FilterOpts, operator []common.Address) (*IAVSOperatorRemovedIterator, error) {

	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _IAVS.contract.FilterLogs(opts, "OperatorRemoved", operatorRule)
	if err != nil {
		return nil, err
	}
	return &IAVSOperatorRemovedIterator{contract: _IAVS.contract, event: "OperatorRemoved", logs: logs, sub: sub}, nil
}

// WatchOperatorRemoved is a free log subscription operation binding the contract event 0x80c0b871b97b595b16a7741c1b06fed0c6f6f558639f18ccbce50724325dc40d.
//
// Solidity: event OperatorRemoved(address indexed operator)
func (_IAVS *IAVSFilterer) WatchOperatorRemoved(opts *bind.WatchOpts, sink chan<- *IAVSOperatorRemoved, operator []common.Address) (event.Subscription, error) {

	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _IAVS.contract.WatchLogs(opts, "OperatorRemoved", operatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IAVSOperatorRemoved)
				if err := _IAVS.contract.UnpackLog(event, "OperatorRemoved", log); err != nil {
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

// ParseOperatorRemoved is a log parse operation binding the contract event 0x80c0b871b97b595b16a7741c1b06fed0c6f6f558639f18ccbce50724325dc40d.
//
// Solidity: event OperatorRemoved(address indexed operator)
func (_IAVS *IAVSFilterer) ParseOperatorRemoved(log types.Log) (*IAVSOperatorRemoved, error) {
	event := new(IAVSOperatorRemoved)
	if err := _IAVS.contract.UnpackLog(event, "OperatorRemoved", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
