package outside

import (
	"errors"
	"onepiece-cron/app/eth"
	"onepiece-cron/app/model"
	"onepiece-cron/utils"
	"strings"
	"time"

	"github.com/shopspring/decimal"
	log "github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

/*
 	统计所有用户资产的接口
	1: 用户token balance 资产
	2: 用户pool 资产
	3: 用户stake 资产
	4: 用户在借贷协议中supply 资产
	上面资产加和 以usd计价
*/
type TvlV1 struct {
	DBConn *gorm.DB
}

func (t *TvlV1) Run() {
	if err := t.BindConn(); err != nil {
		log.Println(err)
		return
	}

	tvl := t.Calctotal()
	if tvl.IsZero() {
		return
	}
	day := time.Now().UTC().Format("2006-01-02")
	log.Println(day, "=====tvl====", tvl)
	data := model.Tvl{
		Day: day,
		Tvl: tvl,
	}
	err := t.DBConn.Where(model.Tvl{Day: day}).Assign(model.Tvl{Tvl: tvl}).FirstOrCreate(&data).Error
	if err != nil {
		log.Println(err)
	}
}

func (t *TvlV1) BindConn() (err error) {
	version := utils.GetConfigString("version", "v1")
	if version == "v2" {
		tx := model.GetDBV1()
		if tx == nil {
			return errors.New("mysqlv1 connect failed!")
		}
		t.DBConn = tx
	} else {
		t.DBConn = model.GetDB()
	}
	return nil
}

func (t *TvlV1) GetTvlData() (tvl decimal.Decimal, err error) {
	if err = t.BindConn(); err != nil {
		return
	}
	tvl = t.Calctotal()

	return
}

func (t *TvlV1) Calctotal() decimal.Decimal {
	var tokenList = make([]model.TokenList, 0)
	t.DBConn.Find(&tokenList)
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

	var balanceList []model.AccountBalance
	t.DBConn.Where("quantity>?", 0).Find(&balanceList)
	for _, v := range balanceList {
		balanceAmount = balanceAmount.Add(v.Quantity.Mul(tokenPriceList[strings.ToLower(v.TokenAddress)]))
	}

	var lendList []model.LendStatistic
	t.DBConn.Where("deposit_amount>?", 0).Find(&lendList)
	for _, v := range lendList {
		tokenDecimal := t.TokenDecimalBig(v.Token, tokenDecimalList)
		singlePrice := v.DepositAmount.Div(tokenDecimal).Mul(tokenPriceList[strings.ToLower(v.Token)])
		depositAmount = depositAmount.Add(singlePrice)
	}

	var stakeList []model.StakeStatistic
	t.DBConn.Where("lp_amount>?", 0).Find(&stakeList)
	for _, v := range stakeList {
		tokenOne := v.AmountOne.Div(t.TokenDecimalBig(v.TokenOne, tokenDecimalList)).Mul(tokenPriceList[strings.ToLower(v.TokenOne)])
		tokenTwo := v.AmountTwo.Div(t.TokenDecimalBig(v.TokenTwo, tokenDecimalList)).Mul(tokenPriceList[strings.ToLower(v.TokenTwo)])
		tokenThree := v.AmountThree.Div(t.TokenDecimalBig(v.TokenThree, tokenDecimalList)).Mul(tokenPriceList[strings.ToLower(v.TokenThree)])
		tokenFour := v.AmountFour.Div(t.TokenDecimalBig(v.TokenFour, tokenDecimalList)).Mul(tokenPriceList[strings.ToLower(v.TokenFour)])
		tokenFive := v.AmountFive.Div(t.TokenDecimalBig(v.TokenFive, tokenDecimalList)).Mul(tokenPriceList[strings.ToLower(v.TokenFive)])
		tokenSix := v.AmountSix.Div(t.TokenDecimalBig(v.TokenSix, tokenDecimalList)).Mul(tokenPriceList[strings.ToLower(v.TokenSix)])
		stakeAmount = stakeAmount.Add(tokenOne).Add(tokenTwo).Add(tokenThree).Add(tokenFour).Add(tokenFive).Add(tokenSix)
	}

	var liquidList []model.LiquidityStatistic
	t.DBConn.Where("lp_amount>?", 0).Find(&liquidList)
	for _, v := range liquidList {
		tokenOne := v.AmountOne.Div(t.TokenDecimalBig(v.TokenOne, tokenDecimalList)).Mul(tokenPriceList[strings.ToLower(v.TokenOne)])
		tokenTwo := v.AmountTwo.Div(t.TokenDecimalBig(v.TokenTwo, tokenDecimalList)).Mul(tokenPriceList[strings.ToLower(v.TokenTwo)])
		tokenThree := v.AmountThree.Div(t.TokenDecimalBig(v.TokenThree, tokenDecimalList)).Mul(tokenPriceList[strings.ToLower(v.TokenThree)])
		tokenFour := v.AmountFour.Div(t.TokenDecimalBig(v.TokenFour, tokenDecimalList)).Mul(tokenPriceList[strings.ToLower(v.TokenFour)])
		tokenFive := v.AmountFive.Div(t.TokenDecimalBig(v.TokenFive, tokenDecimalList)).Mul(tokenPriceList[strings.ToLower(v.TokenFive)])
		tokenSix := v.AmountSix.Div(t.TokenDecimalBig(v.TokenSix, tokenDecimalList)).Mul(tokenPriceList[strings.ToLower(v.TokenSix)])
		liquidAmount = liquidAmount.Add(tokenOne).Add(tokenTwo).Add(tokenThree).Add(tokenFour).Add(tokenFive).Add(tokenSix)
	}

	return balanceAmount.Add(depositAmount).Add(stakeAmount).Add(liquidAmount)
}

func (token *TvlV1) TokenDecimalBig(address string, tokenDecimal map[string]int64) decimal.Decimal {
	if strings.ToLower(address) == eth.UsdtAddress || strings.ToLower(address) == eth.UsdcAddress {
		return decimal.NewFromInt(int64(1e6))
	} else if _, ok := tokenDecimal[strings.ToLower(address)]; ok {
		return decimal.NewFromInt(10).Pow(decimal.NewFromInt(tokenDecimal[strings.ToLower(address)]))
	} else {
		return decimal.NewFromInt(1)
	}
}
