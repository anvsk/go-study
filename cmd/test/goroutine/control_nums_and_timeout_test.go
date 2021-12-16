// 用channel控制并发数量+超时限制
// !!! select case : <-func() chan 等待时无效，会随机选择、得现在外层定义chan后使用！
package goroutine

import (
	"testing"
	"time"
)

func TestDxxx(*testing.T) {
	avi := make(chan int, 5)
	for i := 0; i < 50; i++ {
		avi <- i
		go ttt3(avi)
	}

}

// 不用case:func()的方式(刚开始卡了半天)、外层定义chan
func ttt1(avi chan int) {
	done := make(chan struct{}, 1)
	// chan 引用类型，直接传过去
	// println("before", done)
	go handleFunc(done)
	// println("after", done)

	select {
	case <-done:
		println("suc")
	case <-time.After(2 * time.Second):

		println("timeout")
	}
	<-avi
}

func handleFunc(done chan struct{}) {
	time.Sleep(1 * time.Second)
	done <- struct{}{}
}

// 后来想通了，也可以用func() chan 方式，只需要在func里直接返回chan，然后移步处理handle
func ttt2(avi chan int) {
	done := make(chan struct{}, 1)
	select {
	case <-handleFunc2(done):
		println("suc")
	case <-time.After(2 * time.Second):

		println("timeout")
	}
	<-avi
}

func handleFunc2(done chan struct{}) chan struct{} {
	go handleFunc(done)
	return done
}

// 上面的简化版
// 后来想通了，也可以用func() chan 方式，只需要在func里直接返回chan，然后移步处理handle
func ttt3(avi chan int) {
	done := make(chan struct{}, 1)
	select {
	case <-func() chan struct{} {
		go handleFunc(done)
		return done
	}():
		println("suc")
	case <-time.After(2 * time.Second):

		println("timeout")
	}
	<-avi
}

// var suc int
// var timeout int

// func TestMost(*testing.T) {
// 	for i := 0; i < 5; i++ {
// 		// tmp := time.NewTimer(1 * time.Second)
// 		// tmp, cancel := context.WithTimeout(context.Background(), time.Duration((time.Second)))
// 		tmp, _ := context.WithTimeout(context.Background(), time.Second)
// 		// println(tmp)
// 		// tmp2 := time.After(2 * time.Second)
// 		for {
// 			select {
// 			case <-tmp.Done():
// 				println("timeou22t211")
// 				goto OUT
// 			case <-ttt():
// 				println("suc22222c22211e2ss")
// 				goto OUT
// 			}
// 		}
// 	OUT:
// 		// cancel()
// 		// tmp.Stop()
// 	}

// }

// func ttt() chan struct{} {
// 	time.Sleep(2001 * time.Millisecond)
// 	res := make(chan struct{}, 1)
// 	// println(res)
// 	res <- struct{}{}
// 	return res
// }

// func TestMostG(*testing.T) {
// 	//模拟用户需求业务的数量
// 	task_cnt := 50
// 	//task_cnt := 10

// 	ch := make(chan bool, 8)

// 	for i := 0; i < task_cnt; i++ {

// 		ch <- true
// 		go do(ch, i)
// 	}
// 	println("suc:", suc, "timeout:", timeout)
// }

// func do(ch chan bool, i int) {
// 	defer func() {
// 		<-ch
// 	}()
// 	begin := time.Now().Local().String()
// 	tmp := time.After(1 * time.Second)
// 	for {
// 		select {
// 		case <-tmp:
// 			println("timeout:", i)
// 			timeout++
// 			return
// 		case <-ganhuo(i):
// 			println("success:", i)
// 			println("suc-begin", i, begin)
// 			println("suc-end", i, time.Now().Local().String())
// 			suc++
// 			return
// 		}
// 	}

// }

// func ganhuo(i int) (done chan struct{}) {
// 	<-time.After(2001 * time.Millisecond)
// 	fmt.Println("go func ", i, " goroutine count = ", runtime.NumGoroutine())
// 	done = make(chan struct{}, 1)
// 	done <- struct{}{}
// 	return

// }
