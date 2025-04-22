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

// StakingMetaData contains all meta data concerning the Staking contract.
var StakingMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"constructor\",\"inputs\":[{\"name\":\"_cToken\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"UNLOCK_PERIOD\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"addWhitelist\",\"inputs\":[{\"name\":\"operators\",\"type\":\"address[]\",\"internalType\":\"address[]\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"cToken\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIERC20\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"delegate\",\"inputs\":[{\"name\":\"operator\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"delegations\",\"inputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"initialize\",\"inputs\":[{\"name\":\"initialOwner\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"minOperatorStake\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"operatorStakes\",\"inputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"operatorWhitelist\",\"inputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"owner\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"pause\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"paused\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"removeWhitelist\",\"inputs\":[{\"name\":\"operators\",\"type\":\"address[]\",\"internalType\":\"address[]\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"renounceOwnership\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setMinOperatorStake\",\"inputs\":[{\"name\":\"_minOperatorStake\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"stake\",\"inputs\":[{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"transferOwnership\",\"inputs\":[{\"name\":\"newOwner\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"unpause\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"unstake\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"unstakeRequests\",\"inputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"unlockTime\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"withdrawDelegation\",\"inputs\":[{\"name\":\"operator\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"withdrawStake\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"event\",\"name\":\"DelegationDeposited\",\"inputs\":[{\"name\":\"delegator\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"operator\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"DelegationWithdrawn\",\"inputs\":[{\"name\":\"delegator\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"operator\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Initialized\",\"inputs\":[{\"name\":\"version\",\"type\":\"uint64\",\"indexed\":false,\"internalType\":\"uint64\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"MinOperatorStakeUpdated\",\"inputs\":[{\"name\":\"oldAmount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"newAmount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OwnershipTransferred\",\"inputs\":[{\"name\":\"previousOwner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"newOwner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Paused\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"StakeDeposited\",\"inputs\":[{\"name\":\"operator\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"StakeWithdrawn\",\"inputs\":[{\"name\":\"operator\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Unpaused\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"UnstakeRequested\",\"inputs\":[{\"name\":\"operator\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"unlockTime\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"WhitelistAdded\",\"inputs\":[{\"name\":\"operator\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"WhitelistRemoved\",\"inputs\":[{\"name\":\"operator\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"error\",\"name\":\"EnforcedPause\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"ExpectedPause\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidInitialization\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"NotInitializing\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"OwnableInvalidOwner\",\"inputs\":[{\"name\":\"owner\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"OwnableUnauthorizedAccount\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"ReentrancyGuardReentrantCall\",\"inputs\":[]}]",
	Bin: "0x60a06040523480156200001157600080fd5b5060405162001ad238038062001ad2833981016040819052620000349162000161565b6001600160a01b0381166200008f5760405162461bcd60e51b815260206004820152601f60248201527f5374616b696e673a20496e76616c69642063546f6b656e206164647265737300604482015260640160405180910390fd5b6001600160a01b038116608052620000a6620000ad565b5062000193565b7ff0c57e16840df040f15088dc2f81fe391c3923bec73e23a9662efc9c229c6a00805468010000000000000000900460ff1615620000fe5760405163f92ee8a960e01b815260040160405180910390fd5b80546001600160401b03908116146200015e5780546001600160401b0319166001600160401b0390811782556040519081527fc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d29060200160405180910390a15b50565b6000602082840312156200017457600080fd5b81516001600160a01b03811681146200018c57600080fd5b9392505050565b608051611907620001cb6000396000818161020b015281816104720152818161062c01528181610b0e0152610cfc01526119076000f3fe608060405234801561001057600080fd5b50600436106101425760003560e01c80638456cb59116100b8578063c64814dd1161007c578063c64814dd146102b3578063d775cb61146102de578063e03c8632146102e7578063e0eb4d2e1461030a578063edac985b14610346578063f2fde38b1461035957600080fd5b80638456cb591461024d5780638da5cb5b14610255578063a694fc3a14610285578063bed9d86114610298578063c4d66de8146102a057600080fd5b8063259a28cf1161010a578063259a28cf146101c85780632def6620146101d25780633f4ba83a146101da5780635c975abb146101e257806369e527da14610206578063715018a61461024557600080fd5b806301a5d67e14610147578063026e402b1461015c578063030029021461016f5780631c9856ae1461018257806323245216146101b5575b600080fd5b61015a61015536600461160e565b61036c565b005b61015a61016a36600461160e565b610502565b61015a61017d366004611638565b6106d7565b6101a2610190366004611651565b60026020526000908152604090205481565b6040519081526020015b60405180910390f35b61015a6101c3366004611673565b61079a565b6101a262093a8081565b61015a61082f565b61015a6109f3565b6000805160206118928339815191525460ff165b60405190151581526020016101ac565b61022d7f000000000000000000000000000000000000000000000000000000000000000081565b6040516001600160a01b0390911681526020016101ac565b61015a610a03565b61015a610a15565b7f9016d09d72d40fdae2fd8ceac6b6234c7706214fd39c1cd1e609a0528c199300546001600160a01b031661022d565b61015a610293366004611638565b610a25565b61015a610bbc565b61015a6102ae366004611651565b610d5e565b6101a26102c13660046116e8565b600460209081526000928352604080842090915290825290205481565b6101a260005481565b6101f66102f5366004611651565b60016020526000908152604090205460ff1681565b610331610318366004611651565b6003602052600090815260409020805460019091015482565b604080519283526020830191909152016101ac565b61015a610354366004611673565b610ef1565b61015a610367366004611651565b610fe4565b61037461101f565b61037c611050565b33816103a35760405162461bcd60e51b815260040161039a9061171b565b60405180910390fd5b6001600160a01b038082166000908152600460209081526040808320938716835292905220548211156104285760405162461bcd60e51b815260206004820152602760248201527f5374616b696e673a20496e73756666696369656e742064656c65676174696f6e60448201526608185b5bdd5b9d60ca1b606482015260840161039a565b6001600160a01b0380821660009081526004602090815260408083209387168352929052908120805484929061045f908490611777565b9091555061049990506001600160a01b037f0000000000000000000000000000000000000000000000000000000000000000168284611088565b826001600160a01b0316816001600160a01b03167faa0001b0e2e76b9a1b925257fefa58b756aebc52d2d7c8a85ea5beacc77a2100846040516104de91815260200190565b60405180910390a3506104fe60016000805160206118b283398151915255565b5050565b61050a61101f565b610512611050565b33816105305760405162461bcd60e51b815260040161039a9061171b565b6001600160a01b03831660009081526001602052604090205460ff166105a25760405162461bcd60e51b815260206004820152602160248201527f5374616b696e673a204f70657261746f72206e6f742077686974656c697374656044820152601960fa1b606482015260840161039a565b600080546001600160a01b03851682526002602052604090912054101561061f5760405162461bcd60e51b815260206004820152602b60248201527f5374616b696e673a20496e73756666696369656e74206f70657261746f72207360448201526a1d185ad948185b5bdd5b9d60aa1b606482015260840161039a565b6106546001600160a01b037f0000000000000000000000000000000000000000000000000000000000000000168230856110ff565b6001600160a01b0380821660009081526004602090815260408083209387168352929052908120805484929061068b908490611790565b92505081905550826001600160a01b0316816001600160a01b03167f7816331f779ea2479e6f58bd6f82595b0beefd00953fe2a25702e915976befdc846040516104de91815260200190565b6106df61113d565b600081116107555760405162461bcd60e51b815260206004820152603d60248201527f5374616b696e673a206d696e696d756d206f70657261746f72207374616b652060448201527f616d6f756e74206d7573742062652067726561746572207468616e2030000000606482015260840161039a565b600080549082905560408051828152602081018490527ff6eefd49a83d72df0cc9ddb5557cfa6a047393b8d4a85163a8cefe7aa490a30f910160405180910390a15050565b6107a261113d565b60005b8181101561082a5760008383838181106107c1576107c16117a3565b90506020020160208101906107d69190611651565b6001600160a01b038116600081815260016020526040808220805460ff191690555192935090917fde8cf212af7ce38b2840785a2768d97ff2dbf3c21b516961cec0061e134c2a1e9190a2506001016107a5565b505050565b3360009081526001602052604090205460ff1661085e5760405162461bcd60e51b815260040161039a906117b9565b61086661101f565b61086e611050565b33600081815260026020526040902054806108cb5760405162461bcd60e51b815260206004820152601d60248201527f5374616b696e673a204e6f207374616b6520746f207769746864726177000000604482015260640161039a565b6001600160a01b0382166000908152600360205260409020541561093b5760405162461bcd60e51b815260206004820152602160248201527f5374616b696e673a204578697374696e6720756e7374616b65207265717565736044820152601d60fa1b606482015260840161039a565b6001600160a01b038216600090815260026020526040812081905561096362093a8042611790565b60408051808201825284815260208082018481526001600160a01b0388166000818152600384528590209351845590516001909301929092558251868152908101849052929350917f57e41df54512c76148b5ba9b643d149752b0d35e493b969bd017d0a3fe5228cf91015b60405180910390a25050506109f160016000805160206118b283398151915255565b565b6109fb61113d565b6109f1611198565b610a0b61113d565b6109f160006111f8565b610a1d61113d565b6109f1611269565b3360009081526001602052604090205460ff16610a545760405162461bcd60e51b815260040161039a906117b9565b610a5c61101f565b610a64611050565b3381610a825760405162461bcd60e51b815260040161039a9061171b565b600080546001600160a01b03831682526002602052604090912054610aa8908490611790565b1015610b015760405162461bcd60e51b815260206004820152602260248201527f5374616b696e673a20496e73756666696369656e74207374616b6520616d6f756044820152611b9d60f21b606482015260840161039a565b610b366001600160a01b037f0000000000000000000000000000000000000000000000000000000000000000168230856110ff565b6001600160a01b03811660009081526002602052604081208054849290610b5e908490611790565b90915550506040518281526001600160a01b038216907f0a7bb2e28cc4698aac06db79cf9163bfcc20719286cf59fa7d492ceda1b8edc29060200160405180910390a250610bb960016000805160206118b283398151915255565b50565b3360009081526001602052604090205460ff16610beb5760405162461bcd60e51b815260040161039a906117b9565b610bf361101f565b610bfb611050565b336000818152600360209081526040918290208251808401909352805480845260019091015491830191909152610c805760405162461bcd60e51b815260206004820152602360248201527f5374616b696e673a204e6f2070656e64696e6720756e7374616b652072657175604482015262195cdd60ea1b606482015260840161039a565b8060200151421015610cd45760405162461bcd60e51b815260206004820152601c60248201527f5374616b696e673a20546f6b656e73207374696c6c206c6f636b656400000000604482015260640161039a565b80516001600160a01b03808416600090815260036020526040812081815560010155610d23907f0000000000000000000000000000000000000000000000000000000000000000168483611088565b826001600160a01b03167f8108595eb6bad3acefa9da467d90cc2217686d5c5ac85460f8b7849c840645fc826040516109cf91815260200190565b7ff0c57e16840df040f15088dc2f81fe391c3923bec73e23a9662efc9c229c6a008054600160401b810460ff16159067ffffffffffffffff16600081158015610da45750825b905060008267ffffffffffffffff166001148015610dc15750303b155b905081158015610dcf575080155b15610ded5760405163f92ee8a960e01b815260040160405180910390fd5b845467ffffffffffffffff191660011785558315610e1757845460ff60401b1916600160401b1785555b6001600160a01b038616610e7c5760405162461bcd60e51b815260206004820152602660248201527f5374616b696e673a20496e76616c696420696e697469616c206f776e6572206160448201526564647265737360d01b606482015260840161039a565b610e85866112b2565b610e8d6112c3565b610e956112d3565b6969e10de76676d08000006000558315610ee957845460ff60401b19168555604051600181527fc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d29060200160405180910390a15b505050505050565b610ef961113d565b60005b8181101561082a576000838383818110610f1857610f186117a3565b9050602002016020810190610f2d9190611651565b90506001600160a01b038116610f8f5760405162461bcd60e51b815260206004820152602160248201527f5374616b696e673a20496e76616c6964206f70657261746f72206164647265736044820152607360f81b606482015260840161039a565b6001600160a01b0381166000818152600160208190526040808320805460ff1916909217909155517f4790a4adb426ca2345bb5108f6e454eae852a7bf687544cd66a7270dff3a41d69190a250600101610efc565b610fec61113d565b6001600160a01b03811661101657604051631e4fbdf760e01b81526000600482015260240161039a565b610bb9816111f8565b6000805160206118928339815191525460ff16156109f15760405163d93c066560e01b815260040160405180910390fd5b6000805160206118b283398151915280546001190161108257604051633ee5aeb560e01b815260040160405180910390fd5b60029055565b6040516001600160a01b03831660248201526044810182905261082a90849063a9059cbb60e01b906064015b60408051601f198184030181529190526020810180516001600160e01b03166001600160e01b0319909316929092179091526112e3565b60016000805160206118b283398151915255565b6040516001600160a01b03808516602483015283166044820152606481018290526111379085906323b872dd60e01b906084016110b4565b50505050565b3361116f7f9016d09d72d40fdae2fd8ceac6b6234c7706214fd39c1cd1e609a0528c199300546001600160a01b031690565b6001600160a01b0316146109f15760405163118cdaa760e01b815233600482015260240161039a565b6111a06113b8565b600080516020611892833981519152805460ff191681557f5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa335b6040516001600160a01b03909116815260200160405180910390a150565b7f9016d09d72d40fdae2fd8ceac6b6234c7706214fd39c1cd1e609a0528c19930080546001600160a01b031981166001600160a01b03848116918217845560405192169182907f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e090600090a3505050565b61127161101f565b600080516020611892833981519152805460ff191660011781557f62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258336111da565b6112ba6113e8565b610bb981611431565b6112cb6113e8565b6109f1611439565b6112db6113e8565b6109f161145a565b6000611338826040518060400160405280602081526020017f5361666545524332303a206c6f772d6c6576656c2063616c6c206661696c6564815250856001600160a01b03166114629092919063ffffffff16565b905080516000148061135957508080602001905181019061135991906117fc565b61082a5760405162461bcd60e51b815260206004820152602a60248201527f5361666545524332303a204552433230206f7065726174696f6e20646964206e6044820152691bdd081cdd58d8d9595960b21b606482015260840161039a565b6000805160206118928339815191525460ff166109f157604051638dfc202b60e01b815260040160405180910390fd5b7ff0c57e16840df040f15088dc2f81fe391c3923bec73e23a9662efc9c229c6a0054600160401b900460ff166109f157604051631afcd79f60e31b815260040160405180910390fd5b610fec6113e8565b6114416113e8565b600080516020611892833981519152805460ff19169055565b6110eb6113e8565b60606114718484600085611479565b949350505050565b6060824710156114da5760405162461bcd60e51b815260206004820152602660248201527f416464726573733a20696e73756666696369656e742062616c616e636520666f6044820152651c8818d85b1b60d21b606482015260840161039a565b600080866001600160a01b031685876040516114f69190611842565b60006040518083038185875af1925050503d8060008114611533576040519150601f19603f3d011682016040523d82523d6000602084013e611538565b606091505b509150915061154987838387611554565b979650505050505050565b606083156115c35782516000036115bc576001600160a01b0385163b6115bc5760405162461bcd60e51b815260206004820152601d60248201527f416464726573733a2063616c6c20746f206e6f6e2d636f6e7472616374000000604482015260640161039a565b5081611471565b61147183838151156115d85781518083602001fd5b8060405162461bcd60e51b815260040161039a919061185e565b80356001600160a01b038116811461160957600080fd5b919050565b6000806040838503121561162157600080fd5b61162a836115f2565b946020939093013593505050565b60006020828403121561164a57600080fd5b5035919050565b60006020828403121561166357600080fd5b61166c826115f2565b9392505050565b6000806020838503121561168657600080fd5b823567ffffffffffffffff8082111561169e57600080fd5b818501915085601f8301126116b257600080fd5b8135818111156116c157600080fd5b8660208260051b85010111156116d657600080fd5b60209290920196919550909350505050565b600080604083850312156116fb57600080fd5b611704836115f2565b9150611712602084016115f2565b90509250929050565b60208082526026908201527f5374616b696e673a20416d6f756e74206d75737420626520677265617465722060408201526507468616e20360d41b606082015260800190565b634e487b7160e01b600052601160045260246000fd5b8181038181111561178a5761178a611761565b92915050565b8082018082111561178a5761178a611761565b634e487b7160e01b600052603260045260246000fd5b60208082526023908201527f5374616b696e673a204e6f7420612077686974656c6973746564206f706572616040820152623a37b960e91b606082015260800190565b60006020828403121561180e57600080fd5b8151801515811461166c57600080fd5b60005b83811015611839578181015183820152602001611821565b50506000910152565b6000825161185481846020870161181e565b9190910192915050565b602081526000825180602084015261187d81604085016020870161181e565b601f01601f1916919091016040019291505056fecd5ed15c6e187e77e9aee88184c21f4f2182ab5827cb3b7e07fbedcd63f033009b779b17422d0df92223018b32b4d1fa46e071723d6817e2486d003becc55f00a26469706673582212203de80c6c5493e67e49838c0a609c708748409ace473817e40caf9b30a0c7fa4864736f6c63430008160033",
}

