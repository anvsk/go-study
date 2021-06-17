package sync

import (
    "log"
    "sync"
    "time"
)

var done = false

func read(name string, c *sync.Cond) {
    c.L.Lock()
    log.Println(name, "locking reading")
    for !done {
        c.Wait()
    }
    time.Sleep(time.Second)

    log.Println(name, "starts reading")

    c.L.Unlock()
}

func write(name string, c *sync.Cond) {
    log.Println(name, "starts writing")
    time.Sleep(2 * time.Second)
    c.L.Lock()
    done = true
    c.L.Unlock()
    log.Println(name, "wakes all")
    c.Broadcast()
}

func TestCond() {
    cond := sync.NewCond(&sync.RWMutex{})

    go read("reader1", cond)
    go read("reader2", cond)
    go read("reader3", cond)
    go read("reader4", cond)
    write("writer", cond)

    time.Sleep(time.Second * 7)
}
