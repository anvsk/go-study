package main

import (
	"context"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"

	_ "net/http/pprof"
)

var rr = rand.New(rand.NewSource(time.Now().UnixNano()))
var wg sync.WaitGroup

func main() {
	// util.InitUtil()
	// cache.InitCache()
	// key := "ccc"
	// cache.C.Set(key, 123, 10*time.Second)
	// log.Debug(cache.C.Get(key))
	// db.InitDB()
	// for i := 0; i < 1000000000000000000; i++ {
	//     // <-time.After(2 * time.Millisecond)
	//     log.Debug(i)
	//     go func(ii int) {
	//         time.After(100 * time.Millisecond)
	//         log.Info(ii)
	//     }(i)
	//     // go db.TestMysql()
	//     // go db.TestCH()

	// }
	// <-time.After(1 * time.Hour)
	// a := "aaaa"
	// fmt.Println(a.Len())

	// sync.TestCond()
	// pprof.Testranddomstr()

	// fmt.Println(leetcode.Stradd("98", "55"))
	// ss := "klsadjla"
	// fmt.Println(ss[2])
	// rand.Seed(time.Now().UnixNano())
	// for i := 0; i < 3; i++ {
	// fmt.Println(myrand())
	// fmt.Println(myrand())
	// fmt.Println(myrand())
	// fmt.Println("--====--")
	// }

	// for i := 0; i < 5; i++ {
	//     rand.Seed(time.Now().UnixNano())
	//     fmt.Println(rand.Intn(100))
	// }

	// leetcode.Test()
	go func() {
		http.ListenAndServe(":9008", nil)
	}()
	// f, err := os.Create("trace.out")
	// if err != nil {
	// 	panic(err)
	// }

	// defer f.Close()

	// //启动trace goroutine
	// err = trace.Start(f)
	// if err != nil {
	// 	panic(err)
	// }
	// defer trace.Stop()

	// log.SetFlags(0)
	// log.Println("begin", carbon.Now().ToDateTimeString())
	// a := []int{1}
	// a = append(a, 1, 2, 3)
	// println(a)
	// println(append([]byte("hello "), "world"...))
	// fmt.Println(string(append([]byte("hello "), "world"...)))
	// t1 := time.Now()
	// time.Sleep(time.Second)
	// t2 := time.Since(t1)
	// println(t2)
	// sliceBiggerCapNums()
	// sortprintln(5)
	// SignalHandle()
	// testContext()
	// consume(15)
	// sync2.Comparesyncchan()
	// getslicecap()
	// ctx,_:=context.WithTimeout(context.Background(),time.Second)
	// testclosechan()
	// go gogogo()
	// go gogogo()
	// go gogogo2()
	// go gogogo3()
	// go gogogo3()
	// go gogogo3()
	// go gogogo3()
	// go gogogo3()
	// go gogogo3()
	// go gogogo()
	// gogogo()
	// println(GameCallScore)
	// println(GameCallScore2)
	// println(GameCallScore3)
	// println(GamePlaying)
	// println(GamePlaying2)
	// println(GamePlaying3)
	// println(GamePlaying3)
	// println(aa)
	// println(GameEnd2)
	// s := []int{1, 2, 3, 4, 5, 6}
	// a := s[:3]
	// a = append(a, 7)
	// fmt.Println(s, a) //[6/6]0x14000157f38 [4/6]0x14000157f38 [1 2 3 7 5 6] [1 2 3 7]
	// b := append(a, []int{8, 8, 9}...)
	// fmt.Println(s, a, b)//[1 2 3 7 5 6] [1 2 3 7] [1 2 3 7 8 8 9]
	// a 1,2,3,7,5,6
	// s =a
	// var x int
	// t := runtime.GOMAXPROCS()
	// fmt.Println(runtime.NumCPU())
	// for i := 0; i < 100; i++ {
	// 	go func() {
	// 		for {
	// 			x++
	// 		}
	// 	}()
	// }
	// time.Sleep(time.Second)
	// fmt.Println("x =", x)
	// log.Println("end", carbon.Now().ToDateTimeString())

	// for {
	// }
}

const (
	GameCallScore int = 9

	GameWaitting = iota
	GamePlaying
	GameEnd
)

const (
	GameCallScore2 = iota
	// GameWaitting2  int = 9
	bb
	GamePlaying2
	aa
	GameEnd2
)

const (
	GameWaitting3  int = 9
	GameCallScore3 int = 99
	GamePlaying3       = iota
)

