package sync2

import (
	"fmt"
	"sync"
	"time"
)

func Runs() {
	max = 1000000

	// task1 := make(chan int, 5000)
	task2 := make(chan int, 1000)
	// reschan := make(chan int, 2)

	m1 := make(map[int]int, max)
	m2 := make(map[int]int, max)
	lock := sync.Mutex{}
	// 开携程池
	// for i := 0; i < 20; i++ {
	go func() {
		for {
			tmp := <-task2
			// lock.Lock()
			// defer lock.Unlock()
			m2[tmp] = tmp
		}
	}()
	var ss sync.Mutex{}
	// }

	// for i := 0; i < 50; i++ {
	// 	go func() {
	// 		tmp := <-task1
	// 		lock.Lock()
	// 		defer lock.Unlock()
	// 		m1[tmp] = tmp
	// 	}()
	// }
	// wg:=sync.WaitGroup{}
	// wg.Add(2)

	go func() {
		for {
			if len(m1) >= max {
				fmt.Printf("m1 down==%s\n", time.Now())
				// reschan <- 1
				// wg.Done()
				return
			}
		}
	}()

	go func() {
		for {
			if len(m2) >= max {
				fmt.Printf("m2 down==%s\n", time.Now())
				// reschan <- 2
				// wg.Done()
				return
			}
		}
	}()

	for i := 1; i <= max; i++ {
		tmp := i
		go func() {
			lock.Lock()
			defer lock.Unlock()
			m1[tmp] = tmp
		}()
		go func() {
			task2 <- tmp
		}()
		// task1 <- i
	}
	defer close(task2)
	// defer close(task1)
	// for v := range reschan {
	// 	println(v, " task finished ")
	// }
	// close(reschan)
	for {
	}
	// for len(reschan) != 1 {
	// 	println(len(reschan))
	// }
	println(" script finished ")
}

//
func withlock() {

}
