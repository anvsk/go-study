package main

import (
	"fmt"
	"time"
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

// func main() {
// 	demo1.Demo1()
// 	for {
// 	}
// }

func main() {
	fmt.Println("begin", time.Now())
	// tt := time.NewTicker(time.Second)
	tt := GetTicker(time.Second)
	<-time.After(3 * time.Second)
	go func() {
		for {
			<-tt.C
			ttt := tt.T
			fmt.Println("tt:", time.Now().Format("2006-01-02 15:04:05"), ttt.Format("2006-01-02 15:04:05"))
		}
	}()
	for {
	}
}

type Aa struct {
	C chan struct{}
	T time.Time
}

// change return struct to {chan + time(now)}
func GetTicker(d time.Duration) Aa {
	ch := Aa{
		C: make(chan struct{}),
	}
	go func() {
		for {
			time.Sleep(d)
			// todo
			ch.C <- struct{}{}
			ch.T = time.Now()
		}
	}()
	return ch
}
