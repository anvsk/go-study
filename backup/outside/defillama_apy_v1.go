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

type DefillamaApyV1 struct {
	DBConn *gorm.DB
}

// protocol Tvl apy
func (t *DefillamaApyV1) Run() {
	if err := t.BindConn(); err != nil {
		log.Println(err)
		return
	}

	benqi, aavev3, err := t.Calc()
	if err != nil {
		log.Println(err)
		return
	}

	day := time.Now().UTC().Format("2006-01-02")

	apyBenqi, apyAavev3, err := t.GetMaxApy(day)
	if err != nil {
		log.Println(err)
		return
	}
	log.Println(day, "=====tvl====", benqi, aavev3, "===maxapy===", apyBenqi, apyAavev3)

	dataBenqi := model.DefiLlamaApyStat{
		Strats:   "folding",
		Day:      day,
		Protocol: "benqi",
		Tvl:      benqi,
		Apy:      apyBenqi,
	}
	t.DBConn.Clauses(clause.OnConflict{
		UpdateAll: true,
	}).Create(&dataBenqi)

	dataAavev3 := model.DefiLlamaApyStat{
		Strats:   "folding",
		Day:      day,
		Protocol: "aavev3",
		Tvl:      aavev3,
		Apy:      apyAavev3,
	}
	t.DBConn.Clauses(clause.OnConflict{
		UpdateAll: true,
	}).Create(&dataAavev3)
}

func (t *DefillamaApyV1) GetMaxApy(day string) (benqi, aavev3 decimal.Decimal, err error) {
	var apyInfo model.FoldingApy
	err = t.DBConn.Where("day", day).First(&apyInfo).Error
	if err != nil {
		log.Println("no apy record")
		return
	}

	benqi = apyInfo.Avax
	if apyInfo.Savax.Cmp(apyInfo.Avax) == 1 {
		benqi = apyInfo.Savax
	}
	aavev3 = apyInfo.AvaxAaveV3
	if apyInfo.SavaxAaveV3.Cmp(apyInfo.AvaxAaveV3) == 1 {
		aavev3 = apyInfo.SavaxAaveV3
	}
	return
}

func (t *DefillamaApyV1) GetTvlData() (benqi, aavev3 decimal.Decimal, err error) {
	if err = t.BindConn(); err != nil {
		return
	}
	benqi, aavev3, err = t.Calc()
	if err != nil {
		return
	}
	return
}

func (t *DefillamaApyV1) BindConn() (err error) {
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

func (t *DefillamaApyV1) Calc() (benqi, aavev3 decimal.Decimal, err error) {
	var tokenList = make([]model.TokenList, 0)
	t.DBConn.Find(&tokenList)
	tokenDecimalList := make(map[string]int64)
	tokenPriceList := make(map[string]decimal.Decimal, 0)
	for _, v := range tokenList {
		tokenDecimalList[v.Address] = int64(v.Decimals)
		tokenPriceList[v.Address] = v.Price
	}

	var balanceBenqi []model.AccountBalance
	t.DBConn.Where("user_id in (?)", t.GetUserIDSubQuery().Where("protocol", "benqi")).Where("quantity>?", 0).Find(&balanceBenqi)

	var balanceAavev3 []model.AccountBalance
	t.DBConn.Where("user_id in (?)", t.GetUserIDSubQuery().Where("protocol", "aavev3")).Where("quantity>?", 0).Find(&balanceAavev3)

	var l []model.LendStatistic
	t.DBConn.Where("user_id in (?)", t.GetUserIDSubQuery()).Where("deposit_amount>?", 0).Find(&l)

	var benqiBalance decimal.Decimal
	var aavev3Balance decimal.Decimal
	var benqiLend decimal.Decimal
	var aavev3Lend decimal.Decimal

	for _, v := range balanceBenqi {
		benqiBalance = benqiBalance.Add(v.Quantity.Mul(tokenPriceList[strings.ToLower(v.TokenAddress)]))
	}

	for _, v := range balanceAavev3 {
		aavev3Balance = aavev3Balance.Add(v.Quantity.Mul(tokenPriceList[strings.ToLower(v.TokenAddress)]))
	}

	for _, v := range l {
		tokenDecimal := t.TokenDecimalBig(v.Token, tokenDecimalList)
		singlePrice := v.DepositAmount.Div(tokenDecimal).Mul(tokenPriceList[strings.ToLower(v.Token)])
		if v.Protocol == "benqi" {
			benqiLend = benqiLend.Add(singlePrice)
		} else {
			aavev3Lend = aavev3Lend.Add(singlePrice)
		}
	}
	benqi = benqiBalance.Add(benqiLend)
	aavev3 = aavev3Balance.Add(aavev3Lend)
	return
}

func (t *DefillamaApyV1) GetUserIDSubQuery() *gorm.DB {
	return t.DBConn.Select("user_id").Distinct("user_id").Table("onepiece_account_autofolding")
}

func (*DefillamaApyV1) TokenDecimalBig(address string, tokenDecimal map[string]int64) decimal.Decimal {
	if strings.ToLower(address) == eth.UsdtAddress || strings.ToLower(address) == eth.UsdcAddress {
		return decimal.NewFromInt(int64(1e6))
	} else if _, ok := tokenDecimal[strings.ToLower(address)]; ok {
		return decimal.NewFromInt(10).Pow(decimal.NewFromInt(tokenDecimal[strings.ToLower(address)]))
	} else {
		return decimal.NewFromInt(1)
	}
}
