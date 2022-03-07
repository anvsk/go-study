// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package vote

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

// DemoMetaData contains all meta data concerning the Demo contract.
var DemoMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"bytes32[]\",\"name\":\"proposalNames\",\"type\":\"bytes32[]\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"chairperson\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"delegate\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"voter\",\"type\":\"address\"}],\"name\":\"giveRightToVote\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"proposals\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"name\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"voteCount\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"proposal\",\"type\":\"uint256\"}],\"name\":\"vote\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"voters\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"weight\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"voted\",\"type\":\"bool\"},{\"internalType\":\"address\",\"name\":\"delegate\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"vote\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"winnerName\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"winnerName_\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"winningProposal\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"winningProposal_\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Bin: "0x60806040523480156200001157600080fd5b506040516200146538038062001465833981810160405281019062000037919062000342565b336000806101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555060018060008060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000206000018190555060005b81518110156200017657600260405180604001604052808484815181106200010f576200010e62000393565b5b60200260200101518152602001600081525090806001815401808255809150506001900390600052602060002090600202016000909190919091506000820151816000015560208201518160010155505080806200016d90620003fb565b915050620000e2565b505062000449565b6000604051905090565b600080fd5b600080fd5b600080fd5b6000601f19601f8301169050919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b620001e28262000197565b810181811067ffffffffffffffff82111715620002045762000203620001a8565b5b80604052505050565b6000620002196200017e565b9050620002278282620001d7565b919050565b600067ffffffffffffffff8211156200024a5762000249620001a8565b5b602082029050602081019050919050565b600080fd5b6000819050919050565b620002758162000260565b81146200028157600080fd5b50565b60008151905062000295816200026a565b92915050565b6000620002b2620002ac846200022c565b6200020d565b90508083825260208201905060208402830185811115620002d857620002d76200025b565b5b835b81811015620003055780620002f0888262000284565b845260208401935050602081019050620002da565b5050509392505050565b600082601f83011262000327576200032662000192565b5b8151620003398482602086016200029b565b91505092915050565b6000602082840312156200035b576200035a62000188565b5b600082015167ffffffffffffffff8111156200037c576200037b6200018d565b5b6200038a848285016200030f565b91505092915050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b6000819050919050565b60006200040882620003f1565b91507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8214156200043e576200043d620003c2565b5b600182019050919050565b61100c80620004596000396000f3fe608060405234801561001057600080fd5b50600436106100885760003560e01c8063609ff1bd1161005b578063609ff1bd146101145780639e7b8d6114610132578063a3ec138d1461014e578063e2ba53f01461018157610088565b80630121b93f1461008d578063013cf08b146100a95780632e4176cf146100da5780635c19a95c146100f8575b600080fd5b6100a760048036038101906100a291906109e5565b61019f565b005b6100c360048036038101906100be91906109e5565b6102e6565b6040516100d1929190610a3a565b60405180910390f35b6100e261031a565b6040516100ef9190610aa4565b60405180910390f35b610112600480360381019061010d9190610aeb565b61033e565b005b61011c6106da565b6040516101299190610b18565b60405180910390f35b61014c60048036038101906101479190610aeb565b610762565b005b61016860048036038101906101639190610aeb565b610919565b6040516101789493929190610b4e565b60405180910390f35b610189610976565b6040516101969190610b93565b60405180910390f35b6000600160003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020905060008160000154141561022a576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161022190610c0b565b60405180910390fd5b8060010160009054906101000a900460ff161561027c576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161027390610c77565b60405180910390fd5b60018160010160006101000a81548160ff0219169083151502179055508181600201819055508060000154600283815481106102bb576102ba610c97565b5b906000526020600020906002020160010160008282546102db9190610cf5565b925050819055505050565b600281815481106102f657600080fd5b90600052602060002090600202016000915090508060000154908060010154905082565b60008054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b6000600160003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002090508060010160009054906101000a900460ff16156103d3576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016103ca90610d97565b60405180910390fd5b3373ffffffffffffffffffffffffffffffffffffffff168273ffffffffffffffffffffffffffffffffffffffff161415610442576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161043990610e03565b60405180910390fd5b5b600073ffffffffffffffffffffffffffffffffffffffff16600160008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060010160019054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16146105b257600160008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060010160019054906101000a900473ffffffffffffffffffffffffffffffffffffffff1691503373ffffffffffffffffffffffffffffffffffffffff168273ffffffffffffffffffffffffffffffffffffffff1614156105ad576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016105a490610e6f565b60405180910390fd5b610443565b60018160010160006101000a81548160ff021916908315150217905550818160010160016101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055506000600160008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002090508060010160009054906101000a900460ff16156106b5578160000154600282600201548154811061068957610688610c97565b5b906000526020600020906002020160010160008282546106a99190610cf5565b925050819055506106d5565b81600001548160000160008282546106cd9190610cf5565b925050819055505b505050565b6000806000905060005b60028054905081101561075d57816002828154811061070657610705610c97565b5b906000526020600020906002020160010154111561074a576002818154811061073257610731610c97565b5b90600052602060002090600202016001015491508092505b808061075590610e8f565b9150506106e4565b505090565b60008054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff16146107f0576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016107e790610f4a565b60405180910390fd5b600160008273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060010160009054906101000a900460ff1615610880576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161087790610fb6565b60405180910390fd5b6000600160008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060000154146108cf57600080fd5b60018060008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000206000018190555050565b60016020528060005260406000206000915090508060000154908060010160009054906101000a900460ff16908060010160019054906101000a900473ffffffffffffffffffffffffffffffffffffffff16908060020154905084565b600060026109826106da565b8154811061099357610992610c97565b5b906000526020600020906002020160000154905090565b600080fd5b6000819050919050565b6109c2816109af565b81146109cd57600080fd5b50565b6000813590506109df816109b9565b92915050565b6000602082840312156109fb576109fa6109aa565b5b6000610a09848285016109d0565b91505092915050565b6000819050919050565b610a2581610a12565b82525050565b610a34816109af565b82525050565b6000604082019050610a4f6000830185610a1c565b610a5c6020830184610a2b565b9392505050565b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b6000610a8e82610a63565b9050919050565b610a9e81610a83565b82525050565b6000602082019050610ab96000830184610a95565b92915050565b610ac881610a83565b8114610ad357600080fd5b50565b600081359050610ae581610abf565b92915050565b600060208284031215610b0157610b006109aa565b5b6000610b0f84828501610ad6565b91505092915050565b6000602082019050610b2d6000830184610a2b565b92915050565b60008115159050919050565b610b4881610b33565b82525050565b6000608082019050610b636000830187610a2b565b610b706020830186610b3f565b610b7d6040830185610a95565b610b8a6060830184610a2b565b95945050505050565b6000602082019050610ba86000830184610a1c565b92915050565b600082825260208201905092915050565b7f486173206e6f20726967687420746f20766f7465000000000000000000000000600082015250565b6000610bf5601483610bae565b9150610c0082610bbf565b602082019050919050565b60006020820190508181036000830152610c2481610be8565b9050919050565b7f416c726561647920766f7465642e000000000000000000000000000000000000600082015250565b6000610c61600e83610bae565b9150610c6c82610c2b565b602082019050919050565b60006020820190508181036000830152610c9081610c54565b9050919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b6000610d00826109af565b9150610d0b836109af565b9250827fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff03821115610d4057610d3f610cc6565b5b828201905092915050565b7f596f7520616c726561647920766f7465642e0000000000000000000000000000600082015250565b6000610d81601283610bae565b9150610d8c82610d4b565b602082019050919050565b60006020820190508181036000830152610db081610d74565b9050919050565b7f53656c662d64656c65676174696f6e20697320646973616c6c6f7765642e0000600082015250565b6000610ded601e83610bae565b9150610df882610db7565b602082019050919050565b60006020820190508181036000830152610e1c81610de0565b9050919050565b7f466f756e64206c6f6f7020696e2064656c65676174696f6e2e00000000000000600082015250565b6000610e59601983610bae565b9150610e6482610e23565b602082019050919050565b60006020820190508181036000830152610e8881610e4c565b9050919050565b6000610e9a826109af565b91507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff821415610ecd57610ecc610cc6565b5b600182019050919050565b7f4f6e6c79206368616972706572736f6e2063616e20676976652072696768742060008201527f746f20766f74652e000000000000000000000000000000000000000000000000602082015250565b6000610f34602883610bae565b9150610f3f82610ed8565b604082019050919050565b60006020820190508181036000830152610f6381610f27565b9050919050565b7f54686520766f74657220616c726561647920766f7465642e0000000000000000600082015250565b6000610fa0601883610bae565b9150610fab82610f6a565b602082019050919050565b60006020820190508181036000830152610fcf81610f93565b905091905056fea264697066735822122048f03849ce24a5baa4b95de5e403901d07e4903ad3bf3ad065ce3d462548d53d64736f6c634300080b0033",
}

