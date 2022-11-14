package exams

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

var wg sync.WaitGroup

func TestXxx111(t *testing.T) {
	wg = sync.WaitGroup{}
	nn := 5
	wg.Add(nn)
	fmt.Println("begin")
	stop := make(chan struct{}, 1)
	go gprint(nn, stop)
	<-time.After(5 * time.Second)
	close(stop)
	wg.Wait()
}

// 一个数字n,开n个携程,递增打印1,2,3,4..

func gprint(n int, stop chan struct{}) {
	chs := make([]chan int, n)

	for i := 0; i < n; i++ {
		chs[i] = make(chan int, 1)
		if i == 0 {
			chs[0] <- 1
		}
		go func(m int, ch chan int) {
			for {
				select {
				case <-stop:
					fmt.Println(m, "====out")
					wg.Done()
					return
				default:
					nn := <-ch
					time.Sleep(500 * time.Millisecond)
					fmt.Println(m, ":print=====:", nn)
					if m == n-1 {
						chs[0] <- nn + 1
					} else {
						chs[m+1] <- nn + 1
					}
				}

			}

		}(i, chs[i])
	}

}
