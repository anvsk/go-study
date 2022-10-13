package chans

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

var wg sync.WaitGroup

func TestXxxxxxx(t *testing.T) {
	ch1 := make(chan int, 0)
	ch2 := make(chan int, 0)
	ch3 := make(chan int, 0)
	wg.Add(9)
	start := time.Now().Unix()

	go print("gorouine1", ch1, ch2)
	go print("gorouine2", ch2, ch3)
	go print("gorouine3", ch3, ch1)
	ch1 <- 1

	wg.Wait()
	end := time.Now().Unix()
	fmt.Printf("duration:%d\n", end-start)
}

func print(gorouine string, inputchan chan int, outchan chan int) {
	ii := 0
	for {
		// 模拟内部操作耗时
		time.Sleep(1 * time.Second)
		select {
		case <-inputchan:
			fmt.Println(gorouine, ii)
			if ii == 3 {
				if gorouine != "gorouine3" {
					outchan <- ii
				}
				wg.Done()
				goto OUT
			}
			ii++
			outchan <- ii
			wg.Done()
		}
	}
OUT:
}
