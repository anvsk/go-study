package single_flight

import (
	"fmt"
	"testing"
	"time"
)

func TestGet(t *testing.T) {
	InitEnv()
	for i := 0; i < 10; i++ {
		go func(i int) {
			fmt.Println(i, Get(i, "1"))
			fmt.Println(2000+i, Get(2000+i, "2"))
		}(i)
	}
	<-time.After(2 * time.Second)
}
