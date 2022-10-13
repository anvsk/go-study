package atomicsss

import (
	"fmt"
	"sync"
	"testing"
)

func TestXxx(t *testing.T) {
	// atomic.LoadInt64()
	// atomic.SwapInt32()
	// atomic.CompareAndSwapInt32()
	// var aa atomic.Value
	// aa.Store(1)
	// fmt.Println(aa.Load())
	// aa.Swap(3)
	// fmt.Println(aa.Load())
	// fmt.Println(aa.CompareAndSwap(1, 4))
	// fmt.Println(aa.CompareAndSwap(3, 4))
	// errgroup.Group

	pool := sync.Pool{
		New: func() interface{} {
			fmt.Println("new")
			return 1
		},
	}
	// pool.Put(3)
	for i := 0; i < 10; i++ {
		v := pool.Get().(int)
		fmt.Println(v) // 取出来的值是put进去的，对象复用；如果是新建对象，则取出来的值为0
		pool.Put(i)
	}
}
