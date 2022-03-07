package vote

import (
	"context"
	"crypto/ecdsa"
	"fmt"
	"log"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
)

var eclient *MyClient

func DDeployDemo(ursa string) (address common.Address, tx *types.Transaction, instance *Demo, err error) {
	client, err := ethclient.Dial("http://127.0.0.1:8545")
	// client, err := ethclient.Dial("https://rinkeby.infura.io")
	if err != nil {
		log.Fatal(err)
	}
	eclient = &MyClient{client}
	// privateKey, err := crypto.HexToECDSA(ursa)
	// if err != nil {
	//     log.Fatal(err)
	// }

	// publicKey := privateKey.Public()
	// publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	// if !ok {
	//     log.Fatal("cannot assert type: publicKey is not of type *ecdsa.PublicKey")
	// }

	// fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)
	// nonce, err := client.PendingNonceAt(context.Background(), fromAddress)
	// if err != nil {
	//     log.Fatal(err)
	// }

	// gasPrice, err := client.SuggestGasPrice(context.Background())
	// if err != nil {
	//     log.Fatal(err)
	// }

	// auth := bind.NewKeyedTransactor(privateKey)
	// auth.Nonce = big.NewInt(int64(nonce))
	// auth.Value = big.NewInt(0)     // in wei
	// auth.GasLimit = uint64(0) // in units
	// auth.GasPrice = gasPrice

	adminAuth := eclient.NewTransact(ursa)

	address, tx, instance, err = DeployDemo(adminAuth, client, [][32]byte{
		genOpt("选项1"),
		genOpt("选项2"),
		genOpt("选项3"),
	})
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(address.Hex())   // 0x147B8eb97fD247D06C4006D269c90C1908Fb5D54
	fmt.Println(tx.Hash().Hex()) // 0xdae8ba5444eefdc99f4d45cd0c4f24056cba6a02cefbf78066ef9f4188ff7dc0

	_ = instance
	return
}

func genOpt(s string) (res [32]byte) {
	byt := []byte(s)
	copy(res[:], byt)
	return
}

type MyClient struct {
	*ethclient.Client
}

// type MyClient ethclient.Client

func (client *MyClient) NewTransact(pkey string) (auth *bind.TransactOpts) {
	privateKey, err := crypto.HexToECDSA(pkey)
	if err != nil {
		log.Fatal(err)
	}

	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		log.Fatal("cannot assert type: publicKey is not of type *ecdsa.PublicKey")
	}

	fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)
	nonce, err := client.PendingNonceAt(context.Background(), fromAddress)
	if err != nil {
		log.Fatal(err)
	}

	gasPrice, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		log.Fatal(err)
	}

	auth = bind.NewKeyedTransactor(privateKey)
	auth.Nonce = big.NewInt(int64(nonce))
	auth.Value = big.NewInt(0) // in wei
	auth.GasLimit = uint64(0)  // in units
	auth.GasPrice = gasPrice
	return
}
