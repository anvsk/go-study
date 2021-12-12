package client

import (
	"context"
	"fmt"
	"log"
	"testing"
)

// 建立连接，打印相关信息
func TestNewFunc(*testing.T) {
	client, err := NewClient(localAddr, project)
	if err != nil {
		fmt.Println(err)
		return
	}
	client.GetBanlanAt()
	client.GetEthNums()
	// 头块编号
	header, err := client.HeaderByNumber(context.Background(), nil)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(header.Number.String()) // 5671744

	blockNumber := header.Number
	block, err := client.BlockByNumber(context.Background(), blockNumber)
	if err != nil {
		log.Fatal(err)
	}
	println("区块信息")
	fmt.Println(block.Number().Uint64())     // 5671744
	fmt.Println(block.Time())                // 1527211625
	fmt.Println(block.Difficulty().Uint64()) // 3217000136609065
	fmt.Println(block.Hash().Hex())          // 0x9e8751ebb5069389b855bba72d94902cc385042661498a415979b7b6ee9ba4b9
	fmt.Println(len(block.Transactions()))

	for _, tx := range block.Transactions() {
		fmt.Println(tx.Hash().Hex())        // 0x5d49fcaa394c97ec8a9c3e7bd9e8388d420fb050a52083ca52ff24b3b65bc9c2
		fmt.Println(tx.Value().String())    // 10000000000000000
		fmt.Println(tx.Gas())               // 105000
		fmt.Println(tx.GasPrice().Uint64()) // 102000000000
		fmt.Println(tx.Nonce())             // 110644
		fmt.Println(tx.Data())              // []
		fmt.Println(tx.To().Hex())          // 0x55fE59D8Ad77035154dDd0AD0388D09Dd4047A8e

		chainID, err := client.NetworkID(context.Background())
		if err != nil {
			log.Fatal(err)
		}

		// if msg, err := tx.AsMessage(types.NewEIP155Signer(chainID)); err == nil {
		// 	fmt.Println(msg.From().Hex()) // 0x0fD081e3Bb178dc45c0cb23202069ddA57064258
		// }
		println("chaninID", chainID)

		receipt, err := client.TransactionReceipt(context.Background(), tx.Hash())
		if err != nil {
			log.Fatal(err)
		}

		fmt.Println(receipt.Status) // 1
	}
}

// 转账交易
// 加载私钥获取fromaddr
// 根据公钥获取toaddr
// 计算gas、chainID相关参数
// 获取signTX交易
func TestJiaoYi(*testing.T) {
	client, err := NewClient(localAddr, "0x822b3d337F53BE9f0E6A87fb3D7de913E806F1e4")
	if err != nil {
		fmt.Println(err)
		return
	}
	client.GetBanlanAt()
	client.GetEthNums()
	// if sno, err := client.Transto("201757eded6e4ee1e7067b800a5e7bb953dce9e1c82eefddf3754d24bb1690ca", "0x822b3d337F53BE9f0E6A87fb3D7de913E806F1e4", 5); err != nil {
	// 	log.Fatal(err)
	// } else {
	// 	println(sno)
	// }
}
