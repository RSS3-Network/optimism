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

// RSS3TokenMetaData contains all meta data concerning the RSS3Token contract.
var RSS3TokenMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"constructor\",\"inputs\":[{\"name\":\"_bridge\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_remoteToken\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_name\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"_symbol\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"BRIDGE\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"REMOTE_TOKEN\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"allowance\",\"inputs\":[{\"name\":\"owner\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"spender\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"approve\",\"inputs\":[{\"name\":\"spender\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"balanceOf\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"bridge\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"burn\",\"inputs\":[{\"name\":\"_from\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"decimals\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint8\",\"internalType\":\"uint8\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"decreaseAllowance\",\"inputs\":[{\"name\":\"spender\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"subtractedValue\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"increaseAllowance\",\"inputs\":[{\"name\":\"spender\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"addedValue\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"l1Token\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"l2Bridge\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"mint\",\"inputs\":[{\"name\":\"_to\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"name\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"remoteToken\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"supportsInterface\",\"inputs\":[{\"name\":\"_interfaceId\",\"type\":\"bytes4\",\"internalType\":\"bytes4\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"pure\"},{\"type\":\"function\",\"name\":\"symbol\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"totalSupply\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"transfer\",\"inputs\":[{\"name\":\"to\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"transferFrom\",\"inputs\":[{\"name\":\"from\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"to\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"version\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"view\"},{\"type\":\"event\",\"name\":\"Approval\",\"inputs\":[{\"name\":\"owner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"spender\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"value\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Burn\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Mint\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Transfer\",\"inputs\":[{\"name\":\"from\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"to\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"value\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false}]",
	Bin: "0x60c06040523480156200001157600080fd5b50604051620017393803806200173983398101604081905262000034916200015a565b8181600362000044838262000279565b50600462000053828262000279565b5050506001600160a01b0392831660805250501660a05262000345565b80516001600160a01b03811681146200008857600080fd5b919050565b634e487b7160e01b600052604160045260246000fd5b600082601f830112620000b557600080fd5b81516001600160401b0380821115620000d257620000d26200008d565b604051601f8301601f19908116603f01168101908282118183101715620000fd57620000fd6200008d565b816040528381526020925086838588010111156200011a57600080fd5b600091505b838210156200013e57858201830151818301840152908201906200011f565b83821115620001505760008385830101525b9695505050505050565b600080600080608085870312156200017157600080fd5b6200017c8562000070565b93506200018c6020860162000070565b60408601519093506001600160401b0380821115620001aa57600080fd5b620001b888838901620000a3565b93506060870151915080821115620001cf57600080fd5b50620001de87828801620000a3565b91505092959194509250565b600181811c90821680620001ff57607f821691505b6020821081036200022057634e487b7160e01b600052602260045260246000fd5b50919050565b601f8211156200027457600081815260208120601f850160051c810160208610156200024f5750805b601f850160051c820191505b8181101562000270578281556001016200025b565b5050505b505050565b81516001600160401b038111156200029557620002956200008d565b620002ad81620002a68454620001ea565b8462000226565b602080601f831160018114620002e55760008415620002cc5750858301515b600019600386901b1c1916600185901b17855562000270565b600085815260208120601f198616915b828110156200031657888601518255948401946001909101908401620002f5565b5085821015620003355787850151600019600388901b60f8161c191681555b5050505050600190811b01905550565b60805160a0516113b26200038760003960008181610329015281816103be01528181610603015261073a0152600081816101a9015261034f01526113b26000f3fe608060405234801561001057600080fd5b50600436106101775760003560e01c806370a08231116100d8578063ae1f6aaf1161008c578063dd62ed3e11610066578063dd62ed3e14610373578063e78cea9214610327578063ee9a31a2146103b957600080fd5b8063ae1f6aaf14610327578063c01e1bd61461034d578063d6c0b2c41461034d57600080fd5b80639dc29fac116100bd5780639dc29fac146102ee578063a457c2d714610301578063a9059cbb1461031457600080fd5b806370a08231146102b057806395d89b41146102e657600080fd5b806323b872dd1161012f5780633950935111610114578063395093511461024c57806340c10f191461025f57806354fd4d501461027457600080fd5b806323b872dd1461022a578063313ce5671461023d57600080fd5b806306fdde031161016057806306fdde03146101f0578063095ea7b31461020557806318160ddd1461021857600080fd5b806301ffc9a71461017c578063033964be146101a4575b600080fd5b61018f61018a36600461115b565b6103e0565b60405190151581526020015b60405180910390f35b6101cb7f000000000000000000000000000000000000000000000000000000000000000081565b60405173ffffffffffffffffffffffffffffffffffffffff909116815260200161019b565b6101f86104d1565b60405161019b91906111a4565b61018f610213366004611240565b610563565b6002545b60405190815260200161019b565b61018f61023836600461126a565b61057b565b6040516012815260200161019b565b61018f61025a366004611240565b61059f565b61027261026d366004611240565b6105eb565b005b6101f86040518060400160405280600581526020017f312e332e3000000000000000000000000000000000000000000000000000000081525081565b61021c6102be3660046112a6565b73ffffffffffffffffffffffffffffffffffffffff1660009081526020819052604090205490565b6101f8610713565b6102726102fc366004611240565b610722565b61018f61030f366004611240565b610839565b61018f610322366004611240565b61090a565b7f00000000000000000000000000000000000000000000000000000000000000006101cb565b7f00000000000000000000000000000000000000000000000000000000000000006101cb565b61021c6103813660046112c1565b73ffffffffffffffffffffffffffffffffffffffff918216600090815260016020908152604080832093909416825291909152205490565b6101cb7f000000000000000000000000000000000000000000000000000000000000000081565b60007f01ffc9a7000000000000000000000000000000000000000000000000000000007f1d1d8b63000000000000000000000000000000000000000000000000000000007fec4fc8e3000000000000000000000000000000000000000000000000000000007fffffffff00000000000000000000000000000000000000000000000000000000851683148061049957507fffffffff00000000000000000000000000000000000000000000000000000000858116908316145b806104c857507fffffffff00000000000000000000000000000000000000000000000000000000858116908216145b95945050505050565b6060600380546104e0906112f4565b80601f016020809104026020016040519081016040528092919081815260200182805461050c906112f4565b80156105595780601f1061052e57610100808354040283529160200191610559565b820191906000526020600020905b81548152906001019060200180831161053c57829003601f168201915b5050505050905090565b600033610571818585610918565b5060019392505050565b600033610589858285610acc565b610594858585610ba3565b506001949350505050565b33600081815260016020908152604080832073ffffffffffffffffffffffffffffffffffffffff8716845290915281205490919061057190829086906105e6908790611376565b610918565b3373ffffffffffffffffffffffffffffffffffffffff7f000000000000000000000000000000000000000000000000000000000000000016146106b5576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152603460248201527f4f7074696d69736d4d696e7461626c6545524332303a206f6e6c79206272696460448201527f67652063616e206d696e7420616e64206275726e00000000000000000000000060648201526084015b60405180910390fd5b6106bf8282610e56565b8173ffffffffffffffffffffffffffffffffffffffff167f0f6798a560793a54c3bcfe86a93cde1e73087d944c0ea20544137d41213968858260405161070791815260200190565b60405180910390a25050565b6060600480546104e0906112f4565b3373ffffffffffffffffffffffffffffffffffffffff7f000000000000000000000000000000000000000000000000000000000000000016146107e7576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152603460248201527f4f7074696d69736d4d696e7461626c6545524332303a206f6e6c79206272696460448201527f67652063616e206d696e7420616e64206275726e00000000000000000000000060648201526084016106ac565b6107f18282610f76565b8173ffffffffffffffffffffffffffffffffffffffff167fcc16f5dbb4873280815c1ee09dbd06736cffcc184412cf7a71a0fdb75d397ca58260405161070791815260200190565b33600081815260016020908152604080832073ffffffffffffffffffffffffffffffffffffffff87168452909152812054909190838110156108fd576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602560248201527f45524332303a2064656372656173656420616c6c6f77616e63652062656c6f7760448201527f207a65726f00000000000000000000000000000000000000000000000000000060648201526084016106ac565b6105948286868403610918565b600033610571818585610ba3565b73ffffffffffffffffffffffffffffffffffffffff83166109ba576040517f08c379a0000000000000000000000000000000000000000000000000000000008152602060048201526024808201527f45524332303a20617070726f76652066726f6d20746865207a65726f2061646460448201527f726573730000000000000000000000000000000000000000000000000000000060648201526084016106ac565b73ffffffffffffffffffffffffffffffffffffffff8216610a5d576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602260248201527f45524332303a20617070726f766520746f20746865207a65726f20616464726560448201527f737300000000000000000000000000000000000000000000000000000000000060648201526084016106ac565b73ffffffffffffffffffffffffffffffffffffffff83811660008181526001602090815260408083209487168084529482529182902085905590518481527f8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b92591015b60405180910390a3505050565b73ffffffffffffffffffffffffffffffffffffffff8381166000908152600160209081526040808320938616835292905220547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8114610b9d5781811015610b90576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601d60248201527f45524332303a20696e73756666696369656e7420616c6c6f77616e636500000060448201526064016106ac565b610b9d8484848403610918565b50505050565b73ffffffffffffffffffffffffffffffffffffffff8316610c46576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602560248201527f45524332303a207472616e736665722066726f6d20746865207a65726f20616460448201527f647265737300000000000000000000000000000000000000000000000000000060648201526084016106ac565b73ffffffffffffffffffffffffffffffffffffffff8216610ce9576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602360248201527f45524332303a207472616e7366657220746f20746865207a65726f206164647260448201527f657373000000000000000000000000000000000000000000000000000000000060648201526084016106ac565b73ffffffffffffffffffffffffffffffffffffffff831660009081526020819052604090205481811015610d9f576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602660248201527f45524332303a207472616e7366657220616d6f756e742065786365656473206260448201527f616c616e6365000000000000000000000000000000000000000000000000000060648201526084016106ac565b73ffffffffffffffffffffffffffffffffffffffff808516600090815260208190526040808220858503905591851681529081208054849290610de3908490611376565b925050819055508273ffffffffffffffffffffffffffffffffffffffff168473ffffffffffffffffffffffffffffffffffffffff167fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef84604051610e4991815260200190565b60405180910390a3610b9d565b73ffffffffffffffffffffffffffffffffffffffff8216610ed3576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601f60248201527f45524332303a206d696e7420746f20746865207a65726f20616464726573730060448201526064016106ac565b8060026000828254610ee59190611376565b909155505073ffffffffffffffffffffffffffffffffffffffff821660009081526020819052604081208054839290610f1f908490611376565b909155505060405181815273ffffffffffffffffffffffffffffffffffffffff8316906000907fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef9060200160405180910390a35050565b73ffffffffffffffffffffffffffffffffffffffff8216611019576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602160248201527f45524332303a206275726e2066726f6d20746865207a65726f2061646472657360448201527f730000000000000000000000000000000000000000000000000000000000000060648201526084016106ac565b73ffffffffffffffffffffffffffffffffffffffff8216600090815260208190526040902054818110156110cf576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602260248201527f45524332303a206275726e20616d6f756e7420657863656564732062616c616e60448201527f636500000000000000000000000000000000000000000000000000000000000060648201526084016106ac565b73ffffffffffffffffffffffffffffffffffffffff8316600090815260208190526040812083830390556002805484929061110b90849061138e565b909155505060405182815260009073ffffffffffffffffffffffffffffffffffffffff8516907fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef90602001610abf565b60006020828403121561116d57600080fd5b81357fffffffff000000000000000000000000000000000000000000000000000000008116811461119d57600080fd5b9392505050565b600060208083528351808285015260005b818110156111d1578581018301518582016040015282016111b5565b818111156111e3576000604083870101525b50601f017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe016929092016040019392505050565b803573ffffffffffffffffffffffffffffffffffffffff8116811461123b57600080fd5b919050565b6000806040838503121561125357600080fd5b61125c83611217565b946020939093013593505050565b60008060006060848603121561127f57600080fd5b61128884611217565b925061129660208501611217565b9150604084013590509250925092565b6000602082840312156112b857600080fd5b61119d82611217565b600080604083850312156112d457600080fd5b6112dd83611217565b91506112eb60208401611217565b90509250929050565b600181811c9082168061130857607f821691505b602082108103611341577f4e487b7100000000000000000000000000000000000000000000000000000000600052602260045260246000fd5b50919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b6000821982111561138957611389611347565b500190565b6000828210156113a0576113a0611347565b50039056fea164736f6c634300080f000a",
}

