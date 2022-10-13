package chans

import (
	"fmt"
	"time"
)

func Do() {
	withdrawEventChan := make(chan int)
	go func() {
		i := 0
		for {
			time.Sleep(3 * time.Second)
			withdrawEventChan <- i
			if i > 3 {
				break
			}
			i++
		}
		println("send done")

	}()
	go func() {
		for vLog := range withdrawEventChan {
			fmt.Println(vLog)
		}
		println("recei done")
	}()
	for {
	}
}
