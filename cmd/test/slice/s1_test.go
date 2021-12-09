package slice

import (
	"fmt"
	"testing"
)

func TestAaac(*testing.T) {
	a := []byte{1, 2, 3, 4, 5, 6}
	b := []byte{1, 2, 3, 4, 5, 6}
	// tmp := []byte{0, 0}
	fmt.Println(a)
	// fmt.Println(a[len(a)-2])
	println(a)
	// res := append(a[len(a)-3:len(a)-2], tmp...)
	res := append(a[len(a)-4:len(a)-2], b[len(a)-4:len(a)-2]...)
	println(a)

	fmt.Println("res", res)
	fmt.Println("a", a)
}
