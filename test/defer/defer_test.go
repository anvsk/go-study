package defer_test

import (
	"fmt"
	"testing"
)

func TestXxxx11(*testing.T) {
	fmt.Println("====")
	defer func() {
		if err := recover(); err != nil {
			fmt.Println(err)
		} else {
			fmt.Println("fatal")
		}
	}()

	defer func() {
		panic("defer panic")
	}()
	fmt.Println("---")

	panic("panic")

}

func TestXxxx111(*testing.T) {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("1===", err)
		} else {
			fmt.Println("noerr1")
		}
	}()

	defer func() {
		if err := recover(); err != nil {
			fmt.Println("2====", err)
			panic("2 de error")
		} else {
			fmt.Println("noerr2")
		}
	}()
	panic("3333")
}

func function(index int, value int) int {

	fmt.Println(index)

	return index
}

func TestXsssss(t *testing.T) {
	i := 88
	defer function(1, function(i, 0))
	i = 99
	defer function(2, function(i, 0))
	i = 100
}

func DeferFunc2(i int) int {
	t := i
	defer func() {
		t += 3
	}()
	return t
}

func TestXxx(t *testing.T) {
	fmt.Println(DeferFunc2(1))
}