func gogogo() {
	for i := 0; i < 100; i++ {
		go func() {
			for {
			}
		}()
	}
	for {
	}
}

func gogogo2() {
	for i := 0; i < 100; i++ {
		go func() {
			for {
			}
		}()
	}
	for {
	}
}

func gogogo3() {
	for i := 0; i < 100; i++ {
		go func() {
			for {
			}
		}()
	}
	for {
	}
}

func partition(array []int, i int, j int) int {
	//第一次调用使用数组的第一个元素当作基准元素
	pivot := array[i]
	for i < j {
		for j > i && array[j] > pivot {
			j--
		}
		if j > i {
			array[i] = array[j]
			i++
		}
		for i < j && array[i] < pivot {
			i++
		}
		if i < j {
			array[j] = array[i]
			j--
		}
	}
	array[i] = pivot
	return i
}

func quicksort(array []int, low int, high int) {
	var pivotPos int //划分基准元素索引
	if low < high {
		pivotPos = partition(array, low, high)
		quicksort(array, low, pivotPos-1)
		quicksort(array, pivotPos+1, high)
	}
}

func quick_sort(li []int, left, right int) {
	if left >= right {
		return
	}
	i := left
	j := right
	rand.Seed(time.Now().Unix())
	r := rand.Intn(right-left) + left
	li[i], li[r] = li[r], li[i]
	tmp := li[i]
	for i < j {
		for i < j && li[j] >= tmp {
			j--
		}
		li[i] = li[j]
		for i < j && li[i] <= tmp {
			i++
		}
		li[j] = li[i]
	}
	li[i] = tmp
	quick_sort(li, left, i-1)
	quick_sort(li, i+1, right)
}

func bubble_sort(li []int) {
	for i := 0; i < len(li)-1; i++ {
		exchange := false
		fmt.Println("i=", i)

		for j := 0; j < len(li)-i-1; j++ {
			fmt.Println("j", j)

			if li[j] > li[j+1] {
				fmt.Println("===li[j]")
				fmt.Println(li[j])
				fmt.Println(li[j+1])
				fmt.Println("===li[j]")

				li[j], li[j+1] = li[j+1], li[j]
				exchange = true
			}
		}
		fmt.Println(exchange)

		if !exchange {
			return
		}
	}
}

func getslicecap() {
	a := []byte{}
	a = append(a, 1)
	fmt.Println("cap of a is ", cap(a))

	b := []int{23, 51}
	b = append(b, 4, 5, 6)
	fmt.Println("cap of b is ", cap(b))

	c := []int32{1, 23}
	c = append(c, 2, 5, 6)
	fmt.Println("cap of c is ", cap(c))

	type D struct {
		age  byte
		name string
	}
	d := []D{
		{1, "123"},
		{2, "234"},
	}

	d = append(d, D{4, "456"}, D{5, "567"}, D{6, "678"})
	fmt.Println("cap of d is ", cap(d))
}

func testclosechan() {
	// ch1 := make(chan int, 30)
	// go func() {
	// 	for v := range ch1 {
	// 		time.Sleep(time.Second)
	// 		println("aaaa", v)
	// 	}
	// }()
	// ch1 <- 11
	// ch1 <- 11
	// ch1 <- 11
	// close(ch1)
	// for {
	// }

	// reschan如果没有close会一直阻塞
	// reschan := make(chan int, 10)
	// go func() {
	// 	time.Sleep(2 * time.Second)
	// 	reschan <- 999
	// 	// close(reschan)
	// }()

	// for v := range reschan {
	// 	println(v, " task finished ")
	// }
	// println(" script finished ")
}

type smallobj struct {
	arr [1 << 10]byte
	i   int
}

func consume(n int) {
	ch := make(chan int, n)
	go limitPerSecond(ch, n)
	for i := 1; i < 10000; i++ {
		ch <- i
	}
	for {
	}
}

// 生产者、限流器、消费者模型
// 限流器,分发函数
// 每秒通过的请求数量10？
func limitPerSecond(rec chan int, limit int) {
	for {
		<-time.Tick(time.Second)
		i := 0
		for {
			go myworker(<-rec)
			i++
			if i == limit {
				break
			}
		}
	}
}

func myworker(i int) {
	// time.Sleep(200 * time.Millisecond)
	println(i)
}

func serveHttp() {
	http.HandleFunc("/dhaha", func(rw http.ResponseWriter, r *http.Request) {
		b := "haha"
		rw.Write([]byte(b))
	})

	http.Get("aa")
	http.ListenAndServe(":9999", nil)
	for {
	}
}

