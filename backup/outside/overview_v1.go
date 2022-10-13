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
	"gorm.io/gorm/clause"
)

/*
	[多账户版本\使用v1数据库连接]
 	统计所有用户资产的接口
	1: 用户token balance 资产
	2: 用户pool 资产
	3: 用户stake 资产
	4: 用户在借贷协议中supply 资产
	上面资产加和 以usd计价
*/
type OverviewV1 struct {
	BenqiMap map[int]struct{}
	DBConn   *gorm.DB
}

func (t *OverviewV1) Run() {
	if err := t.BindConn(); err != nil {
		log.Println(err)
		return
	}

	t.GetFoldingBenqiMap()

	day := time.Now().UTC().Format("2006-01-02")
	data := model.OverviewStat{
		Date: time.Now().Unix(),
		Day:  day,
	}
	var err error
	err = t.Calctotal(&data)
	if err != nil {
		log.Println(err)
		return
	}
	err = t.DBConn.Clauses(clause.OnConflict{
		UpdateAll: true,
	}).Create(&data).Error

	if err != nil {
		log.Println(err)
	} else {
		log.Println("create overview success")
	}
}

func (t *OverviewV1) GetData() (data model.OverviewStat, err error) {
	if err = t.BindConn(); err != nil {
		return
	}
	t.GetFoldingBenqiMap()
	err = t.Calctotal(&data)
	if err != nil {
		return
	}

	return
}

func (t *OverviewV1) BindConn() (err error) {
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

func (t *OverviewV1) Calctotal(data *model.OverviewStat) (err error) {
	var tokenList = make([]model.TokenList, 0)
	t.DBConn.Find(&tokenList)

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
	tx := t.DBConn

	tx.
		Select("onepiece_account_balance.*,a.strategy").
		Joins("left join onepiece_account a ON a.id=onepiece_account_balance.user_id").
		Where("quantity>?", 0).
		Find(&balanceList)
	for _, v := range balanceList {
		value := v.Quantity.Mul(tokenPriceList[strings.ToLower(v.TokenAddress)])
		if t.IsStrats(v.Strategy) {
			balanceStratsAmount = balanceStratsAmount.Add(value)
			if v.Strategy == "folding_btc" {
				stratsBtcBenqiAmount = stratsBtcBenqiAmount.Add(value)
			} else {
				if t.IsFoldingBenqi(int(v.UserID)) {
					stratsAvaxBenqiAmount = stratsAvaxBenqiAmount.Add(value)
				} else {
					stratsAvaxAavev3Amount = stratsAvaxAavev3Amount.Add(value)
				}
			}
		} else {
			balanceAmount = balanceAmount.Add(value)
		}
	}

	var lendList []struct {
		model.LendStatistic
		Strategy string
		// Protocol string
	}
	tx.
		Select("onepiece_lend_statistic.*,a.strategy").
		Joins("left join onepiece_account a ON a.id=onepiece_lend_statistic.user_id").
		Where(tx.Where("deposit_amount>?", 0).Or("borrow_amount>?", 0)).
		Find(&lendList)
	for _, v := range lendList {
		tokenDecimal := eth.TokenDecimalBig(v.Token, tokenDecimalList)
		depositAmountTmp := v.DepositAmount.Div(tokenDecimal).Mul(tokenPriceList[strings.ToLower(v.Token)])
		borrowAmountTmp := v.BorrowAmount.Div(tokenDecimal).Mul(tokenPriceList[strings.ToLower(v.Token)])
		if t.IsStrats(v.Strategy) {
			depositStratsAmount = depositStratsAmount.Add(depositAmountTmp)
			// borrowStratsAmount = borrowStratsAmount.Add(borrowAmountTmp)
			if v.Strategy == "folding_btc" {
				stratsBtcBenqiAmount = stratsBtcBenqiAmount.Add(depositAmountTmp)
			} else {
				if t.IsFoldingBenqi(int(v.UserID)) {
					stratsAvaxBenqiAmount = stratsAvaxBenqiAmount.Add(depositAmountTmp)
				} else {
					stratsAvaxAavev3Amount = stratsAvaxAavev3Amount.Add(depositAmountTmp)
				}
			}
		} else {
			depositAmount = depositAmount.Add(depositAmountTmp)
			borrowAmount = borrowAmount.Add(borrowAmountTmp)
		}
	}

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

func (token *OverviewV1) IsStrats(s string) bool {
	if s == "folding" || s == "folding_btc" {
		return true
	}
	return false
}

func (t *OverviewV1) GetFoldingBenqiMap() {
	if t.BenqiMap == nil {
		t.BenqiMap = make(map[int]struct{}, 0)
	}
	var b []model.AccountAutoFolding
	t.DBConn.Where("protocol", "benqi").Group("user_id").Find(&b)
	for _, v := range b {
		t.BenqiMap[int(v.UserID)] = struct{}{}
	}
}

// 查询folding表里事上面协议
func (t *OverviewV1) IsFoldingBenqi(uid int) bool {
	_, ok := t.BenqiMap[uid]
	return ok
}
