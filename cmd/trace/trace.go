package main

import (
	"fmt"
	"os"
	"runtime/trace"
	"time"
)

func main() {

	//创建trace文件
	f, err := os.Create("trace.out")
	if err != nil {
		panic(err)
	}

	defer f.Close()

	//启动trace goroutine
	err = trace.Start(f)
	if err != nil {
		panic(err)
	}
	defer trace.Stop()

	//main
	// for i := 0; i < 5; i++ {
	// 	time.Sleep(time.Second)
	// 	fmt.Println("Hello World")
	// }
	// runtime.GOMAXPROCS(1)
	for i := 0; i < 10; i++ {
		go func(ii int) {
			time.Sleep(time.Duration(700*ii) * time.Millisecond)
			fmt.Println(ii)
		}(i)
	}
	<-time.After(22 * time.Second)
}