func testContext() {
	// ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	// ctx, cancel := context.WithCancel(context.Background())
	ctx, cancel := context.WithDeadline(context.Background(), time.Now().Add(5*time.Second))
	go func() {
		<-time.After(20 * time.Second)
		cancel()
	}()
	select {
	case <-ctx.Done():
		println("aasdsad")
		tt, bool := ctx.Deadline()
		fmt.Println(tt)
		println(bool)
		fmt.Println(ctx.Err())
	}
}

func createM() {
	ch := make(chan int)
	for i := 0; i < 300; i++ {
		go func() {
			time.Sleep(5 * time.Second)
			ch <- 1
		}()
		go func() {
			println("aaa")
		}()
		go func(i int) {
			println(i)
			time.Sleep(3000 * time.Second)
			println(i, "end")

		}(i)
		time.Sleep(80 * time.Millisecond)
	}
	<-ch
	<-time.After(5 * time.Second)
}

func SignalHandle() {
	for {
		ch := make(chan os.Signal)
		signal.Notify(ch, syscall.SIGINT)
		signal.Notify(ch, syscall.SIGKILL)
		sig := <-ch
		fmt.Printf("收到信号：%d %s\n", sig, sig.String())
		switch sig {
		case syscall.SIGINT:
			println("我要开始休眠三秒sigint")
			<-time.After(3 * time.Second)
		case syscall.SIGKILL:
			println("我要开始休眠三秒sigkill")
			<-time.After(3 * time.Second)
			// os.Exit(1)
		default:
			println("我进来default了")
		}
	}
}

func testblock(to time.Duration) int {
	ch := make(chan int)
	wg.Add(1)
	go func() {
		defer wg.Done()
		time.Sleep(time.Second)
		ch <- 1
	}()
	select {
	case <-time.After(to):
		return -1
	case res := <-ch:
		return res
	}

}

// 三个打印函数，要求用三个携程顺序打印各个N次
func sortprintln(n int) {
	ch1 := make(chan int)
	ch2 := make(chan int)
	ch3 := make(chan int)
	closech := make(chan bool)
	go pcat(ch1, ch3, n)
	go pdog(ch2, ch1)
	go pbird(ch3, ch2, closech)

	<-closech
}

func pcat(selfchan, prevchan chan int, times int) {
	i := 0
	for {
		i++
		if i == 1 {
			println("begining... ...")
			println("cat:%d", i, "==")
			selfchan <- 1
			continue
		}
		<-prevchan

		println("cat:%d", i, "==")

		if i == times {
			println("finished... ...ch1")
			close(selfchan)
			break
		}
		selfchan <- 1
	}
}

func pdog(selfchan, prevchan chan int) {
	i := 0
	for {

		tmp := <-prevchan
		if tmp == 0 {
			println("finished... ...ch2")
			close(selfchan)
			break
		}
		i++
		println("dog:%d", i, "==")

		selfchan <- 1
	}
}

func pbird(selfchan, prevchan chan int, closech chan bool) {
	i := 0
	for {

		tmp := <-prevchan
		if tmp == 0 {
			println("finished... ...ch3")
			closech <- true
			break
		}
		i++
		println("bird:%d", i, "==")

		selfchan <- 1
	}
}

func plast(ch1, ch2, ch3 chan struct{}) {
	<-ch2
	println("2:last")
	ch3 <- struct{}{}
}

func slicereal() {
	x := []int{1, 2, 3, 4, 5, 6} //cap 6
	y := x[4:5]                  //cap 2 value[5,6]
	log.Println(cap(y))          //
	log.Println((y))             //
	y[0] = 99
	y = append(y, 10, 10, 10, 10)

	log.Println(x) // 99 ,2,3,4,5,6
	log.Println(y) // 5

	// y = append(y, 10)
	// y[0] = 88
	// log.Println(x) // 88 ,2,3,4,5,6

	// y = append(y, 10)
	// y[0] = 77
	// log.Println(x) // 88 ,2,3,4,5,6

	// y = append(y, 10)

	// y = append(y, 50)// [5,6,50]
	// log.Println(x)//
	// y = append(y, 60)//[5,6,50,60]
	// log.Println(x)//
	// log.Println(y)//
	// y[0] = 20
	// log.Println(y)
}

