package sorts

import (
	"fmt"
	"testing"
)

func TestQuick(t *testing.T) {
	a := []int{9, 8, 7, 1, 2, 3, 4, 77, 0}
	fmt.Println(quick(a))
}

func quick(a []int) []int {
	if len(a) <= 1 {
		return a
	}
	l := []int{}
	r := []int{}
	m := []int{a[0]}
	for i := 1; i < len(a); i++ {
		if a[i] < a[0] {
			l = append(l, a[i])
		} else if a[i] > a[0] {
			r = append(r, a[i])
		} else {
			m = append(m, a[i])
		}
	}
	res := append(append(quick(l), m...), quick(r)...)
	return res
}
