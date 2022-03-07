package slice

import (
	"log"
	"testing"
)

func TestSxxx2(*testing.T) {
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
	s1 := []byte{1}
	log.Println(s1, cap(s1))

	s1 = append(s1, 2)
	log.Println(s1, cap(s1))
}
