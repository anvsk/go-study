package exam

import (
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
