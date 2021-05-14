package main

import (
    "go-ticket/pkg/store/cache"
    "go-ticket/pkg/store/db"
    "go-ticket/pkg/util"
    "time"

    "github.com/pieterclaerhout/go-log"
)

func main() {
    util.InitUtil()
    cache.InitCache()
    key := "ccc"
    // cache.C.Set(key, 123, 10*time.Second)
    log.Debug(cache.C.Get(key))
    db.InitDB()
    // for i := 0; i < 1000000000000000000; i++ {
    //     // <-time.After(2 * time.Millisecond)
    //     log.Debug(i)
    //     go func(ii int) {
    //         time.After(100 * time.Millisecond)
    //         log.Info(ii)
    //     }(i)
    //     // go db.TestMysql()
    //     // go db.TestCH()

    // }
    <-time.After(1 * time.Hour)
}
