package main

import (
    "bytes"
    "fmt"
    "go-ticket/pkg/util"
    "go-ticket/service/ticket"
    "log"
    "os/exec"
    "testing"

    "github.com/robfig/cron/v3"
)

/**************************

实现类似linux 的crontab 更精确的控制：秒级

Field name   | Mandatory? | Allowed values  | Allowed special characters
----------   | ---------- | --------------  | --------------------------
Seconds      | Yes        | 0-59            | * / , -
Minutes      | Yes        | 0-59            | * / , -
Hours        | Yes        | 0-23            | * / , -
Day of month | Yes        | 1-31            | * / , - ?
Month        | Yes        | 1-12 or JAN-DEC | * / , -
Day of week  | Yes        | 0-6 or SUN-SAT  | * / , - ?

Entry                  | Description                                | Equivalent To
-----                  | -----------                                | -------------
@yearly (or @annually) | Run once a year, midnight, Jan. 1st        | 0 0 0 1 1 *
@monthly               | Run once a month, midnight, first of month | 0 0 0 1 * *
@weekly                | Run once a week, midnight between Sat/Sun  | 0 0 0 * * 0
@daily (or @midnight)  | Run once a day, midnight                   | 0 0 0 * * *
@hourly                | Run once an hour, beginning of hour        | 0 0 * * * *

**************************/

func main() {
    runTest()
    // exec_shell("pwd")
    // exec_shell("./tmp/ticket")

}

func runTest() {
    c := cron.New(cron.WithSeconds(), cron.WithChain(
        cron.Recover(cron.DefaultLogger), // or use cron.DefaultLogger
    ))
    c.Start()
    defer c.Stop()
    c.AddFunc("*/1 * * * * *", func() {
        // fmt.Println(time.Now().Clock())
    })
    c.AddFunc("@every 5s", func() { fmt.Println("@every 5s") })
    c.AddFunc("@every 1m", func() { panic("panic") })
    c.AddFunc("0 0 8 * * *", func() {
        util.InitUtil()
        ticket.Bootstrap()
    })
    for {
    }
}

func TestJobPanicRecovery(t *testing.T) {
    // var job cron.DummyJob

    // var buf syncWriter
    // cron := cron.New(cron.WithParser(cron.secondParser),
    //     cron.WithChain(cron.Recover(newBufLogger(&buf))))
    // cron.Start()
    // defer cron.Stop()
    // cron.AddJob("* * * * * ?", job)

    // select {
    // case <-time.After(OneSecond):
    //     if !strings.Contains(buf.String(), "YOLO") {
    //         t.Error("expected a panic to be logged, got none")
    //     }
    //     return
    // }
}

func exec_shell(s string) {
    cmd := exec.Command("/bin/bash", "-c", s)
    var out bytes.Buffer

    cmd.Stdout = &out
    err := cmd.Run()
    if err != nil {
        log.Fatal(err)
    }
    fmt.Printf("%s", out.String())
}
