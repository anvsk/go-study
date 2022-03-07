// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package demopay

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

// DemoPayCar is an auto generated low-level Go binding around an user-defined struct.
type DemoPayCar struct {
	Id          *big.Int
	Goodsdetail []DemoPayGoodsRow
	Members     []common.Address
}

// DemoPayGoods is an auto generated low-level Go binding around an user-defined struct.
type DemoPayGoods struct {
	Id    *big.Int
	Name  string
	Price *big.Int
}

// DemoPayGoodsRow is an auto generated low-level Go binding around an user-defined struct.
type DemoPayGoodsRow struct {
	Goods DemoPayGoods
	Nums  *big.Int
}

// DemoPayOrder is an auto generated low-level Go binding around an user-defined struct.
type DemoPayOrder struct {
	Carnum    *big.Int
	Cardetail DemoPayCar
	Amount    *big.Int
	Confirmed bool
	Payed     bool
}

// DemopayMetaData contains all meta data concerning the Demopay contract.
var DemopayMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"price\",\"type\":\"uint256\"}],\"internalType\":\"structDemoPay.Goods[]\",\"name\":\"goods\",\"type\":\"tuple[]\"},{\"internalType\":\"address\",\"name\":\"bossaddress\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"gid\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"nums\",\"type\":\"uint256\"}],\"name\":\"addGoods\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"confirmOrder\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getCarNum\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"carnum\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getOrder\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"newCar\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"carnum\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"payOrder\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"showOrder\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"carnum\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"components\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"price\",\"type\":\"uint256\"}],\"internalType\":\"structDemoPay.Goods\",\"name\":\"goods\",\"type\":\"tuple\"},{\"internalType\":\"uint256\",\"name\":\"nums\",\"type\":\"uint256\"}],\"internalType\":\"structDemoPay.GoodsRow[]\",\"name\":\"goodsdetail\",\"type\":\"tuple[]\"},{\"internalType\":\"address[]\",\"name\":\"members\",\"type\":\"address[]\"}],\"internalType\":\"structDemoPay.Car\",\"name\":\"cardetail\",\"type\":\"tuple\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"confirmed\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"payed\",\"type\":\"bool\"}],\"internalType\":\"structDemoPay.Order\",\"name\":\"order\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"carnum\",\"type\":\"uint256\"}],\"name\":\"takeCar\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x60806040523480156200001157600080fd5b50604051620021f0380380620021f08339818101604052810190620000379190620005ac565b60005b8251811015620000e2578281815181106200005a576200005962000612565b5b6020026020010151600160008584815181106200007c576200007b62000612565b5b6020026020010151600001518152602001908152602001600020600082015181600001556020820151816001019080519060200190620000be9291906200012b565b50604082015181600201559050508080620000d99062000670565b9150506200003a565b50806000806101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550505062000723565b8280546200013990620006ed565b90600052602060002090601f0160209004810192826200015d5760008555620001a9565b82601f106200017857805160ff1916838001178555620001a9565b82800160010185558215620001a9579182015b82811115620001a85782518255916020019190600101906200018b565b5b509050620001b89190620001bc565b5090565b5b80821115620001d7576000816000905550600101620001bd565b5090565b6000604051905090565b600080fd5b600080fd5b600080fd5b6000601f19601f8301169050919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b6200023f82620001f4565b810181811067ffffffffffffffff8211171562000261576200026062000205565b5b80604052505050565b600062000276620001db565b905062000284828262000234565b919050565b600067ffffffffffffffff821115620002a757620002a662000205565b5b602082029050602081019050919050565b600080fd5b600080fd5b600080fd5b6000819050919050565b620002dc81620002c7565b8114620002e857600080fd5b50565b600081519050620002fc81620002d1565b92915050565b600080fd5b600067ffffffffffffffff82111562000325576200032462000205565b5b6200033082620001f4565b9050602081019050919050565b60005b838110156200035d57808201518184015260208101905062000340565b838111156200036d576000848401525b50505050565b60006200038a620003848462000307565b6200026a565b905082815260208101848484011115620003a957620003a862000302565b5b620003b68482856200033d565b509392505050565b600082601f830112620003d657620003d5620001ef565b5b8151620003e884826020860162000373565b91505092915050565b6000606082840312156200040a5762000409620002bd565b5b6200041660606200026a565b905060006200042884828501620002eb565b600083015250602082015167ffffffffffffffff8111156200044f576200044e620002c2565b5b6200045d84828501620003be565b60208301525060406200047384828501620002eb565b60408301525092915050565b600062000496620004908462000289565b6200026a565b90508083825260208201905060208402830185811115620004bc57620004bb620002b8565b5b835b818110156200050a57805167ffffffffffffffff811115620004e557620004e4620001ef565b5b808601620004f48982620003f1565b85526020850194505050602081019050620004be565b5050509392505050565b600082601f8301126200052c576200052b620001ef565b5b81516200053e8482602086016200047f565b91505092915050565b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b6000620005748262000547565b9050919050565b620005868162000567565b81146200059257600080fd5b50565b600081519050620005a6816200057b565b92915050565b60008060408385031215620005c657620005c5620001e5565b5b600083015167ffffffffffffffff811115620005e757620005e6620001ea565b5b620005f58582860162000514565b9250506020620006088582860162000595565b9150509250929050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b60006200067d82620002c7565b91507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff821415620006b357620006b262000641565b5b600182019050919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052602260045260246000fd5b600060028204905060018216806200070657607f821691505b602082108114156200071d576200071c620006be565b5b50919050565b611abd80620007336000396000f3fe608060405234801561001057600080fd5b50600436106100885760003560e01c8063a00f198a1161005b578063a00f198a146100ef578063a091db2d146100f9578063d36dedd214610115578063fdb9c7931461011f57610088565b8063101b4dd51461008d5780631a4e6df114610097578063212aa138146100b55780637337e3e3146100d3575b600080fd5b61009561013d565b005b61009f610231565b6040516100ac919061153a565b60405180910390f35b6100bd6104db565b6040516100ca919061156b565b60405180910390f35b6100ed60048036038101906100e891906115b7565b610737565b005b6100f76109ba565b005b610113600480360381019061010e91906115f7565b610a51565b005b61011d610bb1565b005b610127610d37565b604051610134919061156b565b60405180910390f35b60006101476104db565b905060006004600083815260200190815260200160002060000154116101a2576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161019990611681565b60405180910390fd5b60016004600083815260200190815260200160002060050160016101000a81548160ff0219169083151502179055503373ffffffffffffffffffffffffffffffffffffffff166108fc60046000848152602001908152602001600020600401549081150290604051600060405180830381858888f1935050505015801561022d573d6000803e3d6000fd5b5050565b610239610e5f565b60006102436104db565b9050600060046000838152602001908152602001600020600001541161029e576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610295906116ed565b60405180910390fd5b600460008281526020019081526020016000206040518060a001604052908160008201548152602001600182016040518060600160405290816000820154815260200160018201805480602002602001604051908101604052809291908181526020016000905b828210156103fb578382906000526020600020906004020160405180604001604052908160008201604051806060016040529081600082015481526020016001820180546103529061173c565b80601f016020809104026020016040519081016040528092919081815260200182805461037e9061173c565b80156103cb5780601f106103a0576101008083540402835291602001916103cb565b820191906000526020600020905b8154815290600101906020018083116103ae57829003601f168201915b50505050508152602001600282015481525050815260200160038201548152505081526020019060010190610305565b5050505081526020016002820180548060200260200160405190810160405280929190818152602001828054801561048857602002820191906000526020600020905b8160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001906001019080831161043e575b5050505050815250508152602001600482015481526020016005820160009054906101000a900460ff161515151581526020016005820160019054906101000a900460ff16151515158152505091505090565b600080600360003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000206040518060600160405290816000820154815260200160018201805480602002602001604051908101604052809291908181526020016000905b8282101561064d578382906000526020600020906004020160405180604001604052908160008201604051806060016040529081600082015481526020016001820180546105a49061173c565b80601f01602080910402602001604051908101604052809291908181526020018280546105d09061173c565b801561061d5780601f106105f25761010080835404028352916020019161061d565b820191906000526020600020905b81548152906001019060200180831161060057829003601f168201915b50505050508152602001600282015481525050815260200160038201548152505081526020019060010190610557565b505050508152602001600282018054806020026020016040519081016040528092919081815260200182805480156106da57602002820191906000526020600020905b8160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019060010190808311610690575b5050505050815250509050600081600001511161072c576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610723906117ba565b60405180910390fd5b806000015191505090565b60006107416104db565b9050600060016000858152602001908152602001600020600001541161079c576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161079390611826565b60405180910390fd5b600082116107df576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016107d690611892565b60405180910390fd5b60026000828152602001908152602001600020600201339080600181540180825580915050600190039060005260206000200160009091909190916101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555060026000828152602001908152602001600020600101604051806040016040528060016000878152602001908152602001600020604051806060016040529081600082015481526020016001820180546108af9061173c565b80601f01602080910402602001604051908101604052809291908181526020018280546108db9061173c565b80156109285780601f106108fd57610100808354040283529160200191610928565b820191906000526020600020905b81548152906001019060200180831161090b57829003601f168201915b50505050508152602001600282015481525050815260200184815250908060018154018082558091505060019003906000526020600020906004020160009091909190915060008201518160000160008201518160000155602082015181600101908051906020019061099c929190610e98565b50604082015181600201555050602082015181600301555050505050565b60006109c46104db565b90506000600460008381526020019081526020016000206000015411610a1f576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610a16906118fe565b60405180910390fd5b60016004600083815260200190815260200160002060050160006101000a81548160ff02191690831515021790555050565b6000600260008381526020019081526020016000206000015411610aaa576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610aa19061196a565b60405180910390fd5b60026000828152602001908152602001600020600201339080600181540180825580915050600190039060005260206000200160009091909190916101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555060026000828152602001908152602001600020600360003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020600082015481600001556001820181600101908054610b93929190610f1e565b506002820181600201908054610baa929190610fbe565b5090505050565b6000610bbb6104db565b90506000600460008381526020019081526020016000206000015414610c16576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610c0d90611681565b60405180910390fd5b8060046000838152602001908152602001600020600001819055506002600082815260200190815260200160002060046000838152602001908152602001600020600101600082015481600001556001820181600101908054610c7a929190610f1e565b506002820181600201908054610c91929190610fbe565b509050506000805b6002600084815260200190815260200160002060010180549050811015610d1757600260008481526020019081526020016000206001018181548110610ce257610ce161198a565b5b90600052602060002090600402016000016002015482610d0291906119e8565b91508080610d0f90611a3e565b915050610c99565b508060046000848152602001908152602001600020600401819055505050565b60006001905080600260008381526020019081526020016000206000018190555060026000828152602001908152602001600020600201339080600181540180825580915050600190039060005260206000200160009091909190916101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555060026000828152602001908152602001600020600360003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020600082015481600001556001820181600101908054610e41929190610f1e565b506002820181600201908054610e58929190610fbe565b5090505090565b6040518060a0016040528060008152602001610e79611010565b8152602001600081526020016000151581526020016000151581525090565b828054610ea49061173c565b90600052602060002090601f016020900481019282610ec65760008555610f0d565b82601f10610edf57805160ff1916838001178555610f0d565b82800160010185558215610f0d579182015b82811115610f0c578251825591602001919060010190610ef1565b5b509050610f1a9190611031565b5090565b828054828255906000526020600020906004028101928215610fad5760005260206000209160040282015b82811115610fac5782826000820181600001600082015481600001556001820181600101908054610f799061173c565b610f8492919061104e565b5060028201548160020155505060038201548160030155505091600401919060040190610f49565b5b509050610fba91906110db565b5090565b828054828255906000526020600020908101928215610fff5760005260206000209182015b82811115610ffe578254825591600101919060010190610fe3565b5b50905061100c9190611031565b5090565b60405180606001604052806000815260200160608152602001606081525090565b5b8082111561104a576000816000905550600101611032565b5090565b82805461105a9061173c565b90600052602060002090601f01602090048101928261107c57600085556110ca565b82601f1061108d57805485556110ca565b828001600101855582156110ca57600052602060002091601f016020900482015b828111156110c95782548255916001019190600101906110ae565b5b5090506110d79190611031565b5090565b5b8082111561111d5760008082016000808201600090556001820160006111029190611121565b600282016000905550506003820160009055506004016110dc565b5090565b50805461112d9061173c565b6000825580601f1061113f575061115e565b601f01602090049060005260206000209081019061115d9190611031565b5b50565b6000819050919050565b61117481611161565b82525050565b600081519050919050565b600082825260208201905092915050565b6000819050602082019050919050565b600081519050919050565b600082825260208201905092915050565b60005b838110156111e05780820151818401526020810190506111c5565b838111156111ef576000848401525b50505050565b6000601f19601f8301169050919050565b6000611211826111a6565b61121b81856111b1565b935061122b8185602086016111c2565b611234816111f5565b840191505092915050565b6000606083016000830151611257600086018261116b565b506020830151848203602086015261126f8282611206565b9150506040830151611284604086018261116b565b508091505092915050565b600060408301600083015184820360008601526112ac828261123f565b91505060208301516112c1602086018261116b565b508091505092915050565b60006112d8838361128f565b905092915050565b6000602082019050919050565b60006112f88261117a565b6113028185611185565b93508360208202850161131485611196565b8060005b85811015611350578484038952815161133185826112cc565b945061133c836112e0565b925060208a01995050600181019050611318565b50829750879550505050505092915050565b600081519050919050565b600082825260208201905092915050565b6000819050602082019050919050565b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b60006113b98261138e565b9050919050565b6113c9816113ae565b82525050565b60006113db83836113c0565b60208301905092915050565b6000602082019050919050565b60006113ff82611362565b611409818561136d565b93506114148361137e565b8060005b8381101561144557815161142c88826113cf565b9750611437836113e7565b925050600181019050611418565b5085935050505092915050565b600060608301600083015161146a600086018261116b565b506020830151848203602086015261148282826112ed565b9150506040830151848203604086015261149c82826113f4565b9150508091505092915050565b60008115159050919050565b6114be816114a9565b82525050565b600060a0830160008301516114dc600086018261116b565b50602083015184820360208601526114f48282611452565b9150506040830151611509604086018261116b565b50606083015161151c60608601826114b5565b50608083015161152f60808601826114b5565b508091505092915050565b6000602082019050818103600083015261155481846114c4565b905092915050565b61156581611161565b82525050565b6000602082019050611580600083018461155c565b92915050565b600080fd5b61159481611161565b811461159f57600080fd5b50565b6000813590506115b18161158b565b92915050565b600080604083850312156115ce576115cd611586565b5b60006115dc858286016115a2565b92505060206115ed858286016115a2565b9150509250929050565b60006020828403121561160d5761160c611586565b5b600061161b848285016115a2565b91505092915050565b600082825260208201905092915050565b7f6f72646572206861642067657474656400000000000000000000000000000000600082015250565b600061166b601083611624565b915061167682611635565b602082019050919050565b6000602082019050818103600083015261169a8161165e565b9050919050565b7f6f726465722069736e6f74206578736974000000000000000000000000000000600082015250565b60006116d7601183611624565b91506116e2826116a1565b602082019050919050565b60006020820190508181036000830152611706816116ca565b9050919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052602260045260246000fd5b6000600282049050600182168061175457607f821691505b602082108114156117685761176761170d565b5b50919050565b7f6361726e756d206973206572726f720000000000000000000000000000000000600082015250565b60006117a4600f83611624565b91506117af8261176e565b602082019050919050565b600060208201905081810360008301526117d381611797565b9050919050565b7f676f6f6473206964206572726f72000000000000000000000000000000000000600082015250565b6000611810600e83611624565b915061181b826117da565b602082019050919050565b6000602082019050818103600083015261183f81611803565b9050919050565b7f206e756d73206d757374206774207a65726f2000000000000000000000000000600082015250565b600061187c601383611624565b915061188782611846565b602082019050919050565b600060208201905081810360008301526118ab8161186f565b9050919050565b7f6f7264657220686173206e6f7420676574746564000000000000000000000000600082015250565b60006118e8601483611624565b91506118f3826118b2565b602082019050919050565b60006020820190508181036000830152611917816118db565b9050919050565b7f636172206e756d73206572726f72000000000000000000000000000000000000600082015250565b6000611954600e83611624565b915061195f8261191e565b602082019050919050565b6000602082019050818103600083015261198381611947565b9050919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b60006119f382611161565b91506119fe83611161565b9250827fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff03821115611a3357611a326119b9565b5b828201905092915050565b6000611a4982611161565b91507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff821415611a7c57611a7b6119b9565b5b60018201905091905056fea26469706673582212200c87630a6537105eb88d6f473f8d03b671a441781b15319ab286cc572a15c86a64736f6c634300080b0033",
}

