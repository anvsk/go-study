package slice

import (
	"fmt"
	"log"
	"testing"
)

func TestSxx222222x2(*testing.T) {
	s1 := []int{1, 2, 3, 4, 5, 6}
	s2 := s1[4:5]
	log.Println(s2)
	s2 = append(s2, 100)
	log.Println(s1)
	log.Println(s2)

	s1[4] = 555
	s2 = append(s2, 200)
	log.Println(s1)
	log.Println(s2)
	s1[4] = 666

	s1 = append(s1, 101)
	log.Println(s1)
	log.Println(s2)

	log.Println(s2)
}

func TestAppendByte(*testing.T) {
	type aa struct {
		c   []string
		a   string
		b   int
		aac map[string]interface{}
	}
	aai := aa{
		[]string{"1"},
		"1",
		0,
		make(map[string]interface{}, 0),
	}
	_ = aai
	// s1 := []byte{1}
	// log.Println(s1, cap(s1))

	// for i := 0; i < 1026; i++ {
	// 	s1 = append(s1, 1)
	// 	if i > 0 {
	// 		log.Println(i, cap(s1))
	// 	}
	// }

	s1 := []aa{aai}
	log.Println(s1, cap(s1))

	for i := 0; i < 1226; i++ {
		s1 = append(s1, aai)
		if i > 0 {
			log.Println(i, cap(s1))
		}
	}
}

func TestX1111xx(t *testing.T) {
	s1 := []int{0}
	fmt.Println(len(s1), cap(s1))
	for i := 0; i < 10; i++ {
		s1 = append(s1, i)
		fmt.Println(len(s1), cap(s1))
	}
	s2 := []byte{0}
	fmt.Println(len(s2), cap(s2))
	for i := 0; i < 10; i++ {
		s2 = append(s2, byte(i))
		fmt.Println(len(s2), cap(s2))
	}
	bools := []bool{true}
	fmt.Println(len(bools), cap(bools))
	for i := 0; i < 10; i++ {
		bools = append(bools, true)
		fmt.Println(len(bools), cap(bools))
	}
}

func TestXxxxx(t *testing.T) {
	fmt.Println(1<<8 == 256)
	fmt.Println(256>>8 == 1)
	fmt.Println(80>>3 == 10)

	a := []int{1, 2, 3, 4, 5}
	b := []int{6, 7, 8, 9, 10}
	c := append(a[2:3], []int{55, 55, 55}...)
	d := append(b[:3], 222)
	_ = c
	_ = d
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println("====")
	c = append(a[1:4], 111)
	d = append(b[2:], 333)
	fmt.Println(a)
	fmt.Println(b)
}

func TestXssss(t *testing.T) {
	s1 := make([]int, 0)
	s2 := []int{}
	var s3 []int
	fmt.Println(s1)
	fmt.Println(s2)
	fmt.Println(s2 == nil)
	fmt.Println(s3 == nil)

}
