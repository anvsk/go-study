package main

import (
    "go-ticket/cmd/test/leetcode"
    "math/rand"
    "time"
)

var rr = rand.New(rand.NewSource(time.Now().UnixNano()))

func main() {
    // util.InitUtil()
    // cache.InitCache()
    // key := "ccc"
    // cache.C.Set(key, 123, 10*time.Second)
    // log.Debug(cache.C.Get(key))
    // db.InitDB()
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
    // <-time.After(1 * time.Hour)
    // a := "aaaa"
    // fmt.Println(a.Len())

    // sync.TestCond()
    // pprof.Testranddomstr()

    // fmt.Println(leetcode.Stradd("98", "55"))
    // ss := "klsadjla"
    // fmt.Println(ss[2])
    // rand.Seed(time.Now().UnixNano())
    // for i := 0; i < 3; i++ {
    // fmt.Println(myrand())
    // fmt.Println(myrand())
    // fmt.Println(myrand())
    // fmt.Println("--====--")
    // }

    // for i := 0; i < 5; i++ {
    //     rand.Seed(time.Now().UnixNano())
    //     fmt.Println(rand.Intn(100))
    // }

    leetcode.Test()

    <-time.After(time.Second)
}

func myrand() int {
    return rr.Intn(9999999)
}
