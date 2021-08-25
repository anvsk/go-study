package sync2

import (
	"fmt"
	"sync"
	"time"
)

type Op struct {
	key int
	val int
}

var lock sync.Mutex

var m1 map[int]int
var m2 map[int]int
var max int = 500000

func update_map_by_mutex(i int) {
	lock.Lock()
	m1[i] = i
	if len(m1) == max {
		fmt.Printf("%s mutex finish\n", time.Now())
	}
	lock.Unlock()
}

var ch chan Op

func update_map_by_chan(i int) {
	ch <- Op{key: i, val: i}
}

func wait_for_chan(m map[int]int) {
	for {
		select {
		case op := <-ch:
			m[op.key] = op.val
			if len(m2) == max {
				fmt.Printf("%s chan finish\n", time.Now())
				return
			}
		}
	}
}

func Comparesyncchan() {

	m1 = make(map[int]int, max)
	m2 = make(map[int]int, max)
	ch = make(chan Op)
	go wait_for_chan(m2)
	for i := 0; i < max; i++ {
		go update_map_by_chan(i)
		go update_map_by_mutex(i)
	}
	for {
	}
}
