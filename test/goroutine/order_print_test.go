package goroutine

import (
	"fmt"
	"testing"
)

// 顺序打印100次abc用channel
func TestOrderAbc(*testing.T) {
	sliChan := make([]chan int, 3)
	num := 0
	for i := 0; i < 3; i++ {
		sliChan[i] = make(chan int, 1)
	}
	for i := 0; i < 3; i++ {
		go func(i int) {
			if num == 0 && i == 0 {
				sliChan[0] <- 1
			}
			for {
				<-sliChan[i]
				dosomething(i)
				num++
				if num < 100 {
					if i == 2 {
						sliChan[0] <- 1
					} else {
						sliChan[i+1] <- 1
					}
				}
			}
		}(i)
	}
	for {
	}
}

func scheduleTask(){
	
}

func dosomething(i int) {
	switch i {
	case 0:
		fmt.Println("a")
	case 1:
		fmt.Println("b")
	case 2:
		fmt.Println("c")

	}
}
