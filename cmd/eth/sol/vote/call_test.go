package vote

import (
	"log"
	"math/big"
	"testing"

	"github.com/ethereum/go-ethereum/common"
	plog "github.com/pieterclaerhout/go-log"
)

var u0 = "0xE280029a7867BA5C9154434886c241775ea87e53"
var u5 = "0xcb98ce2619f90f54052524Fb79b03E0261b01BEE"
var u5rsa = "3b24a4fdf2e6e1375008c387c5456ce00cb0772435ae1938c2fe833103393b9a"
var u4 = "0xb53cC19aD713e00cB71542d064215784c908D387"
var u4rsa = "3af93668029f95d526fc1d2bdefccc120bfe1d26a0462d268e8f6b2f71402ba3"
var u3 = "0x26d8094A90AD1A4440600fa18073730Ef17F5eCE"
var u3rsa = "95ce6122165d94aa51b0fcf51021895b39b0ff291aa640c803d5401bd87894d5"

var u0rsa = "f1b3f8e0d52caec13491368449ab8d90f3d222a3e485aa7f02591bbceb5efba5"

func TestDemo(*testing.T) {

	address, tx, instance, err := DDeployDemo(u0rsa)
	_ = tx
	_ = instance
	if err != nil {
		panic(err)
	}
	demo, err := NewDemo(address, eclient)
	if err != nil {
		panic(err)
	}
	// 开始授权
	plog.Warn("授权前")
	demo.tinfo(u5)
	demo.giveright(u5)
	demo.giveright(u4)
	demo.giveright(u3)
	plog.Warn("授权后")
	demo.tinfo(u5)
	demo.tinfo(u4)
	demo.tinfo(u3)
	demo.windesc()

	// 不授权的投票看看
	// demo.voting(u3rsa, 1)

	// 投票看看
	plog.Warn("u4投票")
	demo.voting(u4rsa, 1)
	demo.windesc()

	// 委托给xx
	plog.Warn("u3委托给u5")
	demo.deleg(u3rsa, u5)
	demo.tinfo(u5)
	demo.tinfo(u4)
	demo.tinfo(u3)
	demo.windesc()

	// 被委托人投票后
	plog.Warn("u5投票")
	demo.voting(u5rsa, 1)
	demo.tinfo(u5)
	demo.tinfo(u4)
	demo.tinfo(u3)
	demo.windesc()

	if maxIndex, err := demo.WinningProposal(nil); err != nil {
		panic(err)
	} else {
		plog.Info("票数最多的选项index:", maxIndex)
		demo.windesc((int(maxIndex.Int64())))
	}
}

// 用户投票情况
func (d *Demo) tinfo(addr string) {
	voters, err := d.Voters(nil, common.HexToAddress(addr))
	if err != nil {
		panic(err)
	}
	log.Println(voters)
}

// 增加投票权限
func (d *Demo) giveright(addr string) {
	// 只能主席授权
	authU0 := eclient.NewTransact(u0rsa)
	tx, err := d.GiveRightToVote(authU0, common.HexToAddress(addr))
	if err != nil {
		panic(err)
	}
	log.Println("GiveRightToVote Tx:", tx)
}

// 获取选项得票情况
func (d *Demo) windesc(args ...int) {
	if len(args) > 0 {
		sindex := args[0]
		Proposals, err := d.Proposals(nil, big.NewInt(int64(sindex)))
		if err != nil {
			log.Println(sindex, "]选项查询错误", err.Error())
		}
		log.Println(sindex, "]==", string(Proposals.Name[:]), Proposals.VoteCount)
	} else {
		// 获取选项
		for i := 0; i < 3; i++ {

			Proposals, err := d.Proposals(nil, big.NewInt(int64(i)))
			if err != nil {
				log.Println(i, "]选项查询错误", err.Error())
			}
			log.Println(i, "]==", string(Proposals.Name[:]), Proposals.VoteCount)
		}
	}
}

// 给选项投票
func (d *Demo) voting(fromkey string, index int) {
	tx, err := d.Vote(eclient.NewTransact(fromkey), big.NewInt(int64(index)))
	if err != nil {
		plog.Info("投票失败", err.Error(), fromkey)
		panic("")
	}
	log.Println("voting tx", tx)
}

// delegate 委托给to
func (d *Demo) deleg(fromkey string, to string) {
	tx, err := d.Delegate(eclient.NewTransact(fromkey), common.HexToAddress(to))
	if err != nil {
		panic(err)
	}
	log.Println("deleg tx", tx)
}
