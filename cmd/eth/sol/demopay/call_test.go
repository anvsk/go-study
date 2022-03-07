package demopay

import (
	"log"
	"math/big"
	"testing"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
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

	address, tx, instance, err := DDeployDemoPay(u0rsa, u3)
	_ = tx
	_ = instance
	if err != nil {
		panic(err)
	}
	demo, err := NewDemopay(address, eclient)
	if err != nil {
		panic(err)
	}
	vone := big.NewInt(int64(1))
	vtwo := big.NewInt(int64(1))
	carnum := big.NewInt(int64(1))
	// u5创建桌子
	_, err = demo.NewCar(eclient.NewTransact(u5rsa))
	HandleError(err)
	// u4加入
	_, err = demo.TakeCar(eclient.NewTransact(u4rsa), carnum)
	HandleError(err, "takecar")
	_, err = demo.AddGoods(eclient.NewTransact(u5rsa), vtwo, vone)
	HandleError(err)
	_, err = demo.AddGoods(eclient.NewTransact(u4rsa), vone, vone)
	_, err = demo.AddGoods(eclient.NewTransact(u4rsa), vone, vone)
	_, err = demo.AddGoods(eclient.NewTransact(u4rsa), vone, vone)
	_, err = demo.AddGoods(eclient.NewTransact(u4rsa), vone, vone)
	// _, err = demo.AddGoods(eclient.NewTransact(u3rsa), vtwo, vone)
	HandleError(err, "addgoods")

	_, err = demo.GetOrder(eclient.NewTransact(u4rsa))
	HandleError(err, "get")
	_, err = demo.ConfirmOrder(eclient.NewTransact(u4rsa))
	HandleError(err, "confirm")
	_, err = demo.PayOrder(eclient.NewTransact(u5rsa))
	HandleError(err, "pay")
	order, err := demo.ShowOrder(&bind.CallOpts{
		From: common.HexToAddress(u5),
	})
	HandleError(err, "show")

	plog.InfoDump(order.Carnum.Int64(), "Carnum")
	plog.InfoDump(order.Amount.Int64(), "Amount")
	// plog.InfoDump(order.Cardetail.Goodsdetail, "Goodsdetail")
	plog.InfoDump(len(order.Cardetail.Members), "members count")

}

func HandleError(args ...interface{}) {
	if args[0] == nil {
		if len(args) > 1 {
			log.Println(args[1].(string), " PASS")
		}
		return
	}
	var title string
	err := args[0].(error)
	if len(args) > 1 {
		title = args[1].(string)
	}
	if err != nil {
		plog.Error("[ERR]", title, err.Error())
	}
}
