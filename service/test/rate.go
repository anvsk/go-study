package test

import (
    "context"
    "log"
    "time"

    "golang.org/x/time/rate"
)

//limit表示每秒产生token数，buret最多存token数
//Allow判断当前是否可以取到token
//Wait阻塞等待知道取到token
//Reserve返回等待时间，再去取token

func Trate() {
    l := rate.NewLimiter(3, 5)
    log.Println(l.Limit(), l.Burst())
    for i := 0; i < 100; i++ {
        //阻塞等待直到，取到一个token
        log.Println("before Wait")
        c, _ := context.WithTimeout(context.Background(), time.Second*1)
        if err := l.Wait(c); err != nil {
            log.Println("limiter wait err:" + err.Error())
        }
        log.Println("after Wait")

        //返回需要等待多久才有新的token,这样就可以等待指定时间执行任务
        r := l.Reserve()
        log.Println("reserve Delay:", r.Delay())

        //判断当前是否可以取到token
        a := l.Allow()
        log.Println("Allow:", a)

        log.Printf("\n")
    }
}
