package sync2

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

var (
	maxWorkers = runtime.GOMAXPROCS(0)
	task       = make([]int, maxWorkers*4)

// 任务数，是worker的四
)

const (
	mutexLocked      = 1 << iota
	mutexWoken                   //2
	mutexStarving                //4
	mutexStarving2               //8
	mutexWaiterShift = iota      //4
	mutexStarving3               //5
	mutexStarving4   = 1 << iota //64
	mutexStarving5               //128
	mutexStarving6   = iota + 1
	mutexStarving7
	mutexStarving8 = iota * 3
	mutexStarving9
)

func main2() {
	// func TestXxxxxxxx(*testing.T) {

	// ctx := context.Background()
	// for i := range task {
	// 	fmt.Println("iiiii====", i)
	// 	// 如果没有worker可用，会阻塞在这里，直到某个worker被释放
	// 	if err := sema.Acquire(ctx, int64(i)); err != nil {
	// 		break
	// 	}
	// 	// 启动worker goroutine
	// 	go func(i int) {
	// 		defer sema.Release(int64(i))
	// 		time.Sleep(100 * time.Millisecond) // 模拟一个耗时操作
	// 		task[i] = i + 1
	// 	}(i)
	// }
	// // 请求所有的worker,这样能确保前面的worker都执行完
	// if err := sema.Acquire(ctx, int64(maxWorkers)); err != nil {
	// 	log.Printf("获取所有的worker失败: %v", err)
	// }
	// fmt.Println(maxWorkers, task)
	// ctx, cancel := context.WithTimeout(context.Background(), time.Second*2)
	// defer cancel()
	// sema := semaphore.NewWeighted(6) //信号量

	// if err := sema.Acquire(ctx, 5); err != nil {
	// 	fmt.Println(err.Error())
	// 	return
	// }
	// if err := sema.Acquire(ctx, 5); err != nil {
	// 	fmt.Println(err.Error())
	// 	// return
	// 	if err := sema.Acquire(ctx, 1); err != nil {
	// 		fmt.Println(err.Error())
	// 		return
	// 	}
	// }
	// sema.Release(4)
	// sema.Release(2)

	i := 0

	go func() {
		i++ // write i
		runtime.Gosched()
	}()
	// sync.Mutex{}
	fmt.Println(i) // read i

}

func loop(ch chan int) {
	for {
		select {
		case i := <-ch:
			fmt.Println("this  value of unbuffer channel", i)
		}
	}
}

// 测试fatalerror
func main() {

	go cc()
	// go func() {
	// 	fmt.Println("comming")
	// 	ch := make(chan int)
	// 	ch <- 1
	// 	fmt.Println("pass")
	// 	go loop(ch)
	// 	time.Sleep(1 * time.Millisecond)
	// }()
	for {
		<-time.After(time.Second)
		fmt.Println("---")
	}
}

func cc() {
	fmt.Println("comming")
	ch := make(chan int)
	ch <- 1
	fmt.Println("pass")
	go loop(ch)
	time.Sleep(1 * time.Millisecond)
}

func wmap() {
	m := map[int]interface{}{}

	wn := 2

	var wg sync.WaitGroup
	wg.Add(wn)

	for i := 0; i < wn; i++ {
		go func() {
			j := 0
			for {
				m[j] = j
				j++
				if j > 100000 {
					break
				}
			}
			wg.Done()
		}()
	}

	wg.Wait()

}
