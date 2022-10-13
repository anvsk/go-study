package sync2

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

func TestSyncCond(t *testing.T) {
	c := sync.NewCond(&sync.Mutex{})    //1
	queue := make([]interface{}, 0, 10) //2

	removeFromQueue := func(delay time.Duration) {
		time.Sleep(delay)
		c.L.Lock() //8
		fmt.Println("Removed ", queue[0], "from queue")
		queue = queue[1:] //9
		c.L.Unlock()      //10
		c.Signal()        //11
	}

	for i := 0; i < 10; i++ {
		c.L.Lock()           //3
		if len(queue) == 2 { //4
			c.Wait() //5
		}
		fmt.Println("Adding", i, " to queue")
		queue = append(queue, struct{}{})
		go removeFromQueue(1 * time.Second) //6
		c.L.Unlock()                        //7
	}
}

// func TestXxxsss(t *testing.T) {
// 	rand := func() interface{} { return rand.Intn(50000000) }

// 	done := make(chan interface{})
// 	defer close(done)

// 	start := time.Now()

// 	randIntStream := toInt(done, repeatFn(done, rand))
// 	fmt.Println("Primes:")
// 	for prime := range take(done, primeFinder(done, randIntStream), 10) {
// 		fmt.Printf("\t%d\n", prime)
// 	}

// 	fmt.Printf("Search took: %v", time.Since(start))
// }

func TestCcountPrimes(*testing.T) {
	n := 10
	isPri := make([]bool, n)
	for i := 2; i*i <= n; i++ {
		if !isPri[i] {
			for j := i * i; j < n; j += i {
				isPri[j] = true
			}
		}
	}
	count := 0
	for k := 1; k < n; k++ {
		if !isPri[k] {
			fmt.Println(k)

			count++
		}
	}
	fmt.Println("===count:", count)
}
