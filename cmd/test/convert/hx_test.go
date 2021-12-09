package convert

import (
	"fmt"
	"testing"
	"time"
)

var LastGetListTime time.Time

func TestMathPow(*testing.T) {
	fmt.Println(LastGetListTime)
	for {
		if time.Since(LastGetListTime) > 3*time.Second {
			println("send...")
			LastGetListTime = time.Now()
		}

		time.Sleep(time.Second)
	}
}

type Aaa struct {
	Tyrespeeds [8]int
}

// tmp, _ := strconv.ParseInt("16", 16, 0)
// fmt.Println(strconv.FormatInt(int64(tmp), 2))
// s := "A"
// tmp, _ := strconv.ParseInt(s, 16, 0)

// fmt.Println(strconv.Itoa(int(tmp)))
// [1 2 22 16 9 10 11 12 13 14 15]
// by := []byte{0x01, 02, 16, 10, 0x09, 0xA, 0x0b, 0x0c, 0x0D, 0x0e, 0x0f}
// fmt.Println(by)
// fmt.Println(byte(0x10))
// s := "总啥宽带哈上来的"
// fmt.Println([]rune("来的"))
// fmt.Println(strings.Contains(s, "来的"))

// by := []byte{178, 100, 101, 102, 3, 4, 5}
// res1 := strings.ToUpper(hex.EncodeToString(by))
// fmt.Println(res1)
// res2, _ := hex.DecodeString(res1)
// fmt.Println(res2)
// ss := time.Now().Add(3 * time.Second)
// s := float64(time.Since(ss).Seconds())
// fmt.Println(s)

// s := "sss"
// mm := sync.Map{}
// mm.Store("demo", s)
// mm.Store("demo", "aa")
// fmt.Println(mm.Load("demo"))
// 	for i := 0; i < 3; i++ {
// 		for j := 0; j < 3; j++ {
// 			println("j", i)
// 			if j == 1 {
// 				println("开始break", j)
// 				goto OUT
// 			}
// 		}
// 		println("i", i)
// 	}
// OUT:
// println("out")
