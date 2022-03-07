package leecode

import (
	"context"
	"log"
	"runtime"
	"sync"
	"testing"
	"time"
)

func TestAsyncdo(*testing.T) {
	log.SetFlags(4)
	log.Println(" Start ")
	nums := 100000000000
	// 不用go 45s
	// if nums > 0 {
	// 	for i := 0; i < nums; i++ {
	// 		_ = i * i
	// 		if i < nums-10 {
	// 			continue
	// 		}
	// 		log.Println(i, ":", i*i)
	// 	}
	// 	return
	// }
	// 装结果
	var res = make(map[int]int, nums) // fatal error: concurrent map writes
	var lock = sync.RWMutex{}
	// 任务通道
	task := make(chan int, runtime.NumCPU()+2)
	// 等待结束
	wg := sync.WaitGroup{}
	// 投递任务
	wg.Add(nums)
	go func() {
		for i := 1; i <= nums; i++ {
			task <- i
		}
	}()
	// cpu个数多线程接收任务并处理，超时打印
	for i := 0; i < runtime.NumCPU(); i++ {
		go func() {
			for {
				// context用法示例
				ctx, cancel := context.WithTimeout(context.Background(), time.Millisecond*1200)
				oneTask, ok := <-task
				if !ok {
					cancel()
					// 协程退出出口
					return
				}
				// 传一个chan接收结果，开始干活
				var done = make(chan int)
				// 模拟耗时动作
				go handle(oneTask, done)
				// 阻塞的通通go起来
				go func() {
					defer wg.Done()
					select {
					// 任务正常完成
					case oneRes := <-done:
						lock.Lock()
						res[oneTask] = oneRes
						lock.Unlock()
					// 一秒超时
					case <-time.After(time.Second):
						// log.Println(oneTask, " time out by after")
					// 从context超时
					case <-ctx.Done():
						// log.Println(oneTask, " time out by ctx")
					}
				}()
			}
		}()
	}
	// 等待全部处理完
	wg.Wait()
	// 展示结果
	for i := 0; i < len(res); i++ {
		if i < len(res)-10 {
			continue
		}
		if res[i] == 0 {
			log.Println(i, ": time out")
			continue
		}
		log.Println(i, ":", res[i])
	}

}

// 模拟实际耗时处理(随机0-1.5秒)
func handle(i int, ch chan int) {
	// rand.Seed(time.Now().UnixNano())
	// time.Sleep(time.Duration(rand.Intn(1500)) * time.Millisecond)
	ch <- i * i
}
