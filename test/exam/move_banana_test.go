package exam

import (
	"log"
	"strconv"
	"testing"
)

const key = "BANANA"

func TestMoveBanana(t *testing.T) {
	ss := "NAANAAXNABABYNNBZBANANA"
	log.Println(Solution(ss))
}

func TestFindMax(t *testing.T) {
	ss := -5000
	log.Println(Solution2(ss))
}

// 遍历去掉字符串中的某个单词
func Solution(s string) int {
	flag := true
	nums := 0
	for flag {
		tmpBan := []string{"B", "A", "N", "A", "N", "A"}
		tmpS := s
		for _, v := range s {
			// if exist --del key if len(tmp)==0 out
			if len(tmpBan) == 0 {
				goto GETONE
			}
			if inarray(tmpBan, string(v)) {
				tmpBan = delarray(tmpBan, string(v))
				tmpS = delstr(tmpS, string(v))
			}
		}
	GETONE:
		// cut s ,run next loop
		if len(tmpBan) == 0 { //full exist
			s = tmpS
			nums++
		} else {
			flag = false
			goto OUT
		}
	}
OUT:
	return nums
}

func inarray(s []string, n string) bool {
	for _, v := range s {
		if v == n {
			return true
		}
	}
	return false
}

func delarray(s []string, n string) (res []string) {
	res = make([]string, 0)
	flag := true
	for _, v := range s {
		if v == n && flag {
			flag = false
			continue
		}
		res = append(res, v)
	}
	return
}

func delstr(s string, n string) (res string) {
	flag := true
	for _, v := range s {
		if string(v) == n && flag {
			flag = false
			continue
		}
		res = res + string(v)
	}
	return
}

// 从一串数字中删除5，拿到最大的数

func Solution2(N int) int {
	str := strconv.Itoa(N)
	flag := false
	max := 0
	for k, v := range str {
		tmpint, _ := strconv.Atoi(string(v))
		if tmpint == 5 {
			tmpstr := delstr2(str, k)
			tmpnumber, _ := strconv.Atoi(tmpstr)
			if !flag || tmpnumber > max {
				max = tmpnumber
			}
			flag = true
		}
	}
	if !flag {
		max = N
	}
	return max
}

func delstr2(s string, k int) (res string) {
	flag := true
	for kk, v := range s {
		if kk == k && flag {
			flag = false
			continue
		}
		res = res + string(v)
	}
	return
}
