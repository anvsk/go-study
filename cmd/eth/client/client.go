package client

import (
	"context"
	"crypto/ecdsa"
	"errors"
	"fmt"
	"math"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
)

var localAddr = "http://localhost:8545"
var project = "0x760855f38bb3CE5c4c8c360e846812878665e477"

type EthClient struct {
	Addr    string
	Project string
	Account common.Address
	// Client  *ethclient.Client
	*ethclient.Client
}

func NewClient(addr, project string) (client *EthClient, err error) {
	e := EthClient{
		Addr:    addr,
		Project: project,
	}
	if err = e.Connect(); err != nil {
		println(err)
		return
	}
	e.Account = e.CommonAccount()
	client = &e
	return
}

// 制定地址建立连接、
func (e *EthClient) Connect() (err error) {
	client, err := ethclient.Dial(e.Addr)
	if err != nil {
		return
	}
	e.Client = client
	return
}

// 制定project获取账户
func (e *EthClient) CommonAccount() common.Address {
	return common.HexToAddress(e.Project)
}

// 获取制定项目的余额
func (e *EthClient) GetBanlanAt() (balance *big.Int, err error) {
	balance, err = e.Client.BalanceAt(context.Background(), e.CommonAccount(), nil)
	if err != nil {
		return
	}
	fmt.Println("BanlanAt", balance.String())
	return
}

// 获取eth数量
func (e *EthClient) GetEthNums() (s big.Float, err error) {
	balance, err := e.GetBanlanAt()
	if err != nil {
		return s, err
	}
	fbalance := new(big.Float)
	fbalance.SetString(balance.String())
	ethValue := new(big.Float).Quo(fbalance, big.NewFloat(math.Pow10(18)))
	fmt.Println("ethValue", ethValue)
	return
}

// 转账交易
// 加载私钥获取fromaddr
// 根据公钥获取toaddr
// 计算gas、chainID相关参数
// 获取signTX交易
func (e *EthClient) Transto(fromPrivateECDSA string, toHexAddress string, howmuch int64) (sno string, err error) {
	// fromaddr
	privateKey, err := crypto.HexToECDSA(fromPrivateECDSA)
	if err != nil {
		return
	}
	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		println("fromaddr assert error")
		err = errors.New("cannot assert type: publicKey is not of type *ecdsa.PublicKey")
		return
	}
	fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)

	// toaddr
	toAddress := common.HexToAddress(toHexAddress)

	// 随机数
	nonce, err := e.Client.PendingNonceAt(context.Background(), fromAddress)
	if err != nil {
		println("nonce error")
		return
	}

	// gas 计算
	gasLimit := uint64(21000) // in units
	gasPrice, err := e.Client.SuggestGasPrice(context.Background())
	if err != nil {
		println("gasPrice", gasPrice)
		return
	}

	// tx
	var data []byte
	value := big.NewInt(howmuch * 1000000000000000000) // in wei (1 eth)
	tx := types.NewTransaction(nonce, toAddress, value, gasLimit, gasPrice, data)

	// chainID
	chainID, err := e.Client.NetworkID(context.Background())
	if err != nil {
		println("chainID", chainID)

		return
	}

	// signTX
	signedTx, err := types.SignTx(tx, types.NewEIP155Signer(chainID), privateKey)
	if err != nil {
		println("signedTx", signedTx)

		return
	}

	// DONE
	err = e.Client.SendTransaction(context.Background(), signedTx)
	if err != nil {
		println("signTx error")
		return
	}
	sno = signedTx.Hash().Hex()
	fmt.Printf("tx sent: %s", sno)
	return
}
