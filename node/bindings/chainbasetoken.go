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

// ChainbaseTokenMetaData contains all meta data concerning the ChainbaseToken contract.
var ChainbaseTokenMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"constructor\",\"inputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"allowance\",\"inputs\":[{\"name\":\"owner\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"spender\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"approve\",\"inputs\":[{\"name\":\"spender\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"balanceOf\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"burn\",\"inputs\":[{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"burnFrom\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"cap\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"decimals\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint8\",\"internalType\":\"uint8\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"decreaseAllowance\",\"inputs\":[{\"name\":\"spender\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"subtractedValue\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"increaseAllowance\",\"inputs\":[{\"name\":\"spender\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"addedValue\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"mint\",\"inputs\":[{\"name\":\"to\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"name\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"owner\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"renounceOwnership\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"symbol\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"totalSupply\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"transfer\",\"inputs\":[{\"name\":\"to\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"transferFrom\",\"inputs\":[{\"name\":\"from\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"to\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"transferOwnership\",\"inputs\":[{\"name\":\"newOwner\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"event\",\"name\":\"Approval\",\"inputs\":[{\"name\":\"owner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"spender\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"value\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OwnershipTransferred\",\"inputs\":[{\"name\":\"previousOwner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"newOwner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Transfer\",\"inputs\":[{\"name\":\"from\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"to\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"value\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false}]",
	Bin: "0x60a06040523480156200001157600080fd5b506c02a68bedbb190931f6500000006040518060400160405280600f81526020016e21b430b4b73130b9b2902a37b5b2b760891b815250604051806040016040528060018152602001604360f81b8152508160039081620000739190620001e7565b506004620000828282620001e7565b50505060008111620000da5760405162461bcd60e51b815260206004820152601560248201527f45524332304361707065643a2063617020697320300000000000000000000000604482015260640160405180910390fd5b608052620000e833620000ee565b620002b3565b600580546001600160a01b038381166001600160a01b0319831681179093556040519116919082907f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e090600090a35050565b634e487b7160e01b600052604160045260246000fd5b600181811c908216806200016b57607f821691505b6020821081036200018c57634e487b7160e01b600052602260045260246000fd5b50919050565b601f821115620001e2576000816000526020600020601f850160051c81016020861015620001bd5750805b601f850160051c820191505b81811015620001de57828155600101620001c9565b5050505b505050565b81516001600160401b0381111562000203576200020362000140565b6200021b8162000214845462000156565b8462000192565b602080601f8311600181146200025357600084156200023a5750858301515b600019600386901b1c1916600185901b178555620001de565b600085815260208120601f198616915b82811015620002845788860151825594840194600190910190840162000263565b5085821015620002a35787850151600019600388901b60f8161c191681555b5050505050600190811b01905550565b608051610d5a620002d6600039600081816101920152610a450152610d5a6000f3fe608060405234801561001057600080fd5b50600436106101165760003560e01c806370a08231116100a257806395d89b411161007157806395d89b4114610250578063a457c2d714610258578063a9059cbb1461026b578063dd62ed3e1461027e578063f2fde38b1461029157600080fd5b806370a08231146101f1578063715018a61461021a57806379cc6790146102225780638da5cb5b1461023557600080fd5b8063313ce567116100e9578063313ce56714610181578063355274ea1461019057806339509351146101b657806340c10f19146101c957806342966c68146101de57600080fd5b806306fdde031461011b578063095ea7b31461013957806318160ddd1461015c57806323b872dd1461016e575b600080fd5b6101236102a4565b6040516101309190610b8a565b60405180910390f35b61014c610147366004610bf5565b610336565b6040519015158152602001610130565b6002545b604051908152602001610130565b61014c61017c366004610c1f565b610350565b60405160128152602001610130565b7f0000000000000000000000000000000000000000000000000000000000000000610160565b61014c6101c4366004610bf5565b610374565b6101dc6101d7366004610bf5565b610396565b005b6101dc6101ec366004610c5b565b6103ac565b6101606101ff366004610c74565b6001600160a01b031660009081526020819052604090205490565b6101dc6103b9565b6101dc610230366004610bf5565b6103cd565b6005546040516001600160a01b039091168152602001610130565b6101236103e2565b61014c610266366004610bf5565b6103f1565b61014c610279366004610bf5565b610471565b61016061028c366004610c96565b61047f565b6101dc61029f366004610c74565b6104aa565b6060600380546102b390610cc9565b80601f01602080910402602001604051908101604052809291908181526020018280546102df90610cc9565b801561032c5780601f106103015761010080835404028352916020019161032c565b820191906000526020600020905b81548152906001019060200180831161030f57829003601f168201915b5050505050905090565b600033610344818585610520565b60019150505b92915050565b60003361035e858285610645565b6103698585856106bf565b506001949350505050565b600033610344818585610387838361047f565b6103919190610d03565b610520565b61039e610863565b6103a882826108bd565b5050565b6103b633826108c7565b50565b6103c1610863565b6103cb60006109f1565b565b6103d8823383610645565b6103a882826108c7565b6060600480546102b390610cc9565b600033816103ff828661047f565b9050838110156104645760405162461bcd60e51b815260206004820152602560248201527f45524332303a2064656372656173656420616c6c6f77616e63652062656c6f77604482015264207a65726f60d81b60648201526084015b60405180910390fd5b6103698286868403610520565b6000336103448185856106bf565b6001600160a01b03918216600090815260016020908152604080832093909416825291909152205490565b6104b2610863565b6001600160a01b0381166105175760405162461bcd60e51b815260206004820152602660248201527f4f776e61626c653a206e6577206f776e657220697320746865207a65726f206160448201526564647265737360d01b606482015260840161045b565b6103b6816109f1565b6001600160a01b0383166105825760405162461bcd60e51b8152602060048201526024808201527f45524332303a20617070726f76652066726f6d20746865207a65726f206164646044820152637265737360e01b606482015260840161045b565b6001600160a01b0382166105e35760405162461bcd60e51b815260206004820152602260248201527f45524332303a20617070726f766520746f20746865207a65726f206164647265604482015261737360f01b606482015260840161045b565b6001600160a01b0383811660008181526001602090815260408083209487168084529482529182902085905590518481527f8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b92591015b60405180910390a3505050565b6000610651848461047f565b905060001981146106b957818110156106ac5760405162461bcd60e51b815260206004820152601d60248201527f45524332303a20696e73756666696369656e7420616c6c6f77616e6365000000604482015260640161045b565b6106b98484848403610520565b50505050565b6001600160a01b0383166107235760405162461bcd60e51b815260206004820152602560248201527f45524332303a207472616e736665722066726f6d20746865207a65726f206164604482015264647265737360d81b606482015260840161045b565b6001600160a01b0382166107855760405162461bcd60e51b815260206004820152602360248201527f45524332303a207472616e7366657220746f20746865207a65726f206164647260448201526265737360e81b606482015260840161045b565b6001600160a01b038316600090815260208190526040902054818110156107fd5760405162461bcd60e51b815260206004820152602660248201527f45524332303a207472616e7366657220616d6f756e7420657863656564732062604482015265616c616e636560d01b606482015260840161045b565b6001600160a01b03848116600081815260208181526040808320878703905593871680835291849020805487019055925185815290927fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef910160405180910390a36106b9565b6005546001600160a01b031633146103cb5760405162461bcd60e51b815260206004820181905260248201527f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e6572604482015260640161045b565b6103a88282610a43565b6001600160a01b0382166109275760405162461bcd60e51b815260206004820152602160248201527f45524332303a206275726e2066726f6d20746865207a65726f206164647265736044820152607360f81b606482015260840161045b565b6001600160a01b0382166000908152602081905260409020548181101561099b5760405162461bcd60e51b815260206004820152602260248201527f45524332303a206275726e20616d6f756e7420657863656564732062616c616e604482015261636560f01b606482015260840161045b565b6001600160a01b0383166000818152602081815260408083208686039055600280548790039055518581529192917fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef9101610638565b600580546001600160a01b038381166001600160a01b0319831681179093556040519116919082907f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e090600090a35050565b7f000000000000000000000000000000000000000000000000000000000000000081610a6e60025490565b610a789190610d03565b1115610ac65760405162461bcd60e51b815260206004820152601960248201527f45524332304361707065643a2063617020657863656564656400000000000000604482015260640161045b565b6103a882826001600160a01b038216610b215760405162461bcd60e51b815260206004820152601f60248201527f45524332303a206d696e7420746f20746865207a65726f206164647265737300604482015260640161045b565b8060026000828254610b339190610d03565b90915550506001600160a01b038216600081815260208181526040808320805486019055518481527fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef910160405180910390a35050565b60006020808352835180602085015260005b81811015610bb857858101830151858201604001528201610b9c565b506000604082860101526040601f19601f8301168501019250505092915050565b80356001600160a01b0381168114610bf057600080fd5b919050565b60008060408385031215610c0857600080fd5b610c1183610bd9565b946020939093013593505050565b600080600060608486031215610c3457600080fd5b610c3d84610bd9565b9250610c4b60208501610bd9565b9150604084013590509250925092565b600060208284031215610c6d57600080fd5b5035919050565b600060208284031215610c8657600080fd5b610c8f82610bd9565b9392505050565b60008060408385031215610ca957600080fd5b610cb283610bd9565b9150610cc060208401610bd9565b90509250929050565b600181811c90821680610cdd57607f821691505b602082108103610cfd57634e487b7160e01b600052602260045260246000fd5b50919050565b8082018082111561034a57634e487b7160e01b600052601160045260246000fdfea2646970667358221220a0271dc67ad8cd244323bd4495c18d7ba684bda1a93c95e581958019dc65220564736f6c63430008160033",
}

