package question

import (
	"fmt"
	"time"
)

func main() {
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
