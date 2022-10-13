package chans

import (
	"fmt"
	"testing"
)

var end = make(chan int, 0)

func TestXssss(*testing.T) {

	chan1 := make(chan int, 0)
	chan2 := make(chan int, 0)
	chan3 := make(chan int, 0)

	go doPrint(chan1, chan2)
	go doPrint(chan2, chan3)
	go doPrint(chan3, chan1)
	chan1 <- 1
	fmt.Println("end", <-end)
}

func doPrint(cur chan int, next chan int) {
	for {
		select {
		case <-end:
			close(cur)
			return
		default:
			value := <-cur
			if value > 100 {
				close(cur)
				close(end)
				return
			}
			fmt.Println(value)
			value++
			// 通知下一个
			next <- value
		}
	}
}
