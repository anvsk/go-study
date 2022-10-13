package funcs

import (
	"fmt"
	"testing"
)

const (
// A int  = 1<<iota
// B int
)

func TestXxx(t *testing.T) {
	sli := []byte{1}
	sli = append(sli, 1)
	sli = append(sli, []byte{1, 2, 3, 4, 5, 6, 7}...)
	_ = sli
	fmt.Println(sli)
}

func foo() (int, int) {
	return 3, 5
}

func TestMultiReturn(t *testing.T) {
	var a = 7
	b := 8
	c := 8
	if a, c = Get(); true {
		fmt.Println("()", a)
		fmt.Println("()", c)
	}
	fmt.Println(a)
	fmt.Println(b)
	// fmt.Println(c)
}

func Get() (int, int) {
	return 88, 99
}
