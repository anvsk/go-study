// 做一些转换测试
package convert

import (
	"fmt"
	"testing"
	"time"
)

func TestTostr(*testing.T) {
	// res, _ := strconv.Atoi("09")
	// res := strings.Split("aaaa", "a")
	// s := "a12387sajd"
	// fmt.Println(s[len(s)-2:])
	a := A{Num1: 1}
	go a.aa()
	tick := time.NewTimer(3 * time.Second)
	for {

		select {
		case <-tick.C:
			fmt.Println("timeout...")
			goto AA
		default:
			if a.Num1 == 0 {
				fmt.Println(" success、...")
				goto AA
			} else {
				fmt.Println("each...")
			}
		}
		time.Sleep(100 * time.Millisecond)
	}
AA:
}

type A struct {
	Num1 int
}

func (a *A) aa() {
	time.Sleep(time.Second)
	a.Num1 = 0
}
