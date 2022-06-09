package pi

import (
	"fmt"
	"runtime"
	"testing"
	"time"
)

var n int64 = 10000000000
var h float64 = 1.0 / float64(n)

func f(a float64) float64 {
	return 4.0 / (1.0 + a*a)
}

func chunk(start, end int64, c chan float64) {
	var sum float64 = 0.0
	for i := start; i < end; i++ {
		x := h * (float64(i) + 0.5)
		sum += f(x)
	}
	c <- sum * h
}

func TestCaculPI(*testing.T) {

	//记录开始时间
	start := time.Now()

	var pi float64
	// np := runtime.NumCPU()
	np := 100
	fmt.Println("np", np)
	runtime.GOMAXPROCS(np)
	c := make(chan float64, np)

	for i := 0; i < np; i++ {
		go chunk(int64(i)*n/int64(np), (int64(i)+1)*n/int64(np), c)
	}

	for i := 0; i < np; i++ {
		pi += <-c
	}

	fmt.Println("Pi: ", pi)

	//记录结束时间
	end := time.Now()

	//输出执行时间，单位为毫秒。
	fmt.Printf("spend time: %vs\n", end.Sub(start).Seconds())
}

func term2(ch chan float64, v float64) {
	k := int64(1) * n / int64(v+1)
	x := h * (float64(k) + 0.5)
	res := f(x) * h
	ch <- res
}
