package question

import (
	"fmt"
	"testing"
	"time"
)

type A struct {
	Mmap map[string]string
	Innt int
}

func TestXxx(t *testing.T) {
	a := A{
		Mmap: map[string]string{"ss": "xx"},
		Innt: 2,
	}
	change(a)
	fmt.Println(a)
	changes(&a)
	fmt.Println(a)
}

func change(a A) {
	a.Mmap = make(map[string]string)
	a.Innt = 3
}

func changes(a *A) {
	a.Mmap = make(map[string]string)
	a.Innt = 3
}

func TestXxx2(t *testing.T) {
	// m1 := map[int]int{9: 9}
	// fmt.Println(&m1[9])
}

func AddElement(slice []int, e int) []int {
	return append(slice, e)
}

func TestXaaaa(t *testing.T) {

	// var slice []int
	// slice = append(slice, 1)

	// newSlice := AddElement(slice, 4)
	// fmt.Println(&slice[0] == &newSlice[0])

	// newSlice2 := AddElement(slice, 3)
	// fmt.Println(&slice[0] == &newSlice2[0])
	// newSlice = AddElement(slice, 2)
	// fmt.Println(&slice[0] == &newSlice[0])
	// newSlice = AddElement(slice, 4)
	// fmt.Println(&slice[0] == &newSlice[0])
	// newSlice = AddElement(slice, 4)
	// fmt.Println(&slice[0] == &newSlice[0])
	fmt.Println("begin", time.Now())
	tt := time.NewTicker(time.Second)
	<-time.After(3 * time.Second)
	go func() {
		for {
			fmt.Println("tt:", <-tt.C)
		}
	}()
	for {
	}
}

func TestXxxxss(t *testing.T) {
	sli := make([]int, 8)

	// sli = append(sli, 1)
	sli[0] = 1
	fmt.Println(sli, cap(sli))
	changeSli(sli)
	fmt.Println(sli, cap(sli))
}

func changeSli(s []int) {
	// ss := make([]int, 8)
	// s = ss
	s = append(s, 1)
	s = append(s, 1)
	s = append(s, 1)
	s = append(s, 1)
	s = append(s, 1)
	s = append(s, 1)
}

func TestXxxxxx(tt *testing.T) {
	m := [...]int{
		// 1: 2,
		'j': 1,
		// 'b': 2,
		// 'c': 3,
	}
	// m['a'] = 3
	fmt.Println(len(m))
}
