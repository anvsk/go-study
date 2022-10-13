package maptest

import (
	"fmt"
	"testing"
)

// key 的定位过程

/****
 key

 hash()		01010101(8)	00000	0001(B)

 高8位找到桶
 低8位在桶里找到keyIndex=tohash[i]

 比较key==keys[keyIndex]

不等继续找overflow

*****/
func TestXxx(t *testing.T) {
	// var m map[interface{}]interface{}
	// m = make(map[interface{}]interface{})
	// for i := 0; i < 10; i++ {
	// 	m[math.NaN()] = i
	// }
	// fmt.Println(m)
	// fmt.Println(m[math.NaN()])
	// // fmt.Println(&m[1])
	// for k, _ := range m {
	// 	delete(m, k)
	// 	fmt.Println(m[k])
	// }
	// var aa [2]*[]*int
	// tmp := 99
	// aaa := &[]*int{&tmp, &tmp, &tmp}
	// aa = [2]*[]*int{aaa, aaa}
	// fmt.Println(*(*aa[1])[0])
	// v := []int{1, 2, 3}
	// changeSli(v)
	// m := map[int]int{0: 1}
	// changeMap(m)
	// fmt.Println(v)
	// fmt.Println(m)
	// fmt.Println(155*18 - 140*18)
	fmt.Println(18.1 / 17.9)
	var aa byte
	_ = aa
}

func changeSli(p []int) {
	p[0] = 99
}

func changeMap(m map[int]int) {
	m[0] = 99
}
