package util

import (
    "time"

    log "github.com/pieterclaerhout/go-log"
)

type TestStruct struct {
    A   int
    B   int
}

// ====[ LOG ]====================================================================
func initLog() {
    tz, _ := time.LoadLocation("Asia/Shanghai")
    log.TimeZone = tz
    log.DebugMode = Config.Debug
    log.DebugSQLMode = true
    log.PrintTimestamp = true
    log.PrintColors = true
    log.TimeFormat = "15:04:05.000" // DefaultTimeFormat = "2006-01-02 15:04:05.000"
}

// ====[ CompanyWeiXin ]====================================================================

type WxRobotContent struct {
    Content string `json:"content" `
}

type WxRobotMsg struct {
    Msgtype  string         `json:"msgtype"`
    Markdown WxRobotContent `json:"markdown"`
}

func Note(s string) {
    if Config.Debug {
        return
    }
    SendWX(s)
}

func SendWX(s string) {
    wc := WxRobotMsg{
        Msgtype: "markdown",
        Markdown: WxRobotContent{
            Content: s,
        },
    }
    Post(Req{
        Url:    Config.CompanyWx,
        Params: wc,
        Res:    nil,
        Headers: map[string]string{
            "Content-Type": "application/json",
        },
    })
}
