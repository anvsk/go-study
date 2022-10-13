package chans

import (
	"context"
	"fmt"
	"testing"
	"time"
)

func TestMaxGo(tt *testing.T) {
	// ch := make(chan int, 10)
	// wg := sync.WaitGroup{}

	// for i := 0; i < 100; i++ {
	// 	ch <- i
	// 	wg.Add(1)
	// 	go func() {
	// 		fmt.Println(<-ch, "done")
	// 		wg.Done()
	// 	}()
	// }
	// wg.Wait()
	// aa := make(map[interface{}]interface{})
	// aa["11"] = 11
	// _ = aa
	// fmt.Println("11")

	// t := Aa{
	// 	1,
	// }
	// t.Set(2)
	// fmt.Println(t)

	ctx, _ := context.WithDeadline(context.Background(), time.Now().Add(3*time.Second))
	_ = ctx
	select {
	case <-ctx.Done():
		fmt.Println("deadline")
	}
}

type Aa struct {
	A int
}

func (t *Aa) Set(i int) {
	t.A = i
}

type student struct {
	id   int32
	name string
}

func TestXxx(t *testing.T) {
	// a := student{id: 1, name: "微客鸟窝"}
	// _ = a

	// // fmt.Printf("a=%v \n", a)  // a=&{1 微客鸟窝}
	// // fmt.Printf("a=%+v \n", a) // a=&{id:1 name:微客鸟窝}
	// // fmt.Printf("a=%#v \n", a) // a=&main.student{id:1, name:"微客鸟窝"}
	// // b := map[string]student{
	// // 	"aaaaaa": *a,
	// // }
	// // var b interface{}
	// b := student{}
	// _ = b
	// // fmt.Printf("%#v", a)
	aa := []rune{}
	for i := 200; i < 1000; i++ {
		aa = append(aa, rune(i))
	}
	fmt.Println(string(aa))
}
