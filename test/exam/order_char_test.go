package exam

import (
	"fmt"
	"sort"
	"testing"
)

func TestOrderChar(*testing.T) {
	s := "sadhjksahd82731jdksahd8"
	sli := []byte(s)
	sort.Slice(sli, func(i, j int) bool {
		return sli[i] < sli[j]
	})
	println(string(sli))
}

func TestMultiCompile(*testing.T) {
	sli1 := []string{"A", "AM", "XU"}
	sli2 := []string{"黄", "绿", "蓝"}
	for _, v1 := range sli1 {
		for _, v2 := range sli2 {
			println(v1 + v2)
		}
	}

}

func TestXxx(t *testing.T) {
	a := []int{1, 2, 3}
	b := []int{1, 1, 1, 1, 1, 1}
	fmt.Println(cap(a))

	fmt.Println(append(a, b...))
	fmt.Println(a)
	fmt.Println(cap(a))
}
