package other

import (
	"fmt"
	"testing"
)

type Aaa chan int
type Maa map[int]Aaa

func TestMapChan(t *testing.T) {
	mm := make(Maa, 3)

	fmt.Printf("%T\n", mm)
	mm[0] = make(Aaa, 2)
	mm[0] <- 1
	mm[0] <- 2
	mm.PrinlnAll()
}

func TestMapChanAlias2(t *testing.T) {
	mm := new(Maa)
	*mm = make(Maa, 3)

	fmt.Printf("%T\n", mm)
	(*mm)[0] = make(Aaa, 2)
	(*mm)[0] <- 1
	(*mm)[0] <- 2
	mm.PrinlnAll()
}

func TestMapChanNormalAlias2(t *testing.T) {
	mm := new(chan int)
	*mm = make(chan int, 3)

	fmt.Printf("%T\n", mm)
	(*mm) <- 1
	(*mm) <- 2
}

func TestChan(t *testing.T) {
	aa := make(Aaa, 3)
	fmt.Printf("%T\n", aa)
	aa <- 2
	aa <- 2
	aa <- 2
	aa.PrinlnAll()
}

func (cc Aaa) PrinlnAll() {
	for v := range cc {
		fmt.Println(v)
	}
}

func (mm Maa) PrinlnAll() {
	for k, v := range mm {
		fmt.Println("===", k, ":")
		for aachan := range v {
			fmt.Println(aachan)
		}
	}
}
