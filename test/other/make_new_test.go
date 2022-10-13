package other

import (
	"fmt"
	"testing"
)

func TestMain(m *testing.M) {
	a := new(chan int)
	*a = make(chan int, 1)
	// go func() {
	*a <- 1
	// }()
	fmt.Println(<-*a)
}

func TestMapChanAlias(t *testing.T) {
	mm := new(Maa)
	*mm = make(Maa, 3)

	fmt.Printf("%T\n", mm)
	(*mm)[0] = make(Aaa, 2)
	(*mm)[0] <- 1
	(*mm)[0] <- 2
	mm.PrinlnAll()
}

func TestMapChanNormalAlias(t *testing.T) {
	mm := new(chan int)
	*mm = make(chan int, 3)

	fmt.Printf("%T\n", mm)
	(*mm) <- 1
	(*mm) <- 2
}

func TestMapChanNonNormalAlias(t *testing.T) {
	// cc := make(*int64, 1)
}