func sliceBiggerCapNums() {
	a := []byte{1, 0}
	a = append(a, 1, 1, 1)
	fmt.Println("cap of a is ", cap(a))

	b := []int{23, 51}
	b = append(b, 4, 1, 1)
	// b = append(b, 4)
	// b = append(b, 4)
	fmt.Println("cap of b is ", cap(b))

	c := []int32{1, 23}
	c = append(c, 2, 5, 6)
	fmt.Println("cap of c is ", cap(c))

	// e := []float64{1, 23}
	// e = append(e, 2, 5, 6)
	// fmt.Println("cap of e is ", cap(e))

	type D struct {
		age  byte
		name string
	}
	d := []D{
		{1, "123"},
		{2, "234"},
	}

	d = append(d, D{4, "456"}, D{5, "567"}, D{6, "678"}, D{6, "678"})
	fmt.Println("cap of d is ", cap(d))

	e := []int32{1, 2, 3}
	fmt.Println("cap of e before:", cap(e))
	e = append(e, 4)
	fmt.Println("cap of e after:", cap(e))

	f := []int{1, 2, 3}
	fmt.Println("cap of f before:", cap(f))
	f = append(f, 4)
	fmt.Println("cap of f after:", cap(f))
}

func myrand() int {
	return rr.Intn(9999999)
}

// 测试有缓冲相关
func nocachechan() {
	ch1 := make(chan int, 90)
	log.Println(<-ch1)
	log.Println(<-ch1)
	log.Println(<-ch1)
}

// 仅用chan实现 同时N线程下载效果
func forrgetchannel() {
	ch1 := make(chan int, 5)
	down := make(chan int)
	go func() {
		log.Println("<-ch1begin")
		<-time.After(1 * time.Second)
		log.Println("<-ch1end")
		for i := 0; i < 10; i++ {
			log.Printf("开启第%d个消费携程", i)
			go func(i int, wg2 *sync.Mutex) {
				for {
					wg2.Lock()

					tmp := <-ch1
					if tmp == 0 {
						log.Println("==消费g", i, "downed")
						down <- 1
						break
					}
					log.Println("==consume==", tmp)
					<-time.After(2 * time.Second)
					wg2.Unlock()
				}
			}(i, &sync.Mutex{})
		}

	}()

	go func() {
		for i := 1; i < 5000; i++ {
			ch1 <- i
			log.Println("==put==", i)
		}
		close(ch1)
	}()
	<-down
}

func selectblock4() {
	ch1 := make(chan int, 0)
	n := 10
	count := 1

	go func() {
		log.Printf("begin handle task... ...")
		<-time.After(1 * time.Second)
		for {
			log.Println("begin for")
			<-time.After(1 * time.Second)

			tmp := <-ch1
			log.Println(tmp)

			if tmp == 0 {
				log.Printf("closed channel")
				break
			}
			log.Println("end for")
			count++
			if count > 15 {
				break
			}
		}
	}()
	for i := 1; i < n; i++ {
		ch1 <- i
	}
	// close(ch1)
	<-time.After(8000 * time.Millisecond)
	log.Printf("finished")
}

func selectblock2() {
	ch1 := make(chan int, 0)

	go func() {
		log.Printf("begin handle task... ...")
		<-time.After(1 * time.Second)
		for {
			log.Println("begin for")

			tmp := <-ch1
			log.Println(tmp)

			if tmp == 0 {
				log.Printf("closed channel")
				break
			}
			log.Println(<-ch1)
			log.Println("next for")
		}
	}()
	for i := 1; i < 10; i++ {
		ch1 <- i
	}
	close(ch1)
	<-time.After(8000 * time.Millisecond)
	log.Printf("finished")
}

func selectblock3() {
	ch1 := make(chan int, 0)

	ch1 <- 1
	go func() {
		time.After(2 * time.Second)
		fmt.Println(<-ch1)
	}()

	time.After(1 * time.Second)
}

func selectblock() {
	ch1 := make(chan int)
	ch2 := make(chan int)
	ctx, cancel := context.WithCancel(context.Background())
	fmt.Println(00)
	go func() {
		ch2 <- 2
		ch1 <- 1
		ch2 <- 2
		<-time.After(4 * time.Second)

		cancel()
	}()

	for {
		select {
		case <-ctx.Done():
			fmt.Println("ctx.Done")
			goto aa
		case <-ch1:
			fmt.Println(11)
		case <-ch2:
			fmt.Println(22)
		default:
			fmt.Println("default")
			<-time.After(500 * time.Millisecond)
		}
	}
aa:

	log.Println(<-ctx.Done())
	log.Println(<-ctx.Done())
	log.Println(<-ctx.Done())
	log.Println(<-ctx.Done())
	log.Println(<-ctx.Done())
	log.Println(<-ctx.Done())
	fmt.Println(99)
}
