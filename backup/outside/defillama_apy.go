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

type DefillamaApy struct{}

// protocol Tvl apy
func (t *DefillamaApy) Run() {
	v1 := DefillamaApyV1{}
	benqiV1, aavev3V1, err := v1.GetTvlData()
	if err != nil {
		log.Println("DefillamaApy dataV1获取失败error:", err)
		return
	}
	log.Printf("=====v1DefillamaApy:%+v", benqiV1, "=====", aavev3V1)

	benqi, aavev3, err := t.Calc()
	if err != nil {
		log.Println(err)
		return
	}
	log.Printf("======v2DefillamaApy:%+v", benqi, "=====", aavev3)

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
		Tvl:      benqi.Add(benqiV1),
		Apy:      apyBenqi,
	}
	model.GetDB().Clauses(clause.OnConflict{
		UpdateAll: true,
	}).Create(&dataBenqi)

	dataAavev3 := model.DefiLlamaApyStat{
		Strats:   "folding",
		Day:      day,
		Protocol: "aavev3",
		Tvl:      aavev3.Add(aavev3V1),
		Apy:      apyAavev3,
	}
	model.GetDB().Clauses(clause.OnConflict{
		UpdateAll: true,
	}).Create(&dataAavev3)
}

func (t *DefillamaApy) GetMaxApy(day string) (benqi, aavev3 decimal.Decimal, err error) {
	var apyInfo model.FoldingApy
	err = model.GetDB().Where("day", day).First(&apyInfo).Error
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

func (t *DefillamaApy) Calc() (benqi, aavev3 decimal.Decimal, err error) {
	tokenList, err := dao.GetTokenDao().GetTokenList(map[string]interface{}{})
	if err != nil {
		log.Println(err)
		return
	}
	tokenDecimalList := make(map[string]int64)
	tokenPriceList := make(map[string]decimal.Decimal, 0)
	for _, v := range tokenList {
		tokenDecimalList[v.Address] = int64(v.Decimals)
		tokenPriceList[v.Address] = v.Price
	}

	db := model.GetDB()

	var balanceBenqi []model.AccountBalanceSub
	db.Where("protocol", "benqi").Where("quantity>?", 0).Find(&balanceBenqi)

	var balanceAavev3 []model.AccountBalanceSub
	db.Where("protocol", "aavev3").Where("quantity>?", 0).Find(&balanceAavev3)

	var l []model.LendStatisticSub
	db.Where("deposit_amount>?", 0).Find(&l)

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

func (*DefillamaApy) TokenDecimalBig(address string, tokenDecimal map[string]int64) decimal.Decimal {
	if strings.ToLower(address) == eth.UsdtAddress || strings.ToLower(address) == eth.UsdcAddress {
		return decimal.NewFromInt(int64(1e6))
	} else if _, ok := tokenDecimal[strings.ToLower(address)]; ok {
		return decimal.NewFromInt(10).Pow(decimal.NewFromInt(tokenDecimal[strings.ToLower(address)]))
	} else {
		return decimal.NewFromInt(1)
	}
}