// RSS3TokenABI is the input ABI used to generate the binding from.
// Deprecated: Use RSS3TokenMetaData.ABI instead.
var RSS3TokenABI = RSS3TokenMetaData.ABI

// RSS3TokenBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use RSS3TokenMetaData.Bin instead.
var RSS3TokenBin = RSS3TokenMetaData.Bin

// DeployRSS3Token deploys a new Ethereum contract, binding an instance of RSS3Token to it.
func DeployRSS3Token(auth *bind.TransactOpts, backend bind.ContractBackend, _bridge common.Address, _remoteToken common.Address, _name string, _symbol string) (common.Address, *types.Transaction, *RSS3Token, error) {
	parsed, err := RSS3TokenMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(RSS3TokenBin), backend, _bridge, _remoteToken, _name, _symbol)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &RSS3Token{RSS3TokenCaller: RSS3TokenCaller{contract: contract}, RSS3TokenTransactor: RSS3TokenTransactor{contract: contract}, RSS3TokenFilterer: RSS3TokenFilterer{contract: contract}}, nil
}

// RSS3Token is an auto generated Go binding around an Ethereum contract.
type RSS3Token struct {
	RSS3TokenCaller     // Read-only binding to the contract
	RSS3TokenTransactor // Write-only binding to the contract
	RSS3TokenFilterer   // Log filterer for contract events
}

