package goroutine

import (
	"fmt"
	"testing"
	"time"
)

func TestChannelWait(t *testing.T) {
	chans := make([]chan int, 10)
	for i := 0; i < 10; i++ {
		ch := make(chan int)
		chans = append(chans, ch)
		go Pro(ch)
	}
	for _, v := range chans {
		fmt.Println(<-v)
	}
}

func Pro(ch chan int) {
	time.Sleep(time.Second)
	ch <- 1
}
