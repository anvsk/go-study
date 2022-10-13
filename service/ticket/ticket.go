package ticket

import (
	"errors"
	"fmt"
	"go-study/pkg/util"
	"sync"

	log "github.com/pieterclaerhout/go-log"
)

const pre = "https://www.ssky123.com/api/v2/"

var token string
var passengers []OrderPassenger
var uid int
var loc *sync.RWMutex
var wg *sync.WaitGroup

func Bootstrap() {
	loc = new(sync.RWMutex)
	wg = &sync.WaitGroup{}
	// load flag
	initFlagArgs()
	// refreshtoken
	tickFreshToken()
	// checkToken()
	// passengers
	getPassengers()
	// 买票
	// for range time.Tick(2000 * time.Millisecond) {
	handle(list())
	// }
	wg.Wait()
}

type listArgs struct {
	F int    `json:"startPortNo"`
	T int    `json:"endPortNo"`
	D string `json:"startDate"`
}

func list() *[]ListData {
	res := ListRes{}
	if Config.Customization.Date == "" {
		Config.Customization.Date = util.WeekDate(6)
	}
	args := listArgs{
		Config.PortNo[Config.Customization.From],
		Config.PortNo[Config.Customization.To],
		Config.Customization.Date,
	}
	// 登录
	err := util.Post(util.Req{
		Url:     fmt.Sprintf("%sline/ship/enq", pre),
		Params:  args,
		Res:     &res,
		Headers: getHeader(),
		Check:   ResCheck,
	})
	if err != nil {
		return nil
	}

	return &res.Data
}

func handle(lists *[]ListData) {
	if len(*lists) == 0 {
		log.Debug("search res empty!")
	}
	for _, v := range *lists {
		if checkLine(&v) == false {
			continue
		}
		log.Debug("LINE")
		log.Warnf("[Line]%s [Bus]%s [Ship]%s", fmt.Sprintf("%d%d", v.LineNum, v.Sx), v.BusStartTime, v.SailTime)
		for _, vv := range v.SeatClasses {
			if flag := checkSeat(&vv); flag == false {
				continue
			}
			log.Debug(" ")
			log.Infof("[%s] nums:%s price:%s", vv.ClassName, vv.PubCurrentCount, vv.TotalPrice)
			wg.Add(1)
			go order(&v, &vv)
		}
	}
}

func order(ll *ListData, ss *SeatClasses) (err error) {
	defer wg.Done()
	siglePrice := int(ss.TotalPrice)
	orderItemRequests := []OrderItemRequests{}
	for _, v := range passengers {
		tmp := OrderItemRequests{
			OrderPassenger: v,
			SeatClassName:  ss.ClassName,
			SeatClass:      ss.ClassNum,
			TicketFee:      siglePrice,
			RealFee:        siglePrice,
			FreeChildCount: 0,
		}
		orderItemRequests = append(orderItemRequests, tmp)
	}
	orderArgs := Order{
		AccountTypeID:     "0",
		BuyTicketType:     1, //2:摆渡车
		UserID:            uid,
		ContactNum:        Config.Mobile,
		ListData:          *ll,
		OrderItemRequests: orderItemRequests,
		TotalFee:          siglePrice * len(passengers),
		TotalPayFee:       siglePrice * len(passengers),
	}
	orderArgs.SailDate = Config.Date
	res := OrderRes{}
	if Config.Debug {
		err = errors.New("debug!!!")
	} else {
		if err = util.Post(util.Req{
			Url:     fmt.Sprintf("%sholding/save", pre),
			Params:  orderArgs,
			Res:     &res,
			Headers: getHeader(),
			Check:   ResCheck,
			Async:   true,
		}); err != nil {
			return
		}
	}
	if res.Data.OrderID != "" {
		util.Note(" Success ")
	} else {
		err = errors.New(res.Message)
		return
	}
	return nil
}
