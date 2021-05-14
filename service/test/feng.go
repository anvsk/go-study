package test

import (
    "context"
    "time"

    "github.com/pieterclaerhout/go-log"
)

func Tfeng() error {
    ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
    defer cancel()
    // fn := func() error {
    //     <-time.After(2 * time.Second)
    //     return nil
    // }
    // create channel with buffer size 1 to avoid goroutine leak
    done := make(chan error)
    panicChan := make(chan interface{}, 1)
    go func() {
        defer func() {
            if p := recover(); p != nil {
                panicChan <- p
            }
        }()
        // done <- fn()
    }()

    select {
    case p := <-panicChan:
        panic(p)
    case err := <-done:
        log.Debug("<-down")
        return err
    case <-ctx.Done():
        log.Debug("<-timeout")
        return ctx.Err()
    }
}