// DemoABI is the input ABI used to generate the binding from.
// Deprecated: Use DemoMetaData.ABI instead.
var DemoABI = DemoMetaData.ABI

// DemoBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use DemoMetaData.Bin instead.
var DemoBin = DemoMetaData.Bin

// DeployDemo deploys a new Ethereum contract, binding an instance of Demo to it.
func DeployDemo(auth *bind.TransactOpts, backend bind.ContractBackend, proposalNames [][32]byte) (common.Address, *types.Transaction, *Demo, error) {
	parsed, err := DemoMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(DemoBin), backend, proposalNames)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Demo{DemoCaller: DemoCaller{contract: contract}, DemoTransactor: DemoTransactor{contract: contract}, DemoFilterer: DemoFilterer{contract: contract}}, nil
}

// Demo is an auto generated Go binding around an Ethereum contract.
type Demo struct {
	DemoCaller     // Read-only binding to the contract
	DemoTransactor // Write-only binding to the contract
	DemoFilterer   // Log filterer for contract events
}

// DemoCaller is an auto generated read-only Go binding around an Ethereum contract.
type DemoCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DemoTransactor is an auto generated write-only Go binding around an Ethereum contract.
type DemoTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DemoFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type DemoFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DemoSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type DemoSession struct {
	Contract     *Demo             // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// DemoCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type DemoCallerSession struct {
	Contract *DemoCaller   // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// DemoTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type DemoTransactorSession struct {
	Contract     *DemoTransactor   // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// DemoRaw is an auto generated low-level Go binding around an Ethereum contract.
type DemoRaw struct {
	Contract *Demo // Generic contract binding to access the raw methods on
}

// DemoCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type DemoCallerRaw struct {
	Contract *DemoCaller // Generic read-only contract binding to access the raw methods on
}

// DemoTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type DemoTransactorRaw struct {
	Contract *DemoTransactor // Generic write-only contract binding to access the raw methods on
}

// NewDemo creates a new instance of Demo, bound to a specific deployed contract.
func NewDemo(address common.Address, backend bind.ContractBackend) (*Demo, error) {
	contract, err := bindDemo(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Demo{DemoCaller: DemoCaller{contract: contract}, DemoTransactor: DemoTransactor{contract: contract}, DemoFilterer: DemoFilterer{contract: contract}}, nil
}

// NewDemoCaller creates a new read-only instance of Demo, bound to a specific deployed contract.
func NewDemoCaller(address common.Address, caller bind.ContractCaller) (*DemoCaller, error) {
	contract, err := bindDemo(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &DemoCaller{contract: contract}, nil
}

// NewDemoTransactor creates a new write-only instance of Demo, bound to a specific deployed contract.
func NewDemoTransactor(address common.Address, transactor bind.ContractTransactor) (*DemoTransactor, error) {
	contract, err := bindDemo(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &DemoTransactor{contract: contract}, nil
}

// NewDemoFilterer creates a new log filterer instance of Demo, bound to a specific deployed contract.
func NewDemoFilterer(address common.Address, filterer bind.ContractFilterer) (*DemoFilterer, error) {
	contract, err := bindDemo(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &DemoFilterer{contract: contract}, nil
}

// bindDemo binds a generic wrapper to an already deployed contract.
func bindDemo(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(DemoABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Demo *DemoRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Demo.Contract.DemoCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Demo *DemoRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Demo.Contract.DemoTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Demo *DemoRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Demo.Contract.DemoTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Demo *DemoCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Demo.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Demo *DemoTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Demo.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Demo *DemoTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Demo.Contract.contract.Transact(opts, method, params...)
}

// Chairperson is a free data retrieval call binding the contract method 0x2e4176cf.
//
// Solidity: function chairperson() view returns(address)
func (_Demo *DemoCaller) Chairperson(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Demo.contract.Call(opts, &out, "chairperson")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Chairperson is a free data retrieval call binding the contract method 0x2e4176cf.
//
// Solidity: function chairperson() view returns(address)
func (_Demo *DemoSession) Chairperson() (common.Address, error) {
	return _Demo.Contract.Chairperson(&_Demo.CallOpts)
}

// Chairperson is a free data retrieval call binding the contract method 0x2e4176cf.
//
// Solidity: function chairperson() view returns(address)
func (_Demo *DemoCallerSession) Chairperson() (common.Address, error) {
	return _Demo.Contract.Chairperson(&_Demo.CallOpts)
}

// Proposals is a free data retrieval call binding the contract method 0x013cf08b.
//
// Solidity: function proposals(uint256 ) view returns(bytes32 name, uint256 voteCount)
func (_Demo *DemoCaller) Proposals(opts *bind.CallOpts, arg0 *big.Int) (struct {
	Name      [32]byte
	VoteCount *big.Int
}, error) {
	var out []interface{}
	err := _Demo.contract.Call(opts, &out, "proposals", arg0)

	outstruct := new(struct {
		Name      [32]byte
		VoteCount *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Name = *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)
	outstruct.VoteCount = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// Proposals is a free data retrieval call binding the contract method 0x013cf08b.
//
// Solidity: function proposals(uint256 ) view returns(bytes32 name, uint256 voteCount)
func (_Demo *DemoSession) Proposals(arg0 *big.Int) (struct {
	Name      [32]byte
	VoteCount *big.Int
}, error) {
	return _Demo.Contract.Proposals(&_Demo.CallOpts, arg0)
}

// Proposals is a free data retrieval call binding the contract method 0x013cf08b.
//
// Solidity: function proposals(uint256 ) view returns(bytes32 name, uint256 voteCount)
func (_Demo *DemoCallerSession) Proposals(arg0 *big.Int) (struct {
	Name      [32]byte
	VoteCount *big.Int
}, error) {
	return _Demo.Contract.Proposals(&_Demo.CallOpts, arg0)
}

// Voters is a free data retrieval call binding the contract method 0xa3ec138d.
//
// Solidity: function voters(address ) view returns(uint256 weight, bool voted, address delegate, uint256 vote)
func (_Demo *DemoCaller) Voters(opts *bind.CallOpts, arg0 common.Address) (struct {
	Weight   *big.Int
	Voted    bool
	Delegate common.Address
	Vote     *big.Int
}, error) {
	var out []interface{}
	err := _Demo.contract.Call(opts, &out, "voters", arg0)

	outstruct := new(struct {
		Weight   *big.Int
		Voted    bool
		Delegate common.Address
		Vote     *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Weight = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.Voted = *abi.ConvertType(out[1], new(bool)).(*bool)
	outstruct.Delegate = *abi.ConvertType(out[2], new(common.Address)).(*common.Address)
	outstruct.Vote = *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// Voters is a free data retrieval call binding the contract method 0xa3ec138d.
//
// Solidity: function voters(address ) view returns(uint256 weight, bool voted, address delegate, uint256 vote)
func (_Demo *DemoSession) Voters(arg0 common.Address) (struct {
	Weight   *big.Int
	Voted    bool
	Delegate common.Address
	Vote     *big.Int
}, error) {
	return _Demo.Contract.Voters(&_Demo.CallOpts, arg0)
}

// Voters is a free data retrieval call binding the contract method 0xa3ec138d.
//
// Solidity: function voters(address ) view returns(uint256 weight, bool voted, address delegate, uint256 vote)
func (_Demo *DemoCallerSession) Voters(arg0 common.Address) (struct {
	Weight   *big.Int
	Voted    bool
	Delegate common.Address
	Vote     *big.Int
}, error) {
	return _Demo.Contract.Voters(&_Demo.CallOpts, arg0)
}

// WinnerName is a free data retrieval call binding the contract method 0xe2ba53f0.
//
// Solidity: function winnerName() view returns(bytes32 winnerName_)
func (_Demo *DemoCaller) WinnerName(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Demo.contract.Call(opts, &out, "winnerName")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// WinnerName is a free data retrieval call binding the contract method 0xe2ba53f0.
//
// Solidity: function winnerName() view returns(bytes32 winnerName_)
func (_Demo *DemoSession) WinnerName() ([32]byte, error) {
	return _Demo.Contract.WinnerName(&_Demo.CallOpts)
}

// WinnerName is a free data retrieval call binding the contract method 0xe2ba53f0.
//
// Solidity: function winnerName() view returns(bytes32 winnerName_)
func (_Demo *DemoCallerSession) WinnerName() ([32]byte, error) {
	return _Demo.Contract.WinnerName(&_Demo.CallOpts)
}

// WinningProposal is a free data retrieval call binding the contract method 0x609ff1bd.
//
// Solidity: function winningProposal() view returns(uint256 winningProposal_)
func (_Demo *DemoCaller) WinningProposal(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Demo.contract.Call(opts, &out, "winningProposal")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// WinningProposal is a free data retrieval call binding the contract method 0x609ff1bd.
//
// Solidity: function winningProposal() view returns(uint256 winningProposal_)
func (_Demo *DemoSession) WinningProposal() (*big.Int, error) {
	return _Demo.Contract.WinningProposal(&_Demo.CallOpts)
}

// WinningProposal is a free data retrieval call binding the contract method 0x609ff1bd.
//
// Solidity: function winningProposal() view returns(uint256 winningProposal_)
func (_Demo *DemoCallerSession) WinningProposal() (*big.Int, error) {
	return _Demo.Contract.WinningProposal(&_Demo.CallOpts)
}

// Delegate is a paid mutator transaction binding the contract method 0x5c19a95c.
//
// Solidity: function delegate(address to) returns()
func (_Demo *DemoTransactor) Delegate(opts *bind.TransactOpts, to common.Address) (*types.Transaction, error) {
	return _Demo.contract.Transact(opts, "delegate", to)
}

// Delegate is a paid mutator transaction binding the contract method 0x5c19a95c.
//
// Solidity: function delegate(address to) returns()
func (_Demo *DemoSession) Delegate(to common.Address) (*types.Transaction, error) {
	return _Demo.Contract.Delegate(&_Demo.TransactOpts, to)
}

// Delegate is a paid mutator transaction binding the contract method 0x5c19a95c.
//
// Solidity: function delegate(address to) returns()
func (_Demo *DemoTransactorSession) Delegate(to common.Address) (*types.Transaction, error) {
	return _Demo.Contract.Delegate(&_Demo.TransactOpts, to)
}

// GiveRightToVote is a paid mutator transaction binding the contract method 0x9e7b8d61.
//
// Solidity: function giveRightToVote(address voter) returns()
func (_Demo *DemoTransactor) GiveRightToVote(opts *bind.TransactOpts, voter common.Address) (*types.Transaction, error) {
	return _Demo.contract.Transact(opts, "giveRightToVote", voter)
}

// GiveRightToVote is a paid mutator transaction binding the contract method 0x9e7b8d61.
//
// Solidity: function giveRightToVote(address voter) returns()
func (_Demo *DemoSession) GiveRightToVote(voter common.Address) (*types.Transaction, error) {
	return _Demo.Contract.GiveRightToVote(&_Demo.TransactOpts, voter)
}

// GiveRightToVote is a paid mutator transaction binding the contract method 0x9e7b8d61.
//
// Solidity: function giveRightToVote(address voter) returns()
func (_Demo *DemoTransactorSession) GiveRightToVote(voter common.Address) (*types.Transaction, error) {
	return _Demo.Contract.GiveRightToVote(&_Demo.TransactOpts, voter)
}

// Vote is a paid mutator transaction binding the contract method 0x0121b93f.
//
// Solidity: function vote(uint256 proposal) returns()
func (_Demo *DemoTransactor) Vote(opts *bind.TransactOpts, proposal *big.Int) (*types.Transaction, error) {
	return _Demo.contract.Transact(opts, "vote", proposal)
}

// Vote is a paid mutator transaction binding the contract method 0x0121b93f.
//
// Solidity: function vote(uint256 proposal) returns()
func (_Demo *DemoSession) Vote(proposal *big.Int) (*types.Transaction, error) {
	return _Demo.Contract.Vote(&_Demo.TransactOpts, proposal)
}

// Vote is a paid mutator transaction binding the contract method 0x0121b93f.
//
// Solidity: function vote(uint256 proposal) returns()
func (_Demo *DemoTransactorSession) Vote(proposal *big.Int) (*types.Transaction, error) {
	return _Demo.Contract.Vote(&_Demo.TransactOpts, proposal)
}