// ChainbaseTokenABI is the input ABI used to generate the binding from.
// Deprecated: Use ChainbaseTokenMetaData.ABI instead.
var ChainbaseTokenABI = ChainbaseTokenMetaData.ABI

// ChainbaseTokenBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use ChainbaseTokenMetaData.Bin instead.
var ChainbaseTokenBin = ChainbaseTokenMetaData.Bin

// DeployChainbaseToken deploys a new Ethereum contract, binding an instance of ChainbaseToken to it.
func DeployChainbaseToken(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *ChainbaseToken, error) {
	parsed, err := ChainbaseTokenMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(ChainbaseTokenBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &ChainbaseToken{ChainbaseTokenCaller: ChainbaseTokenCaller{contract: contract}, ChainbaseTokenTransactor: ChainbaseTokenTransactor{contract: contract}, ChainbaseTokenFilterer: ChainbaseTokenFilterer{contract: contract}}, nil
}

// ChainbaseToken is an auto generated Go binding around an Ethereum contract.
type ChainbaseToken struct {
	ChainbaseTokenCaller     // Read-only binding to the contract
	ChainbaseTokenTransactor // Write-only binding to the contract
	ChainbaseTokenFilterer   // Log filterer for contract events
}

// ChainbaseTokenCaller is an auto generated read-only Go binding around an Ethereum contract.
type ChainbaseTokenCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ChainbaseTokenTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ChainbaseTokenTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ChainbaseTokenFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ChainbaseTokenFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ChainbaseTokenSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ChainbaseTokenSession struct {
	Contract     *ChainbaseToken   // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ChainbaseTokenCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ChainbaseTokenCallerSession struct {
	Contract *ChainbaseTokenCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts         // Call options to use throughout this session
}

// ChainbaseTokenTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ChainbaseTokenTransactorSession struct {
	Contract     *ChainbaseTokenTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts         // Transaction auth options to use throughout this session
}

