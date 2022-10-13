package outside

import (
	"onepiece-cron/app/dao"
	"onepiece-cron/app/eth"
	"onepiece-cron/app/model"
	"strings"
	"time"

	"github.com/shopspring/decimal"
	log "github.com/sirupsen/logrus"
)

/*
 	统计所有用户资产的接口
	1: 用户token balance 资产
	2: 用户pool 资产
	3: 用户stake 资产
	4: 用户在借贷协议中supply 资产
	上面资产加和 以usd计价
*/
type Tvl struct{}

func (t *Tvl) Run() {
	v1 := TvlV1{}
	tvlV1, err := v1.GetTvlData()
	if err != nil {
		log.Println("tvl dataV1获取失败error:", err)
		return
	}
	log.Printf("=====v1tvl:%+v", tvlV1)

	tvl := t.Calctotal()
	if tvl.IsZero() {
		return
	}
	day := time.Now().UTC().Format("2006-01-02")
	log.Println("=====v2tvl:%+v", tvl)

	tvlAmount := tvl.Add(tvlV1)
	data := model.Tvl{
		Day: day,
		Tvl: tvlAmount,
	}
	err = model.GetDB().Where(model.Tvl{Day: day}).Assign(model.Tvl{Tvl: tvlAmount}).FirstOrCreate(&data).Error
	if err != nil {
		log.Println(err)
	}
}

func (token *Tvl) Calctotal() decimal.Decimal {
	tokenList, err := dao.GetTokenDao().GetTokenList(map[string]interface{}{})
	if err != nil {
		log.Println(err)
		return decimal.Zero
	}
	tokenDecimalList := make(map[string]int64)
	tokenPriceList := make(map[string]decimal.Decimal, 0)
	for _, v := range tokenList {
		tokenDecimalList[v.Address] = int64(v.Decimals)
		tokenPriceList[v.Address] = v.Price
	}

	var balanceAmount decimal.Decimal
	var depositAmount decimal.Decimal
	var stakeAmount decimal.Decimal
	var liquidAmount decimal.Decimal

	var balanceSubAmount decimal.Decimal
	var depositSubAmount decimal.Decimal

	var balanceList []model.AccountBalance
	model.GetDB().Where("quantity>?", 0).Find(&balanceList)
	for _, v := range balanceList {
		balanceAmount = balanceAmount.Add(v.Quantity.Mul(tokenPriceList[strings.ToLower(v.TokenAddress)]))
	}

	var balanceSubList []model.AccountBalanceSub
	model.GetDB().Where("quantity>?", 0).Find(&balanceSubList)
	for _, v := range balanceSubList {
		balanceSubAmount = balanceSubAmount.Add(v.Quantity.Mul(tokenPriceList[strings.ToLower(v.TokenAddress)]))
	}

	var lendList []model.LendStatistic
	model.GetDB().Where("deposit_amount>?", 0).Find(&lendList)
	for _, v := range lendList {
		tokenDecimal := token.TokenDecimalBig(v.Token, tokenDecimalList)
		singlePrice := v.DepositAmount.Div(tokenDecimal).Mul(tokenPriceList[strings.ToLower(v.Token)])
		depositAmount = depositAmount.Add(singlePrice)
	}

	var lendSubList []model.LendStatisticSub
	model.GetDB().Where("deposit_amount>?", 0).Find(&lendSubList)
	for _, v := range lendSubList {
		tokenDecimal := token.TokenDecimalBig(v.Token, tokenDecimalList)
		singlePrice := v.DepositAmount.Div(tokenDecimal).Mul(tokenPriceList[strings.ToLower(v.Token)])
		depositSubAmount = depositSubAmount.Add(singlePrice)
	}

	var stakeList []model.StakeStatistic
	model.GetDB().Where("lp_amount>?", 0).Find(&stakeList)
	for _, v := range stakeList {
		tokenOne := v.AmountOne.Div(token.TokenDecimalBig(v.TokenOne, tokenDecimalList)).Mul(tokenPriceList[strings.ToLower(v.TokenOne)])
		tokenTwo := v.AmountTwo.Div(token.TokenDecimalBig(v.TokenTwo, tokenDecimalList)).Mul(tokenPriceList[strings.ToLower(v.TokenTwo)])
		tokenThree := v.AmountThree.Div(token.TokenDecimalBig(v.TokenThree, tokenDecimalList)).Mul(tokenPriceList[strings.ToLower(v.TokenThree)])
		tokenFour := v.AmountFour.Div(token.TokenDecimalBig(v.TokenFour, tokenDecimalList)).Mul(tokenPriceList[strings.ToLower(v.TokenFour)])
		tokenFive := v.AmountFive.Div(token.TokenDecimalBig(v.TokenFive, tokenDecimalList)).Mul(tokenPriceList[strings.ToLower(v.TokenFive)])
		tokenSix := v.AmountSix.Div(token.TokenDecimalBig(v.TokenSix, tokenDecimalList)).Mul(tokenPriceList[strings.ToLower(v.TokenSix)])
		stakeAmount = stakeAmount.Add(tokenOne).Add(tokenTwo).Add(tokenThree).Add(tokenFour).Add(tokenFive).Add(tokenSix)
	}

	var liquidList []model.LiquidityStatistic
	model.GetDB().Where("lp_amount>?", 0).Find(&liquidList)
	for _, v := range liquidList {
		tokenOne := v.AmountOne.Div(token.TokenDecimalBig(v.TokenOne, tokenDecimalList)).Mul(tokenPriceList[strings.ToLower(v.TokenOne)])
		tokenTwo := v.AmountTwo.Div(token.TokenDecimalBig(v.TokenTwo, tokenDecimalList)).Mul(tokenPriceList[strings.ToLower(v.TokenTwo)])
		tokenThree := v.AmountThree.Div(token.TokenDecimalBig(v.TokenThree, tokenDecimalList)).Mul(tokenPriceList[strings.ToLower(v.TokenThree)])
		tokenFour := v.AmountFour.Div(token.TokenDecimalBig(v.TokenFour, tokenDecimalList)).Mul(tokenPriceList[strings.ToLower(v.TokenFour)])
		tokenFive := v.AmountFive.Div(token.TokenDecimalBig(v.TokenFive, tokenDecimalList)).Mul(tokenPriceList[strings.ToLower(v.TokenFive)])
		tokenSix := v.AmountSix.Div(token.TokenDecimalBig(v.TokenSix, tokenDecimalList)).Mul(tokenPriceList[strings.ToLower(v.TokenSix)])
		liquidAmount = liquidAmount.Add(tokenOne).Add(tokenTwo).Add(tokenThree).Add(tokenFour).Add(tokenFive).Add(tokenSix)
	}

	return balanceAmount.Add(depositAmount).Add(stakeAmount).Add(liquidAmount).Add(balanceSubAmount).Add(depositSubAmount)
}

func (token *Tvl) TokenDecimalBig(address string, tokenDecimal map[string]int64) decimal.Decimal {
	if strings.ToLower(address) == eth.UsdtAddress || strings.ToLower(address) == eth.UsdcAddress {
		return decimal.NewFromInt(int64(1e6))
	} else if _, ok := tokenDecimal[strings.ToLower(address)]; ok {
		return decimal.NewFromInt(10).Pow(decimal.NewFromInt(tokenDecimal[strings.ToLower(address)]))
	} else {
		return decimal.NewFromInt(1)
	}
}