// RSS3TokenCaller is an auto generated read-only Go binding around an Ethereum contract.
type RSS3TokenCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RSS3TokenTransactor is an auto generated write-only Go binding around an Ethereum contract.
type RSS3TokenTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RSS3TokenFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type RSS3TokenFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RSS3TokenSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type RSS3TokenSession struct {
	Contract     *RSS3Token        // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// RSS3TokenCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type RSS3TokenCallerSession struct {
	Contract *RSS3TokenCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts    // Call options to use throughout this session
}

// RSS3TokenTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type RSS3TokenTransactorSession struct {
	Contract     *RSS3TokenTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// RSS3TokenRaw is an auto generated low-level Go binding around an Ethereum contract.
type RSS3TokenRaw struct {
	Contract *RSS3Token // Generic contract binding to access the raw methods on
}

// RSS3TokenCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type RSS3TokenCallerRaw struct {
	Contract *RSS3TokenCaller // Generic read-only contract binding to access the raw methods on
}

// RSS3TokenTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type RSS3TokenTransactorRaw struct {
	Contract *RSS3TokenTransactor // Generic write-only contract binding to access the raw methods on
}

// NewRSS3Token creates a new instance of RSS3Token, bound to a specific deployed contract.
func NewRSS3Token(address common.Address, backend bind.ContractBackend) (*RSS3Token, error) {
	contract, err := bindRSS3Token(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &RSS3Token{RSS3TokenCaller: RSS3TokenCaller{contract: contract}, RSS3TokenTransactor: RSS3TokenTransactor{contract: contract}, RSS3TokenFilterer: RSS3TokenFilterer{contract: contract}}, nil
}

// NewRSS3TokenCaller creates a new read-only instance of RSS3Token, bound to a specific deployed contract.
func NewRSS3TokenCaller(address common.Address, caller bind.ContractCaller) (*RSS3TokenCaller, error) {
	contract, err := bindRSS3Token(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &RSS3TokenCaller{contract: contract}, nil
}

// NewRSS3TokenTransactor creates a new write-only instance of RSS3Token, bound to a specific deployed contract.
func NewRSS3TokenTransactor(address common.Address, transactor bind.ContractTransactor) (*RSS3TokenTransactor, error) {
	contract, err := bindRSS3Token(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &RSS3TokenTransactor{contract: contract}, nil
}

// NewRSS3TokenFilterer creates a new log filterer instance of RSS3Token, bound to a specific deployed contract.
func NewRSS3TokenFilterer(address common.Address, filterer bind.ContractFilterer) (*RSS3TokenFilterer, error) {
	contract, err := bindRSS3Token(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &RSS3TokenFilterer{contract: contract}, nil
}

// bindRSS3Token binds a generic wrapper to an already deployed contract.
func bindRSS3Token(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(RSS3TokenABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_RSS3Token *RSS3TokenRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _RSS3Token.Contract.RSS3TokenCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_RSS3Token *RSS3TokenRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _RSS3Token.Contract.RSS3TokenTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_RSS3Token *RSS3TokenRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _RSS3Token.Contract.RSS3TokenTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_RSS3Token *RSS3TokenCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _RSS3Token.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_RSS3Token *RSS3TokenTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _RSS3Token.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_RSS3Token *RSS3TokenTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _RSS3Token.Contract.contract.Transact(opts, method, params...)
}

// BRIDGE is a free data retrieval call binding the contract method 0xee9a31a2.
//
// Solidity: function BRIDGE() view returns(address)
func (_RSS3Token *RSS3TokenCaller) BRIDGE(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _RSS3Token.contract.Call(opts, &out, "BRIDGE")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// BRIDGE is a free data retrieval call binding the contract method 0xee9a31a2.
//
// Solidity: function BRIDGE() view returns(address)
func (_RSS3Token *RSS3TokenSession) BRIDGE() (common.Address, error) {
	return _RSS3Token.Contract.BRIDGE(&_RSS3Token.CallOpts)
}

// BRIDGE is a free data retrieval call binding the contract method 0xee9a31a2.
//
// Solidity: function BRIDGE() view returns(address)
func (_RSS3Token *RSS3TokenCallerSession) BRIDGE() (common.Address, error) {
	return _RSS3Token.Contract.BRIDGE(&_RSS3Token.CallOpts)
}

// REMOTETOKEN is a free data retrieval call binding the contract method 0x033964be.
//
// Solidity: function REMOTE_TOKEN() view returns(address)
func (_RSS3Token *RSS3TokenCaller) REMOTETOKEN(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _RSS3Token.contract.Call(opts, &out, "REMOTE_TOKEN")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// REMOTETOKEN is a free data retrieval call binding the contract method 0x033964be.
//
// Solidity: function REMOTE_TOKEN() view returns(address)
func (_RSS3Token *RSS3TokenSession) REMOTETOKEN() (common.Address, error) {
	return _RSS3Token.Contract.REMOTETOKEN(&_RSS3Token.CallOpts)
}

// REMOTETOKEN is a free data retrieval call binding the contract method 0x033964be.
//
// Solidity: function REMOTE_TOKEN() view returns(address)
func (_RSS3Token *RSS3TokenCallerSession) REMOTETOKEN() (common.Address, error) {
	return _RSS3Token.Contract.REMOTETOKEN(&_RSS3Token.CallOpts)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_RSS3Token *RSS3TokenCaller) Allowance(opts *bind.CallOpts, owner common.Address, spender common.Address) (*big.Int, error) {
	var out []interface{}
	err := _RSS3Token.contract.Call(opts, &out, "allowance", owner, spender)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_RSS3Token *RSS3TokenSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _RSS3Token.Contract.Allowance(&_RSS3Token.CallOpts, owner, spender)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_RSS3Token *RSS3TokenCallerSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _RSS3Token.Contract.Allowance(&_RSS3Token.CallOpts, owner, spender)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_RSS3Token *RSS3TokenCaller) BalanceOf(opts *bind.CallOpts, account common.Address) (*big.Int, error) {
	var out []interface{}
	err := _RSS3Token.contract.Call(opts, &out, "balanceOf", account)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_RSS3Token *RSS3TokenSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _RSS3Token.Contract.BalanceOf(&_RSS3Token.CallOpts, account)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_RSS3Token *RSS3TokenCallerSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _RSS3Token.Contract.BalanceOf(&_RSS3Token.CallOpts, account)
}

// Bridge is a free data retrieval call binding the contract method 0xe78cea92.
//
// Solidity: function bridge() view returns(address)
func (_RSS3Token *RSS3TokenCaller) Bridge(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _RSS3Token.contract.Call(opts, &out, "bridge")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Bridge is a free data retrieval call binding the contract method 0xe78cea92.
//
// Solidity: function bridge() view returns(address)
func (_RSS3Token *RSS3TokenSession) Bridge() (common.Address, error) {
	return _RSS3Token.Contract.Bridge(&_RSS3Token.CallOpts)
}

// Bridge is a free data retrieval call binding the contract method 0xe78cea92.
//
// Solidity: function bridge() view returns(address)
func (_RSS3Token *RSS3TokenCallerSession) Bridge() (common.Address, error) {
	return _RSS3Token.Contract.Bridge(&_RSS3Token.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_RSS3Token *RSS3TokenCaller) Decimals(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _RSS3Token.contract.Call(opts, &out, "decimals")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_RSS3Token *RSS3TokenSession) Decimals() (uint8, error) {
	return _RSS3Token.Contract.Decimals(&_RSS3Token.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_RSS3Token *RSS3TokenCallerSession) Decimals() (uint8, error) {
	return _RSS3Token.Contract.Decimals(&_RSS3Token.CallOpts)
}

// L1Token is a free data retrieval call binding the contract method 0xc01e1bd6.
//
// Solidity: function l1Token() view returns(address)
func (_RSS3Token *RSS3TokenCaller) L1Token(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _RSS3Token.contract.Call(opts, &out, "l1Token")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// L1Token is a free data retrieval call binding the contract method 0xc01e1bd6.
//
// Solidity: function l1Token() view returns(address)
func (_RSS3Token *RSS3TokenSession) L1Token() (common.Address, error) {
	return _RSS3Token.Contract.L1Token(&_RSS3Token.CallOpts)
}

// L1Token is a free data retrieval call binding the contract method 0xc01e1bd6.
//
// Solidity: function l1Token() view returns(address)
func (_RSS3Token *RSS3TokenCallerSession) L1Token() (common.Address, error) {
	return _RSS3Token.Contract.L1Token(&_RSS3Token.CallOpts)
}

// L2Bridge is a free data retrieval call binding the contract method 0xae1f6aaf.
//
// Solidity: function l2Bridge() view returns(address)
func (_RSS3Token *RSS3TokenCaller) L2Bridge(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _RSS3Token.contract.Call(opts, &out, "l2Bridge")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// L2Bridge is a free data retrieval call binding the contract method 0xae1f6aaf.
//
// Solidity: function l2Bridge() view returns(address)
func (_RSS3Token *RSS3TokenSession) L2Bridge() (common.Address, error) {
	return _RSS3Token.Contract.L2Bridge(&_RSS3Token.CallOpts)
}

// L2Bridge is a free data retrieval call binding the contract method 0xae1f6aaf.
//
// Solidity: function l2Bridge() view returns(address)
func (_RSS3Token *RSS3TokenCallerSession) L2Bridge() (common.Address, error) {
	return _RSS3Token.Contract.L2Bridge(&_RSS3Token.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_RSS3Token *RSS3TokenCaller) Name(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _RSS3Token.contract.Call(opts, &out, "name")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_RSS3Token *RSS3TokenSession) Name() (string, error) {
	return _RSS3Token.Contract.Name(&_RSS3Token.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_RSS3Token *RSS3TokenCallerSession) Name() (string, error) {
	return _RSS3Token.Contract.Name(&_RSS3Token.CallOpts)
}

// RemoteToken is a free data retrieval call binding the contract method 0xd6c0b2c4.
//
// Solidity: function remoteToken() view returns(address)
func (_RSS3Token *RSS3TokenCaller) RemoteToken(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _RSS3Token.contract.Call(opts, &out, "remoteToken")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// RemoteToken is a free data retrieval call binding the contract method 0xd6c0b2c4.
//
// Solidity: function remoteToken() view returns(address)
func (_RSS3Token *RSS3TokenSession) RemoteToken() (common.Address, error) {
	return _RSS3Token.Contract.RemoteToken(&_RSS3Token.CallOpts)
}

// RemoteToken is a free data retrieval call binding the contract method 0xd6c0b2c4.
//
// Solidity: function remoteToken() view returns(address)
func (_RSS3Token *RSS3TokenCallerSession) RemoteToken() (common.Address, error) {
	return _RSS3Token.Contract.RemoteToken(&_RSS3Token.CallOpts)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 _interfaceId) pure returns(bool)
func (_RSS3Token *RSS3TokenCaller) SupportsInterface(opts *bind.CallOpts, _interfaceId [4]byte) (bool, error) {
	var out []interface{}
	err := _RSS3Token.contract.Call(opts, &out, "supportsInterface", _interfaceId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 _interfaceId) pure returns(bool)
func (_RSS3Token *RSS3TokenSession) SupportsInterface(_interfaceId [4]byte) (bool, error) {
	return _RSS3Token.Contract.SupportsInterface(&_RSS3Token.CallOpts, _interfaceId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 _interfaceId) pure returns(bool)
func (_RSS3Token *RSS3TokenCallerSession) SupportsInterface(_interfaceId [4]byte) (bool, error) {
	return _RSS3Token.Contract.SupportsInterface(&_RSS3Token.CallOpts, _interfaceId)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_RSS3Token *RSS3TokenCaller) Symbol(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _RSS3Token.contract.Call(opts, &out, "symbol")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_RSS3Token *RSS3TokenSession) Symbol() (string, error) {
	return _RSS3Token.Contract.Symbol(&_RSS3Token.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_RSS3Token *RSS3TokenCallerSession) Symbol() (string, error) {
	return _RSS3Token.Contract.Symbol(&_RSS3Token.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_RSS3Token *RSS3TokenCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _RSS3Token.contract.Call(opts, &out, "totalSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_RSS3Token *RSS3TokenSession) TotalSupply() (*big.Int, error) {
	return _RSS3Token.Contract.TotalSupply(&_RSS3Token.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_RSS3Token *RSS3TokenCallerSession) TotalSupply() (*big.Int, error) {
	return _RSS3Token.Contract.TotalSupply(&_RSS3Token.CallOpts)
}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() view returns(string)
func (_RSS3Token *RSS3TokenCaller) Version(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _RSS3Token.contract.Call(opts, &out, "version")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() view returns(string)
func (_RSS3Token *RSS3TokenSession) Version() (string, error) {
	return _RSS3Token.Contract.Version(&_RSS3Token.CallOpts)
}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() view returns(string)
func (_RSS3Token *RSS3TokenCallerSession) Version() (string, error) {
	return _RSS3Token.Contract.Version(&_RSS3Token.CallOpts)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_RSS3Token *RSS3TokenTransactor) Approve(opts *bind.TransactOpts, spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _RSS3Token.contract.Transact(opts, "approve", spender, amount)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_RSS3Token *RSS3TokenSession) Approve(spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _RSS3Token.Contract.Approve(&_RSS3Token.TransactOpts, spender, amount)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_RSS3Token *RSS3TokenTransactorSession) Approve(spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _RSS3Token.Contract.Approve(&_RSS3Token.TransactOpts, spender, amount)
}

// Burn is a paid mutator transaction binding the contract method 0x9dc29fac.
//
// Solidity: function burn(address _from, uint256 _amount) returns()
func (_RSS3Token *RSS3TokenTransactor) Burn(opts *bind.TransactOpts, _from common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _RSS3Token.contract.Transact(opts, "burn", _from, _amount)
}

// Burn is a paid mutator transaction binding the contract method 0x9dc29fac.
//
// Solidity: function burn(address _from, uint256 _amount) returns()
func (_RSS3Token *RSS3TokenSession) Burn(_from common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _RSS3Token.Contract.Burn(&_RSS3Token.TransactOpts, _from, _amount)
}

// Burn is a paid mutator transaction binding the contract method 0x9dc29fac.
//
// Solidity: function burn(address _from, uint256 _amount) returns()
func (_RSS3Token *RSS3TokenTransactorSession) Burn(_from common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _RSS3Token.Contract.Burn(&_RSS3Token.TransactOpts, _from, _amount)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address spender, uint256 subtractedValue) returns(bool)
func (_RSS3Token *RSS3TokenTransactor) DecreaseAllowance(opts *bind.TransactOpts, spender common.Address, subtractedValue *big.Int) (*types.Transaction, error) {
	return _RSS3Token.contract.Transact(opts, "decreaseAllowance", spender, subtractedValue)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address spender, uint256 subtractedValue) returns(bool)
func (_RSS3Token *RSS3TokenSession) DecreaseAllowance(spender common.Address, subtractedValue *big.Int) (*types.Transaction, error) {
	return _RSS3Token.Contract.DecreaseAllowance(&_RSS3Token.TransactOpts, spender, subtractedValue)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address spender, uint256 subtractedValue) returns(bool)
func (_RSS3Token *RSS3TokenTransactorSession) DecreaseAllowance(spender common.Address, subtractedValue *big.Int) (*types.Transaction, error) {
	return _RSS3Token.Contract.DecreaseAllowance(&_RSS3Token.TransactOpts, spender, subtractedValue)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address spender, uint256 addedValue) returns(bool)
func (_RSS3Token *RSS3TokenTransactor) IncreaseAllowance(opts *bind.TransactOpts, spender common.Address, addedValue *big.Int) (*types.Transaction, error) {
	return _RSS3Token.contract.Transact(opts, "increaseAllowance", spender, addedValue)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address spender, uint256 addedValue) returns(bool)
func (_RSS3Token *RSS3TokenSession) IncreaseAllowance(spender common.Address, addedValue *big.Int) (*types.Transaction, error) {
	return _RSS3Token.Contract.IncreaseAllowance(&_RSS3Token.TransactOpts, spender, addedValue)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address spender, uint256 addedValue) returns(bool)
func (_RSS3Token *RSS3TokenTransactorSession) IncreaseAllowance(spender common.Address, addedValue *big.Int) (*types.Transaction, error) {
	return _RSS3Token.Contract.IncreaseAllowance(&_RSS3Token.TransactOpts, spender, addedValue)
}

// Mint is a paid mutator transaction binding the contract method 0x40c10f19.
//
// Solidity: function mint(address _to, uint256 _amount) returns()
func (_RSS3Token *RSS3TokenTransactor) Mint(opts *bind.TransactOpts, _to common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _RSS3Token.contract.Transact(opts, "mint", _to, _amount)
}

// Mint is a paid mutator transaction binding the contract method 0x40c10f19.
//
// Solidity: function mint(address _to, uint256 _amount) returns()
func (_RSS3Token *RSS3TokenSession) Mint(_to common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _RSS3Token.Contract.Mint(&_RSS3Token.TransactOpts, _to, _amount)
}

// Mint is a paid mutator transaction binding the contract method 0x40c10f19.
//
// Solidity: function mint(address _to, uint256 _amount) returns()
func (_RSS3Token *RSS3TokenTransactorSession) Mint(_to common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _RSS3Token.Contract.Mint(&_RSS3Token.TransactOpts, _to, _amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 amount) returns(bool)
func (_RSS3Token *RSS3TokenTransactor) Transfer(opts *bind.TransactOpts, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _RSS3Token.contract.Transact(opts, "transfer", to, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 amount) returns(bool)
func (_RSS3Token *RSS3TokenSession) Transfer(to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _RSS3Token.Contract.Transfer(&_RSS3Token.TransactOpts, to, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 amount) returns(bool)
func (_RSS3Token *RSS3TokenTransactorSession) Transfer(to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _RSS3Token.Contract.Transfer(&_RSS3Token.TransactOpts, to, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 amount) returns(bool)
func (_RSS3Token *RSS3TokenTransactor) TransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _RSS3Token.contract.Transact(opts, "transferFrom", from, to, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 amount) returns(bool)
func (_RSS3Token *RSS3TokenSession) TransferFrom(from common.Address, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _RSS3Token.Contract.TransferFrom(&_RSS3Token.TransactOpts, from, to, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 amount) returns(bool)
func (_RSS3Token *RSS3TokenTransactorSession) TransferFrom(from common.Address, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _RSS3Token.Contract.TransferFrom(&_RSS3Token.TransactOpts, from, to, amount)
}

// RSS3TokenApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the RSS3Token contract.
type RSS3TokenApprovalIterator struct {
	Event *RSS3TokenApproval // Event containing the contract specifics and raw log

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
func (it *RSS3TokenApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RSS3TokenApproval)
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
		it.Event = new(RSS3TokenApproval)
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
func (it *RSS3TokenApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RSS3TokenApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RSS3TokenApproval represents a Approval event raised by the RSS3Token contract.
type RSS3TokenApproval struct {
	Owner   common.Address
	Spender common.Address
	Value   *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_RSS3Token *RSS3TokenFilterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, spender []common.Address) (*RSS3TokenApprovalIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _RSS3Token.contract.FilterLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return &RSS3TokenApprovalIterator{contract: _RSS3Token.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_RSS3Token *RSS3TokenFilterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *RSS3TokenApproval, owner []common.Address, spender []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _RSS3Token.contract.WatchLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RSS3TokenApproval)
				if err := _RSS3Token.contract.UnpackLog(event, "Approval", log); err != nil {
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
func (_RSS3Token *RSS3TokenFilterer) ParseApproval(log types.Log) (*RSS3TokenApproval, error) {
	event := new(RSS3TokenApproval)
	if err := _RSS3Token.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// RSS3TokenBurnIterator is returned from FilterBurn and is used to iterate over the raw logs and unpacked data for Burn events raised by the RSS3Token contract.
type RSS3TokenBurnIterator struct {
	Event *RSS3TokenBurn // Event containing the contract specifics and raw log

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
func (it *RSS3TokenBurnIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RSS3TokenBurn)
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
		it.Event = new(RSS3TokenBurn)
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
func (it *RSS3TokenBurnIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RSS3TokenBurnIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RSS3TokenBurn represents a Burn event raised by the RSS3Token contract.
type RSS3TokenBurn struct {
	Account common.Address
	Amount  *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterBurn is a free log retrieval operation binding the contract event 0xcc16f5dbb4873280815c1ee09dbd06736cffcc184412cf7a71a0fdb75d397ca5.
//
// Solidity: event Burn(address indexed account, uint256 amount)
func (_RSS3Token *RSS3TokenFilterer) FilterBurn(opts *bind.FilterOpts, account []common.Address) (*RSS3TokenBurnIterator, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _RSS3Token.contract.FilterLogs(opts, "Burn", accountRule)
	if err != nil {
		return nil, err
	}
	return &RSS3TokenBurnIterator{contract: _RSS3Token.contract, event: "Burn", logs: logs, sub: sub}, nil
}

// WatchBurn is a free log subscription operation binding the contract event 0xcc16f5dbb4873280815c1ee09dbd06736cffcc184412cf7a71a0fdb75d397ca5.
//
// Solidity: event Burn(address indexed account, uint256 amount)
func (_RSS3Token *RSS3TokenFilterer) WatchBurn(opts *bind.WatchOpts, sink chan<- *RSS3TokenBurn, account []common.Address) (event.Subscription, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _RSS3Token.contract.WatchLogs(opts, "Burn", accountRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RSS3TokenBurn)
				if err := _RSS3Token.contract.UnpackLog(event, "Burn", log); err != nil {
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

// ParseBurn is a log parse operation binding the contract event 0xcc16f5dbb4873280815c1ee09dbd06736cffcc184412cf7a71a0fdb75d397ca5.
//
// Solidity: event Burn(address indexed account, uint256 amount)
func (_RSS3Token *RSS3TokenFilterer) ParseBurn(log types.Log) (*RSS3TokenBurn, error) {
	event := new(RSS3TokenBurn)
	if err := _RSS3Token.contract.UnpackLog(event, "Burn", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// RSS3TokenMintIterator is returned from FilterMint and is used to iterate over the raw logs and unpacked data for Mint events raised by the RSS3Token contract.
type RSS3TokenMintIterator struct {
	Event *RSS3TokenMint // Event containing the contract specifics and raw log

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
func (it *RSS3TokenMintIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RSS3TokenMint)
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
		it.Event = new(RSS3TokenMint)
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
func (it *RSS3TokenMintIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RSS3TokenMintIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RSS3TokenMint represents a Mint event raised by the RSS3Token contract.
type RSS3TokenMint struct {
	Account common.Address
	Amount  *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterMint is a free log retrieval operation binding the contract event 0x0f6798a560793a54c3bcfe86a93cde1e73087d944c0ea20544137d4121396885.
//
// Solidity: event Mint(address indexed account, uint256 amount)
func (_RSS3Token *RSS3TokenFilterer) FilterMint(opts *bind.FilterOpts, account []common.Address) (*RSS3TokenMintIterator, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _RSS3Token.contract.FilterLogs(opts, "Mint", accountRule)
	if err != nil {
		return nil, err
	}
	return &RSS3TokenMintIterator{contract: _RSS3Token.contract, event: "Mint", logs: logs, sub: sub}, nil
}

// WatchMint is a free log subscription operation binding the contract event 0x0f6798a560793a54c3bcfe86a93cde1e73087d944c0ea20544137d4121396885.
//
// Solidity: event Mint(address indexed account, uint256 amount)
func (_RSS3Token *RSS3TokenFilterer) WatchMint(opts *bind.WatchOpts, sink chan<- *RSS3TokenMint, account []common.Address) (event.Subscription, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _RSS3Token.contract.WatchLogs(opts, "Mint", accountRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RSS3TokenMint)
				if err := _RSS3Token.contract.UnpackLog(event, "Mint", log); err != nil {
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

// ParseMint is a log parse operation binding the contract event 0x0f6798a560793a54c3bcfe86a93cde1e73087d944c0ea20544137d4121396885.
//
// Solidity: event Mint(address indexed account, uint256 amount)
func (_RSS3Token *RSS3TokenFilterer) ParseMint(log types.Log) (*RSS3TokenMint, error) {
	event := new(RSS3TokenMint)
	if err := _RSS3Token.contract.UnpackLog(event, "Mint", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// RSS3TokenTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the RSS3Token contract.
type RSS3TokenTransferIterator struct {
	Event *RSS3TokenTransfer // Event containing the contract specifics and raw log

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
func (it *RSS3TokenTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RSS3TokenTransfer)
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
		it.Event = new(RSS3TokenTransfer)
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
func (it *RSS3TokenTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RSS3TokenTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RSS3TokenTransfer represents a Transfer event raised by the RSS3Token contract.
type RSS3TokenTransfer struct {
	From  common.Address
	To    common.Address
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_RSS3Token *RSS3TokenFilterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*RSS3TokenTransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _RSS3Token.contract.FilterLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &RSS3TokenTransferIterator{contract: _RSS3Token.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_RSS3Token *RSS3TokenFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *RSS3TokenTransfer, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _RSS3Token.contract.WatchLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RSS3TokenTransfer)
				if err := _RSS3Token.contract.UnpackLog(event, "Transfer", log); err != nil {
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
func (_RSS3Token *RSS3TokenFilterer) ParseTransfer(log types.Log) (*RSS3TokenTransfer, error) {
	event := new(RSS3TokenTransfer)
	if err := _RSS3Token.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
