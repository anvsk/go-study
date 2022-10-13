package outside

import (
	"onepiece-cron/app/dao"
	"onepiece-cron/app/eth"
	"onepiece-cron/app/model"
	"strings"
	"time"

	"github.com/shopspring/decimal"
	log "github.com/sirupsen/logrus"
	"gorm.io/gorm/clause"
)

/*
	[子账户版本\使用默认数据库连接]
 	统计所有用户资产的接口
	1: 用户token balance 资产
	2: 用户pool 资产
	3: 用户stake 资产
	4: 用户在借贷协议中supply 资产
	上面资产加和 以usd计价
*/
type Overview struct {
}

func (t *Overview) Run() {
	var err error

	v1 := OverviewV1{}
	dataV1, err := v1.GetData()
	if err != nil {
		log.Println("Overview dataV1获取失败error:", err)
		return
	}
	log.Printf("v1Overview:%+v", dataV1)
	day := time.Now().UTC().Format("2006-01-02")
	data := model.OverviewStat{
		Date: time.Now().Unix(),
		Day:  day,
	}
	err = t.Calctotal(&data)
	if err != nil {
		log.Println(err)
		return
	}
	log.Printf("v2Overview:%+v", data)

	sumData := t.CombineData(dataV1, data)

	err = model.GetDB().Clauses(clause.OnConflict{
		UpdateAll: true,
	}).Create(&sumData).Error

	if err != nil {
		log.Println(err)
	} else {
		log.Println("create overview success")
	}
}

