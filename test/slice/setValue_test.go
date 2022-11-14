package slice

import (
	"fmt"
	"testing"
)

func TestX(t *testing.T) {
	a := []int{1, 2, 3}
	// changeSli(a)
	b := a
	// b[0] = 99
	changeSli(b)
	fmt.Println(a)
}

func changeSli(a []int) {
	a[0] = 99
}
