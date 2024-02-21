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
)

// L1StandardBridgeMetaData contains all meta data concerning the L1StandardBridge contract.
var L1StandardBridgeMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"constructor\",\"inputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"MESSENGER\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractCrossDomainMessenger\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"OTHER_BRIDGE\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractStandardBridge\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"bridgeERC20\",\"inputs\":[{\"name\":\"_localToken\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_remoteToken\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"_minGasLimit\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"_extraData\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"bridgeERC20To\",\"inputs\":[{\"name\":\"_localToken\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_remoteToken\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_to\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"_minGasLimit\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"_extraData\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"depositERC20\",\"inputs\":[{\"name\":\"_l1Token\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_l2Token\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"_minGasLimit\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"_extraData\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"depositERC20To\",\"inputs\":[{\"name\":\"_l1Token\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_l2Token\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_to\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"_minGasLimit\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"_extraData\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"deposits\",\"inputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"finalizeBridgeERC20\",\"inputs\":[{\"name\":\"_localToken\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_remoteToken\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_from\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_to\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"_extraData\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"finalizeERC20Withdrawal\",\"inputs\":[{\"name\":\"_l1Token\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_l2Token\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_from\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_to\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"_extraData\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"initialize\",\"inputs\":[{\"name\":\"_messenger\",\"type\":\"address\",\"internalType\":\"contractCrossDomainMessenger\"},{\"name\":\"_superchainConfig\",\"type\":\"address\",\"internalType\":\"contractSuperchainConfig\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"l2TokenBridge\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"messenger\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractCrossDomainMessenger\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"otherBridge\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractStandardBridge\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"paused\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"superchainConfig\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractSuperchainConfig\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"version\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"view\"},{\"type\":\"event\",\"name\":\"ERC20BridgeFinalized\",\"inputs\":[{\"name\":\"localToken\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"remoteToken\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"from\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"to\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"extraData\",\"type\":\"bytes\",\"indexed\":false,\"internalType\":\"bytes\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"ERC20BridgeInitiated\",\"inputs\":[{\"name\":\"localToken\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"remoteToken\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"from\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"to\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"extraData\",\"type\":\"bytes\",\"indexed\":false,\"internalType\":\"bytes\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"ERC20DepositInitiated\",\"inputs\":[{\"name\":\"l1Token\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"l2Token\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"from\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"to\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"extraData\",\"type\":\"bytes\",\"indexed\":false,\"internalType\":\"bytes\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"ERC20WithdrawalFinalized\",\"inputs\":[{\"name\":\"l1Token\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"l2Token\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"from\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"to\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"extraData\",\"type\":\"bytes\",\"indexed\":false,\"internalType\":\"bytes\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Initialized\",\"inputs\":[{\"name\":\"version\",\"type\":\"uint8\",\"indexed\":false,\"internalType\":\"uint8\"}],\"anonymous\":false}]",
	Bin: "0x60806040523480156200001157600080fd5b506200001f60008062000025565b62000234565b600054610100900460ff1615808015620000465750600054600160ff909116105b8062000076575062000063306200018a60201b62000bd71760201c565b15801562000076575060005460ff166001145b620000df5760405162461bcd60e51b815260206004820152602e60248201527f496e697469616c697a61626c653a20636f6e747261637420697320616c72656160448201526d191e481a5b9a5d1a585b1a5e995960921b60648201526084015b60405180910390fd5b6000805460ff19166001179055801562000103576000805461ff0019166101001790555b603280546001600160a01b0319166001600160a01b0384161790556200013e8373420000000000000000000000000000000000001062000199565b801562000185576000805461ff0019169055604051600181527f7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb38474024989060200160405180910390a15b505050565b6001600160a01b03163b151590565b600054610100900460ff16620002065760405162461bcd60e51b815260206004820152602b60248201527f496e697469616c697a61626c653a20636f6e7472616374206973206e6f74206960448201526a6e697469616c697a696e6760a81b6064820152608401620000d6565b600380546001600160a01b039384166001600160a01b03199182161790915560048054929093169116179055565b611dbe80620002446000396000f3fe608060405234801561001057600080fd5b50600436106101005760003560e01c80637f46ddb21161009757806391c49bf81161006657806391c49bf81461021e578063927ede2d1461029b578063a9f9e675146102b9578063c89701a2146102cc57600080fd5b80637f46ddb21461021e578063838b25201461023c578063870876231461024f5780638f601f661461026257600080fd5b8063540abf73116100d3578063540abf731461019757806354fd4d50146101aa57806358a997f6146101f35780635c975abb1461020657600080fd5b80630166a07a1461010557806335e80ab31461011a5780633cb747bf14610164578063485cc95514610184575b600080fd5b6101186101133660046119ad565b6102ec565b005b60325461013a9073ffffffffffffffffffffffffffffffffffffffff1681565b60405173ffffffffffffffffffffffffffffffffffffffff90911681526020015b60405180910390f35b60035461013a9073ffffffffffffffffffffffffffffffffffffffff1681565b610118610192366004611a45565b61070b565b6101186101a5366004611a97565b6108f5565b6101e66040518060400160405280600581526020017f322e312e3000000000000000000000000000000000000000000000000000000081525081565b60405161015b9190611b84565b610118610201366004611b97565b61093a565b61020e610a16565b604051901515815260200161015b565b60045473ffffffffffffffffffffffffffffffffffffffff1661013a565b61011861024a366004611a97565b610aaf565b61011861025d366004611b97565b610af4565b61028d610270366004611a45565b600260209081526000928352604080842090915290825290205481565b60405190815260200161015b565b60035473ffffffffffffffffffffffffffffffffffffffff1661013a565b6101186102c73660046119ad565b610bc8565b60045461013a9073ffffffffffffffffffffffffffffffffffffffff1681565b60035473ffffffffffffffffffffffffffffffffffffffff16331480156103bf575060048054600354604080517f6e296e45000000000000000000000000000000000000000000000000000000008152905173ffffffffffffffffffffffffffffffffffffffff938416949390921692636e296e459282820192602092908290030181865afa158015610383573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906103a79190611c1a565b73ffffffffffffffffffffffffffffffffffffffff16145b610476576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152604160248201527f5374616e646172644272696467653a2066756e6374696f6e2063616e206f6e6c60448201527f792062652063616c6c65642066726f6d20746865206f7468657220627269646760648201527f6500000000000000000000000000000000000000000000000000000000000000608482015260a4015b60405180910390fd5b61047e610a16565b156104e5576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601660248201527f5374616e646172644272696467653a2070617573656400000000000000000000604482015260640161046d565b6104ee87610bf3565b1561063c576104fd8787610c55565b6105af576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152604a60248201527f5374616e646172644272696467653a2077726f6e672072656d6f746520746f6b60448201527f656e20666f72204f7074696d69736d204d696e7461626c65204552433230206c60648201527f6f63616c20746f6b656e00000000000000000000000000000000000000000000608482015260a40161046d565b6040517f40c10f1900000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff8581166004830152602482018590528816906340c10f1990604401600060405180830381600087803b15801561061f57600080fd5b505af1158015610633573d6000803e3d6000fd5b505050506106be565b73ffffffffffffffffffffffffffffffffffffffff8088166000908152600260209081526040808320938a168352929052205461067a908490611c66565b73ffffffffffffffffffffffffffffffffffffffff8089166000818152600260209081526040808320948c16835293905291909120919091556106be908585610d75565b610702878787878787878080601f016020809104026020016040519081016040528093929190818152602001838380828437600092019190915250610e4992505050565b50505050505050565b600054610100900460ff161580801561072b5750600054600160ff909116105b806107455750303b158015610745575060005460ff166001145b6107d1576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602e60248201527f496e697469616c697a61626c653a20636f6e747261637420697320616c72656160448201527f647920696e697469616c697a6564000000000000000000000000000000000000606482015260840161046d565b600080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00166001179055801561082f57600080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00ff166101001790555b603280547fffffffffffffffffffffffff00000000000000000000000000000000000000001673ffffffffffffffffffffffffffffffffffffffff841617905561088d83734200000000000000000000000000000000000010610ed7565b80156108f057600080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00ff169055604051600181527f7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb38474024989060200160405180910390a15b505050565b61070287873388888888888080601f016020809104026020016040519081016040528093929190818152602001838380828437600092019190915250610fc192505050565b333b156109c9576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152603760248201527f5374616e646172644272696467653a2066756e6374696f6e2063616e206f6e6c60448201527f792062652063616c6c65642066726f6d20616e20454f41000000000000000000606482015260840161046d565b610a0e86863333888888888080601f0160208091040260200160405190810160405280939291908181526020018383808284376000920191909152506112ec92505050565b505050505050565b603254604080517f5c975abb000000000000000000000000000000000000000000000000000000008152905160009273ffffffffffffffffffffffffffffffffffffffff1691635c975abb9160048083019260209291908290030181865afa158015610a86573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610aaa9190611c7d565b905090565b61070287873388888888888080601f0160208091040260200160405190810160405280939291908181526020018383808284376000920191909152506112ec92505050565b333b15610b83576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152603760248201527f5374616e646172644272696467653a2066756e6374696f6e2063616e206f6e6c60448201527f792062652063616c6c65642066726f6d20616e20454f41000000000000000000606482015260840161046d565b610a0e86863333888888888080601f016020809104026020016040519081016040528093929190818152602001838380828437600092019190915250610fc192505050565b610702878787878787876102ec565b73ffffffffffffffffffffffffffffffffffffffff163b151590565b6000610c1f827f1d1d8b63000000000000000000000000000000000000000000000000000000006112fb565b80610c4f5750610c4f827fec4fc8e3000000000000000000000000000000000000000000000000000000006112fb565b92915050565b6000610c81837f1d1d8b63000000000000000000000000000000000000000000000000000000006112fb565b15610d2a578273ffffffffffffffffffffffffffffffffffffffff1663c01e1bd66040518163ffffffff1660e01b8152600401602060405180830381865afa158015610cd1573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610cf59190611c1a565b73ffffffffffffffffffffffffffffffffffffffff168273ffffffffffffffffffffffffffffffffffffffff16149050610c4f565b8273ffffffffffffffffffffffffffffffffffffffff1663d6c0b2c46040518163ffffffff1660e01b8152600401602060405180830381865afa158015610cd1573d6000803e3d6000fd5b60405173ffffffffffffffffffffffffffffffffffffffff83166024820152604481018290526108f09084907fa9059cbb00000000000000000000000000000000000000000000000000000000906064015b604080517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe08184030181529190526020810180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff167fffffffff000000000000000000000000000000000000000000000000000000009093169290921790915261131e565b8373ffffffffffffffffffffffffffffffffffffffff168573ffffffffffffffffffffffffffffffffffffffff168773ffffffffffffffffffffffffffffffffffffffff167f3ceee06c1e37648fcbb6ed52e17b3e1f275a1f8c7b22a84b2b84732431e046b3868686604051610ec193929190611c9f565b60405180910390a4610a0e86868686868661142a565b600054610100900460ff16610f6e576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602b60248201527f496e697469616c697a61626c653a20636f6e7472616374206973206e6f74206960448201527f6e697469616c697a696e67000000000000000000000000000000000000000000606482015260840161046d565b6003805473ffffffffffffffffffffffffffffffffffffffff9384167fffffffffffffffffffffffff00000000000000000000000000000000000000009182161790915560048054929093169116179055565b610fca87610bf3565b1561111857610fd98787610c55565b61108b576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152604a60248201527f5374616e646172644272696467653a2077726f6e672072656d6f746520746f6b60448201527f656e20666f72204f7074696d69736d204d696e7461626c65204552433230206c60648201527f6f63616c20746f6b656e00000000000000000000000000000000000000000000608482015260a40161046d565b6040517f9dc29fac00000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff868116600483015260248201859052881690639dc29fac90604401600060405180830381600087803b1580156110fb57600080fd5b505af115801561110f573d6000803e3d6000fd5b505050506111ac565b61113a73ffffffffffffffffffffffffffffffffffffffff88168630866114b2565b73ffffffffffffffffffffffffffffffffffffffff8088166000908152600260209081526040808320938a1683529290522054611178908490611cdd565b73ffffffffffffffffffffffffffffffffffffffff8089166000908152600260209081526040808320938b16835292905220555b6111ba878787878786611516565b60035460045460405173ffffffffffffffffffffffffffffffffffffffff92831692633dbb202b9216907f0166a07a000000000000000000000000000000000000000000000000000000009061121e908b908d908c908c908c908b90602401611cf5565b604080517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe08184030181529181526020820180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff167fffffffff000000000000000000000000000000000000000000000000000000009485161790525160e085901b90921682526112b192918790600401611d50565b600060405180830381600087803b1580156112cb57600080fd5b505af11580156112df573d6000803e3d6000fd5b5050505050505050505050565b61070287878787878787610fc1565b6000611306836115a4565b801561131757506113178383611608565b9392505050565b6000611380826040518060400160405280602081526020017f5361666545524332303a206c6f772d6c6576656c2063616c6c206661696c65648152508573ffffffffffffffffffffffffffffffffffffffff166116d79092919063ffffffff16565b8051909150156108f0578080602001905181019061139e9190611c7d565b6108f0576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602a60248201527f5361666545524332303a204552433230206f7065726174696f6e20646964206e60448201527f6f74207375636365656400000000000000000000000000000000000000000000606482015260840161046d565b8373ffffffffffffffffffffffffffffffffffffffff168573ffffffffffffffffffffffffffffffffffffffff168773ffffffffffffffffffffffffffffffffffffffff167fd59c65b35445225835c83f50b6ede06a7be047d22e357073e250d9af537518cd8686866040516114a293929190611c9f565b60405180910390a4505050505050565b60405173ffffffffffffffffffffffffffffffffffffffff808516602483015283166044820152606481018290526115109085907f23b872dd0000000000000000000000000000000000000000000000000000000090608401610dc7565b50505050565b8373ffffffffffffffffffffffffffffffffffffffff168573ffffffffffffffffffffffffffffffffffffffff168773ffffffffffffffffffffffffffffffffffffffff167f718594027abd4eaed59f95162563e0cc6d0e8d5b86b1c7be8b1b0ac3343d039686868660405161158e93929190611c9f565b60405180910390a4610a0e8686868686866116ee565b60006115d0827f01ffc9a700000000000000000000000000000000000000000000000000000000611608565b8015610c4f5750611601827fffffffff00000000000000000000000000000000000000000000000000000000611608565b1592915050565b604080517fffffffff000000000000000000000000000000000000000000000000000000008316602480830191909152825180830390910181526044909101909152602080820180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff167f01ffc9a700000000000000000000000000000000000000000000000000000000178152825160009392849283928392918391908a617530fa92503d915060005190508280156116c0575060208210155b80156116cc5750600081115b979650505050505050565b60606116e68484600085611766565b949350505050565b8373ffffffffffffffffffffffffffffffffffffffff168573ffffffffffffffffffffffffffffffffffffffff168773ffffffffffffffffffffffffffffffffffffffff167f7ff126db8024424bbfd9826e8ab82ff59136289ea440b04b39a0df1b03b9cabf8686866040516114a293929190611c9f565b6060824710156117f8576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602660248201527f416464726573733a20696e73756666696369656e742062616c616e636520666f60448201527f722063616c6c0000000000000000000000000000000000000000000000000000606482015260840161046d565b73ffffffffffffffffffffffffffffffffffffffff85163b611876576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601d60248201527f416464726573733a2063616c6c20746f206e6f6e2d636f6e7472616374000000604482015260640161046d565b6000808673ffffffffffffffffffffffffffffffffffffffff16858760405161189f9190611d95565b60006040518083038185875af1925050503d80600081146118dc576040519150601f19603f3d011682016040523d82523d6000602084013e6118e1565b606091505b50915091506116cc828286606083156118fb575081611317565b82511561190b5782518084602001fd5b816040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161046d9190611b84565b73ffffffffffffffffffffffffffffffffffffffff8116811461196157600080fd5b50565b60008083601f84011261197657600080fd5b50813567ffffffffffffffff81111561198e57600080fd5b6020830191508360208285010111156119a657600080fd5b9250929050565b600080600080600080600060c0888a0312156119c857600080fd5b87356119d38161193f565b965060208801356119e38161193f565b955060408801356119f38161193f565b94506060880135611a038161193f565b93506080880135925060a088013567ffffffffffffffff811115611a2657600080fd5b611a328a828b01611964565b989b979a50959850939692959293505050565b60008060408385031215611a5857600080fd5b8235611a638161193f565b91506020830135611a738161193f565b809150509250929050565b803563ffffffff81168114611a9257600080fd5b919050565b600080600080600080600060c0888a031215611ab257600080fd5b8735611abd8161193f565b96506020880135611acd8161193f565b95506040880135611add8161193f565b945060608801359350611af260808901611a7e565b925060a088013567ffffffffffffffff811115611a2657600080fd5b60005b83811015611b29578181015183820152602001611b11565b838111156115105750506000910152565b60008151808452611b52816020860160208601611b0e565b601f017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0169290920160200192915050565b6020815260006113176020830184611b3a565b60008060008060008060a08789031215611bb057600080fd5b8635611bbb8161193f565b95506020870135611bcb8161193f565b945060408701359350611be060608801611a7e565b9250608087013567ffffffffffffffff811115611bfc57600080fd5b611c0889828a01611964565b979a9699509497509295939492505050565b600060208284031215611c2c57600080fd5b81516113178161193f565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b600082821015611c7857611c78611c37565b500390565b600060208284031215611c8f57600080fd5b8151801515811461131757600080fd5b73ffffffffffffffffffffffffffffffffffffffff84168152826020820152606060408201526000611cd46060830184611b3a565b95945050505050565b60008219821115611cf057611cf0611c37565b500190565b600073ffffffffffffffffffffffffffffffffffffffff80891683528088166020840152808716604084015280861660608401525083608083015260c060a0830152611d4460c0830184611b3a565b98975050505050505050565b73ffffffffffffffffffffffffffffffffffffffff84168152606060208201526000611d7f6060830185611b3a565b905063ffffffff83166040830152949350505050565b60008251611da7818460208701611b0e565b919091019291505056fea164736f6c634300080f000a",
}