// DemopayABI is the input ABI used to generate the binding from.
// Deprecated: Use DemopayMetaData.ABI instead.
var DemopayABI = DemopayMetaData.ABI

// DemopayBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use DemopayMetaData.Bin instead.
var DemopayBin = DemopayMetaData.Bin

// DeployDemopay deploys a new Ethereum contract, binding an instance of Demopay to it.
func DeployDemopay(auth *bind.TransactOpts, backend bind.ContractBackend, goods []DemoPayGoods, bossaddress common.Address) (common.Address, *types.Transaction, *Demopay, error) {
	parsed, err := DemopayMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(DemopayBin), backend, goods, bossaddress)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Demopay{DemopayCaller: DemopayCaller{contract: contract}, DemopayTransactor: DemopayTransactor{contract: contract}, DemopayFilterer: DemopayFilterer{contract: contract}}, nil
}

// Demopay is an auto generated Go binding around an Ethereum contract.
type Demopay struct {
	DemopayCaller     // Read-only binding to the contract
	DemopayTransactor // Write-only binding to the contract
	DemopayFilterer   // Log filterer for contract events
}

// DemopayCaller is an auto generated read-only Go binding around an Ethereum contract.
type DemopayCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DemopayTransactor is an auto generated write-only Go binding around an Ethereum contract.
type DemopayTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DemopayFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type DemopayFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DemopaySession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type DemopaySession struct {
	Contract     *Demopay          // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// DemopayCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type DemopayCallerSession struct {
	Contract *DemopayCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts  // Call options to use throughout this session
}

// DemopayTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type DemopayTransactorSession struct {
	Contract     *DemopayTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// DemopayRaw is an auto generated low-level Go binding around an Ethereum contract.
type DemopayRaw struct {
	Contract *Demopay // Generic contract binding to access the raw methods on
}

// DemopayCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type DemopayCallerRaw struct {
	Contract *DemopayCaller // Generic read-only contract binding to access the raw methods on
}

// DemopayTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type DemopayTransactorRaw struct {
	Contract *DemopayTransactor // Generic write-only contract binding to access the raw methods on
}

// NewDemopay creates a new instance of Demopay, bound to a specific deployed contract.
func NewDemopay(address common.Address, backend bind.ContractBackend) (*Demopay, error) {
	contract, err := bindDemopay(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Demopay{DemopayCaller: DemopayCaller{contract: contract}, DemopayTransactor: DemopayTransactor{contract: contract}, DemopayFilterer: DemopayFilterer{contract: contract}}, nil
}

// NewDemopayCaller creates a new read-only instance of Demopay, bound to a specific deployed contract.
func NewDemopayCaller(address common.Address, caller bind.ContractCaller) (*DemopayCaller, error) {
	contract, err := bindDemopay(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &DemopayCaller{contract: contract}, nil
}

// NewDemopayTransactor creates a new write-only instance of Demopay, bound to a specific deployed contract.
func NewDemopayTransactor(address common.Address, transactor bind.ContractTransactor) (*DemopayTransactor, error) {
	contract, err := bindDemopay(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &DemopayTransactor{contract: contract}, nil
}

// NewDemopayFilterer creates a new log filterer instance of Demopay, bound to a specific deployed contract.
func NewDemopayFilterer(address common.Address, filterer bind.ContractFilterer) (*DemopayFilterer, error) {
	contract, err := bindDemopay(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &DemopayFilterer{contract: contract}, nil
}

// bindDemopay binds a generic wrapper to an already deployed contract.
func bindDemopay(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(DemopayABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Demopay *DemopayRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Demopay.Contract.DemopayCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Demopay *DemopayRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Demopay.Contract.DemopayTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Demopay *DemopayRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Demopay.Contract.DemopayTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Demopay *DemopayCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Demopay.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Demopay *DemopayTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Demopay.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Demopay *DemopayTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Demopay.Contract.contract.Transact(opts, method, params...)
}

// GetCarNum is a free data retrieval call binding the contract method 0x212aa138.
//
// Solidity: function getCarNum() view returns(uint256 carnum)
func (_Demopay *DemopayCaller) GetCarNum(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Demopay.contract.Call(opts, &out, "getCarNum")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetCarNum is a free data retrieval call binding the contract method 0x212aa138.
//
// Solidity: function getCarNum() view returns(uint256 carnum)
func (_Demopay *DemopaySession) GetCarNum() (*big.Int, error) {
	return _Demopay.Contract.GetCarNum(&_Demopay.CallOpts)
}

// GetCarNum is a free data retrieval call binding the contract method 0x212aa138.
//
// Solidity: function getCarNum() view returns(uint256 carnum)
func (_Demopay *DemopayCallerSession) GetCarNum() (*big.Int, error) {
	return _Demopay.Contract.GetCarNum(&_Demopay.CallOpts)
}

// ShowOrder is a free data retrieval call binding the contract method 0x1a4e6df1.
//
// Solidity: function showOrder() view returns((uint256,(uint256,((uint256,string,uint256),uint256)[],address[]),uint256,bool,bool) order)
func (_Demopay *DemopayCaller) ShowOrder(opts *bind.CallOpts) (DemoPayOrder, error) {
	var out []interface{}
	err := _Demopay.contract.Call(opts, &out, "showOrder")

	if err != nil {
		return *new(DemoPayOrder), err
	}

	out0 := *abi.ConvertType(out[0], new(DemoPayOrder)).(*DemoPayOrder)

	return out0, err

}

// ShowOrder is a free data retrieval call binding the contract method 0x1a4e6df1.
//
// Solidity: function showOrder() view returns((uint256,(uint256,((uint256,string,uint256),uint256)[],address[]),uint256,bool,bool) order)
func (_Demopay *DemopaySession) ShowOrder() (DemoPayOrder, error) {
	return _Demopay.Contract.ShowOrder(&_Demopay.CallOpts)
}

// ShowOrder is a free data retrieval call binding the contract method 0x1a4e6df1.
//
// Solidity: function showOrder() view returns((uint256,(uint256,((uint256,string,uint256),uint256)[],address[]),uint256,bool,bool) order)
func (_Demopay *DemopayCallerSession) ShowOrder() (DemoPayOrder, error) {
	return _Demopay.Contract.ShowOrder(&_Demopay.CallOpts)
}

// AddGoods is a paid mutator transaction binding the contract method 0x7337e3e3.
//
// Solidity: function addGoods(uint256 gid, uint256 nums) returns()
func (_Demopay *DemopayTransactor) AddGoods(opts *bind.TransactOpts, gid *big.Int, nums *big.Int) (*types.Transaction, error) {
	return _Demopay.contract.Transact(opts, "addGoods", gid, nums)
}

// AddGoods is a paid mutator transaction binding the contract method 0x7337e3e3.
//
// Solidity: function addGoods(uint256 gid, uint256 nums) returns()
func (_Demopay *DemopaySession) AddGoods(gid *big.Int, nums *big.Int) (*types.Transaction, error) {
	return _Demopay.Contract.AddGoods(&_Demopay.TransactOpts, gid, nums)
}

// AddGoods is a paid mutator transaction binding the contract method 0x7337e3e3.
//
// Solidity: function addGoods(uint256 gid, uint256 nums) returns()
func (_Demopay *DemopayTransactorSession) AddGoods(gid *big.Int, nums *big.Int) (*types.Transaction, error) {
	return _Demopay.Contract.AddGoods(&_Demopay.TransactOpts, gid, nums)
}

// ConfirmOrder is a paid mutator transaction binding the contract method 0xa00f198a.
//
// Solidity: function confirmOrder() returns()
func (_Demopay *DemopayTransactor) ConfirmOrder(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Demopay.contract.Transact(opts, "confirmOrder")
}

// ConfirmOrder is a paid mutator transaction binding the contract method 0xa00f198a.
//
// Solidity: function confirmOrder() returns()
func (_Demopay *DemopaySession) ConfirmOrder() (*types.Transaction, error) {
	return _Demopay.Contract.ConfirmOrder(&_Demopay.TransactOpts)
}

// ConfirmOrder is a paid mutator transaction binding the contract method 0xa00f198a.
//
// Solidity: function confirmOrder() returns()
func (_Demopay *DemopayTransactorSession) ConfirmOrder() (*types.Transaction, error) {
	return _Demopay.Contract.ConfirmOrder(&_Demopay.TransactOpts)
}

// GetOrder is a paid mutator transaction binding the contract method 0xd36dedd2.
//
// Solidity: function getOrder() returns()
func (_Demopay *DemopayTransactor) GetOrder(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Demopay.contract.Transact(opts, "getOrder")
}

// GetOrder is a paid mutator transaction binding the contract method 0xd36dedd2.
//
// Solidity: function getOrder() returns()
func (_Demopay *DemopaySession) GetOrder() (*types.Transaction, error) {
	return _Demopay.Contract.GetOrder(&_Demopay.TransactOpts)
}

// GetOrder is a paid mutator transaction binding the contract method 0xd36dedd2.
//
// Solidity: function getOrder() returns()
func (_Demopay *DemopayTransactorSession) GetOrder() (*types.Transaction, error) {
	return _Demopay.Contract.GetOrder(&_Demopay.TransactOpts)
}

// NewCar is a paid mutator transaction binding the contract method 0xfdb9c793.
//
// Solidity: function newCar() returns(uint256 carnum)
func (_Demopay *DemopayTransactor) NewCar(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Demopay.contract.Transact(opts, "newCar")
}

// NewCar is a paid mutator transaction binding the contract method 0xfdb9c793.
//
// Solidity: function newCar() returns(uint256 carnum)
func (_Demopay *DemopaySession) NewCar() (*types.Transaction, error) {
	return _Demopay.Contract.NewCar(&_Demopay.TransactOpts)
}

// NewCar is a paid mutator transaction binding the contract method 0xfdb9c793.
//
// Solidity: function newCar() returns(uint256 carnum)
func (_Demopay *DemopayTransactorSession) NewCar() (*types.Transaction, error) {
	return _Demopay.Contract.NewCar(&_Demopay.TransactOpts)
}

// PayOrder is a paid mutator transaction binding the contract method 0x101b4dd5.
//
// Solidity: function payOrder() returns()
func (_Demopay *DemopayTransactor) PayOrder(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Demopay.contract.Transact(opts, "payOrder")
}

// PayOrder is a paid mutator transaction binding the contract method 0x101b4dd5.
//
// Solidity: function payOrder() returns()
func (_Demopay *DemopaySession) PayOrder() (*types.Transaction, error) {
	return _Demopay.Contract.PayOrder(&_Demopay.TransactOpts)
}

// PayOrder is a paid mutator transaction binding the contract method 0x101b4dd5.
//
// Solidity: function payOrder() returns()
func (_Demopay *DemopayTransactorSession) PayOrder() (*types.Transaction, error) {
	return _Demopay.Contract.PayOrder(&_Demopay.TransactOpts)
}

// TakeCar is a paid mutator transaction binding the contract method 0xa091db2d.
//
// Solidity: function takeCar(uint256 carnum) returns()
func (_Demopay *DemopayTransactor) TakeCar(opts *bind.TransactOpts, carnum *big.Int) (*types.Transaction, error) {
	return _Demopay.contract.Transact(opts, "takeCar", carnum)
}

// TakeCar is a paid mutator transaction binding the contract method 0xa091db2d.
//
// Solidity: function takeCar(uint256 carnum) returns()
func (_Demopay *DemopaySession) TakeCar(carnum *big.Int) (*types.Transaction, error) {
	return _Demopay.Contract.TakeCar(&_Demopay.TransactOpts, carnum)
}

// TakeCar is a paid mutator transaction binding the contract method 0xa091db2d.
//
// Solidity: function takeCar(uint256 carnum) returns()
func (_Demopay *DemopayTransactorSession) TakeCar(carnum *big.Int) (*types.Transaction, error) {
	return _Demopay.Contract.TakeCar(&_Demopay.TransactOpts, carnum)
}