// StakingABI is the input ABI used to generate the binding from.
// Deprecated: Use StakingMetaData.ABI instead.
var StakingABI = StakingMetaData.ABI

// StakingBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use StakingMetaData.Bin instead.
var StakingBin = StakingMetaData.Bin

// DeployStaking deploys a new Ethereum contract, binding an instance of Staking to it.
func DeployStaking(auth *bind.TransactOpts, backend bind.ContractBackend, _cToken common.Address) (common.Address, *types.Transaction, *Staking, error) {
	parsed, err := StakingMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(StakingBin), backend, _cToken)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Staking{StakingCaller: StakingCaller{contract: contract}, StakingTransactor: StakingTransactor{contract: contract}, StakingFilterer: StakingFilterer{contract: contract}}, nil
}

// Staking is an auto generated Go binding around an Ethereum contract.
type Staking struct {
	StakingCaller     // Read-only binding to the contract
	StakingTransactor // Write-only binding to the contract
	StakingFilterer   // Log filterer for contract events
}

// StakingCaller is an auto generated read-only Go binding around an Ethereum contract.
type StakingCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StakingTransactor is an auto generated write-only Go binding around an Ethereum contract.
type StakingTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StakingFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type StakingFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StakingSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type StakingSession struct {
	Contract     *Staking          // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// StakingCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type StakingCallerSession struct {
	Contract *StakingCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts  // Call options to use throughout this session
}

// StakingTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type StakingTransactorSession struct {
	Contract     *StakingTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// StakingRaw is an auto generated low-level Go binding around an Ethereum contract.
type StakingRaw struct {
	Contract *Staking // Generic contract binding to access the raw methods on
}

// StakingCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type StakingCallerRaw struct {
	Contract *StakingCaller // Generic read-only contract binding to access the raw methods on
}

// StakingTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type StakingTransactorRaw struct {
	Contract *StakingTransactor // Generic write-only contract binding to access the raw methods on
}

// NewStaking creates a new instance of Staking, bound to a specific deployed contract.
func NewStaking(address common.Address, backend bind.ContractBackend) (*Staking, error) {
	contract, err := bindStaking(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Staking{StakingCaller: StakingCaller{contract: contract}, StakingTransactor: StakingTransactor{contract: contract}, StakingFilterer: StakingFilterer{contract: contract}}, nil
}

// NewStakingCaller creates a new read-only instance of Staking, bound to a specific deployed contract.
func NewStakingCaller(address common.Address, caller bind.ContractCaller) (*StakingCaller, error) {
	contract, err := bindStaking(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &StakingCaller{contract: contract}, nil
}

// NewStakingTransactor creates a new write-only instance of Staking, bound to a specific deployed contract.
func NewStakingTransactor(address common.Address, transactor bind.ContractTransactor) (*StakingTransactor, error) {
	contract, err := bindStaking(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &StakingTransactor{contract: contract}, nil
}

// NewStakingFilterer creates a new log filterer instance of Staking, bound to a specific deployed contract.
func NewStakingFilterer(address common.Address, filterer bind.ContractFilterer) (*StakingFilterer, error) {
	contract, err := bindStaking(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &StakingFilterer{contract: contract}, nil
}

// bindStaking binds a generic wrapper to an already deployed contract.
func bindStaking(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := StakingMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Staking *StakingRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Staking.Contract.StakingCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Staking *StakingRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Staking.Contract.StakingTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Staking *StakingRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Staking.Contract.StakingTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Staking *StakingCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Staking.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Staking *StakingTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Staking.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Staking *StakingTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Staking.Contract.contract.Transact(opts, method, params...)
}

// UNLOCKPERIOD is a free data retrieval call binding the contract method 0x259a28cf.
//
// Solidity: function UNLOCK_PERIOD() view returns(uint256)
func (_Staking *StakingCaller) UNLOCKPERIOD(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Staking.contract.Call(opts, &out, "UNLOCK_PERIOD")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// UNLOCKPERIOD is a free data retrieval call binding the contract method 0x259a28cf.
//
// Solidity: function UNLOCK_PERIOD() view returns(uint256)
func (_Staking *StakingSession) UNLOCKPERIOD() (*big.Int, error) {
	return _Staking.Contract.UNLOCKPERIOD(&_Staking.CallOpts)
}

// UNLOCKPERIOD is a free data retrieval call binding the contract method 0x259a28cf.
//
// Solidity: function UNLOCK_PERIOD() view returns(uint256)
func (_Staking *StakingCallerSession) UNLOCKPERIOD() (*big.Int, error) {
	return _Staking.Contract.UNLOCKPERIOD(&_Staking.CallOpts)
}

// CToken is a free data retrieval call binding the contract method 0x69e527da.
//
// Solidity: function cToken() view returns(address)
func (_Staking *StakingCaller) CToken(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Staking.contract.Call(opts, &out, "cToken")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// CToken is a free data retrieval call binding the contract method 0x69e527da.
//
// Solidity: function cToken() view returns(address)
func (_Staking *StakingSession) CToken() (common.Address, error) {
	return _Staking.Contract.CToken(&_Staking.CallOpts)
}

// CToken is a free data retrieval call binding the contract method 0x69e527da.
//
// Solidity: function cToken() view returns(address)
func (_Staking *StakingCallerSession) CToken() (common.Address, error) {
	return _Staking.Contract.CToken(&_Staking.CallOpts)
}

// Delegations is a free data retrieval call binding the contract method 0xc64814dd.
//
// Solidity: function delegations(address , address ) view returns(uint256)
func (_Staking *StakingCaller) Delegations(opts *bind.CallOpts, arg0 common.Address, arg1 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Staking.contract.Call(opts, &out, "delegations", arg0, arg1)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Delegations is a free data retrieval call binding the contract method 0xc64814dd.
//
// Solidity: function delegations(address , address ) view returns(uint256)
func (_Staking *StakingSession) Delegations(arg0 common.Address, arg1 common.Address) (*big.Int, error) {
	return _Staking.Contract.Delegations(&_Staking.CallOpts, arg0, arg1)
}

// Delegations is a free data retrieval call binding the contract method 0xc64814dd.
//
// Solidity: function delegations(address , address ) view returns(uint256)
func (_Staking *StakingCallerSession) Delegations(arg0 common.Address, arg1 common.Address) (*big.Int, error) {
	return _Staking.Contract.Delegations(&_Staking.CallOpts, arg0, arg1)
}

// MinOperatorStake is a free data retrieval call binding the contract method 0xd775cb61.
//
// Solidity: function minOperatorStake() view returns(uint256)
func (_Staking *StakingCaller) MinOperatorStake(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Staking.contract.Call(opts, &out, "minOperatorStake")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MinOperatorStake is a free data retrieval call binding the contract method 0xd775cb61.
//
// Solidity: function minOperatorStake() view returns(uint256)
func (_Staking *StakingSession) MinOperatorStake() (*big.Int, error) {
	return _Staking.Contract.MinOperatorStake(&_Staking.CallOpts)
}

// MinOperatorStake is a free data retrieval call binding the contract method 0xd775cb61.
//
// Solidity: function minOperatorStake() view returns(uint256)
func (_Staking *StakingCallerSession) MinOperatorStake() (*big.Int, error) {
	return _Staking.Contract.MinOperatorStake(&_Staking.CallOpts)
}

// OperatorStakes is a free data retrieval call binding the contract method 0x1c9856ae.
//
// Solidity: function operatorStakes(address ) view returns(uint256)
func (_Staking *StakingCaller) OperatorStakes(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Staking.contract.Call(opts, &out, "operatorStakes", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// OperatorStakes is a free data retrieval call binding the contract method 0x1c9856ae.
//
// Solidity: function operatorStakes(address ) view returns(uint256)
func (_Staking *StakingSession) OperatorStakes(arg0 common.Address) (*big.Int, error) {
	return _Staking.Contract.OperatorStakes(&_Staking.CallOpts, arg0)
}

// OperatorStakes is a free data retrieval call binding the contract method 0x1c9856ae.
//
// Solidity: function operatorStakes(address ) view returns(uint256)
func (_Staking *StakingCallerSession) OperatorStakes(arg0 common.Address) (*big.Int, error) {
	return _Staking.Contract.OperatorStakes(&_Staking.CallOpts, arg0)
}

// OperatorWhitelist is a free data retrieval call binding the contract method 0xe03c8632.
//
// Solidity: function operatorWhitelist(address ) view returns(bool)
func (_Staking *StakingCaller) OperatorWhitelist(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var out []interface{}
	err := _Staking.contract.Call(opts, &out, "operatorWhitelist", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// OperatorWhitelist is a free data retrieval call binding the contract method 0xe03c8632.
//
// Solidity: function operatorWhitelist(address ) view returns(bool)
func (_Staking *StakingSession) OperatorWhitelist(arg0 common.Address) (bool, error) {
	return _Staking.Contract.OperatorWhitelist(&_Staking.CallOpts, arg0)
}

// OperatorWhitelist is a free data retrieval call binding the contract method 0xe03c8632.
//
// Solidity: function operatorWhitelist(address ) view returns(bool)
func (_Staking *StakingCallerSession) OperatorWhitelist(arg0 common.Address) (bool, error) {
	return _Staking.Contract.OperatorWhitelist(&_Staking.CallOpts, arg0)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Staking *StakingCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Staking.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Staking *StakingSession) Owner() (common.Address, error) {
	return _Staking.Contract.Owner(&_Staking.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Staking *StakingCallerSession) Owner() (common.Address, error) {
	return _Staking.Contract.Owner(&_Staking.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_Staking *StakingCaller) Paused(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _Staking.contract.Call(opts, &out, "paused")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_Staking *StakingSession) Paused() (bool, error) {
	return _Staking.Contract.Paused(&_Staking.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_Staking *StakingCallerSession) Paused() (bool, error) {
	return _Staking.Contract.Paused(&_Staking.CallOpts)
}

// UnstakeRequests is a free data retrieval call binding the contract method 0xe0eb4d2e.
//
// Solidity: function unstakeRequests(address ) view returns(uint256 amount, uint256 unlockTime)
func (_Staking *StakingCaller) UnstakeRequests(opts *bind.CallOpts, arg0 common.Address) (struct {
	Amount     *big.Int
	UnlockTime *big.Int
}, error) {
	var out []interface{}
	err := _Staking.contract.Call(opts, &out, "unstakeRequests", arg0)

	outstruct := new(struct {
		Amount     *big.Int
		UnlockTime *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Amount = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.UnlockTime = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// UnstakeRequests is a free data retrieval call binding the contract method 0xe0eb4d2e.
//
// Solidity: function unstakeRequests(address ) view returns(uint256 amount, uint256 unlockTime)
func (_Staking *StakingSession) UnstakeRequests(arg0 common.Address) (struct {
	Amount     *big.Int
	UnlockTime *big.Int
}, error) {
	return _Staking.Contract.UnstakeRequests(&_Staking.CallOpts, arg0)
}

// UnstakeRequests is a free data retrieval call binding the contract method 0xe0eb4d2e.
//
// Solidity: function unstakeRequests(address ) view returns(uint256 amount, uint256 unlockTime)
func (_Staking *StakingCallerSession) UnstakeRequests(arg0 common.Address) (struct {
	Amount     *big.Int
	UnlockTime *big.Int
}, error) {
	return _Staking.Contract.UnstakeRequests(&_Staking.CallOpts, arg0)
}

// AddWhitelist is a paid mutator transaction binding the contract method 0xedac985b.
//
// Solidity: function addWhitelist(address[] operators) returns()
func (_Staking *StakingTransactor) AddWhitelist(opts *bind.TransactOpts, operators []common.Address) (*types.Transaction, error) {
	return _Staking.contract.Transact(opts, "addWhitelist", operators)
}

// AddWhitelist is a paid mutator transaction binding the contract method 0xedac985b.
//
// Solidity: function addWhitelist(address[] operators) returns()
func (_Staking *StakingSession) AddWhitelist(operators []common.Address) (*types.Transaction, error) {
	return _Staking.Contract.AddWhitelist(&_Staking.TransactOpts, operators)
}

// AddWhitelist is a paid mutator transaction binding the contract method 0xedac985b.
//
// Solidity: function addWhitelist(address[] operators) returns()
func (_Staking *StakingTransactorSession) AddWhitelist(operators []common.Address) (*types.Transaction, error) {
	return _Staking.Contract.AddWhitelist(&_Staking.TransactOpts, operators)
}

// Delegate is a paid mutator transaction binding the contract method 0x026e402b.
//
// Solidity: function delegate(address operator, uint256 amount) returns()
func (_Staking *StakingTransactor) Delegate(opts *bind.TransactOpts, operator common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Staking.contract.Transact(opts, "delegate", operator, amount)
}

// Delegate is a paid mutator transaction binding the contract method 0x026e402b.
//
// Solidity: function delegate(address operator, uint256 amount) returns()
func (_Staking *StakingSession) Delegate(operator common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Staking.Contract.Delegate(&_Staking.TransactOpts, operator, amount)
}

// Delegate is a paid mutator transaction binding the contract method 0x026e402b.
//
// Solidity: function delegate(address operator, uint256 amount) returns()
func (_Staking *StakingTransactorSession) Delegate(operator common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Staking.Contract.Delegate(&_Staking.TransactOpts, operator, amount)
}

// Initialize is a paid mutator transaction binding the contract method 0xc4d66de8.
//
// Solidity: function initialize(address initialOwner) returns()
func (_Staking *StakingTransactor) Initialize(opts *bind.TransactOpts, initialOwner common.Address) (*types.Transaction, error) {
	return _Staking.contract.Transact(opts, "initialize", initialOwner)
}

// Initialize is a paid mutator transaction binding the contract method 0xc4d66de8.
//
// Solidity: function initialize(address initialOwner) returns()
func (_Staking *StakingSession) Initialize(initialOwner common.Address) (*types.Transaction, error) {
	return _Staking.Contract.Initialize(&_Staking.TransactOpts, initialOwner)
}

// Initialize is a paid mutator transaction binding the contract method 0xc4d66de8.
//
// Solidity: function initialize(address initialOwner) returns()
func (_Staking *StakingTransactorSession) Initialize(initialOwner common.Address) (*types.Transaction, error) {
	return _Staking.Contract.Initialize(&_Staking.TransactOpts, initialOwner)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_Staking *StakingTransactor) Pause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Staking.contract.Transact(opts, "pause")
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_Staking *StakingSession) Pause() (*types.Transaction, error) {
	return _Staking.Contract.Pause(&_Staking.TransactOpts)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_Staking *StakingTransactorSession) Pause() (*types.Transaction, error) {
	return _Staking.Contract.Pause(&_Staking.TransactOpts)
}

// RemoveWhitelist is a paid mutator transaction binding the contract method 0x23245216.
//
// Solidity: function removeWhitelist(address[] operators) returns()
func (_Staking *StakingTransactor) RemoveWhitelist(opts *bind.TransactOpts, operators []common.Address) (*types.Transaction, error) {
	return _Staking.contract.Transact(opts, "removeWhitelist", operators)
}

// RemoveWhitelist is a paid mutator transaction binding the contract method 0x23245216.
//
// Solidity: function removeWhitelist(address[] operators) returns()
func (_Staking *StakingSession) RemoveWhitelist(operators []common.Address) (*types.Transaction, error) {
	return _Staking.Contract.RemoveWhitelist(&_Staking.TransactOpts, operators)
}

// RemoveWhitelist is a paid mutator transaction binding the contract method 0x23245216.
//
// Solidity: function removeWhitelist(address[] operators) returns()
func (_Staking *StakingTransactorSession) RemoveWhitelist(operators []common.Address) (*types.Transaction, error) {
	return _Staking.Contract.RemoveWhitelist(&_Staking.TransactOpts, operators)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Staking *StakingTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Staking.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Staking *StakingSession) RenounceOwnership() (*types.Transaction, error) {
	return _Staking.Contract.RenounceOwnership(&_Staking.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Staking *StakingTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _Staking.Contract.RenounceOwnership(&_Staking.TransactOpts)
}

// SetMinOperatorStake is a paid mutator transaction binding the contract method 0x03002902.
//
// Solidity: function setMinOperatorStake(uint256 _minOperatorStake) returns()
func (_Staking *StakingTransactor) SetMinOperatorStake(opts *bind.TransactOpts, _minOperatorStake *big.Int) (*types.Transaction, error) {
	return _Staking.contract.Transact(opts, "setMinOperatorStake", _minOperatorStake)
}

// SetMinOperatorStake is a paid mutator transaction binding the contract method 0x03002902.
//
// Solidity: function setMinOperatorStake(uint256 _minOperatorStake) returns()
func (_Staking *StakingSession) SetMinOperatorStake(_minOperatorStake *big.Int) (*types.Transaction, error) {
	return _Staking.Contract.SetMinOperatorStake(&_Staking.TransactOpts, _minOperatorStake)
}

// SetMinOperatorStake is a paid mutator transaction binding the contract method 0x03002902.
//
// Solidity: function setMinOperatorStake(uint256 _minOperatorStake) returns()
func (_Staking *StakingTransactorSession) SetMinOperatorStake(_minOperatorStake *big.Int) (*types.Transaction, error) {
	return _Staking.Contract.SetMinOperatorStake(&_Staking.TransactOpts, _minOperatorStake)
}

// Stake is a paid mutator transaction binding the contract method 0xa694fc3a.
//
// Solidity: function stake(uint256 amount) returns()
func (_Staking *StakingTransactor) Stake(opts *bind.TransactOpts, amount *big.Int) (*types.Transaction, error) {
	return _Staking.contract.Transact(opts, "stake", amount)
}

// Stake is a paid mutator transaction binding the contract method 0xa694fc3a.
//
// Solidity: function stake(uint256 amount) returns()
func (_Staking *StakingSession) Stake(amount *big.Int) (*types.Transaction, error) {
	return _Staking.Contract.Stake(&_Staking.TransactOpts, amount)
}

// Stake is a paid mutator transaction binding the contract method 0xa694fc3a.
//
// Solidity: function stake(uint256 amount) returns()
func (_Staking *StakingTransactorSession) Stake(amount *big.Int) (*types.Transaction, error) {
	return _Staking.Contract.Stake(&_Staking.TransactOpts, amount)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Staking *StakingTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _Staking.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Staking *StakingSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Staking.Contract.TransferOwnership(&_Staking.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Staking *StakingTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Staking.Contract.TransferOwnership(&_Staking.TransactOpts, newOwner)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_Staking *StakingTransactor) Unpause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Staking.contract.Transact(opts, "unpause")
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_Staking *StakingSession) Unpause() (*types.Transaction, error) {
	return _Staking.Contract.Unpause(&_Staking.TransactOpts)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_Staking *StakingTransactorSession) Unpause() (*types.Transaction, error) {
	return _Staking.Contract.Unpause(&_Staking.TransactOpts)
}

// Unstake is a paid mutator transaction binding the contract method 0x2def6620.
//
// Solidity: function unstake() returns()
func (_Staking *StakingTransactor) Unstake(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Staking.contract.Transact(opts, "unstake")
}

// Unstake is a paid mutator transaction binding the contract method 0x2def6620.
//
// Solidity: function unstake() returns()
func (_Staking *StakingSession) Unstake() (*types.Transaction, error) {
	return _Staking.Contract.Unstake(&_Staking.TransactOpts)
}

// Unstake is a paid mutator transaction binding the contract method 0x2def6620.
//
// Solidity: function unstake() returns()
func (_Staking *StakingTransactorSession) Unstake() (*types.Transaction, error) {
	return _Staking.Contract.Unstake(&_Staking.TransactOpts)
}

// WithdrawDelegation is a paid mutator transaction binding the contract method 0x01a5d67e.
//
// Solidity: function withdrawDelegation(address operator, uint256 amount) returns()
func (_Staking *StakingTransactor) WithdrawDelegation(opts *bind.TransactOpts, operator common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Staking.contract.Transact(opts, "withdrawDelegation", operator, amount)
}

// WithdrawDelegation is a paid mutator transaction binding the contract method 0x01a5d67e.
//
// Solidity: function withdrawDelegation(address operator, uint256 amount) returns()
func (_Staking *StakingSession) WithdrawDelegation(operator common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Staking.Contract.WithdrawDelegation(&_Staking.TransactOpts, operator, amount)
}

// WithdrawDelegation is a paid mutator transaction binding the contract method 0x01a5d67e.
//
// Solidity: function withdrawDelegation(address operator, uint256 amount) returns()
func (_Staking *StakingTransactorSession) WithdrawDelegation(operator common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Staking.Contract.WithdrawDelegation(&_Staking.TransactOpts, operator, amount)
}

// WithdrawStake is a paid mutator transaction binding the contract method 0xbed9d861.
//
// Solidity: function withdrawStake() returns()
func (_Staking *StakingTransactor) WithdrawStake(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Staking.contract.Transact(opts, "withdrawStake")
}

// WithdrawStake is a paid mutator transaction binding the contract method 0xbed9d861.
//
// Solidity: function withdrawStake() returns()
func (_Staking *StakingSession) WithdrawStake() (*types.Transaction, error) {
	return _Staking.Contract.WithdrawStake(&_Staking.TransactOpts)
}

// WithdrawStake is a paid mutator transaction binding the contract method 0xbed9d861.
//
// Solidity: function withdrawStake() returns()
func (_Staking *StakingTransactorSession) WithdrawStake() (*types.Transaction, error) {
	return _Staking.Contract.WithdrawStake(&_Staking.TransactOpts)
}

// StakingDelegationDepositedIterator is returned from FilterDelegationDeposited and is used to iterate over the raw logs and unpacked data for DelegationDeposited events raised by the Staking contract.
type StakingDelegationDepositedIterator struct {
	Event *StakingDelegationDeposited // Event containing the contract specifics and raw log

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
func (it *StakingDelegationDepositedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StakingDelegationDeposited)
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
		it.Event = new(StakingDelegationDeposited)
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
func (it *StakingDelegationDepositedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StakingDelegationDepositedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StakingDelegationDeposited represents a DelegationDeposited event raised by the Staking contract.
type StakingDelegationDeposited struct {
	Delegator common.Address
	Operator  common.Address
	Amount    *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterDelegationDeposited is a free log retrieval operation binding the contract event 0x7816331f779ea2479e6f58bd6f82595b0beefd00953fe2a25702e915976befdc.
//
// Solidity: event DelegationDeposited(address indexed delegator, address indexed operator, uint256 amount)
func (_Staking *StakingFilterer) FilterDelegationDeposited(opts *bind.FilterOpts, delegator []common.Address, operator []common.Address) (*StakingDelegationDepositedIterator, error) {

	var delegatorRule []interface{}
	for _, delegatorItem := range delegator {
		delegatorRule = append(delegatorRule, delegatorItem)
	}
	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _Staking.contract.FilterLogs(opts, "DelegationDeposited", delegatorRule, operatorRule)
	if err != nil {
		return nil, err
	}
	return &StakingDelegationDepositedIterator{contract: _Staking.contract, event: "DelegationDeposited", logs: logs, sub: sub}, nil
}

// WatchDelegationDeposited is a free log subscription operation binding the contract event 0x7816331f779ea2479e6f58bd6f82595b0beefd00953fe2a25702e915976befdc.
//
// Solidity: event DelegationDeposited(address indexed delegator, address indexed operator, uint256 amount)
func (_Staking *StakingFilterer) WatchDelegationDeposited(opts *bind.WatchOpts, sink chan<- *StakingDelegationDeposited, delegator []common.Address, operator []common.Address) (event.Subscription, error) {

	var delegatorRule []interface{}
	for _, delegatorItem := range delegator {
		delegatorRule = append(delegatorRule, delegatorItem)
	}
	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _Staking.contract.WatchLogs(opts, "DelegationDeposited", delegatorRule, operatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StakingDelegationDeposited)
				if err := _Staking.contract.UnpackLog(event, "DelegationDeposited", log); err != nil {
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

// ParseDelegationDeposited is a log parse operation binding the contract event 0x7816331f779ea2479e6f58bd6f82595b0beefd00953fe2a25702e915976befdc.
//
// Solidity: event DelegationDeposited(address indexed delegator, address indexed operator, uint256 amount)
func (_Staking *StakingFilterer) ParseDelegationDeposited(log types.Log) (*StakingDelegationDeposited, error) {
	event := new(StakingDelegationDeposited)
	if err := _Staking.contract.UnpackLog(event, "DelegationDeposited", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StakingDelegationWithdrawnIterator is returned from FilterDelegationWithdrawn and is used to iterate over the raw logs and unpacked data for DelegationWithdrawn events raised by the Staking contract.
type StakingDelegationWithdrawnIterator struct {
	Event *StakingDelegationWithdrawn // Event containing the contract specifics and raw log

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
func (it *StakingDelegationWithdrawnIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StakingDelegationWithdrawn)
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
		it.Event = new(StakingDelegationWithdrawn)
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
func (it *StakingDelegationWithdrawnIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StakingDelegationWithdrawnIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StakingDelegationWithdrawn represents a DelegationWithdrawn event raised by the Staking contract.
type StakingDelegationWithdrawn struct {
	Delegator common.Address
	Operator  common.Address
	Amount    *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterDelegationWithdrawn is a free log retrieval operation binding the contract event 0xaa0001b0e2e76b9a1b925257fefa58b756aebc52d2d7c8a85ea5beacc77a2100.
//
// Solidity: event DelegationWithdrawn(address indexed delegator, address indexed operator, uint256 amount)
func (_Staking *StakingFilterer) FilterDelegationWithdrawn(opts *bind.FilterOpts, delegator []common.Address, operator []common.Address) (*StakingDelegationWithdrawnIterator, error) {

	var delegatorRule []interface{}
	for _, delegatorItem := range delegator {
		delegatorRule = append(delegatorRule, delegatorItem)
	}
	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _Staking.contract.FilterLogs(opts, "DelegationWithdrawn", delegatorRule, operatorRule)
	if err != nil {
		return nil, err
	}
	return &StakingDelegationWithdrawnIterator{contract: _Staking.contract, event: "DelegationWithdrawn", logs: logs, sub: sub}, nil
}

// WatchDelegationWithdrawn is a free log subscription operation binding the contract event 0xaa0001b0e2e76b9a1b925257fefa58b756aebc52d2d7c8a85ea5beacc77a2100.
//
// Solidity: event DelegationWithdrawn(address indexed delegator, address indexed operator, uint256 amount)
func (_Staking *StakingFilterer) WatchDelegationWithdrawn(opts *bind.WatchOpts, sink chan<- *StakingDelegationWithdrawn, delegator []common.Address, operator []common.Address) (event.Subscription, error) {

	var delegatorRule []interface{}
	for _, delegatorItem := range delegator {
		delegatorRule = append(delegatorRule, delegatorItem)
	}
	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _Staking.contract.WatchLogs(opts, "DelegationWithdrawn", delegatorRule, operatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StakingDelegationWithdrawn)
				if err := _Staking.contract.UnpackLog(event, "DelegationWithdrawn", log); err != nil {
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

// ParseDelegationWithdrawn is a log parse operation binding the contract event 0xaa0001b0e2e76b9a1b925257fefa58b756aebc52d2d7c8a85ea5beacc77a2100.
//
// Solidity: event DelegationWithdrawn(address indexed delegator, address indexed operator, uint256 amount)
func (_Staking *StakingFilterer) ParseDelegationWithdrawn(log types.Log) (*StakingDelegationWithdrawn, error) {
	event := new(StakingDelegationWithdrawn)
	if err := _Staking.contract.UnpackLog(event, "DelegationWithdrawn", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StakingInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the Staking contract.
type StakingInitializedIterator struct {
	Event *StakingInitialized // Event containing the contract specifics and raw log

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
func (it *StakingInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StakingInitialized)
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
		it.Event = new(StakingInitialized)
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
func (it *StakingInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StakingInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StakingInitialized represents a Initialized event raised by the Staking contract.
type StakingInitialized struct {
	Version uint64
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0xc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2.
//
// Solidity: event Initialized(uint64 version)
func (_Staking *StakingFilterer) FilterInitialized(opts *bind.FilterOpts) (*StakingInitializedIterator, error) {

	logs, sub, err := _Staking.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &StakingInitializedIterator{contract: _Staking.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0xc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2.
//
// Solidity: event Initialized(uint64 version)
func (_Staking *StakingFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *StakingInitialized) (event.Subscription, error) {

	logs, sub, err := _Staking.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StakingInitialized)
				if err := _Staking.contract.UnpackLog(event, "Initialized", log); err != nil {
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

// ParseInitialized is a log parse operation binding the contract event 0xc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2.
//
// Solidity: event Initialized(uint64 version)
func (_Staking *StakingFilterer) ParseInitialized(log types.Log) (*StakingInitialized, error) {
	event := new(StakingInitialized)
	if err := _Staking.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StakingMinOperatorStakeUpdatedIterator is returned from FilterMinOperatorStakeUpdated and is used to iterate over the raw logs and unpacked data for MinOperatorStakeUpdated events raised by the Staking contract.
type StakingMinOperatorStakeUpdatedIterator struct {
	Event *StakingMinOperatorStakeUpdated // Event containing the contract specifics and raw log

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
func (it *StakingMinOperatorStakeUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StakingMinOperatorStakeUpdated)
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
		it.Event = new(StakingMinOperatorStakeUpdated)
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
func (it *StakingMinOperatorStakeUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StakingMinOperatorStakeUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StakingMinOperatorStakeUpdated represents a MinOperatorStakeUpdated event raised by the Staking contract.
type StakingMinOperatorStakeUpdated struct {
	OldAmount *big.Int
	NewAmount *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterMinOperatorStakeUpdated is a free log retrieval operation binding the contract event 0xf6eefd49a83d72df0cc9ddb5557cfa6a047393b8d4a85163a8cefe7aa490a30f.
//
// Solidity: event MinOperatorStakeUpdated(uint256 oldAmount, uint256 newAmount)
func (_Staking *StakingFilterer) FilterMinOperatorStakeUpdated(opts *bind.FilterOpts) (*StakingMinOperatorStakeUpdatedIterator, error) {

	logs, sub, err := _Staking.contract.FilterLogs(opts, "MinOperatorStakeUpdated")
	if err != nil {
		return nil, err
	}
	return &StakingMinOperatorStakeUpdatedIterator{contract: _Staking.contract, event: "MinOperatorStakeUpdated", logs: logs, sub: sub}, nil
}

// WatchMinOperatorStakeUpdated is a free log subscription operation binding the contract event 0xf6eefd49a83d72df0cc9ddb5557cfa6a047393b8d4a85163a8cefe7aa490a30f.
//
// Solidity: event MinOperatorStakeUpdated(uint256 oldAmount, uint256 newAmount)
func (_Staking *StakingFilterer) WatchMinOperatorStakeUpdated(opts *bind.WatchOpts, sink chan<- *StakingMinOperatorStakeUpdated) (event.Subscription, error) {

	logs, sub, err := _Staking.contract.WatchLogs(opts, "MinOperatorStakeUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StakingMinOperatorStakeUpdated)
				if err := _Staking.contract.UnpackLog(event, "MinOperatorStakeUpdated", log); err != nil {
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

// ParseMinOperatorStakeUpdated is a log parse operation binding the contract event 0xf6eefd49a83d72df0cc9ddb5557cfa6a047393b8d4a85163a8cefe7aa490a30f.
//
// Solidity: event MinOperatorStakeUpdated(uint256 oldAmount, uint256 newAmount)
func (_Staking *StakingFilterer) ParseMinOperatorStakeUpdated(log types.Log) (*StakingMinOperatorStakeUpdated, error) {
	event := new(StakingMinOperatorStakeUpdated)
	if err := _Staking.contract.UnpackLog(event, "MinOperatorStakeUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StakingOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the Staking contract.
type StakingOwnershipTransferredIterator struct {
	Event *StakingOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *StakingOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StakingOwnershipTransferred)
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
		it.Event = new(StakingOwnershipTransferred)
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
func (it *StakingOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StakingOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StakingOwnershipTransferred represents a OwnershipTransferred event raised by the Staking contract.
type StakingOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Staking *StakingFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*StakingOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Staking.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &StakingOwnershipTransferredIterator{contract: _Staking.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Staking *StakingFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *StakingOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Staking.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StakingOwnershipTransferred)
				if err := _Staking.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_Staking *StakingFilterer) ParseOwnershipTransferred(log types.Log) (*StakingOwnershipTransferred, error) {
	event := new(StakingOwnershipTransferred)
	if err := _Staking.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StakingPausedIterator is returned from FilterPaused and is used to iterate over the raw logs and unpacked data for Paused events raised by the Staking contract.
type StakingPausedIterator struct {
	Event *StakingPaused // Event containing the contract specifics and raw log

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
func (it *StakingPausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StakingPaused)
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
		it.Event = new(StakingPaused)
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
func (it *StakingPausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StakingPausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StakingPaused represents a Paused event raised by the Staking contract.
type StakingPaused struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterPaused is a free log retrieval operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_Staking *StakingFilterer) FilterPaused(opts *bind.FilterOpts) (*StakingPausedIterator, error) {

	logs, sub, err := _Staking.contract.FilterLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return &StakingPausedIterator{contract: _Staking.contract, event: "Paused", logs: logs, sub: sub}, nil
}

// WatchPaused is a free log subscription operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_Staking *StakingFilterer) WatchPaused(opts *bind.WatchOpts, sink chan<- *StakingPaused) (event.Subscription, error) {

	logs, sub, err := _Staking.contract.WatchLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StakingPaused)
				if err := _Staking.contract.UnpackLog(event, "Paused", log); err != nil {
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

// ParsePaused is a log parse operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_Staking *StakingFilterer) ParsePaused(log types.Log) (*StakingPaused, error) {
	event := new(StakingPaused)
	if err := _Staking.contract.UnpackLog(event, "Paused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StakingStakeDepositedIterator is returned from FilterStakeDeposited and is used to iterate over the raw logs and unpacked data for StakeDeposited events raised by the Staking contract.
type StakingStakeDepositedIterator struct {
	Event *StakingStakeDeposited // Event containing the contract specifics and raw log

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
func (it *StakingStakeDepositedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StakingStakeDeposited)
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
		it.Event = new(StakingStakeDeposited)
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
func (it *StakingStakeDepositedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StakingStakeDepositedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StakingStakeDeposited represents a StakeDeposited event raised by the Staking contract.
type StakingStakeDeposited struct {
	Operator common.Address
	Amount   *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterStakeDeposited is a free log retrieval operation binding the contract event 0x0a7bb2e28cc4698aac06db79cf9163bfcc20719286cf59fa7d492ceda1b8edc2.
//
// Solidity: event StakeDeposited(address indexed operator, uint256 amount)
func (_Staking *StakingFilterer) FilterStakeDeposited(opts *bind.FilterOpts, operator []common.Address) (*StakingStakeDepositedIterator, error) {

	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _Staking.contract.FilterLogs(opts, "StakeDeposited", operatorRule)
	if err != nil {
		return nil, err
	}
	return &StakingStakeDepositedIterator{contract: _Staking.contract, event: "StakeDeposited", logs: logs, sub: sub}, nil
}

// WatchStakeDeposited is a free log subscription operation binding the contract event 0x0a7bb2e28cc4698aac06db79cf9163bfcc20719286cf59fa7d492ceda1b8edc2.
//
// Solidity: event StakeDeposited(address indexed operator, uint256 amount)
func (_Staking *StakingFilterer) WatchStakeDeposited(opts *bind.WatchOpts, sink chan<- *StakingStakeDeposited, operator []common.Address) (event.Subscription, error) {

	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _Staking.contract.WatchLogs(opts, "StakeDeposited", operatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StakingStakeDeposited)
				if err := _Staking.contract.UnpackLog(event, "StakeDeposited", log); err != nil {
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

// ParseStakeDeposited is a log parse operation binding the contract event 0x0a7bb2e28cc4698aac06db79cf9163bfcc20719286cf59fa7d492ceda1b8edc2.
//
// Solidity: event StakeDeposited(address indexed operator, uint256 amount)
func (_Staking *StakingFilterer) ParseStakeDeposited(log types.Log) (*StakingStakeDeposited, error) {
	event := new(StakingStakeDeposited)
	if err := _Staking.contract.UnpackLog(event, "StakeDeposited", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StakingStakeWithdrawnIterator is returned from FilterStakeWithdrawn and is used to iterate over the raw logs and unpacked data for StakeWithdrawn events raised by the Staking contract.
type StakingStakeWithdrawnIterator struct {
	Event *StakingStakeWithdrawn // Event containing the contract specifics and raw log

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
func (it *StakingStakeWithdrawnIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StakingStakeWithdrawn)
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
		it.Event = new(StakingStakeWithdrawn)
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
func (it *StakingStakeWithdrawnIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StakingStakeWithdrawnIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StakingStakeWithdrawn represents a StakeWithdrawn event raised by the Staking contract.
type StakingStakeWithdrawn struct {
	Operator common.Address
	Amount   *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterStakeWithdrawn is a free log retrieval operation binding the contract event 0x8108595eb6bad3acefa9da467d90cc2217686d5c5ac85460f8b7849c840645fc.
//
// Solidity: event StakeWithdrawn(address indexed operator, uint256 amount)
func (_Staking *StakingFilterer) FilterStakeWithdrawn(opts *bind.FilterOpts, operator []common.Address) (*StakingStakeWithdrawnIterator, error) {

	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _Staking.contract.FilterLogs(opts, "StakeWithdrawn", operatorRule)
	if err != nil {
		return nil, err
	}
	return &StakingStakeWithdrawnIterator{contract: _Staking.contract, event: "StakeWithdrawn", logs: logs, sub: sub}, nil
}

// WatchStakeWithdrawn is a free log subscription operation binding the contract event 0x8108595eb6bad3acefa9da467d90cc2217686d5c5ac85460f8b7849c840645fc.
//
// Solidity: event StakeWithdrawn(address indexed operator, uint256 amount)
func (_Staking *StakingFilterer) WatchStakeWithdrawn(opts *bind.WatchOpts, sink chan<- *StakingStakeWithdrawn, operator []common.Address) (event.Subscription, error) {

	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _Staking.contract.WatchLogs(opts, "StakeWithdrawn", operatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StakingStakeWithdrawn)
				if err := _Staking.contract.UnpackLog(event, "StakeWithdrawn", log); err != nil {
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

// ParseStakeWithdrawn is a log parse operation binding the contract event 0x8108595eb6bad3acefa9da467d90cc2217686d5c5ac85460f8b7849c840645fc.
//
// Solidity: event StakeWithdrawn(address indexed operator, uint256 amount)
func (_Staking *StakingFilterer) ParseStakeWithdrawn(log types.Log) (*StakingStakeWithdrawn, error) {
	event := new(StakingStakeWithdrawn)
	if err := _Staking.contract.UnpackLog(event, "StakeWithdrawn", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StakingUnpausedIterator is returned from FilterUnpaused and is used to iterate over the raw logs and unpacked data for Unpaused events raised by the Staking contract.
type StakingUnpausedIterator struct {
	Event *StakingUnpaused // Event containing the contract specifics and raw log

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
func (it *StakingUnpausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StakingUnpaused)
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
		it.Event = new(StakingUnpaused)
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
func (it *StakingUnpausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StakingUnpausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StakingUnpaused represents a Unpaused event raised by the Staking contract.
type StakingUnpaused struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterUnpaused is a free log retrieval operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_Staking *StakingFilterer) FilterUnpaused(opts *bind.FilterOpts) (*StakingUnpausedIterator, error) {

	logs, sub, err := _Staking.contract.FilterLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return &StakingUnpausedIterator{contract: _Staking.contract, event: "Unpaused", logs: logs, sub: sub}, nil
}

// WatchUnpaused is a free log subscription operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_Staking *StakingFilterer) WatchUnpaused(opts *bind.WatchOpts, sink chan<- *StakingUnpaused) (event.Subscription, error) {

	logs, sub, err := _Staking.contract.WatchLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StakingUnpaused)
				if err := _Staking.contract.UnpackLog(event, "Unpaused", log); err != nil {
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

// ParseUnpaused is a log parse operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_Staking *StakingFilterer) ParseUnpaused(log types.Log) (*StakingUnpaused, error) {
	event := new(StakingUnpaused)
	if err := _Staking.contract.UnpackLog(event, "Unpaused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StakingUnstakeRequestedIterator is returned from FilterUnstakeRequested and is used to iterate over the raw logs and unpacked data for UnstakeRequested events raised by the Staking contract.
type StakingUnstakeRequestedIterator struct {
	Event *StakingUnstakeRequested // Event containing the contract specifics and raw log

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
func (it *StakingUnstakeRequestedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StakingUnstakeRequested)
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
		it.Event = new(StakingUnstakeRequested)
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
func (it *StakingUnstakeRequestedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StakingUnstakeRequestedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StakingUnstakeRequested represents a UnstakeRequested event raised by the Staking contract.
type StakingUnstakeRequested struct {
	Operator   common.Address
	Amount     *big.Int
	UnlockTime *big.Int
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterUnstakeRequested is a free log retrieval operation binding the contract event 0x57e41df54512c76148b5ba9b643d149752b0d35e493b969bd017d0a3fe5228cf.
//
// Solidity: event UnstakeRequested(address indexed operator, uint256 amount, uint256 unlockTime)
func (_Staking *StakingFilterer) FilterUnstakeRequested(opts *bind.FilterOpts, operator []common.Address) (*StakingUnstakeRequestedIterator, error) {

	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _Staking.contract.FilterLogs(opts, "UnstakeRequested", operatorRule)
	if err != nil {
		return nil, err
	}
	return &StakingUnstakeRequestedIterator{contract: _Staking.contract, event: "UnstakeRequested", logs: logs, sub: sub}, nil
}

// WatchUnstakeRequested is a free log subscription operation binding the contract event 0x57e41df54512c76148b5ba9b643d149752b0d35e493b969bd017d0a3fe5228cf.
//
// Solidity: event UnstakeRequested(address indexed operator, uint256 amount, uint256 unlockTime)
func (_Staking *StakingFilterer) WatchUnstakeRequested(opts *bind.WatchOpts, sink chan<- *StakingUnstakeRequested, operator []common.Address) (event.Subscription, error) {

	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _Staking.contract.WatchLogs(opts, "UnstakeRequested", operatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StakingUnstakeRequested)
				if err := _Staking.contract.UnpackLog(event, "UnstakeRequested", log); err != nil {
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

// ParseUnstakeRequested is a log parse operation binding the contract event 0x57e41df54512c76148b5ba9b643d149752b0d35e493b969bd017d0a3fe5228cf.
//
// Solidity: event UnstakeRequested(address indexed operator, uint256 amount, uint256 unlockTime)
func (_Staking *StakingFilterer) ParseUnstakeRequested(log types.Log) (*StakingUnstakeRequested, error) {
	event := new(StakingUnstakeRequested)
	if err := _Staking.contract.UnpackLog(event, "UnstakeRequested", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StakingWhitelistAddedIterator is returned from FilterWhitelistAdded and is used to iterate over the raw logs and unpacked data for WhitelistAdded events raised by the Staking contract.
type StakingWhitelistAddedIterator struct {
	Event *StakingWhitelistAdded // Event containing the contract specifics and raw log

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
func (it *StakingWhitelistAddedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StakingWhitelistAdded)
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
		it.Event = new(StakingWhitelistAdded)
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
func (it *StakingWhitelistAddedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StakingWhitelistAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StakingWhitelistAdded represents a WhitelistAdded event raised by the Staking contract.
type StakingWhitelistAdded struct {
	Operator common.Address
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterWhitelistAdded is a free log retrieval operation binding the contract event 0x4790a4adb426ca2345bb5108f6e454eae852a7bf687544cd66a7270dff3a41d6.
//
// Solidity: event WhitelistAdded(address indexed operator)
func (_Staking *StakingFilterer) FilterWhitelistAdded(opts *bind.FilterOpts, operator []common.Address) (*StakingWhitelistAddedIterator, error) {

	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _Staking.contract.FilterLogs(opts, "WhitelistAdded", operatorRule)
	if err != nil {
		return nil, err
	}
	return &StakingWhitelistAddedIterator{contract: _Staking.contract, event: "WhitelistAdded", logs: logs, sub: sub}, nil
}

// WatchWhitelistAdded is a free log subscription operation binding the contract event 0x4790a4adb426ca2345bb5108f6e454eae852a7bf687544cd66a7270dff3a41d6.
//
// Solidity: event WhitelistAdded(address indexed operator)
func (_Staking *StakingFilterer) WatchWhitelistAdded(opts *bind.WatchOpts, sink chan<- *StakingWhitelistAdded, operator []common.Address) (event.Subscription, error) {

	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _Staking.contract.WatchLogs(opts, "WhitelistAdded", operatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StakingWhitelistAdded)
				if err := _Staking.contract.UnpackLog(event, "WhitelistAdded", log); err != nil {
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

// ParseWhitelistAdded is a log parse operation binding the contract event 0x4790a4adb426ca2345bb5108f6e454eae852a7bf687544cd66a7270dff3a41d6.
//
// Solidity: event WhitelistAdded(address indexed operator)
func (_Staking *StakingFilterer) ParseWhitelistAdded(log types.Log) (*StakingWhitelistAdded, error) {
	event := new(StakingWhitelistAdded)
	if err := _Staking.contract.UnpackLog(event, "WhitelistAdded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StakingWhitelistRemovedIterator is returned from FilterWhitelistRemoved and is used to iterate over the raw logs and unpacked data for WhitelistRemoved events raised by the Staking contract.
type StakingWhitelistRemovedIterator struct {
	Event *StakingWhitelistRemoved // Event containing the contract specifics and raw log

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
func (it *StakingWhitelistRemovedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StakingWhitelistRemoved)
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
		it.Event = new(StakingWhitelistRemoved)
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
func (it *StakingWhitelistRemovedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StakingWhitelistRemovedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StakingWhitelistRemoved represents a WhitelistRemoved event raised by the Staking contract.
type StakingWhitelistRemoved struct {
	Operator common.Address
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterWhitelistRemoved is a free log retrieval operation binding the contract event 0xde8cf212af7ce38b2840785a2768d97ff2dbf3c21b516961cec0061e134c2a1e.
//
// Solidity: event WhitelistRemoved(address indexed operator)
func (_Staking *StakingFilterer) FilterWhitelistRemoved(opts *bind.FilterOpts, operator []common.Address) (*StakingWhitelistRemovedIterator, error) {

	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _Staking.contract.FilterLogs(opts, "WhitelistRemoved", operatorRule)
	if err != nil {
		return nil, err
	}
	return &StakingWhitelistRemovedIterator{contract: _Staking.contract, event: "WhitelistRemoved", logs: logs, sub: sub}, nil
}

// WatchWhitelistRemoved is a free log subscription operation binding the contract event 0xde8cf212af7ce38b2840785a2768d97ff2dbf3c21b516961cec0061e134c2a1e.
//
// Solidity: event WhitelistRemoved(address indexed operator)
func (_Staking *StakingFilterer) WatchWhitelistRemoved(opts *bind.WatchOpts, sink chan<- *StakingWhitelistRemoved, operator []common.Address) (event.Subscription, error) {

	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _Staking.contract.WatchLogs(opts, "WhitelistRemoved", operatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StakingWhitelistRemoved)
				if err := _Staking.contract.UnpackLog(event, "WhitelistRemoved", log); err != nil {
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

// ParseWhitelistRemoved is a log parse operation binding the contract event 0xde8cf212af7ce38b2840785a2768d97ff2dbf3c21b516961cec0061e134c2a1e.
//
// Solidity: event WhitelistRemoved(address indexed operator)
func (_Staking *StakingFilterer) ParseWhitelistRemoved(log types.Log) (*StakingWhitelistRemoved, error) {
	event := new(StakingWhitelistRemoved)
	if err := _Staking.contract.UnpackLog(event, "WhitelistRemoved", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