// L1StandardBridgeABI is the input ABI used to generate the binding from.
// Deprecated: Use L1StandardBridgeMetaData.ABI instead.
var L1StandardBridgeABI = L1StandardBridgeMetaData.ABI

// L1StandardBridgeBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use L1StandardBridgeMetaData.Bin instead.
var L1StandardBridgeBin = L1StandardBridgeMetaData.Bin

// DeployL1StandardBridge deploys a new Ethereum contract, binding an instance of L1StandardBridge to it.
func DeployL1StandardBridge(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *L1StandardBridge, error) {
	parsed, err := L1StandardBridgeMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(L1StandardBridgeBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &L1StandardBridge{L1StandardBridgeCaller: L1StandardBridgeCaller{contract: contract}, L1StandardBridgeTransactor: L1StandardBridgeTransactor{contract: contract}, L1StandardBridgeFilterer: L1StandardBridgeFilterer{contract: contract}}, nil
}

// L1StandardBridge is an auto generated Go binding around an Ethereum contract.
type L1StandardBridge struct {
	L1StandardBridgeCaller     // Read-only binding to the contract
	L1StandardBridgeTransactor // Write-only binding to the contract
	L1StandardBridgeFilterer   // Log filterer for contract events
}

// L1StandardBridgeCaller is an auto generated read-only Go binding around an Ethereum contract.
type L1StandardBridgeCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// L1StandardBridgeTransactor is an auto generated write-only Go binding around an Ethereum contract.
type L1StandardBridgeTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// L1StandardBridgeFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type L1StandardBridgeFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// L1StandardBridgeSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type L1StandardBridgeSession struct {
	Contract     *L1StandardBridge // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// L1StandardBridgeCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type L1StandardBridgeCallerSession struct {
	Contract *L1StandardBridgeCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts           // Call options to use throughout this session
}

