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

// ERC20MockMetaData contains all meta data concerning the ERC20Mock contract.
var ERC20MockMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"function\",\"name\":\"allowance\",\"inputs\":[{\"name\":\"owner\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"spender\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"approve\",\"inputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"balanceOf\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"mint\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"totalSupply\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"transfer\",\"inputs\":[{\"name\":\"to\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"transferFrom\",\"inputs\":[{\"name\":\"from\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"to\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"event\",\"name\":\"Approval\",\"inputs\":[{\"name\":\"owner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"spender\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"value\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Transfer\",\"inputs\":[{\"name\":\"from\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"to\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"value\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false}]",
	Bin: "0x608060405234801561001057600080fd5b50610528806100206000396000f3fe608060405234801561001057600080fd5b506004361061007d5760003560e01c806340c10f191161005b57806340c10f19146100cf57806370a08231146100e4578063a9059cbb1461010d578063dd62ed3e1461012057600080fd5b8063095ea7b31461008257806318160ddd146100aa57806323b872dd146100bc575b600080fd5b610095610090366004610416565b610159565b60405190151581526020015b60405180910390f35b6002545b6040519081526020016100a1565b6100956100ca366004610440565b610162565b6100e26100dd366004610416565b610179565b005b6100ae6100f236600461047c565b6001600160a01b031660009081526020819052604090205490565b61009561011b366004610416565b610187565b6100ae61012e36600461049e565b6001600160a01b03918216600090815260016020908152604080832093909416825291909152205490565b60015b92915050565b600061016f848484610191565b5060019392505050565b610183828261033b565b5050565b60003361016f8185855b6001600160a01b0383166101fa5760405162461bcd60e51b815260206004820152602560248201527f45524332303a207472616e736665722066726f6d20746865207a65726f206164604482015264647265737360d81b60648201526084015b60405180910390fd5b6001600160a01b03821661025c5760405162461bcd60e51b815260206004820152602360248201527f45524332303a207472616e7366657220746f20746865207a65726f206164647260448201526265737360e81b60648201526084016101f1565b6001600160a01b0383166000908152602081905260409020548111156102d35760405162461bcd60e51b815260206004820152602660248201527f45524332303a207472616e7366657220616d6f756e7420657863656564732062604482015265616c616e636560d01b60648201526084016101f1565b6001600160a01b038381166000818152602081815260408083208054879003905593861680835291849020805486019055925184815290927fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef910160405180910390a3505050565b6001600160a01b0382166103915760405162461bcd60e51b815260206004820152601f60248201527f45524332303a206d696e7420746f20746865207a65726f20616464726573730060448201526064016101f1565b80600260008282546103a391906104d1565b90915550506001600160a01b038216600081815260208181526040808320805486019055518481527fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef910160405180910390a35050565b80356001600160a01b038116811461041157600080fd5b919050565b6000806040838503121561042957600080fd5b610432836103fa565b946020939093013593505050565b60008060006060848603121561045557600080fd5b61045e846103fa565b925061046c602085016103fa565b9150604084013590509250925092565b60006020828403121561048e57600080fd5b610497826103fa565b9392505050565b600080604083850312156104b157600080fd5b6104ba836103fa565b91506104c8602084016103fa565b90509250929050565b8082018082111561015c57634e487b7160e01b600052601160045260246000fdfea264697066735822122064da433f3ffdbe8e26ba04580cd422ddcc96de63727aeff6422f4745a3a2612a64736f6c63430008180033",
}

// ERC20MockABI is the input ABI used to generate the binding from.
// Deprecated: Use ERC20MockMetaData.ABI instead.
var ERC20MockABI = ERC20MockMetaData.ABI

// ERC20MockBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use ERC20MockMetaData.Bin instead.
var ERC20MockBin = ERC20MockMetaData.Bin

// DeployERC20Mock deploys a new Ethereum contract, binding an instance of ERC20Mock to it.
func DeployERC20Mock(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *ERC20Mock, error) {
	parsed, err := ERC20MockMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(ERC20MockBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &ERC20Mock{ERC20MockCaller: ERC20MockCaller{contract: contract}, ERC20MockTransactor: ERC20MockTransactor{contract: contract}, ERC20MockFilterer: ERC20MockFilterer{contract: contract}}, nil
}

// ERC20Mock is an auto generated Go binding around an Ethereum contract.
type ERC20Mock struct {
	ERC20MockCaller     // Read-only binding to the contract
	ERC20MockTransactor // Write-only binding to the contract
	ERC20MockFilterer   // Log filterer for contract events
}

// ERC20MockCaller is an auto generated read-only Go binding around an Ethereum contract.
type ERC20MockCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ERC20MockTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ERC20MockTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ERC20MockFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ERC20MockFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ERC20MockSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ERC20MockSession struct {
	Contract     *ERC20Mock        // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ERC20MockCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ERC20MockCallerSession struct {
	Contract *ERC20MockCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts    // Call options to use throughout this session
}

// ERC20MockTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ERC20MockTransactorSession struct {
	Contract     *ERC20MockTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// ERC20MockRaw is an auto generated low-level Go binding around an Ethereum contract.
type ERC20MockRaw struct {
	Contract *ERC20Mock // Generic contract binding to access the raw methods on
}

// ERC20MockCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ERC20MockCallerRaw struct {
	Contract *ERC20MockCaller // Generic read-only contract binding to access the raw methods on
}

