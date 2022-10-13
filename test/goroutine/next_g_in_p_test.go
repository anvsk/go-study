package goroutine

import (
	"bytes"
	"fmt"
	"runtime"
	"strconv"
	"testing"
	"time"
)

// P 的结构体包含一个gnext：最后一个g
func TestXxx(t *testing.T) {
	runtime.GOMAXPROCS(1)
	for i := 0; i < 20; i++ {
		// fmt.Println(GetGID())

		go func(ii int) {
			fmt.Println(GetGID(), "=====", ii)
		}(i)
	}
	<-time.After(time.Second)
	x, y, z := 1, 2, 3
	_ = x
	_ = y
	_ = z
}

func GetGID() uint64 {
	b := make([]byte, 64)
	b = b[:runtime.Stack(b, false)]
	b = bytes.TrimPrefix(b, []byte("goroutine "))
	b = b[:bytes.IndexByte(b, ' ')]
	n, _ := strconv.ParseUint(string(b), 10, 64)
	return n
}