// ChainbaseTokenRaw is an auto generated low-level Go binding around an Ethereum contract.
type ChainbaseTokenRaw struct {
	Contract *ChainbaseToken // Generic contract binding to access the raw methods on
}

// ChainbaseTokenCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ChainbaseTokenCallerRaw struct {
	Contract *ChainbaseTokenCaller // Generic read-only contract binding to access the raw methods on
}

// ChainbaseTokenTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ChainbaseTokenTransactorRaw struct {
	Contract *ChainbaseTokenTransactor // Generic write-only contract binding to access the raw methods on
}

// NewChainbaseToken creates a new instance of ChainbaseToken, bound to a specific deployed contract.
func NewChainbaseToken(address common.Address, backend bind.ContractBackend) (*ChainbaseToken, error) {
	contract, err := bindChainbaseToken(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ChainbaseToken{ChainbaseTokenCaller: ChainbaseTokenCaller{contract: contract}, ChainbaseTokenTransactor: ChainbaseTokenTransactor{contract: contract}, ChainbaseTokenFilterer: ChainbaseTokenFilterer{contract: contract}}, nil
}

// NewChainbaseTokenCaller creates a new read-only instance of ChainbaseToken, bound to a specific deployed contract.
func NewChainbaseTokenCaller(address common.Address, caller bind.ContractCaller) (*ChainbaseTokenCaller, error) {
	contract, err := bindChainbaseToken(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ChainbaseTokenCaller{contract: contract}, nil
}

// NewChainbaseTokenTransactor creates a new write-only instance of ChainbaseToken, bound to a specific deployed contract.
func NewChainbaseTokenTransactor(address common.Address, transactor bind.ContractTransactor) (*ChainbaseTokenTransactor, error) {
	contract, err := bindChainbaseToken(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ChainbaseTokenTransactor{contract: contract}, nil
}

// NewChainbaseTokenFilterer creates a new log filterer instance of ChainbaseToken, bound to a specific deployed contract.
func NewChainbaseTokenFilterer(address common.Address, filterer bind.ContractFilterer) (*ChainbaseTokenFilterer, error) {
	contract, err := bindChainbaseToken(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ChainbaseTokenFilterer{contract: contract}, nil
}

// bindChainbaseToken binds a generic wrapper to an already deployed contract.
func bindChainbaseToken(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := ChainbaseTokenMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ChainbaseToken *ChainbaseTokenRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ChainbaseToken.Contract.ChainbaseTokenCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ChainbaseToken *ChainbaseTokenRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ChainbaseToken.Contract.ChainbaseTokenTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ChainbaseToken *ChainbaseTokenRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ChainbaseToken.Contract.ChainbaseTokenTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ChainbaseToken *ChainbaseTokenCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ChainbaseToken.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ChainbaseToken *ChainbaseTokenTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ChainbaseToken.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ChainbaseToken *ChainbaseTokenTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ChainbaseToken.Contract.contract.Transact(opts, method, params...)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_ChainbaseToken *ChainbaseTokenCaller) Allowance(opts *bind.CallOpts, owner common.Address, spender common.Address) (*big.Int, error) {
	var out []interface{}
	err := _ChainbaseToken.contract.Call(opts, &out, "allowance", owner, spender)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_ChainbaseToken *ChainbaseTokenSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _ChainbaseToken.Contract.Allowance(&_ChainbaseToken.CallOpts, owner, spender)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_ChainbaseToken *ChainbaseTokenCallerSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _ChainbaseToken.Contract.Allowance(&_ChainbaseToken.CallOpts, owner, spender)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_ChainbaseToken *ChainbaseTokenCaller) BalanceOf(opts *bind.CallOpts, account common.Address) (*big.Int, error) {
	var out []interface{}
	err := _ChainbaseToken.contract.Call(opts, &out, "balanceOf", account)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_ChainbaseToken *ChainbaseTokenSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _ChainbaseToken.Contract.BalanceOf(&_ChainbaseToken.CallOpts, account)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_ChainbaseToken *ChainbaseTokenCallerSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _ChainbaseToken.Contract.BalanceOf(&_ChainbaseToken.CallOpts, account)
}

// Cap is a free data retrieval call binding the contract method 0x355274ea.
//
// Solidity: function cap() view returns(uint256)
func (_ChainbaseToken *ChainbaseTokenCaller) Cap(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _ChainbaseToken.contract.Call(opts, &out, "cap")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Cap is a free data retrieval call binding the contract method 0x355274ea.
//
// Solidity: function cap() view returns(uint256)
func (_ChainbaseToken *ChainbaseTokenSession) Cap() (*big.Int, error) {
	return _ChainbaseToken.Contract.Cap(&_ChainbaseToken.CallOpts)
}

// Cap is a free data retrieval call binding the contract method 0x355274ea.
//
// Solidity: function cap() view returns(uint256)
func (_ChainbaseToken *ChainbaseTokenCallerSession) Cap() (*big.Int, error) {
	return _ChainbaseToken.Contract.Cap(&_ChainbaseToken.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_ChainbaseToken *ChainbaseTokenCaller) Decimals(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _ChainbaseToken.contract.Call(opts, &out, "decimals")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_ChainbaseToken *ChainbaseTokenSession) Decimals() (uint8, error) {
	return _ChainbaseToken.Contract.Decimals(&_ChainbaseToken.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_ChainbaseToken *ChainbaseTokenCallerSession) Decimals() (uint8, error) {
	return _ChainbaseToken.Contract.Decimals(&_ChainbaseToken.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_ChainbaseToken *ChainbaseTokenCaller) Name(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _ChainbaseToken.contract.Call(opts, &out, "name")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_ChainbaseToken *ChainbaseTokenSession) Name() (string, error) {
	return _ChainbaseToken.Contract.Name(&_ChainbaseToken.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_ChainbaseToken *ChainbaseTokenCallerSession) Name() (string, error) {
	return _ChainbaseToken.Contract.Name(&_ChainbaseToken.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_ChainbaseToken *ChainbaseTokenCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ChainbaseToken.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_ChainbaseToken *ChainbaseTokenSession) Owner() (common.Address, error) {
	return _ChainbaseToken.Contract.Owner(&_ChainbaseToken.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_ChainbaseToken *ChainbaseTokenCallerSession) Owner() (common.Address, error) {
	return _ChainbaseToken.Contract.Owner(&_ChainbaseToken.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_ChainbaseToken *ChainbaseTokenCaller) Symbol(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _ChainbaseToken.contract.Call(opts, &out, "symbol")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_ChainbaseToken *ChainbaseTokenSession) Symbol() (string, error) {
	return _ChainbaseToken.Contract.Symbol(&_ChainbaseToken.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_ChainbaseToken *ChainbaseTokenCallerSession) Symbol() (string, error) {
	return _ChainbaseToken.Contract.Symbol(&_ChainbaseToken.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_ChainbaseToken *ChainbaseTokenCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _ChainbaseToken.contract.Call(opts, &out, "totalSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_ChainbaseToken *ChainbaseTokenSession) TotalSupply() (*big.Int, error) {
	return _ChainbaseToken.Contract.TotalSupply(&_ChainbaseToken.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_ChainbaseToken *ChainbaseTokenCallerSession) TotalSupply() (*big.Int, error) {
	return _ChainbaseToken.Contract.TotalSupply(&_ChainbaseToken.CallOpts)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_ChainbaseToken *ChainbaseTokenTransactor) Approve(opts *bind.TransactOpts, spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _ChainbaseToken.contract.Transact(opts, "approve", spender, amount)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_ChainbaseToken *ChainbaseTokenSession) Approve(spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _ChainbaseToken.Contract.Approve(&_ChainbaseToken.TransactOpts, spender, amount)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_ChainbaseToken *ChainbaseTokenTransactorSession) Approve(spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _ChainbaseToken.Contract.Approve(&_ChainbaseToken.TransactOpts, spender, amount)
}

// Burn is a paid mutator transaction binding the contract method 0x42966c68.
//
// Solidity: function burn(uint256 amount) returns()
func (_ChainbaseToken *ChainbaseTokenTransactor) Burn(opts *bind.TransactOpts, amount *big.Int) (*types.Transaction, error) {
	return _ChainbaseToken.contract.Transact(opts, "burn", amount)
}

// Burn is a paid mutator transaction binding the contract method 0x42966c68.
//
// Solidity: function burn(uint256 amount) returns()
func (_ChainbaseToken *ChainbaseTokenSession) Burn(amount *big.Int) (*types.Transaction, error) {
	return _ChainbaseToken.Contract.Burn(&_ChainbaseToken.TransactOpts, amount)
}

// Burn is a paid mutator transaction binding the contract method 0x42966c68.
//
// Solidity: function burn(uint256 amount) returns()
func (_ChainbaseToken *ChainbaseTokenTransactorSession) Burn(amount *big.Int) (*types.Transaction, error) {
	return _ChainbaseToken.Contract.Burn(&_ChainbaseToken.TransactOpts, amount)
}

// BurnFrom is a paid mutator transaction binding the contract method 0x79cc6790.
//
// Solidity: function burnFrom(address account, uint256 amount) returns()
func (_ChainbaseToken *ChainbaseTokenTransactor) BurnFrom(opts *bind.TransactOpts, account common.Address, amount *big.Int) (*types.Transaction, error) {
	return _ChainbaseToken.contract.Transact(opts, "burnFrom", account, amount)
}

// BurnFrom is a paid mutator transaction binding the contract method 0x79cc6790.
//
// Solidity: function burnFrom(address account, uint256 amount) returns()
func (_ChainbaseToken *ChainbaseTokenSession) BurnFrom(account common.Address, amount *big.Int) (*types.Transaction, error) {
	return _ChainbaseToken.Contract.BurnFrom(&_ChainbaseToken.TransactOpts, account, amount)
}

// BurnFrom is a paid mutator transaction binding the contract method 0x79cc6790.
//
// Solidity: function burnFrom(address account, uint256 amount) returns()
func (_ChainbaseToken *ChainbaseTokenTransactorSession) BurnFrom(account common.Address, amount *big.Int) (*types.Transaction, error) {
	return _ChainbaseToken.Contract.BurnFrom(&_ChainbaseToken.TransactOpts, account, amount)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address spender, uint256 subtractedValue) returns(bool)
func (_ChainbaseToken *ChainbaseTokenTransactor) DecreaseAllowance(opts *bind.TransactOpts, spender common.Address, subtractedValue *big.Int) (*types.Transaction, error) {
	return _ChainbaseToken.contract.Transact(opts, "decreaseAllowance", spender, subtractedValue)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address spender, uint256 subtractedValue) returns(bool)
func (_ChainbaseToken *ChainbaseTokenSession) DecreaseAllowance(spender common.Address, subtractedValue *big.Int) (*types.Transaction, error) {
	return _ChainbaseToken.Contract.DecreaseAllowance(&_ChainbaseToken.TransactOpts, spender, subtractedValue)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address spender, uint256 subtractedValue) returns(bool)
func (_ChainbaseToken *ChainbaseTokenTransactorSession) DecreaseAllowance(spender common.Address, subtractedValue *big.Int) (*types.Transaction, error) {
	return _ChainbaseToken.Contract.DecreaseAllowance(&_ChainbaseToken.TransactOpts, spender, subtractedValue)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address spender, uint256 addedValue) returns(bool)
func (_ChainbaseToken *ChainbaseTokenTransactor) IncreaseAllowance(opts *bind.TransactOpts, spender common.Address, addedValue *big.Int) (*types.Transaction, error) {
	return _ChainbaseToken.contract.Transact(opts, "increaseAllowance", spender, addedValue)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address spender, uint256 addedValue) returns(bool)
func (_ChainbaseToken *ChainbaseTokenSession) IncreaseAllowance(spender common.Address, addedValue *big.Int) (*types.Transaction, error) {
	return _ChainbaseToken.Contract.IncreaseAllowance(&_ChainbaseToken.TransactOpts, spender, addedValue)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address spender, uint256 addedValue) returns(bool)
func (_ChainbaseToken *ChainbaseTokenTransactorSession) IncreaseAllowance(spender common.Address, addedValue *big.Int) (*types.Transaction, error) {
	return _ChainbaseToken.Contract.IncreaseAllowance(&_ChainbaseToken.TransactOpts, spender, addedValue)
}

// Mint is a paid mutator transaction binding the contract method 0x40c10f19.
//
// Solidity: function mint(address to, uint256 amount) returns()
func (_ChainbaseToken *ChainbaseTokenTransactor) Mint(opts *bind.TransactOpts, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _ChainbaseToken.contract.Transact(opts, "mint", to, amount)
}

// Mint is a paid mutator transaction binding the contract method 0x40c10f19.
//
// Solidity: function mint(address to, uint256 amount) returns()
func (_ChainbaseToken *ChainbaseTokenSession) Mint(to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _ChainbaseToken.Contract.Mint(&_ChainbaseToken.TransactOpts, to, amount)
}

// Mint is a paid mutator transaction binding the contract method 0x40c10f19.
//
// Solidity: function mint(address to, uint256 amount) returns()
func (_ChainbaseToken *ChainbaseTokenTransactorSession) Mint(to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _ChainbaseToken.Contract.Mint(&_ChainbaseToken.TransactOpts, to, amount)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_ChainbaseToken *ChainbaseTokenTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ChainbaseToken.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_ChainbaseToken *ChainbaseTokenSession) RenounceOwnership() (*types.Transaction, error) {
	return _ChainbaseToken.Contract.RenounceOwnership(&_ChainbaseToken.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_ChainbaseToken *ChainbaseTokenTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _ChainbaseToken.Contract.RenounceOwnership(&_ChainbaseToken.TransactOpts)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 amount) returns(bool)
func (_ChainbaseToken *ChainbaseTokenTransactor) Transfer(opts *bind.TransactOpts, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _ChainbaseToken.contract.Transact(opts, "transfer", to, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 amount) returns(bool)
func (_ChainbaseToken *ChainbaseTokenSession) Transfer(to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _ChainbaseToken.Contract.Transfer(&_ChainbaseToken.TransactOpts, to, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 amount) returns(bool)
func (_ChainbaseToken *ChainbaseTokenTransactorSession) Transfer(to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _ChainbaseToken.Contract.Transfer(&_ChainbaseToken.TransactOpts, to, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 amount) returns(bool)
func (_ChainbaseToken *ChainbaseTokenTransactor) TransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _ChainbaseToken.contract.Transact(opts, "transferFrom", from, to, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 amount) returns(bool)
func (_ChainbaseToken *ChainbaseTokenSession) TransferFrom(from common.Address, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _ChainbaseToken.Contract.TransferFrom(&_ChainbaseToken.TransactOpts, from, to, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 amount) returns(bool)
func (_ChainbaseToken *ChainbaseTokenTransactorSession) TransferFrom(from common.Address, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _ChainbaseToken.Contract.TransferFrom(&_ChainbaseToken.TransactOpts, from, to, amount)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_ChainbaseToken *ChainbaseTokenTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _ChainbaseToken.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_ChainbaseToken *ChainbaseTokenSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _ChainbaseToken.Contract.TransferOwnership(&_ChainbaseToken.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_ChainbaseToken *ChainbaseTokenTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _ChainbaseToken.Contract.TransferOwnership(&_ChainbaseToken.TransactOpts, newOwner)
}

// ChainbaseTokenApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the ChainbaseToken contract.
type ChainbaseTokenApprovalIterator struct {
	Event *ChainbaseTokenApproval // Event containing the contract specifics and raw log

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
func (it *ChainbaseTokenApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ChainbaseTokenApproval)
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
		it.Event = new(ChainbaseTokenApproval)
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
func (it *ChainbaseTokenApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ChainbaseTokenApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ChainbaseTokenApproval represents a Approval event raised by the ChainbaseToken contract.
type ChainbaseTokenApproval struct {
	Owner   common.Address
	Spender common.Address
	Value   *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_ChainbaseToken *ChainbaseTokenFilterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, spender []common.Address) (*ChainbaseTokenApprovalIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _ChainbaseToken.contract.FilterLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return &ChainbaseTokenApprovalIterator{contract: _ChainbaseToken.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_ChainbaseToken *ChainbaseTokenFilterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *ChainbaseTokenApproval, owner []common.Address, spender []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _ChainbaseToken.contract.WatchLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ChainbaseTokenApproval)
				if err := _ChainbaseToken.contract.UnpackLog(event, "Approval", log); err != nil {
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
func (_ChainbaseToken *ChainbaseTokenFilterer) ParseApproval(log types.Log) (*ChainbaseTokenApproval, error) {
	event := new(ChainbaseTokenApproval)
	if err := _ChainbaseToken.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ChainbaseTokenOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the ChainbaseToken contract.
type ChainbaseTokenOwnershipTransferredIterator struct {
	Event *ChainbaseTokenOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *ChainbaseTokenOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ChainbaseTokenOwnershipTransferred)
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
		it.Event = new(ChainbaseTokenOwnershipTransferred)
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
func (it *ChainbaseTokenOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ChainbaseTokenOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ChainbaseTokenOwnershipTransferred represents a OwnershipTransferred event raised by the ChainbaseToken contract.
type ChainbaseTokenOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_ChainbaseToken *ChainbaseTokenFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*ChainbaseTokenOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _ChainbaseToken.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &ChainbaseTokenOwnershipTransferredIterator{contract: _ChainbaseToken.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_ChainbaseToken *ChainbaseTokenFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *ChainbaseTokenOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _ChainbaseToken.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ChainbaseTokenOwnershipTransferred)
				if err := _ChainbaseToken.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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

// ParseOwnershipTransferred is a log parse operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_ChainbaseToken *ChainbaseTokenFilterer) ParseOwnershipTransferred(log types.Log) (*ChainbaseTokenOwnershipTransferred, error) {
	event := new(ChainbaseTokenOwnershipTransferred)
	if err := _ChainbaseToken.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ChainbaseTokenTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the ChainbaseToken contract.
type ChainbaseTokenTransferIterator struct {
	Event *ChainbaseTokenTransfer // Event containing the contract specifics and raw log

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
func (it *ChainbaseTokenTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ChainbaseTokenTransfer)
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
		it.Event = new(ChainbaseTokenTransfer)
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
func (it *ChainbaseTokenTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ChainbaseTokenTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ChainbaseTokenTransfer represents a Transfer event raised by the ChainbaseToken contract.
type ChainbaseTokenTransfer struct {
	From  common.Address
	To    common.Address
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_ChainbaseToken *ChainbaseTokenFilterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*ChainbaseTokenTransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _ChainbaseToken.contract.FilterLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &ChainbaseTokenTransferIterator{contract: _ChainbaseToken.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_ChainbaseToken *ChainbaseTokenFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *ChainbaseTokenTransfer, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _ChainbaseToken.contract.WatchLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ChainbaseTokenTransfer)
				if err := _ChainbaseToken.contract.UnpackLog(event, "Transfer", log); err != nil {
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
func (_ChainbaseToken *ChainbaseTokenFilterer) ParseTransfer(log types.Log) (*ChainbaseTokenTransfer, error) {
	event := new(ChainbaseTokenTransfer)
	if err := _ChainbaseToken.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