func (t *Overview) Calctotal(data *model.OverviewStat) (err error) {
	tokenList, err := dao.GetTokenDao().GetTokenList(map[string]interface{}{})
	if err != nil {
		return
	}
	tokenDecimalList := make(map[string]int64)
	tokenPriceList := make(map[string]decimal.Decimal, 0)
	for _, v := range tokenList {
		tokenDecimalList[v.Address] = int64(v.Decimals)
		tokenPriceList[v.Address] = v.Price
	}

	var balanceAmount decimal.Decimal
	var balanceStratsAmount decimal.Decimal
	var depositAmount decimal.Decimal
	var depositStratsAmount decimal.Decimal
	var stratsAvaxBenqiAmount decimal.Decimal
	var stratsAvaxAavev3Amount decimal.Decimal
	var stratsBtcBenqiAmount decimal.Decimal
	var borrowAmount decimal.Decimal
	var borrowStratsAmount decimal.Decimal
	var stakeAmount decimal.Decimal
	var liquidAmount decimal.Decimal

	_ = borrowStratsAmount

	var balanceList []struct {
		model.AccountBalance
		Strategy string
	}
	tx := model.GetDB()
	tx.
		Select("onepiece_account_balance.*,a.strategy").
		Joins("left join onepiece_account a ON a.id=onepiece_account_balance.user_id").
		Where("quantity>?", 0).
		Find(&balanceList)
	for _, v := range balanceList {
		value := v.Quantity.Mul(tokenPriceList[strings.ToLower(v.TokenAddress)])
		balanceAmount = balanceAmount.Add(value)
	}

	var balanceSubList []struct {
		model.AccountBalanceSub
		Strategy string
		Protocol string
	}
	tx.
		Select("onepiece_account_balance_sub.*,a.strategy,a.protocol").
		Joins("left join onepiece_account_sub a ON a.id=onepiece_account_balance_sub.sub_id").
		Where("quantity>?", 0).
		Find(&balanceSubList)
	for _, v := range balanceSubList {
		value := v.Quantity.Mul(tokenPriceList[strings.ToLower(v.TokenAddress)])
		balanceStratsAmount = balanceStratsAmount.Add(value)
		if v.Strategy == "folding" {
			if v.Protocol == "benqi" {
				stratsAvaxBenqiAmount = stratsAvaxBenqiAmount.Add(value)
			} else {
				stratsAvaxAavev3Amount = stratsAvaxAavev3Amount.Add(value)
			}
		} else {
			stratsBtcBenqiAmount = stratsBtcBenqiAmount.Add(value)
		}
	}

	// lend

	var lendList []struct {
		model.LendStatistic
	}
	tx.
		Where(tx.Where("deposit_amount>?", 0).Or("borrow_amount>?", 0)).
		Find(&lendList)
	for _, v := range lendList {
		tokenDecimal := eth.TokenDecimalBig(v.Token, tokenDecimalList)
		depositAmountTmp := v.DepositAmount.Div(tokenDecimal).Mul(tokenPriceList[strings.ToLower(v.Token)])
		borrowAmountTmp := v.BorrowAmount.Div(tokenDecimal).Mul(tokenPriceList[strings.ToLower(v.Token)])

		depositAmount = depositAmount.Add(depositAmountTmp)
		borrowAmount = borrowAmount.Add(borrowAmountTmp)
	}

	var lendSubList []struct {
		model.LendStatisticSub
		Strategy string
		Protocol string
	}
	tx.
		Select("onepiece_lend_statistic_sub.*,a.protocol").
		Joins("left join onepiece_account_sub a ON a.id=onepiece_lend_statistic_sub.sub_id").
		Where(tx.Where("deposit_amount>?", 0).Or("borrow_amount>?", 0)).
		Find(&lendSubList)
	for _, v := range lendSubList {
		tokenDecimal := eth.TokenDecimalBig(v.Token, tokenDecimalList)
		depositAmountTmp := v.DepositAmount.Div(tokenDecimal).Mul(tokenPriceList[strings.ToLower(v.Token)])
		// borrowAmountTmp := v.BorrowAmount.Div(tokenDecimal).Mul(tokenPriceList[strings.ToLower(v.Token)])

		depositStratsAmount = depositStratsAmount.Add(depositAmountTmp)
		// borrowStratsAmount = borrowStratsAmount.Add(borrowAmountTmp)

		if v.Strategy == "folding" {
			if v.Protocol == "benqi" {
				stratsAvaxBenqiAmount = stratsAvaxBenqiAmount.Add(depositAmountTmp)
			} else {
				stratsAvaxAavev3Amount = stratsAvaxAavev3Amount.Add(depositAmountTmp)
			}
		} else {
			stratsBtcBenqiAmount = stratsBtcBenqiAmount.Add(depositAmountTmp)
		}
	}

	// stake

	var stakeList []model.StakeStatistic
	tx.Where("lp_amount>?", 0).Find(&stakeList)
	for _, v := range stakeList {
		tokenOne := v.AmountOne.Div(eth.TokenDecimalBig(v.TokenOne, tokenDecimalList)).Mul(tokenPriceList[strings.ToLower(v.TokenOne)])
		tokenTwo := v.AmountTwo.Div(eth.TokenDecimalBig(v.TokenTwo, tokenDecimalList)).Mul(tokenPriceList[strings.ToLower(v.TokenTwo)])
		tokenThree := v.AmountThree.Div(eth.TokenDecimalBig(v.TokenThree, tokenDecimalList)).Mul(tokenPriceList[strings.ToLower(v.TokenThree)])
		tokenFour := v.AmountFour.Div(eth.TokenDecimalBig(v.TokenFour, tokenDecimalList)).Mul(tokenPriceList[strings.ToLower(v.TokenFour)])
		tokenFive := v.AmountFive.Div(eth.TokenDecimalBig(v.TokenFive, tokenDecimalList)).Mul(tokenPriceList[strings.ToLower(v.TokenFive)])
		tokenSix := v.AmountSix.Div(eth.TokenDecimalBig(v.TokenSix, tokenDecimalList)).Mul(tokenPriceList[strings.ToLower(v.TokenSix)])
		stakeAmount = stakeAmount.Add(tokenOne).Add(tokenTwo).Add(tokenThree).Add(tokenFour).Add(tokenFive).Add(tokenSix)
	}

	var liquidList []model.LiquidityStatistic
	tx.Where("lp_amount>?", 0).Find(&liquidList)
	for _, v := range liquidList {
		tokenOne := v.AmountOne.Div(eth.TokenDecimalBig(v.TokenOne, tokenDecimalList)).Mul(tokenPriceList[strings.ToLower(v.TokenOne)])
		tokenTwo := v.AmountTwo.Div(eth.TokenDecimalBig(v.TokenTwo, tokenDecimalList)).Mul(tokenPriceList[strings.ToLower(v.TokenTwo)])
		tokenThree := v.AmountThree.Div(eth.TokenDecimalBig(v.TokenThree, tokenDecimalList)).Mul(tokenPriceList[strings.ToLower(v.TokenThree)])
		tokenFour := v.AmountFour.Div(eth.TokenDecimalBig(v.TokenFour, tokenDecimalList)).Mul(tokenPriceList[strings.ToLower(v.TokenFour)])
		tokenFive := v.AmountFive.Div(eth.TokenDecimalBig(v.TokenFive, tokenDecimalList)).Mul(tokenPriceList[strings.ToLower(v.TokenFive)])
		tokenSix := v.AmountSix.Div(eth.TokenDecimalBig(v.TokenSix, tokenDecimalList)).Mul(tokenPriceList[strings.ToLower(v.TokenSix)])
		liquidAmount = liquidAmount.Add(tokenOne).Add(tokenTwo).Add(tokenThree).Add(tokenFour).Add(tokenFive).Add(tokenSix)
	}

	var accounts int64
	tx.Where("deleted", 0).Model(&model.Account{}).Count(&accounts)
	data.Accounts = accounts

	strategies := balanceStratsAmount.Add(depositStratsAmount)
	tvl := balanceAmount.Add(depositAmount).Add(stakeAmount).Add(liquidAmount).Add(strategies)
	data.Balance = balanceAmount
	data.Supply = depositAmount
	data.Borrow = borrowAmount
	data.Liquidity = liquidAmount
	data.Farm = stakeAmount
	data.Strategies = strategies
	data.StratAvaxAavev3 = stratsAvaxAavev3Amount
	data.StratAvaxBenqi = stratsAvaxBenqiAmount
	data.StratBtcBenqi = stratsBtcBenqiAmount
	data.Tvl = tvl
	return
}

func (t *Overview) CombineData(v1, data model.OverviewStat) model.OverviewStat {
	data.Accounts += v1.Accounts
	data.Tvl = v1.Tvl.Add(data.Tvl)
	data.Supply = v1.Supply.Add(data.Supply)
	data.Borrow = v1.Borrow.Add(data.Borrow)
	data.Liquidity = v1.Liquidity.Add(data.Liquidity)
	data.Farm = v1.Farm.Add(data.Farm)
	data.StratAvaxAavev3 = v1.StratAvaxAavev3.Add(data.StratAvaxAavev3)
	data.StratAvaxBenqi = v1.StratAvaxBenqi.Add(data.StratAvaxBenqi)
	data.StratBtcBenqi = v1.StratBtcBenqi.Add(data.StratBtcBenqi)
	data.Strategies = v1.Strategies.Add(data.Strategies)
	data.Balance = v1.Balance.Add(data.Balance)
	return data
}