// L1StandardBridgeTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type L1StandardBridgeTransactorSession struct {
	Contract     *L1StandardBridgeTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts           // Transaction auth options to use throughout this session
}

// L1StandardBridgeRaw is an auto generated low-level Go binding around an Ethereum contract.
type L1StandardBridgeRaw struct {
	Contract *L1StandardBridge // Generic contract binding to access the raw methods on
}

// L1StandardBridgeCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type L1StandardBridgeCallerRaw struct {
	Contract *L1StandardBridgeCaller // Generic read-only contract binding to access the raw methods on
}

// L1StandardBridgeTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type L1StandardBridgeTransactorRaw struct {
	Contract *L1StandardBridgeTransactor // Generic write-only contract binding to access the raw methods on
}

// NewL1StandardBridge creates a new instance of L1StandardBridge, bound to a specific deployed contract.
func NewL1StandardBridge(address common.Address, backend bind.ContractBackend) (*L1StandardBridge, error) {
	contract, err := bindL1StandardBridge(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &L1StandardBridge{L1StandardBridgeCaller: L1StandardBridgeCaller{contract: contract}, L1StandardBridgeTransactor: L1StandardBridgeTransactor{contract: contract}, L1StandardBridgeFilterer: L1StandardBridgeFilterer{contract: contract}}, nil
}

// NewL1StandardBridgeCaller creates a new read-only instance of L1StandardBridge, bound to a specific deployed contract.
func NewL1StandardBridgeCaller(address common.Address, caller bind.ContractCaller) (*L1StandardBridgeCaller, error) {
	contract, err := bindL1StandardBridge(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &L1StandardBridgeCaller{contract: contract}, nil
}

// NewL1StandardBridgeTransactor creates a new write-only instance of L1StandardBridge, bound to a specific deployed contract.
func NewL1StandardBridgeTransactor(address common.Address, transactor bind.ContractTransactor) (*L1StandardBridgeTransactor, error) {
	contract, err := bindL1StandardBridge(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &L1StandardBridgeTransactor{contract: contract}, nil
}

// NewL1StandardBridgeFilterer creates a new log filterer instance of L1StandardBridge, bound to a specific deployed contract.
func NewL1StandardBridgeFilterer(address common.Address, filterer bind.ContractFilterer) (*L1StandardBridgeFilterer, error) {
	contract, err := bindL1StandardBridge(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &L1StandardBridgeFilterer{contract: contract}, nil
}

// bindL1StandardBridge binds a generic wrapper to an already deployed contract.
func bindL1StandardBridge(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(L1StandardBridgeABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_L1StandardBridge *L1StandardBridgeRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _L1StandardBridge.Contract.L1StandardBridgeCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_L1StandardBridge *L1StandardBridgeRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _L1StandardBridge.Contract.L1StandardBridgeTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_L1StandardBridge *L1StandardBridgeRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _L1StandardBridge.Contract.L1StandardBridgeTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_L1StandardBridge *L1StandardBridgeCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _L1StandardBridge.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_L1StandardBridge *L1StandardBridgeTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _L1StandardBridge.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_L1StandardBridge *L1StandardBridgeTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _L1StandardBridge.Contract.contract.Transact(opts, method, params...)
}

// MESSENGER is a free data retrieval call binding the contract method 0x927ede2d.
//
// Solidity: function MESSENGER() view returns(address)
func (_L1StandardBridge *L1StandardBridgeCaller) MESSENGER(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _L1StandardBridge.contract.Call(opts, &out, "MESSENGER")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// MESSENGER is a free data retrieval call binding the contract method 0x927ede2d.
//
// Solidity: function MESSENGER() view returns(address)
func (_L1StandardBridge *L1StandardBridgeSession) MESSENGER() (common.Address, error) {
	return _L1StandardBridge.Contract.MESSENGER(&_L1StandardBridge.CallOpts)
}

// MESSENGER is a free data retrieval call binding the contract method 0x927ede2d.
//
// Solidity: function MESSENGER() view returns(address)
func (_L1StandardBridge *L1StandardBridgeCallerSession) MESSENGER() (common.Address, error) {
	return _L1StandardBridge.Contract.MESSENGER(&_L1StandardBridge.CallOpts)
}

// OTHERBRIDGE is a free data retrieval call binding the contract method 0x7f46ddb2.
//
// Solidity: function OTHER_BRIDGE() view returns(address)
func (_L1StandardBridge *L1StandardBridgeCaller) OTHERBRIDGE(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _L1StandardBridge.contract.Call(opts, &out, "OTHER_BRIDGE")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// OTHERBRIDGE is a free data retrieval call binding the contract method 0x7f46ddb2.
//
// Solidity: function OTHER_BRIDGE() view returns(address)
func (_L1StandardBridge *L1StandardBridgeSession) OTHERBRIDGE() (common.Address, error) {
	return _L1StandardBridge.Contract.OTHERBRIDGE(&_L1StandardBridge.CallOpts)
}

// OTHERBRIDGE is a free data retrieval call binding the contract method 0x7f46ddb2.
//
// Solidity: function OTHER_BRIDGE() view returns(address)
func (_L1StandardBridge *L1StandardBridgeCallerSession) OTHERBRIDGE() (common.Address, error) {
	return _L1StandardBridge.Contract.OTHERBRIDGE(&_L1StandardBridge.CallOpts)
}

// Deposits is a free data retrieval call binding the contract method 0x8f601f66.
//
// Solidity: function deposits(address , address ) view returns(uint256)
func (_L1StandardBridge *L1StandardBridgeCaller) Deposits(opts *bind.CallOpts, arg0 common.Address, arg1 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _L1StandardBridge.contract.Call(opts, &out, "deposits", arg0, arg1)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Deposits is a free data retrieval call binding the contract method 0x8f601f66.
//
// Solidity: function deposits(address , address ) view returns(uint256)
func (_L1StandardBridge *L1StandardBridgeSession) Deposits(arg0 common.Address, arg1 common.Address) (*big.Int, error) {
	return _L1StandardBridge.Contract.Deposits(&_L1StandardBridge.CallOpts, arg0, arg1)
}

// Deposits is a free data retrieval call binding the contract method 0x8f601f66.
//
// Solidity: function deposits(address , address ) view returns(uint256)
func (_L1StandardBridge *L1StandardBridgeCallerSession) Deposits(arg0 common.Address, arg1 common.Address) (*big.Int, error) {
	return _L1StandardBridge.Contract.Deposits(&_L1StandardBridge.CallOpts, arg0, arg1)
}

// L2TokenBridge is a free data retrieval call binding the contract method 0x91c49bf8.
//
// Solidity: function l2TokenBridge() view returns(address)
func (_L1StandardBridge *L1StandardBridgeCaller) L2TokenBridge(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _L1StandardBridge.contract.Call(opts, &out, "l2TokenBridge")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// L2TokenBridge is a free data retrieval call binding the contract method 0x91c49bf8.
//
// Solidity: function l2TokenBridge() view returns(address)
func (_L1StandardBridge *L1StandardBridgeSession) L2TokenBridge() (common.Address, error) {
	return _L1StandardBridge.Contract.L2TokenBridge(&_L1StandardBridge.CallOpts)
}

// L2TokenBridge is a free data retrieval call binding the contract method 0x91c49bf8.
//
// Solidity: function l2TokenBridge() view returns(address)
func (_L1StandardBridge *L1StandardBridgeCallerSession) L2TokenBridge() (common.Address, error) {
	return _L1StandardBridge.Contract.L2TokenBridge(&_L1StandardBridge.CallOpts)
}

// Messenger is a free data retrieval call binding the contract method 0x3cb747bf.
//
// Solidity: function messenger() view returns(address)
func (_L1StandardBridge *L1StandardBridgeCaller) Messenger(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _L1StandardBridge.contract.Call(opts, &out, "messenger")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Messenger is a free data retrieval call binding the contract method 0x3cb747bf.
//
// Solidity: function messenger() view returns(address)
func (_L1StandardBridge *L1StandardBridgeSession) Messenger() (common.Address, error) {
	return _L1StandardBridge.Contract.Messenger(&_L1StandardBridge.CallOpts)
}

// Messenger is a free data retrieval call binding the contract method 0x3cb747bf.
//
// Solidity: function messenger() view returns(address)
func (_L1StandardBridge *L1StandardBridgeCallerSession) Messenger() (common.Address, error) {
	return _L1StandardBridge.Contract.Messenger(&_L1StandardBridge.CallOpts)
}

// OtherBridge is a free data retrieval call binding the contract method 0xc89701a2.
//
// Solidity: function otherBridge() view returns(address)
func (_L1StandardBridge *L1StandardBridgeCaller) OtherBridge(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _L1StandardBridge.contract.Call(opts, &out, "otherBridge")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// OtherBridge is a free data retrieval call binding the contract method 0xc89701a2.
//
// Solidity: function otherBridge() view returns(address)
func (_L1StandardBridge *L1StandardBridgeSession) OtherBridge() (common.Address, error) {
	return _L1StandardBridge.Contract.OtherBridge(&_L1StandardBridge.CallOpts)
}

// OtherBridge is a free data retrieval call binding the contract method 0xc89701a2.
//
// Solidity: function otherBridge() view returns(address)
func (_L1StandardBridge *L1StandardBridgeCallerSession) OtherBridge() (common.Address, error) {
	return _L1StandardBridge.Contract.OtherBridge(&_L1StandardBridge.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_L1StandardBridge *L1StandardBridgeCaller) Paused(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _L1StandardBridge.contract.Call(opts, &out, "paused")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_L1StandardBridge *L1StandardBridgeSession) Paused() (bool, error) {
	return _L1StandardBridge.Contract.Paused(&_L1StandardBridge.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_L1StandardBridge *L1StandardBridgeCallerSession) Paused() (bool, error) {
	return _L1StandardBridge.Contract.Paused(&_L1StandardBridge.CallOpts)
}

// SuperchainConfig is a free data retrieval call binding the contract method 0x35e80ab3.
//
// Solidity: function superchainConfig() view returns(address)
func (_L1StandardBridge *L1StandardBridgeCaller) SuperchainConfig(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _L1StandardBridge.contract.Call(opts, &out, "superchainConfig")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// SuperchainConfig is a free data retrieval call binding the contract method 0x35e80ab3.
//
// Solidity: function superchainConfig() view returns(address)
func (_L1StandardBridge *L1StandardBridgeSession) SuperchainConfig() (common.Address, error) {
	return _L1StandardBridge.Contract.SuperchainConfig(&_L1StandardBridge.CallOpts)
}

// SuperchainConfig is a free data retrieval call binding the contract method 0x35e80ab3.
//
// Solidity: function superchainConfig() view returns(address)
func (_L1StandardBridge *L1StandardBridgeCallerSession) SuperchainConfig() (common.Address, error) {
	return _L1StandardBridge.Contract.SuperchainConfig(&_L1StandardBridge.CallOpts)
}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() view returns(string)
func (_L1StandardBridge *L1StandardBridgeCaller) Version(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _L1StandardBridge.contract.Call(opts, &out, "version")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() view returns(string)
func (_L1StandardBridge *L1StandardBridgeSession) Version() (string, error) {
	return _L1StandardBridge.Contract.Version(&_L1StandardBridge.CallOpts)
}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() view returns(string)
func (_L1StandardBridge *L1StandardBridgeCallerSession) Version() (string, error) {
	return _L1StandardBridge.Contract.Version(&_L1StandardBridge.CallOpts)
}

// BridgeERC20 is a paid mutator transaction binding the contract method 0x87087623.
//
// Solidity: function bridgeERC20(address _localToken, address _remoteToken, uint256 _amount, uint32 _minGasLimit, bytes _extraData) returns()
func (_L1StandardBridge *L1StandardBridgeTransactor) BridgeERC20(opts *bind.TransactOpts, _localToken common.Address, _remoteToken common.Address, _amount *big.Int, _minGasLimit uint32, _extraData []byte) (*types.Transaction, error) {
	return _L1StandardBridge.contract.Transact(opts, "bridgeERC20", _localToken, _remoteToken, _amount, _minGasLimit, _extraData)
}

// BridgeERC20 is a paid mutator transaction binding the contract method 0x87087623.
//
// Solidity: function bridgeERC20(address _localToken, address _remoteToken, uint256 _amount, uint32 _minGasLimit, bytes _extraData) returns()
func (_L1StandardBridge *L1StandardBridgeSession) BridgeERC20(_localToken common.Address, _remoteToken common.Address, _amount *big.Int, _minGasLimit uint32, _extraData []byte) (*types.Transaction, error) {
	return _L1StandardBridge.Contract.BridgeERC20(&_L1StandardBridge.TransactOpts, _localToken, _remoteToken, _amount, _minGasLimit, _extraData)
}

// BridgeERC20 is a paid mutator transaction binding the contract method 0x87087623.
//
// Solidity: function bridgeERC20(address _localToken, address _remoteToken, uint256 _amount, uint32 _minGasLimit, bytes _extraData) returns()
func (_L1StandardBridge *L1StandardBridgeTransactorSession) BridgeERC20(_localToken common.Address, _remoteToken common.Address, _amount *big.Int, _minGasLimit uint32, _extraData []byte) (*types.Transaction, error) {
	return _L1StandardBridge.Contract.BridgeERC20(&_L1StandardBridge.TransactOpts, _localToken, _remoteToken, _amount, _minGasLimit, _extraData)
}

// BridgeERC20To is a paid mutator transaction binding the contract method 0x540abf73.
//
// Solidity: function bridgeERC20To(address _localToken, address _remoteToken, address _to, uint256 _amount, uint32 _minGasLimit, bytes _extraData) returns()
func (_L1StandardBridge *L1StandardBridgeTransactor) BridgeERC20To(opts *bind.TransactOpts, _localToken common.Address, _remoteToken common.Address, _to common.Address, _amount *big.Int, _minGasLimit uint32, _extraData []byte) (*types.Transaction, error) {
	return _L1StandardBridge.contract.Transact(opts, "bridgeERC20To", _localToken, _remoteToken, _to, _amount, _minGasLimit, _extraData)
}

// BridgeERC20To is a paid mutator transaction binding the contract method 0x540abf73.
//
// Solidity: function bridgeERC20To(address _localToken, address _remoteToken, address _to, uint256 _amount, uint32 _minGasLimit, bytes _extraData) returns()
func (_L1StandardBridge *L1StandardBridgeSession) BridgeERC20To(_localToken common.Address, _remoteToken common.Address, _to common.Address, _amount *big.Int, _minGasLimit uint32, _extraData []byte) (*types.Transaction, error) {
	return _L1StandardBridge.Contract.BridgeERC20To(&_L1StandardBridge.TransactOpts, _localToken, _remoteToken, _to, _amount, _minGasLimit, _extraData)
}

// BridgeERC20To is a paid mutator transaction binding the contract method 0x540abf73.
//
// Solidity: function bridgeERC20To(address _localToken, address _remoteToken, address _to, uint256 _amount, uint32 _minGasLimit, bytes _extraData) returns()
func (_L1StandardBridge *L1StandardBridgeTransactorSession) BridgeERC20To(_localToken common.Address, _remoteToken common.Address, _to common.Address, _amount *big.Int, _minGasLimit uint32, _extraData []byte) (*types.Transaction, error) {
	return _L1StandardBridge.Contract.BridgeERC20To(&_L1StandardBridge.TransactOpts, _localToken, _remoteToken, _to, _amount, _minGasLimit, _extraData)
}

// DepositERC20 is a paid mutator transaction binding the contract method 0x58a997f6.
//
// Solidity: function depositERC20(address _l1Token, address _l2Token, uint256 _amount, uint32 _minGasLimit, bytes _extraData) returns()
func (_L1StandardBridge *L1StandardBridgeTransactor) DepositERC20(opts *bind.TransactOpts, _l1Token common.Address, _l2Token common.Address, _amount *big.Int, _minGasLimit uint32, _extraData []byte) (*types.Transaction, error) {
	return _L1StandardBridge.contract.Transact(opts, "depositERC20", _l1Token, _l2Token, _amount, _minGasLimit, _extraData)
}

// DepositERC20 is a paid mutator transaction binding the contract method 0x58a997f6.
//
// Solidity: function depositERC20(address _l1Token, address _l2Token, uint256 _amount, uint32 _minGasLimit, bytes _extraData) returns()
func (_L1StandardBridge *L1StandardBridgeSession) DepositERC20(_l1Token common.Address, _l2Token common.Address, _amount *big.Int, _minGasLimit uint32, _extraData []byte) (*types.Transaction, error) {
	return _L1StandardBridge.Contract.DepositERC20(&_L1StandardBridge.TransactOpts, _l1Token, _l2Token, _amount, _minGasLimit, _extraData)
}

// DepositERC20 is a paid mutator transaction binding the contract method 0x58a997f6.
//
// Solidity: function depositERC20(address _l1Token, address _l2Token, uint256 _amount, uint32 _minGasLimit, bytes _extraData) returns()
func (_L1StandardBridge *L1StandardBridgeTransactorSession) DepositERC20(_l1Token common.Address, _l2Token common.Address, _amount *big.Int, _minGasLimit uint32, _extraData []byte) (*types.Transaction, error) {
	return _L1StandardBridge.Contract.DepositERC20(&_L1StandardBridge.TransactOpts, _l1Token, _l2Token, _amount, _minGasLimit, _extraData)
}

// DepositERC20To is a paid mutator transaction binding the contract method 0x838b2520.
//
// Solidity: function depositERC20To(address _l1Token, address _l2Token, address _to, uint256 _amount, uint32 _minGasLimit, bytes _extraData) returns()
func (_L1StandardBridge *L1StandardBridgeTransactor) DepositERC20To(opts *bind.TransactOpts, _l1Token common.Address, _l2Token common.Address, _to common.Address, _amount *big.Int, _minGasLimit uint32, _extraData []byte) (*types.Transaction, error) {
	return _L1StandardBridge.contract.Transact(opts, "depositERC20To", _l1Token, _l2Token, _to, _amount, _minGasLimit, _extraData)
}

// DepositERC20To is a paid mutator transaction binding the contract method 0x838b2520.
//
// Solidity: function depositERC20To(address _l1Token, address _l2Token, address _to, uint256 _amount, uint32 _minGasLimit, bytes _extraData) returns()
func (_L1StandardBridge *L1StandardBridgeSession) DepositERC20To(_l1Token common.Address, _l2Token common.Address, _to common.Address, _amount *big.Int, _minGasLimit uint32, _extraData []byte) (*types.Transaction, error) {
	return _L1StandardBridge.Contract.DepositERC20To(&_L1StandardBridge.TransactOpts, _l1Token, _l2Token, _to, _amount, _minGasLimit, _extraData)
}

// DepositERC20To is a paid mutator transaction binding the contract method 0x838b2520.
//
// Solidity: function depositERC20To(address _l1Token, address _l2Token, address _to, uint256 _amount, uint32 _minGasLimit, bytes _extraData) returns()
func (_L1StandardBridge *L1StandardBridgeTransactorSession) DepositERC20To(_l1Token common.Address, _l2Token common.Address, _to common.Address, _amount *big.Int, _minGasLimit uint32, _extraData []byte) (*types.Transaction, error) {
	return _L1StandardBridge.Contract.DepositERC20To(&_L1StandardBridge.TransactOpts, _l1Token, _l2Token, _to, _amount, _minGasLimit, _extraData)
}

// FinalizeBridgeERC20 is a paid mutator transaction binding the contract method 0x0166a07a.
//
// Solidity: function finalizeBridgeERC20(address _localToken, address _remoteToken, address _from, address _to, uint256 _amount, bytes _extraData) returns()
func (_L1StandardBridge *L1StandardBridgeTransactor) FinalizeBridgeERC20(opts *bind.TransactOpts, _localToken common.Address, _remoteToken common.Address, _from common.Address, _to common.Address, _amount *big.Int, _extraData []byte) (*types.Transaction, error) {
	return _L1StandardBridge.contract.Transact(opts, "finalizeBridgeERC20", _localToken, _remoteToken, _from, _to, _amount, _extraData)
}

// FinalizeBridgeERC20 is a paid mutator transaction binding the contract method 0x0166a07a.
//
// Solidity: function finalizeBridgeERC20(address _localToken, address _remoteToken, address _from, address _to, uint256 _amount, bytes _extraData) returns()
func (_L1StandardBridge *L1StandardBridgeSession) FinalizeBridgeERC20(_localToken common.Address, _remoteToken common.Address, _from common.Address, _to common.Address, _amount *big.Int, _extraData []byte) (*types.Transaction, error) {
	return _L1StandardBridge.Contract.FinalizeBridgeERC20(&_L1StandardBridge.TransactOpts, _localToken, _remoteToken, _from, _to, _amount, _extraData)
}

// FinalizeBridgeERC20 is a paid mutator transaction binding the contract method 0x0166a07a.
//
// Solidity: function finalizeBridgeERC20(address _localToken, address _remoteToken, address _from, address _to, uint256 _amount, bytes _extraData) returns()
func (_L1StandardBridge *L1StandardBridgeTransactorSession) FinalizeBridgeERC20(_localToken common.Address, _remoteToken common.Address, _from common.Address, _to common.Address, _amount *big.Int, _extraData []byte) (*types.Transaction, error) {
	return _L1StandardBridge.Contract.FinalizeBridgeERC20(&_L1StandardBridge.TransactOpts, _localToken, _remoteToken, _from, _to, _amount, _extraData)
}

// FinalizeERC20Withdrawal is a paid mutator transaction binding the contract method 0xa9f9e675.
//
// Solidity: function finalizeERC20Withdrawal(address _l1Token, address _l2Token, address _from, address _to, uint256 _amount, bytes _extraData) returns()
func (_L1StandardBridge *L1StandardBridgeTransactor) FinalizeERC20Withdrawal(opts *bind.TransactOpts, _l1Token common.Address, _l2Token common.Address, _from common.Address, _to common.Address, _amount *big.Int, _extraData []byte) (*types.Transaction, error) {
	return _L1StandardBridge.contract.Transact(opts, "finalizeERC20Withdrawal", _l1Token, _l2Token, _from, _to, _amount, _extraData)
}

// FinalizeERC20Withdrawal is a paid mutator transaction binding the contract method 0xa9f9e675.
//
// Solidity: function finalizeERC20Withdrawal(address _l1Token, address _l2Token, address _from, address _to, uint256 _amount, bytes _extraData) returns()
func (_L1StandardBridge *L1StandardBridgeSession) FinalizeERC20Withdrawal(_l1Token common.Address, _l2Token common.Address, _from common.Address, _to common.Address, _amount *big.Int, _extraData []byte) (*types.Transaction, error) {
	return _L1StandardBridge.Contract.FinalizeERC20Withdrawal(&_L1StandardBridge.TransactOpts, _l1Token, _l2Token, _from, _to, _amount, _extraData)
}

// FinalizeERC20Withdrawal is a paid mutator transaction binding the contract method 0xa9f9e675.
//
// Solidity: function finalizeERC20Withdrawal(address _l1Token, address _l2Token, address _from, address _to, uint256 _amount, bytes _extraData) returns()
func (_L1StandardBridge *L1StandardBridgeTransactorSession) FinalizeERC20Withdrawal(_l1Token common.Address, _l2Token common.Address, _from common.Address, _to common.Address, _amount *big.Int, _extraData []byte) (*types.Transaction, error) {
	return _L1StandardBridge.Contract.FinalizeERC20Withdrawal(&_L1StandardBridge.TransactOpts, _l1Token, _l2Token, _from, _to, _amount, _extraData)
}

// Initialize is a paid mutator transaction binding the contract method 0x485cc955.
//
// Solidity: function initialize(address _messenger, address _superchainConfig) returns()
func (_L1StandardBridge *L1StandardBridgeTransactor) Initialize(opts *bind.TransactOpts, _messenger common.Address, _superchainConfig common.Address) (*types.Transaction, error) {
	return _L1StandardBridge.contract.Transact(opts, "initialize", _messenger, _superchainConfig)
}

// Initialize is a paid mutator transaction binding the contract method 0x485cc955.
//
// Solidity: function initialize(address _messenger, address _superchainConfig) returns()
func (_L1StandardBridge *L1StandardBridgeSession) Initialize(_messenger common.Address, _superchainConfig common.Address) (*types.Transaction, error) {
	return _L1StandardBridge.Contract.Initialize(&_L1StandardBridge.TransactOpts, _messenger, _superchainConfig)
}

// Initialize is a paid mutator transaction binding the contract method 0x485cc955.
//
// Solidity: function initialize(address _messenger, address _superchainConfig) returns()
func (_L1StandardBridge *L1StandardBridgeTransactorSession) Initialize(_messenger common.Address, _superchainConfig common.Address) (*types.Transaction, error) {
	return _L1StandardBridge.Contract.Initialize(&_L1StandardBridge.TransactOpts, _messenger, _superchainConfig)
}

// L1StandardBridgeERC20BridgeFinalizedIterator is returned from FilterERC20BridgeFinalized and is used to iterate over the raw logs and unpacked data for ERC20BridgeFinalized events raised by the L1StandardBridge contract.
type L1StandardBridgeERC20BridgeFinalizedIterator struct {
	Event *L1StandardBridgeERC20BridgeFinalized // Event containing the contract specifics and raw log

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
func (it *L1StandardBridgeERC20BridgeFinalizedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(L1StandardBridgeERC20BridgeFinalized)
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
		it.Event = new(L1StandardBridgeERC20BridgeFinalized)
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
func (it *L1StandardBridgeERC20BridgeFinalizedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *L1StandardBridgeERC20BridgeFinalizedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// L1StandardBridgeERC20BridgeFinalized represents a ERC20BridgeFinalized event raised by the L1StandardBridge contract.
type L1StandardBridgeERC20BridgeFinalized struct {
	LocalToken  common.Address
	RemoteToken common.Address
	From        common.Address
	To          common.Address
	Amount      *big.Int
	ExtraData   []byte
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterERC20BridgeFinalized is a free log retrieval operation binding the contract event 0xd59c65b35445225835c83f50b6ede06a7be047d22e357073e250d9af537518cd.
//
// Solidity: event ERC20BridgeFinalized(address indexed localToken, address indexed remoteToken, address indexed from, address to, uint256 amount, bytes extraData)
func (_L1StandardBridge *L1StandardBridgeFilterer) FilterERC20BridgeFinalized(opts *bind.FilterOpts, localToken []common.Address, remoteToken []common.Address, from []common.Address) (*L1StandardBridgeERC20BridgeFinalizedIterator, error) {

	var localTokenRule []interface{}
	for _, localTokenItem := range localToken {
		localTokenRule = append(localTokenRule, localTokenItem)
	}
	var remoteTokenRule []interface{}
	for _, remoteTokenItem := range remoteToken {
		remoteTokenRule = append(remoteTokenRule, remoteTokenItem)
	}
	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}

	logs, sub, err := _L1StandardBridge.contract.FilterLogs(opts, "ERC20BridgeFinalized", localTokenRule, remoteTokenRule, fromRule)
	if err != nil {
		return nil, err
	}
	return &L1StandardBridgeERC20BridgeFinalizedIterator{contract: _L1StandardBridge.contract, event: "ERC20BridgeFinalized", logs: logs, sub: sub}, nil
}

// WatchERC20BridgeFinalized is a free log subscription operation binding the contract event 0xd59c65b35445225835c83f50b6ede06a7be047d22e357073e250d9af537518cd.
//
// Solidity: event ERC20BridgeFinalized(address indexed localToken, address indexed remoteToken, address indexed from, address to, uint256 amount, bytes extraData)
func (_L1StandardBridge *L1StandardBridgeFilterer) WatchERC20BridgeFinalized(opts *bind.WatchOpts, sink chan<- *L1StandardBridgeERC20BridgeFinalized, localToken []common.Address, remoteToken []common.Address, from []common.Address) (event.Subscription, error) {

	var localTokenRule []interface{}
	for _, localTokenItem := range localToken {
		localTokenRule = append(localTokenRule, localTokenItem)
	}
	var remoteTokenRule []interface{}
	for _, remoteTokenItem := range remoteToken {
		remoteTokenRule = append(remoteTokenRule, remoteTokenItem)
	}
	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}

	logs, sub, err := _L1StandardBridge.contract.WatchLogs(opts, "ERC20BridgeFinalized", localTokenRule, remoteTokenRule, fromRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(L1StandardBridgeERC20BridgeFinalized)
				if err := _L1StandardBridge.contract.UnpackLog(event, "ERC20BridgeFinalized", log); err != nil {
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

// ParseERC20BridgeFinalized is a log parse operation binding the contract event 0xd59c65b35445225835c83f50b6ede06a7be047d22e357073e250d9af537518cd.
//
// Solidity: event ERC20BridgeFinalized(address indexed localToken, address indexed remoteToken, address indexed from, address to, uint256 amount, bytes extraData)
func (_L1StandardBridge *L1StandardBridgeFilterer) ParseERC20BridgeFinalized(log types.Log) (*L1StandardBridgeERC20BridgeFinalized, error) {
	event := new(L1StandardBridgeERC20BridgeFinalized)
	if err := _L1StandardBridge.contract.UnpackLog(event, "ERC20BridgeFinalized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// L1StandardBridgeERC20BridgeInitiatedIterator is returned from FilterERC20BridgeInitiated and is used to iterate over the raw logs and unpacked data for ERC20BridgeInitiated events raised by the L1StandardBridge contract.
type L1StandardBridgeERC20BridgeInitiatedIterator struct {
	Event *L1StandardBridgeERC20BridgeInitiated // Event containing the contract specifics and raw log

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
func (it *L1StandardBridgeERC20BridgeInitiatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(L1StandardBridgeERC20BridgeInitiated)
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
		it.Event = new(L1StandardBridgeERC20BridgeInitiated)
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
func (it *L1StandardBridgeERC20BridgeInitiatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *L1StandardBridgeERC20BridgeInitiatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// L1StandardBridgeERC20BridgeInitiated represents a ERC20BridgeInitiated event raised by the L1StandardBridge contract.
type L1StandardBridgeERC20BridgeInitiated struct {
	LocalToken  common.Address
	RemoteToken common.Address
	From        common.Address
	To          common.Address
	Amount      *big.Int
	ExtraData   []byte
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterERC20BridgeInitiated is a free log retrieval operation binding the contract event 0x7ff126db8024424bbfd9826e8ab82ff59136289ea440b04b39a0df1b03b9cabf.
//
// Solidity: event ERC20BridgeInitiated(address indexed localToken, address indexed remoteToken, address indexed from, address to, uint256 amount, bytes extraData)
func (_L1StandardBridge *L1StandardBridgeFilterer) FilterERC20BridgeInitiated(opts *bind.FilterOpts, localToken []common.Address, remoteToken []common.Address, from []common.Address) (*L1StandardBridgeERC20BridgeInitiatedIterator, error) {

	var localTokenRule []interface{}
	for _, localTokenItem := range localToken {
		localTokenRule = append(localTokenRule, localTokenItem)
	}
	var remoteTokenRule []interface{}
	for _, remoteTokenItem := range remoteToken {
		remoteTokenRule = append(remoteTokenRule, remoteTokenItem)
	}
	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}

	logs, sub, err := _L1StandardBridge.contract.FilterLogs(opts, "ERC20BridgeInitiated", localTokenRule, remoteTokenRule, fromRule)
	if err != nil {
		return nil, err
	}
	return &L1StandardBridgeERC20BridgeInitiatedIterator{contract: _L1StandardBridge.contract, event: "ERC20BridgeInitiated", logs: logs, sub: sub}, nil
}

// WatchERC20BridgeInitiated is a free log subscription operation binding the contract event 0x7ff126db8024424bbfd9826e8ab82ff59136289ea440b04b39a0df1b03b9cabf.
//
// Solidity: event ERC20BridgeInitiated(address indexed localToken, address indexed remoteToken, address indexed from, address to, uint256 amount, bytes extraData)
func (_L1StandardBridge *L1StandardBridgeFilterer) WatchERC20BridgeInitiated(opts *bind.WatchOpts, sink chan<- *L1StandardBridgeERC20BridgeInitiated, localToken []common.Address, remoteToken []common.Address, from []common.Address) (event.Subscription, error) {

	var localTokenRule []interface{}
	for _, localTokenItem := range localToken {
		localTokenRule = append(localTokenRule, localTokenItem)
	}
	var remoteTokenRule []interface{}
	for _, remoteTokenItem := range remoteToken {
		remoteTokenRule = append(remoteTokenRule, remoteTokenItem)
	}
	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}

	logs, sub, err := _L1StandardBridge.contract.WatchLogs(opts, "ERC20BridgeInitiated", localTokenRule, remoteTokenRule, fromRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(L1StandardBridgeERC20BridgeInitiated)
				if err := _L1StandardBridge.contract.UnpackLog(event, "ERC20BridgeInitiated", log); err != nil {
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

// ParseERC20BridgeInitiated is a log parse operation binding the contract event 0x7ff126db8024424bbfd9826e8ab82ff59136289ea440b04b39a0df1b03b9cabf.
//
// Solidity: event ERC20BridgeInitiated(address indexed localToken, address indexed remoteToken, address indexed from, address to, uint256 amount, bytes extraData)
func (_L1StandardBridge *L1StandardBridgeFilterer) ParseERC20BridgeInitiated(log types.Log) (*L1StandardBridgeERC20BridgeInitiated, error) {
	event := new(L1StandardBridgeERC20BridgeInitiated)
	if err := _L1StandardBridge.contract.UnpackLog(event, "ERC20BridgeInitiated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// L1StandardBridgeERC20DepositInitiatedIterator is returned from FilterERC20DepositInitiated and is used to iterate over the raw logs and unpacked data for ERC20DepositInitiated events raised by the L1StandardBridge contract.
type L1StandardBridgeERC20DepositInitiatedIterator struct {
	Event *L1StandardBridgeERC20DepositInitiated // Event containing the contract specifics and raw log

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
func (it *L1StandardBridgeERC20DepositInitiatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(L1StandardBridgeERC20DepositInitiated)
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
		it.Event = new(L1StandardBridgeERC20DepositInitiated)
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
func (it *L1StandardBridgeERC20DepositInitiatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *L1StandardBridgeERC20DepositInitiatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// L1StandardBridgeERC20DepositInitiated represents a ERC20DepositInitiated event raised by the L1StandardBridge contract.
type L1StandardBridgeERC20DepositInitiated struct {
	L1Token   common.Address
	L2Token   common.Address
	From      common.Address
	To        common.Address
	Amount    *big.Int
	ExtraData []byte
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterERC20DepositInitiated is a free log retrieval operation binding the contract event 0x718594027abd4eaed59f95162563e0cc6d0e8d5b86b1c7be8b1b0ac3343d0396.
//
// Solidity: event ERC20DepositInitiated(address indexed l1Token, address indexed l2Token, address indexed from, address to, uint256 amount, bytes extraData)
func (_L1StandardBridge *L1StandardBridgeFilterer) FilterERC20DepositInitiated(opts *bind.FilterOpts, l1Token []common.Address, l2Token []common.Address, from []common.Address) (*L1StandardBridgeERC20DepositInitiatedIterator, error) {

	var l1TokenRule []interface{}
	for _, l1TokenItem := range l1Token {
		l1TokenRule = append(l1TokenRule, l1TokenItem)
	}
	var l2TokenRule []interface{}
	for _, l2TokenItem := range l2Token {
		l2TokenRule = append(l2TokenRule, l2TokenItem)
	}
	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}

	logs, sub, err := _L1StandardBridge.contract.FilterLogs(opts, "ERC20DepositInitiated", l1TokenRule, l2TokenRule, fromRule)
	if err != nil {
		return nil, err
	}
	return &L1StandardBridgeERC20DepositInitiatedIterator{contract: _L1StandardBridge.contract, event: "ERC20DepositInitiated", logs: logs, sub: sub}, nil
}

// WatchERC20DepositInitiated is a free log subscription operation binding the contract event 0x718594027abd4eaed59f95162563e0cc6d0e8d5b86b1c7be8b1b0ac3343d0396.
//
// Solidity: event ERC20DepositInitiated(address indexed l1Token, address indexed l2Token, address indexed from, address to, uint256 amount, bytes extraData)
func (_L1StandardBridge *L1StandardBridgeFilterer) WatchERC20DepositInitiated(opts *bind.WatchOpts, sink chan<- *L1StandardBridgeERC20DepositInitiated, l1Token []common.Address, l2Token []common.Address, from []common.Address) (event.Subscription, error) {

	var l1TokenRule []interface{}
	for _, l1TokenItem := range l1Token {
		l1TokenRule = append(l1TokenRule, l1TokenItem)
	}
	var l2TokenRule []interface{}
	for _, l2TokenItem := range l2Token {
		l2TokenRule = append(l2TokenRule, l2TokenItem)
	}
	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}

	logs, sub, err := _L1StandardBridge.contract.WatchLogs(opts, "ERC20DepositInitiated", l1TokenRule, l2TokenRule, fromRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(L1StandardBridgeERC20DepositInitiated)
				if err := _L1StandardBridge.contract.UnpackLog(event, "ERC20DepositInitiated", log); err != nil {
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

// ParseERC20DepositInitiated is a log parse operation binding the contract event 0x718594027abd4eaed59f95162563e0cc6d0e8d5b86b1c7be8b1b0ac3343d0396.
//
// Solidity: event ERC20DepositInitiated(address indexed l1Token, address indexed l2Token, address indexed from, address to, uint256 amount, bytes extraData)
func (_L1StandardBridge *L1StandardBridgeFilterer) ParseERC20DepositInitiated(log types.Log) (*L1StandardBridgeERC20DepositInitiated, error) {
	event := new(L1StandardBridgeERC20DepositInitiated)
	if err := _L1StandardBridge.contract.UnpackLog(event, "ERC20DepositInitiated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// L1StandardBridgeERC20WithdrawalFinalizedIterator is returned from FilterERC20WithdrawalFinalized and is used to iterate over the raw logs and unpacked data for ERC20WithdrawalFinalized events raised by the L1StandardBridge contract.
type L1StandardBridgeERC20WithdrawalFinalizedIterator struct {
	Event *L1StandardBridgeERC20WithdrawalFinalized // Event containing the contract specifics and raw log

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
func (it *L1StandardBridgeERC20WithdrawalFinalizedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(L1StandardBridgeERC20WithdrawalFinalized)
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
		it.Event = new(L1StandardBridgeERC20WithdrawalFinalized)
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
func (it *L1StandardBridgeERC20WithdrawalFinalizedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *L1StandardBridgeERC20WithdrawalFinalizedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// L1StandardBridgeERC20WithdrawalFinalized represents a ERC20WithdrawalFinalized event raised by the L1StandardBridge contract.
type L1StandardBridgeERC20WithdrawalFinalized struct {
	L1Token   common.Address
	L2Token   common.Address
	From      common.Address
	To        common.Address
	Amount    *big.Int
	ExtraData []byte
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterERC20WithdrawalFinalized is a free log retrieval operation binding the contract event 0x3ceee06c1e37648fcbb6ed52e17b3e1f275a1f8c7b22a84b2b84732431e046b3.
//
// Solidity: event ERC20WithdrawalFinalized(address indexed l1Token, address indexed l2Token, address indexed from, address to, uint256 amount, bytes extraData)
func (_L1StandardBridge *L1StandardBridgeFilterer) FilterERC20WithdrawalFinalized(opts *bind.FilterOpts, l1Token []common.Address, l2Token []common.Address, from []common.Address) (*L1StandardBridgeERC20WithdrawalFinalizedIterator, error) {

	var l1TokenRule []interface{}
	for _, l1TokenItem := range l1Token {
		l1TokenRule = append(l1TokenRule, l1TokenItem)
	}
	var l2TokenRule []interface{}
	for _, l2TokenItem := range l2Token {
		l2TokenRule = append(l2TokenRule, l2TokenItem)
	}
	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}

	logs, sub, err := _L1StandardBridge.contract.FilterLogs(opts, "ERC20WithdrawalFinalized", l1TokenRule, l2TokenRule, fromRule)
	if err != nil {
		return nil, err
	}
	return &L1StandardBridgeERC20WithdrawalFinalizedIterator{contract: _L1StandardBridge.contract, event: "ERC20WithdrawalFinalized", logs: logs, sub: sub}, nil
}

// WatchERC20WithdrawalFinalized is a free log subscription operation binding the contract event 0x3ceee06c1e37648fcbb6ed52e17b3e1f275a1f8c7b22a84b2b84732431e046b3.
//
// Solidity: event ERC20WithdrawalFinalized(address indexed l1Token, address indexed l2Token, address indexed from, address to, uint256 amount, bytes extraData)
func (_L1StandardBridge *L1StandardBridgeFilterer) WatchERC20WithdrawalFinalized(opts *bind.WatchOpts, sink chan<- *L1StandardBridgeERC20WithdrawalFinalized, l1Token []common.Address, l2Token []common.Address, from []common.Address) (event.Subscription, error) {

	var l1TokenRule []interface{}
	for _, l1TokenItem := range l1Token {
		l1TokenRule = append(l1TokenRule, l1TokenItem)
	}
	var l2TokenRule []interface{}
	for _, l2TokenItem := range l2Token {
		l2TokenRule = append(l2TokenRule, l2TokenItem)
	}
	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}

	logs, sub, err := _L1StandardBridge.contract.WatchLogs(opts, "ERC20WithdrawalFinalized", l1TokenRule, l2TokenRule, fromRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(L1StandardBridgeERC20WithdrawalFinalized)
				if err := _L1StandardBridge.contract.UnpackLog(event, "ERC20WithdrawalFinalized", log); err != nil {
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

// ParseERC20WithdrawalFinalized is a log parse operation binding the contract event 0x3ceee06c1e37648fcbb6ed52e17b3e1f275a1f8c7b22a84b2b84732431e046b3.
//
// Solidity: event ERC20WithdrawalFinalized(address indexed l1Token, address indexed l2Token, address indexed from, address to, uint256 amount, bytes extraData)
func (_L1StandardBridge *L1StandardBridgeFilterer) ParseERC20WithdrawalFinalized(log types.Log) (*L1StandardBridgeERC20WithdrawalFinalized, error) {
	event := new(L1StandardBridgeERC20WithdrawalFinalized)
	if err := _L1StandardBridge.contract.UnpackLog(event, "ERC20WithdrawalFinalized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// L1StandardBridgeInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the L1StandardBridge contract.
type L1StandardBridgeInitializedIterator struct {
	Event *L1StandardBridgeInitialized // Event containing the contract specifics and raw log

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
func (it *L1StandardBridgeInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(L1StandardBridgeInitialized)
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
		it.Event = new(L1StandardBridgeInitialized)
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
func (it *L1StandardBridgeInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *L1StandardBridgeInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// L1StandardBridgeInitialized represents a Initialized event raised by the L1StandardBridge contract.
type L1StandardBridgeInitialized struct {
	Version uint8
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_L1StandardBridge *L1StandardBridgeFilterer) FilterInitialized(opts *bind.FilterOpts) (*L1StandardBridgeInitializedIterator, error) {

	logs, sub, err := _L1StandardBridge.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &L1StandardBridgeInitializedIterator{contract: _L1StandardBridge.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_L1StandardBridge *L1StandardBridgeFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *L1StandardBridgeInitialized) (event.Subscription, error) {

	logs, sub, err := _L1StandardBridge.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(L1StandardBridgeInitialized)
				if err := _L1StandardBridge.contract.UnpackLog(event, "Initialized", log); err != nil {
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

// ParseInitialized is a log parse operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_L1StandardBridge *L1StandardBridgeFilterer) ParseInitialized(log types.Log) (*L1StandardBridgeInitialized, error) {
	event := new(L1StandardBridgeInitialized)
	if err := _L1StandardBridge.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
