package sorts

import (
	"context"
	"fmt"
	"sync"
	"testing"
	"time"
)

func TestBinaryFind(t *testing.T) {
	arr := make([]int, 1024*1024, 1024*1024)
	for i := 0; i < 1024*1024; i++ {
		arr[i] = i + 1
	}
	id := bin_search(arr, 1024)
	if id != -1 {
		fmt.Println("find:===", id, arr[id])
	} else {
		fmt.Println("没有找到数据")
	}
}

func bin_search(a []int, t int) int {
	low := 0
	high := len(a) - 1
	mid := (high + low) / 2
	for low <= high {
		mid = (high + low) / 2
		fmt.Println(mid)
		if a[mid] > t {
			high = mid - 1
		} else if a[mid] < t {
			low = mid + 1
		} else {
			return mid
		}
	}
	return -1
}

func TestPool(t *testing.T) {
	wg := sync.WaitGroup{}
	ch := make(chan int, 0)
	for i := 0; i < 5; i++ {
		go func() {
			for {
				if v, ok := <-ch; ok {
					fmt.Println(v)
					wg.Done()
				} else {
					fmt.Println("closed")
					wg.Done()
					return
				}

			}
		}()
	}
	wg.Add(5)

	for i := 0; i < 30; i++ {
		wg.Add(1)
		ch <- i
	}
	close(ch)

	wg.Wait()

	ctx := context.WithValue(context.Background(), 1, 1)
	_ = ctx
	ctx2, cancel := context.WithDeadline(ctx, time.Now())
	_ = ctx2
	cancel()
}
