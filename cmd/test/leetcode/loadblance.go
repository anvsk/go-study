package leetcode

import (
    "fmt"
    "math/rand"
    "time"
)

type lb struct {
    servers []int
    r       *rand.Rand
    index   int
}

type DisMode int

const (
    RANDOM DisMode = iota
    LUNXUN
)

type Dis interface {
    Get()
}

func NewLoadBalance() *lb {
    return &lb{
        servers: []int{1, 2, 3, 4, 5, 6, 7, 8, 9},
        r:       rand.New(rand.NewSource(time.Now().UnixNano())),
    }
}

func (bb *lb) Get(mode DisMode) int {
    switch mode {
    case RANDOM:
        return bb.servers[bb.r.Intn(len(bb.servers))]
    case LUNXUN:
        bb.index += 1
        return bb.servers[bb.index%(len(bb.servers))]
    }
    return 0
}

func Test() {
    bb := NewLoadBalance()
    n := 20
    for i := 0; i < n; i++ {
        fmt.Println(bb.Get(RANDOM))
    }
    fmt.Println("======")
    for i := 0; i < n; i++ {
        fmt.Println(bb.Get(LUNXUN))
    }
}
