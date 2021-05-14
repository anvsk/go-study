package test

import (
    "sync"
    "time"

    "github.com/pieterclaerhout/go-log"
)

func gen(nums ...int) <-chan int {
    out := make(chan int)
    go func() {
        for _, n := range nums {
            // log.Info("aa")
            out <- n
        }
        close(out)
    }()
    return out
}

func sq(in <-chan int) <-chan int {
    out := make(chan int)
    go func() {
        for n := range in {
            // time.Sleep(2 * time.Second)
            out <- n * n
        }
        close(out)
    }()
    return out
}

// func merge(cs ...<-chan int) <-chan int {
//     var wg sync.WaitGroup
//     out := make(chan int)

//     // Start an output goroutine for each input channel in cs.  output
//     // copies values from c to out until c is closed, then calls wg.Done.
//     output := func(c <-chan int) {
//         for n := range c {
//             out <- n
//         }
//         wg.Done()
//     }
//     wg.Add(len(cs))
//     for _, c := range cs {
//         go output(c)
//     }

//     // Start a goroutine to close out once all the output goroutines are
//     // done.  This must start after the wg.Add call.
//     go func() {
//         wg.Wait()
//         close(out)
//     }()
//     return out
// }

func merge(cc ...<-chan int) <-chan int {
    wg := sync.WaitGroup{}
    outchan := make(chan int)
    outfunc := func(c <-chan int) {
        defer wg.Done()
        for v := range c {
            outchan <- v
        }
    }
    for _, v := range cc {
        wg.Add(1)
        go outfunc(v)
    }
    go func() {
        wg.Wait()
        close(outchan)
    }()
    return outchan
}

func putjob(n int) <-chan int {
    ch := make(chan int, 50)
    go func() {
        i := 0
        for range time.Tick(1 * time.Millisecond) {
            i++
            ch <- i
            if i >= 10 {
                log.Info("生产完毕。。。")
                break
            }
        }
        close(ch)
    }()
    return ch
}

func dojob(ch <-chan int) {
    wg := sync.WaitGroup{}
    dst := func(n int) {
        defer wg.Done()
        time.Sleep(3 * time.Second)
        log.Warn("已处理", n)
    }
    for v := range ch {
        wg.Add(1)
        go dst(v)
    }
    wg.Wait()
    log.Warn("接收完毕")
}

func dojob2(ch <-chan int) {
    wg := sync.WaitGroup{}
    dst := func(ch <-chan int, i int) {
        defer wg.Done()
        num := 0
        for {
            n, ok := <-ch
            if ok == false {
                log.Warn("协程", i, "接收完毕", n)
                break
            } else {
                time.Sleep(1 * time.Second)

                log.DebugSeparator("协程", i)
                log.Warn("协程", i, "在处理", n)
                num += 1
            }
        }
        log.Info("协程", i, "已处理", num, "个任务")
    }
    task := 2
    wg.Add(task)
    for i := 0; i < task; i++ {
        go dst(ch, i)
    }
    wg.Wait()

    log.Warn("end")

}
