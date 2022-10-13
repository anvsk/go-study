package chans

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

func TestDo(*testing.T) {
	Do()
}

// 未初始化的channel，读写阻塞，close->panic
// 初始化的 close 读零值\写panic
// 		[无缓冲]channel，读写阻塞
// 		[有缓冲]channel，读阻塞写不阻塞
func TestNoSpecific(*testing.T) {
	var aa chan int
	aa = make(chan int, 1)

	fmt.Println("<-aa=")
	// fmt.Println("<-aa=", <-aa)
	go func() {
		fmt.Println(<-aa)
	}()
	aa <- 1
	fmt.Println("toClose")
	close(aa)
	// aa <- 1
	if t1, ok := <-aa; ok {
		fmt.Println(t1)
	} else {
		fmt.Println("close")
	}
}

func TestXxxxx(t *testing.T) {

	lens := 30
	ch := make(chan bool, 30)
	var wg sync.WaitGroup
	for i := 0; i < lens; i++ {
		wg.Add(1)
		go func() {
			time.Sleep(time.Millisecond * 1)
			ch <- true
			wg.Done()
		}()
	}
	/*
		   go func() {
			 wg.Wait()
			 close(ch)
		  }()*/
	/*hjkdshfkjdsklfs、
	  wg.Wait()
	  close(ch)
	*/

	var s []bool
	for i := range ch {
		s = append(s, i)
	}
	wg.Wait()
	close(ch)
	fmt.Println(len(s))
	fmt.Println("~over~")
}
