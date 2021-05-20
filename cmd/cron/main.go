package main

import (
    "fmt"
    "testing"
    "time"

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
}

func runTest() {
    c := cron.New(cron.WithSeconds(), cron.WithChain(
        cron.Recover(cron.DefaultLogger), // or use cron.DefaultLogger
    ))
    c.Start()
    // defer c.Stop()
    c.AddFunc("*/1 * * * * *", func() {
        <-time.After(time.Minute)
        fmt.Println(time.Now().Clock())
    })
    c.AddFunc("*/2 * * * * *", func() { fmt.Println("Every 2 second ") })
    c.AddFunc("@every 1s", func() { fmt.Println("@every 1s") })
    c.AddFunc("0 30 * * * *", func() { fmt.Println("Every hour on the half hour") })
    c.AddFunc("@hourly", func() { fmt.Println("Every hour") })
    c.AddFunc("@every 1h30m", func() { fmt.Println("Every hour thirty") })
    c.AddFunc("@daily", func() { fmt.Println("Every day") })
    <-time.After(time.Minute)
    // c.Stop() // Stop the scheduler (does not stop any jobs already running).
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
