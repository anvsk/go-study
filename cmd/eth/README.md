## 区块链相关

#### client：配合 ganache、客户端调用 eth 相关方法 demo

#### sol：智能合约相关 demo

生成 abi 和 bin 文件
abi 定义方法和参数
bin 用于合约部署 deploy
`solc -o ./ --bin --abi --overwrite vote.sol`
生成可以调用相关方法的 go 文件
`abigen --abi Ballot.abi --pkg vote --bin Ballot.bin --out gen_vote.go`

合约：

- vote
  [投票]

  - 它实现了一个投票合约。当然，电子投票的主要问题是如何将投票权分配给正确的人以及如何防止操纵。我们不会在这里解决所有问题，但至少我们将展示如何进行委托投票，以便同时自动计算投票并且完全透明。
  - 这个想法是为每张选票创建一份合同，为每个选项提供一个短名称。然后，担任主席的合约的创建者将分别授予每个地址的投票权。
  - 然后，地址背后的人可以选择自己投票或将投票委托给他们信任的人。
  - 在投票时间结束时，winningProposal() 将返回得票最多的提案。

- auction
  [公开拍卖]
  [升级盲拍]

- demopay
  [消费付账]

  - n 个人消费(比如吃饭)之后使用合约进行支付账单
  - 所有人都可以查看账单
  - 任何一个人都可以发起账单
  - 一个人确认账单即可进行支付
  - 合约中支付账单的资金不用考虑谁来出

  `solc -o ./ --bin --abi --overwrite demopay.sol`

  `abigen --abi DemoPay.abi --pkg demopay --bin DemoPay.bin --out demopay.go`