// ERC20MockTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ERC20MockTransactorRaw struct {
	Contract *ERC20MockTransactor // Generic write-only contract binding to access the raw methods on
}

// NewERC20Mock creates a new instance of ERC20Mock, bound to a specific deployed contract.
func NewERC20Mock(address common.Address, backend bind.ContractBackend) (*ERC20Mock, error) {
	contract, err := bindERC20Mock(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ERC20Mock{ERC20MockCaller: ERC20MockCaller{contract: contract}, ERC20MockTransactor: ERC20MockTransactor{contract: contract}, ERC20MockFilterer: ERC20MockFilterer{contract: contract}}, nil
}

// NewERC20MockCaller creates a new read-only instance of ERC20Mock, bound to a specific deployed contract.
func NewERC20MockCaller(address common.Address, caller bind.ContractCaller) (*ERC20MockCaller, error) {
	contract, err := bindERC20Mock(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ERC20MockCaller{contract: contract}, nil
}

// NewERC20MockTransactor creates a new write-only instance of ERC20Mock, bound to a specific deployed contract.
func NewERC20MockTransactor(address common.Address, transactor bind.ContractTransactor) (*ERC20MockTransactor, error) {
	contract, err := bindERC20Mock(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ERC20MockTransactor{contract: contract}, nil
}

// NewERC20MockFilterer creates a new log filterer instance of ERC20Mock, bound to a specific deployed contract.
func NewERC20MockFilterer(address common.Address, filterer bind.ContractFilterer) (*ERC20MockFilterer, error) {
	contract, err := bindERC20Mock(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ERC20MockFilterer{contract: contract}, nil
}

// bindERC20Mock binds a generic wrapper to an already deployed contract.
func bindERC20Mock(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := ERC20MockMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ERC20Mock *ERC20MockRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ERC20Mock.Contract.ERC20MockCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ERC20Mock *ERC20MockRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ERC20Mock.Contract.ERC20MockTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ERC20Mock *ERC20MockRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ERC20Mock.Contract.ERC20MockTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ERC20Mock *ERC20MockCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ERC20Mock.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ERC20Mock *ERC20MockTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ERC20Mock.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ERC20Mock *ERC20MockTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ERC20Mock.Contract.contract.Transact(opts, method, params...)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_ERC20Mock *ERC20MockCaller) Allowance(opts *bind.CallOpts, owner common.Address, spender common.Address) (*big.Int, error) {
	var out []interface{}
	err := _ERC20Mock.contract.Call(opts, &out, "allowance", owner, spender)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_ERC20Mock *ERC20MockSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _ERC20Mock.Contract.Allowance(&_ERC20Mock.CallOpts, owner, spender)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_ERC20Mock *ERC20MockCallerSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _ERC20Mock.Contract.Allowance(&_ERC20Mock.CallOpts, owner, spender)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_ERC20Mock *ERC20MockCaller) BalanceOf(opts *bind.CallOpts, account common.Address) (*big.Int, error) {
	var out []interface{}
	err := _ERC20Mock.contract.Call(opts, &out, "balanceOf", account)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_ERC20Mock *ERC20MockSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _ERC20Mock.Contract.BalanceOf(&_ERC20Mock.CallOpts, account)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_ERC20Mock *ERC20MockCallerSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _ERC20Mock.Contract.BalanceOf(&_ERC20Mock.CallOpts, account)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_ERC20Mock *ERC20MockCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _ERC20Mock.contract.Call(opts, &out, "totalSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_ERC20Mock *ERC20MockSession) TotalSupply() (*big.Int, error) {
	return _ERC20Mock.Contract.TotalSupply(&_ERC20Mock.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_ERC20Mock *ERC20MockCallerSession) TotalSupply() (*big.Int, error) {
	return _ERC20Mock.Contract.TotalSupply(&_ERC20Mock.CallOpts)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address , uint256 ) returns(bool)
func (_ERC20Mock *ERC20MockTransactor) Approve(opts *bind.TransactOpts, arg0 common.Address, arg1 *big.Int) (*types.Transaction, error) {
	return _ERC20Mock.contract.Transact(opts, "approve", arg0, arg1)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address , uint256 ) returns(bool)
func (_ERC20Mock *ERC20MockSession) Approve(arg0 common.Address, arg1 *big.Int) (*types.Transaction, error) {
	return _ERC20Mock.Contract.Approve(&_ERC20Mock.TransactOpts, arg0, arg1)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address , uint256 ) returns(bool)
func (_ERC20Mock *ERC20MockTransactorSession) Approve(arg0 common.Address, arg1 *big.Int) (*types.Transaction, error) {
	return _ERC20Mock.Contract.Approve(&_ERC20Mock.TransactOpts, arg0, arg1)
}

// Mint is a paid mutator transaction binding the contract method 0x40c10f19.
//
// Solidity: function mint(address account, uint256 amount) returns()
func (_ERC20Mock *ERC20MockTransactor) Mint(opts *bind.TransactOpts, account common.Address, amount *big.Int) (*types.Transaction, error) {
	return _ERC20Mock.contract.Transact(opts, "mint", account, amount)
}

// Mint is a paid mutator transaction binding the contract method 0x40c10f19.
//
// Solidity: function mint(address account, uint256 amount) returns()
func (_ERC20Mock *ERC20MockSession) Mint(account common.Address, amount *big.Int) (*types.Transaction, error) {
	return _ERC20Mock.Contract.Mint(&_ERC20Mock.TransactOpts, account, amount)
}

// Mint is a paid mutator transaction binding the contract method 0x40c10f19.
//
// Solidity: function mint(address account, uint256 amount) returns()
func (_ERC20Mock *ERC20MockTransactorSession) Mint(account common.Address, amount *big.Int) (*types.Transaction, error) {
	return _ERC20Mock.Contract.Mint(&_ERC20Mock.TransactOpts, account, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 amount) returns(bool)
func (_ERC20Mock *ERC20MockTransactor) Transfer(opts *bind.TransactOpts, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _ERC20Mock.contract.Transact(opts, "transfer", to, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 amount) returns(bool)
func (_ERC20Mock *ERC20MockSession) Transfer(to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _ERC20Mock.Contract.Transfer(&_ERC20Mock.TransactOpts, to, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 amount) returns(bool)
func (_ERC20Mock *ERC20MockTransactorSession) Transfer(to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _ERC20Mock.Contract.Transfer(&_ERC20Mock.TransactOpts, to, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 amount) returns(bool)
func (_ERC20Mock *ERC20MockTransactor) TransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _ERC20Mock.contract.Transact(opts, "transferFrom", from, to, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 amount) returns(bool)
func (_ERC20Mock *ERC20MockSession) TransferFrom(from common.Address, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _ERC20Mock.Contract.TransferFrom(&_ERC20Mock.TransactOpts, from, to, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 amount) returns(bool)
func (_ERC20Mock *ERC20MockTransactorSession) TransferFrom(from common.Address, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _ERC20Mock.Contract.TransferFrom(&_ERC20Mock.TransactOpts, from, to, amount)
}

// ERC20MockApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the ERC20Mock contract.
type ERC20MockApprovalIterator struct {
	Event *ERC20MockApproval // Event containing the contract specifics and raw log

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
func (it *ERC20MockApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ERC20MockApproval)
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
		it.Event = new(ERC20MockApproval)
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
func (it *ERC20MockApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ERC20MockApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ERC20MockApproval represents a Approval event raised by the ERC20Mock contract.
type ERC20MockApproval struct {
	Owner   common.Address
	Spender common.Address
	Value   *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_ERC20Mock *ERC20MockFilterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, spender []common.Address) (*ERC20MockApprovalIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _ERC20Mock.contract.FilterLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return &ERC20MockApprovalIterator{contract: _ERC20Mock.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_ERC20Mock *ERC20MockFilterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *ERC20MockApproval, owner []common.Address, spender []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _ERC20Mock.contract.WatchLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ERC20MockApproval)
				if err := _ERC20Mock.contract.UnpackLog(event, "Approval", log); err != nil {
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

// ParseApproval is a log parse operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_ERC20Mock *ERC20MockFilterer) ParseApproval(log types.Log) (*ERC20MockApproval, error) {
	event := new(ERC20MockApproval)
	if err := _ERC20Mock.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ERC20MockTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the ERC20Mock contract.
type ERC20MockTransferIterator struct {
	Event *ERC20MockTransfer // Event containing the contract specifics and raw log

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
func (it *ERC20MockTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ERC20MockTransfer)
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
		it.Event = new(ERC20MockTransfer)
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
func (it *ERC20MockTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ERC20MockTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ERC20MockTransfer represents a Transfer event raised by the ERC20Mock contract.
type ERC20MockTransfer struct {
	From  common.Address
	To    common.Address
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_ERC20Mock *ERC20MockFilterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*ERC20MockTransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _ERC20Mock.contract.FilterLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &ERC20MockTransferIterator{contract: _ERC20Mock.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_ERC20Mock *ERC20MockFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *ERC20MockTransfer, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _ERC20Mock.contract.WatchLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ERC20MockTransfer)
				if err := _ERC20Mock.contract.UnpackLog(event, "Transfer", log); err != nil {
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

// ParseTransfer is a log parse operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_ERC20Mock *ERC20MockFilterer) ParseTransfer(log types.Log) (*ERC20MockTransfer, error) {
	event := new(ERC20MockTransfer)
	if err := _ERC20Mock.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
