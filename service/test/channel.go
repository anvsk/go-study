package test

import (
    "context"
    "sync"
    "time"

    "github.com/pieterclaerhout/go-log"
)

func Tselect() {
    wg := sync.WaitGroup{}

    ctx, _ := context.WithTimeout(context.Background(), 1*time.Second)
    // defer cancel()
    a := make(chan int)
    b := make(chan int)
    funcc := func(c chan int, i int, ctx context.Context) {
        defer wg.Done()
        done := make(chan int)
        go func() {
            time.Sleep(time.Duration(i) * time.Second)
            c <- i
            done <- i
        }()
        go func() {

            select {
            case <-ctx.Done():
                log.Debug("gorouting-done")
            case <-done:
            }
        }()
    }
    wg.Add(2)
    go funcc(a, 2, ctx)
    go funcc(b, 4, ctx)
    var flag bool
    for {

        select {
        case <-a:
            log.Debug("a")
        case <-b:
            log.Debug("b")
        case <-ctx.Done():
            // case <-time.After(time.Second):
            // log.Debug("beats!!!")
            // default:
            log.Debug("default-doen")
            // return
            flag = true
            break
        }
        if flag {
            break
        }
    }
    wg.Wait()
    log.Debug("end")
}
