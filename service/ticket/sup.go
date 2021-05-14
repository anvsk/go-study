// 辅助方法
package ticket

import (
    "encoding/json"
    "errors"
    "fmt"
    "go-ticket/pkg/util"
    "time"
)

func checkLine(ll *ListData) bool {
    if Config.LineNum != "" && fmt.Sprintf("%d%d", ll.LineNum, ll.Sx) != Config.LineNum {
        return false
    }
    if Config.MinShipTime != "" {
        if flag, err := util.CompareTime(Config.MinShipTime, ll.SailTime); flag == false && err == nil {
            return false
        }
    }
    if Config.LatestShipTime != "" {
        if flag, err := util.CompareTime(ll.SailTime, Config.LatestShipTime); flag == false && err == nil {
            return false
        }
    }
    return true
}

func checkSeat(ss *SeatClasses) bool {
    if Config.Class != "" && fmt.Sprintf("%s", ss.ClassName) != Config.Class {
        return false
    }
    if ss.PubCurrentCount < 1 {
        return false
    }
    return true
}

func checkToken() error {
    res := CommonRes{}
    err := util.Get(util.Req{
        Url:     fmt.Sprintf("%suser/tokenCheck", pre),
        Res:     &res,
        Headers: getHeader(),
        Check:   ResCheck,
    })
    if err != nil {
        return err
    }
    return nil
}

func getHeader() map[string]string {
    return map[string]string{
        "authentication": Config.User.Authentication,
        "token":          getToken(),
        "Content-Type":   "application/json",
    }
}

func ResCheck(body []byte) error {
    tmp := CommonRes{}
    json.Unmarshal(body, &tmp)

    switch tmp.Code {
    case 500:
        if tmp.Message == "用户未登陆" {
            freshToken()
            return errors.New("token expired")
        }
    case 300:
        return errors.New(tmp.Message)
    }
    return nil
}

func tickFreshToken() {
    freshToken()
    go func() {
        for range time.Tick(60 * time.Second) {
            freshToken()
        }
    }()
}

func freshToken() {
    loc.Lock()
    res := LoginRes{}
    util.Post(util.Req{
        Url:   fmt.Sprintf("%suser/passLogin?phoneNum=%s&passwd=%s&deviceType=2", pre, Config.Mobile, Config.Password),
        Res:   &res,
        Check: ResCheck,
    })

    uid = res.Data.UserID
    token = res.Data.Token
    loc.Unlock()
}

func getToken() string {
    loc.RLock()
    tmp := token
    loc.RUnlock()
    return tmp
}

func getPassengers() error {
    res := PassengerRes{}
    err := util.Get(util.Req{
        Url:     fmt.Sprintf("%suser/passenger/list", pre),
        Res:     &res,
        Headers: getHeader(),
        Check:   ResCheck,
        Params: map[string][]string{
            "aa": {"aa", "aaa"},
            "bb": {"bb", "bbbb"},
            "cc": {"cc"},
        },
    })
    if err != nil {
        return err
    }
    flag := false
    if len(Config.User.Passengers) == 0 {
        flag = true
    }
    for _, v := range res.Data {
        if !flag {
            if exist, _ := util.InArray(v.PassName, Config.User.Passengers); exist {
                flag = true
            }
        }
        if flag {
            passengers = append(passengers, OrderPassenger{
                PassName:       v.PassName,
                PassID:         v.ID,
                CredentialType: v.CredentialType,
            })
        }
    }
    return nil
}
