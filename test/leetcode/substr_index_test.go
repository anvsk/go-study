package leetcode

import (
	"fmt"
	"testing"

	"github.com/jinzhu/copier"
)

/**
给定一个字符串 s 和一个字符串数组 words。 words 中所有字符串 #长度相同。

 s 中的 串联子串 是指一个包含  words 中所有字符串以任意顺序排列连接起来的子串。

    例如，如果 words = ["ab","cd","ef"]， 那么 "abcdef"， "abefcd"，"cdabef"， "cdefab"，"efabcd"， 和 "efcdab" 都是串联子串。 "acdbef" 不是串联子串，因为他不是任何 words 排列的连接。

返回所有串联字串在 s 中的开始索引。你可以以 任意顺序 返回答案。
**/

func TestSubstrIndex(t *testing.T) {
	s := "barfoofoobarthebarfoobarman"
	words := []string{"bar", "bar", "the"}
	res2 := d2(s, words)
	fmt.Println(res2)

	res := d1(s, words)
	fmt.Println(res)
}

func BenchmarkDeleteSlice1(b *testing.B) {
	s := "barfoofoobarthebarfoobarmanbarfoofoobarthebarfoobarmanbarfoofoobarthebarfoobarmanbarfoofoobarthebarfoobarmanbarfoofoobarthebarfoobarmanbarfoofoobarthebarfoobarmanbarfoofoobarthebarfoobarmanbarfoofoobarthebarfoobarmanbarfoofoobarthebarfoobarmanbarfoofoobarthebarfoobarmanbarfoofoobarthebarfoobarmanbarfoofoobarthebarfoobarmanbarfoofoobarthebarfoobarmanbarfoofoobarthebarfoobarmanbarfoofoobarthebarfoobarmanbarfoofoobarthebarfoobarmanbarfoofoobarthebarfoobarmanbarfoofoobarthebarfoobarmanbarfoofoobarthebarfoobarmanbarfoofoobarthebarfoobarmanbarfoofoobarthebarfoobarmanbarfoofoobarthebarfoobarmanbarfoofoobarthebarfoobarmanbarfoofoobarthebarfoobarmanbarfoofoobarthebarfoobarmanbarfoofoobarthebarfoobarmanbarfoofoobarthebarfoobarmanbarfoofoobarthebarfoobarmanbarfoofoobarthebarfoobarmanbarfoofoobarthebarfoobarmanbarfoofoobarthebarfoobarmanbarfoofoobarthebarfoobarmanbarfoofoobarthebarfoobarmanbarfoofoobarthebarfoobarmanbarfoofoobarthebarfoobarmanbarfoofoobarthebarfoobarmanbarfoofoobarthebarfoobarmanbarfoofoobarthebarfoobarmanbarfoofoobarthebarfoobarmanbarfoofoobarthebarfoobarmanbarfoofoobarthebarfoobarmanbarfoofoobarthebarfoobarmanbarfoofoobarthebarfoobarmanbarfoofoobarthebarfoobarmanbarfoofoobarthebarfoobarmanbarfoofoobarthebarfoobarmanbarfoofoobarthebarfoobarmanbarfoofoobarthebarfoobarmanbarfoofoobarthebarfoobarmanbarfoofoobarthebarfoobarmanbarfoofoobarthebarfoobarmanbarfoofoobarthebarfoobarmanbarfoofoobarthebarfoobarmanbarfoofoobarthebarfoobarmanbarfoofoobarthebarfoobarmanbarfoofoobarthebarfoobarmanbarfoofoobarthebarfoobarmanbarfoofoobarthebarfoobarmanbarfoofoobarthebarfoobarmanbarfoofoobarthebarfoobarmanbarfoofoobarthebarfoobarmanbarfoofoobarthebarfoobarmanbarfoofoobarthebarfoobarmanbarfoofoobarthebarfoobarmanbarfoofoobarthebarfoobarmanbarfoofoobarthebarfoobarmanbarfoofoobarthebarfoobarmanbarfoofoobarthebarfoobarmanbarfoofoobarthebarfoobarmanbarfoofoobarthebarfoobarmanbarfoofoobarthebarfoobarmanbarfoofoobarthebarfoobarmanbarfoofoobarthebarfoobarmanbarfoofoobarthebarfoobarmanbarfoofoobarthebarfoobarmanbarfoofoobarthebarfoobarmanbarfoofoobarthebarfoobarmanbarfoofoobarthebarfoobarmanbarfoofoobarthebarfoobarmanbarfoofoobarthebarfoobarmanbarfoofoobarthebarfoobarmanbarfoofoobarthebarfoobarmanbarfoofoobarthebarfoobarmanbarfoofoobarthebarfoobarmanbarfoofoobarthebarfoobarmanbarfoofoobarthebarfoobarmanbarfoofoobarthebarfoobarmanbarfoofoobarthebarfoobarmanbarfoofoobarthebarfoobarmanbarfoofoobarthebarfoobarmanbarfoofoobarthebarfoobarmanbarfoofoobarthebarfoobarmanbarfoofoobarthebarfoobarmanbarfoofoobarthebarfoobarmanbarfoofoobarthebarfoobarmanbarfoofoobarthebarfoobarmanbarfoofoobarthebarfoobarmanbarfoofoobarthebarfoobarmanbarfoofoobarthebarfoobarmanbarfoofoobarthebarfoobarmanbarfoofoobarthebarfoobarmanbarfoofoobarthebarfoobarmanbarfoofoobarthebarfoobarmanbarfoofoobarthebarfoobarmanbarfoofoobarthebarfoobarmanbarfoofoobarthebarfoobarmanbarfoofoobarthebarfoobarmanbarfoofoobarthebarfoobarmanbarfoofoobarthebarfoobarmanbarfoofoobarthebarfoobarmanbarfoofoobarthebarfoobarmanbarfoofoobarthebarfoobarmanbarfoofoobarthebarfoobarmanbarfoofoobarthebarfoobarmanbarfoofoobarthebarfoobarmanbarfoofoobarthebarfoobarmanbarfoofoobarthebarfoobarmanbarfoofoobarthebarfoobarmanbarfoofoobarthebarfoobarmanbarfoofoobarthebarfoobarmanbarfoofoobarthebarfoobarmanbarfoofoobarthebarfoobarmanbarfoofoobarthebarfoobarmanbarfoofoobarthebarfoobarmanbarfoofoobarthebarfoobarmanbarfoofoobarthebarfoobarmanbarfoofoobarthebarfoobarmanbarfoofoobarthebarfoobarmanbarfoofoobarthebarfoobarmanbarfoofoobarthebarfoobarmanbarfoofoobarthebarfoobarmanbarfoofoobarthebarfoobarmanbarfoofoobarthebarfoobarmanbarfoofoobarthebarfoobarmanbarfoofoobarthebarfoobarmanbarfoofoobarthebarfoobarmanbarfoofoobarthebarfoobarmanbarfoofoobarthebarfoobarmanbarfoofoobarthebarfoobarmanbarfoofoobarthebarfoobarmanbarfoofoobarthebarfoobarmanbarfoofoobarthebarfoobarmanbarfoofoobarthebarfoobarmanbarfoofoobarthebarfoobarmanbarfoofoobarthebarfoobarmanbarfoofoobarthebarfoobarmanbarfoofoobarthebarfoobarmanbarfoofoobarthebarfoobarmanbarfoofoobarthebarfoobarmanbarfoofoobarthebarfoobarmanbarfoofoobarthebarfoobarmanbarfoofoobarthebarfoobarmanbarfoofoobarthebarfoobarmanbarfoofoobarthebarfoobarmanbarfoofoobarthebarfoobarmanbarfoofoobarthebarfoobarmanbarfoofoobarthebarfoobarmanbarfoofoobarthebarfoobarmanbarfoofoobarthebarfoobarmanbarfoofoobarthebarfoobarmanbarfoofoobarthebarfoobarmanbarfoofoobarthebarfoobarmanbarfoofoobarthebarfoobarmanbarfoofoobarthebarfoobarmanbarfoofoobarthebarfoobarmanbarfoofoobarthebarfoobarmanbarfoofoobarthebarfoobarmanbarfoofoobarthebarfoobarmanbarfoofoobarthebarfoobarmanbarfoofoobarthebarfoobarmanbarfoofoobarthebarfoobarmanbarfoofoobarthebarfoobarmanbarfoofoobarthebarfoobarmanbarfoofoobarthebarfoobarmanbarfoofoobarthebarfoobarmanbarfoofoobarthebarfoobarmanbarfoofoobarthebarfoobarmanbarfoofoobarthebarfoobarmanbarfoofoobarthebarfoobarmanbarfoofoobarthebarfoobarmanbarfoofoobarthebarfoobarmanbarfoofoobarthebarfoobarmanbarfoofoobarthebarfoobarmanbarfoofoobarthebarfoobarmanbarfoofoobarthebarfoobarmanbarfoofoobarthebarfoobarmanbarfoofoobarthebarfoobarmanbarfoofoobarthebarfoobarmanbarfoofoobarthebarfoobarmanbarfoofoobarthebarfoobarmanbarfoofoobarthebarfoobarmanbarfoofoobarthebarfoobarmanbarfoofoobarthebarfoobarmanbarfoofoobarthebarfoobarmanbarfoofoobarthebarfoobarmanbarfoofoobarthebarfoobarmanbarfoofoobarthebarfoobarmanbarfoofoobarthebarfoobarmanbarfoofoobarthebarfoobarmanbarfoofoobarthebarfoobarmanbarfoofoobarthebarfoobarmanbarfoofoobarthebarfoobarmanbarfoofoobarthebarfoobarmanbarfoofoobarthebarfoobarmanbarfoofoobarthebarfoobarmanbarfoofoobarthebarfoobarmanbarfoofoobarthebarfoobarmanbarfoofoobarthebarfoobarmanbarfoofoobarthebarfoobarmanbarfoofoobarthebarfoobarmanbarfoofoobarthebarfoobarmanbarfoofoobarthebarfoobarmanbarfoofoobarthebarfoobarmanbarfoofoobarthebarfoobarmanbarfoofoobarthebarfoobarmanbarfoofoobarthebarfoobarman"
	words := []string{"bar", "foo", "the", "the", "arf", "oob", "arm"}
	for i := 0; i < b.N; i++ {
		_ = d1(s, words)
	}
}

// 885878	      1283 ns/op	     656 B/op	      26 allocs/op
// 891974	      1290 ns/op	     656 B/op	      26 allocs/op
// 864404	      1289 ns/op	     656 B/op	      26 allocs/op

// 10000	    103571 ns/op	    9704 B/op	      64 allocs/op
// 10000	    103105 ns/op	    9704 B/op	      64 allocs/op
// 10000	    103731 ns/op	    9704 B/op	      64 allocs/op

// 96900	     11951 ns/op	    6136 B/op	     263 allocs/op
// 96840	     11917 ns/op	    6136 B/op	     263 allocs/op
// 100603	     11861 ns/op	    6136 B/op	     263 allocs/op

// 100458	     11904 ns/op	    6136 B/op	     263 allocs/op
// 99656	     12072 ns/op	    6136 B/op	     263 allocs/op
// 99620	     12109 ns/op	    6136 B/op	     263 allocs/op
func BenchmarkDeleteSlice2(b *testing.B) {
	s := "barfoofoobarthebarfoobarmanbarfoofoobarthebarfoobarmanbarfoofoobarthebarfoobarmanbarfoofoobarthebarfoobarmanbarfoofoobarthebarfoobarmanbarfoofoobarthebarfoobarmanbarfoofoobarthebarfoobarmanbarfoofoobarthebarfoobarmanbarfoofoobarthebarfoobarmanbarfoofoobarthebarfoobarmanbarfoofoobarthebarfoobarmanbarfoofoobarthebarfoobarmanbarfoofoobarthebarfoobarmanbarfoofoobarthebarfoobarmanbarfoofoobarthebarfoobarmanbarfoofoobarthebarfoobarmanbarfoofoobarthebarfoobarmanbarfoofoobarthebarfoobarmanbarfoofoobarthebarfoobarmanbarfoofoobarthebarfoobarmanbarfoofoobarthebarfoobarmanbarfoofoobarthebarfoobarmanbarfoofoobarthebarfoobarmanbarfoofoobarthebarfoobarmanbarfoofoobarthebarfoobarmanbarfoofoobarthebarfoobarmanbarfoofoobarthebarfoobarmanbarfoofoobarthebarfoobarmanbarfoofoobarthebarfoobarmanbarfoofoobarthebarfoobarmanbarfoofoobarthebarfoobarmanbarfoofoobarthebarfoobarmanbarfoofoobarthebarfoobarmanbarfoofoobarthebarfoobarmanbarfoofoobarthebarfoobarmanbarfoofoobarthebarfoobarmanbarfoofoobarthebarfoobarmanbarfoofoobarthebarfoobarmanbarfoofoobarthebarfoobarmanbarfoofoobarthebarfoobarmanbarfoofoobarthebarfoobarmanbarfoofoobarthebarfoobarmanbarfoofoobarthebarfoobarmanbarfoofoobarthebarfoobarmanbarfoofoobarthebarfoobarmanbarfoofoobarthebarfoobarmanbarfoofoobarthebarfoobarmanbarfoofoobarthebarfoobarmanbarfoofoobarthebarfoobarmanbarfoofoobarthebarfoobarmanbarfoofoobarthebarfoobarmanbarfoofoobarthebarfoobarmanbarfoofoobarthebarfoobarmanbarfoofoobarthebarfoobarmanbarfoofoobarthebarfoobarmanbarfoofoobarthebarfoobarmanbarfoofoobarthebarfoobarmanbarfoofoobarthebarfoobarmanbarfoofoobarthebarfoobarmanbarfoofoobarthebarfoobarmanbarfoofoobarthebarfoobarmanbarfoofoobarthebarfoobarmanbarfoofoobarthebarfoobarmanbarfoofoobarthebarfoobarmanbarfoofoobarthebarfoobarmanbarfoofoobarthebarfoobarmanbarfoofoobarthebarfoobarmanbarfoofoobarthebarfoobarmanbarfoofoobarthebarfoobarmanbarfoofoobarthebarfoobarmanbarfoofoobarthebarfoobarmanbarfoofoobarthebarfoobarmanbarfoofoobarthebarfoobarmanbarfoofoobarthebarfoobarmanbarfoofoobarthebarfoobarmanbarfoofoobarthebarfoobarmanbarfoofoobarthebarfoobarmanbarfoofoobarthebarfoobarmanbarfoofoobarthebarfoobarmanbarfoofoobarthebarfoobarmanbarfoofoobarthebarfoobarmanbarfoofoobarthebarfoobarmanbarfoofoobarthebarfoobarmanbarfoofoobarthebarfoobarmanbarfoofoobarthebarfoobarmanbarfoofoobarthebarfoobarmanbarfoofoobarthebarfoobarmanbarfoofoobarthebarfoobarmanbarfoofoobarthebarfoobarmanbarfoofoobarthebarfoobarmanbarfoofoobarthebarfoobarmanbarfoofoobarthebarfoobarmanbarfoofoobarthebarfoobarmanbarfoofoobarthebarfoobarmanbarfoofoobarthebarfoobarmanbarfoofoobarthebarfoobarmanbarfoofoobarthebarfoobarmanbarfoofoobarthebarfoobarmanbarfoofoobarthebarfoobarmanbarfoofoobarthebarfoobarmanbarfoofoobarthebarfoobarmanbarfoofoobarthebarfoobarmanbarfoofoobarthebarfoobarmanbarfoofoobarthebarfoobarmanbarfoofoobarthebarfoobarmanbarfoofoobarthebarfoobarmanbarfoofoobarthebarfoobarmanbarfoofoobarthebarfoobarmanbarfoofoobarthebarfoobarmanbarfoofoobarthebarfoobarmanbarfoofoobarthebarfoobarmanbarfoofoobarthebarfoobarmanbarfoofoobarthebarfoobarmanbarfoofoobarthebarfoobarmanbarfoofoobarthebarfoobarmanbarfoofoobarthebarfoobarmanbarfoofoobarthebarfoobarmanbarfoofoobarthebarfoobarmanbarfoofoobarthebarfoobarmanbarfoofoobarthebarfoobarmanbarfoofoobarthebarfoobarmanbarfoofoobarthebarfoobarmanbarfoofoobarthebarfoobarmanbarfoofoobarthebarfoobarmanbarfoofoobarthebarfoobarmanbarfoofoobarthebarfoobarmanbarfoofoobarthebarfoobarmanbarfoofoobarthebarfoobarmanbarfoofoobarthebarfoobarmanbarfoofoobarthebarfoobarmanbarfoofoobarthebarfoobarmanbarfoofoobarthebarfoobarmanbarfoofoobarthebarfoobarmanbarfoofoobarthebarfoobarmanbarfoofoobarthebarfoobarmanbarfoofoobarthebarfoobarmanbarfoofoobarthebarfoobarmanbarfoofoobarthebarfoobarmanbarfoofoobarthebarfoobarmanbarfoofoobarthebarfoobarmanbarfoofoobarthebarfoobarmanbarfoofoobarthebarfoobarmanbarfoofoobarthebarfoobarmanbarfoofoobarthebarfoobarmanbarfoofoobarthebarfoobarmanbarfoofoobarthebarfoobarmanbarfoofoobarthebarfoobarmanbarfoofoobarthebarfoobarmanbarfoofoobarthebarfoobarmanbarfoofoobarthebarfoobarmanbarfoofoobarthebarfoobarmanbarfoofoobarthebarfoobarmanbarfoofoobarthebarfoobarmanbarfoofoobarthebarfoobarmanbarfoofoobarthebarfoobarmanbarfoofoobarthebarfoobarmanbarfoofoobarthebarfoobarmanbarfoofoobarthebarfoobarmanbarfoofoobarthebarfoobarmanbarfoofoobarthebarfoobarmanbarfoofoobarthebarfoobarmanbarfoofoobarthebarfoobarmanbarfoofoobarthebarfoobarmanbarfoofoobarthebarfoobarmanbarfoofoobarthebarfoobarmanbarfoofoobarthebarfoobarmanbarfoofoobarthebarfoobarmanbarfoofoobarthebarfoobarmanbarfoofoobarthebarfoobarmanbarfoofoobarthebarfoobarmanbarfoofoobarthebarfoobarmanbarfoofoobarthebarfoobarmanbarfoofoobarthebarfoobarmanbarfoofoobarthebarfoobarmanbarfoofoobarthebarfoobarmanbarfoofoobarthebarfoobarmanbarfoofoobarthebarfoobarmanbarfoofoobarthebarfoobarmanbarfoofoobarthebarfoobarmanbarfoofoobarthebarfoobarmanbarfoofoobarthebarfoobarmanbarfoofoobarthebarfoobarmanbarfoofoobarthebarfoobarmanbarfoofoobarthebarfoobarmanbarfoofoobarthebarfoobarmanbarfoofoobarthebarfoobarmanbarfoofoobarthebarfoobarmanbarfoofoobarthebarfoobarmanbarfoofoobarthebarfoobarmanbarfoofoobarthebarfoobarmanbarfoofoobarthebarfoobarmanbarfoofoobarthebarfoobarmanbarfoofoobarthebarfoobarmanbarfoofoobarthebarfoobarmanbarfoofoobarthebarfoobarmanbarfoofoobarthebarfoobarmanbarfoofoobarthebarfoobarmanbarfoofoobarthebarfoobarmanbarfoofoobarthebarfoobarmanbarfoofoobarthebarfoobarmanbarfoofoobarthebarfoobarmanbarfoofoobarthebarfoobarmanbarfoofoobarthebarfoobarmanbarfoofoobarthebarfoobarmanbarfoofoobarthebarfoobarmanbarfoofoobarthebarfoobarmanbarfoofoobarthebarfoobarmanbarfoofoobarthebarfoobarmanbarfoofoobarthebarfoobarmanbarfoofoobarthebarfoobarmanbarfoofoobarthebarfoobarmanbarfoofoobarthebarfoobarmanbarfoofoobarthebarfoobarmanbarfoofoobarthebarfoobarmanbarfoofoobarthebarfoobarmanbarfoofoobarthebarfoobarman"
	words := []string{"bar", "foo", "the", "the", "arf", "oob", "arm"}
	for i := 0; i < b.N; i++ {
		_ = d2(s, words)
	}
}

// 1. 找出所有的排列组合
// 取指定长度的串进行比对
// 存到map里面,然后取len(arr)长度的串比对

// 取指定长度的串,轮训比对,一样则tmparr删除
// 在紧接着做同样操作
// 成功次数==len(arr) 记录index,否则 index++重复

func d1(s string, words []string) []int {
	arr := divideArrs(words)
	m := make(map[string]struct{})
	_ = m
	clauseLen := len(words) * len(words[0])
	res := []int{}
	// for _, v := range arr {
	// 	clauseLen = len(v)
	// 	m[v] = struct{}{}
	// }
	for i := 0; i < len(s); i++ {
		startIndex := i
		endIndex := i + clauseLen
		if endIndex >= len(s) {
			break
		}
		curstr := s[startIndex:endIndex]
		// if _, ok := m[curstr]; ok {
		// 	res = append(res, i)
		// }
		for _, v := range arr {
			if v == curstr {
				res = append(res, i)
			}
		}
	}
	return res
}

func d2(s string, words []string) []int {
	res := []int{}
	m := make(map[string]struct{})
	for _, v := range words {
		m[v] = struct{}{}
	}
	clauseLen := len(words[0]) * len(words)
	wordLen := len(words[0])
	for i := 0; i < len(s); i++ {
		startIndex := i
		endIndex := i + clauseLen
		if endIndex >= len(s) {
			break
		}
		tmpM := []string{}
		copier.Copy(&tmpM, words)
		curstr := s[startIndex:endIndex]
		for j := 0; j < len(words); j++ {
			wordStart := j * wordLen
			wordEnd := j*wordLen + wordLen
			curword := curstr[wordStart:wordEnd]
			flag := false
			for _, v := range tmpM {
				if v == curword {
					flag = true
					tmpM = deletes(tmpM, v)
					goto OUT2
				}
			}
		OUT2:
			if !flag {
				goto OUT
			}
		}
		res = append(res, i)
	OUT:
	}
	return res
}

func deletes(sli []string, s string) []string {
	t := []string{}
	flag := false
	for _, v := range sli {
		if v == s && !flag {
			flag = true
		} else {
			t = append(t, v)
		}
	}
	return t
}

// 传入一个单词数组,返回所有可能组成的句子
func divideArrs(words []string) (clause []string) {
	if len(words) == 0 {
		return []string{}
	}
	// 存放所有结果组合
	allwords := [][]string{words}
	for i := 0; i < len(words)-1; i++ {
		for _, v := range allwords {
			for j := i + 1; j < len(words); j++ {
				if v[i] == v[j] {
					continue
				}
				var tmp []string
				copier.Copy(&tmp, v)
				tmp[i], tmp[j] = tmp[j], tmp[i]
				allwords = append(allwords, tmp)
			}
		}
	}
	res := []string{}
	for _, v := range allwords {
		tmp := ""
		for _, vv := range v {
			tmp += vv
		}
		res = append(res, tmp)
	}
	return res
}

func TestXxx(t *testing.T) {
	m := []int{1, 1, 2, 3, 4, 5, 5}
	n := DeleteSlice4(m, 1)
	fmt.Println(n)
}

// DeleteSlice4 删除指定元素。
func DeleteSlice4(a []int, elem int) []int {
	tgt := a[:0]
	for _, v := range a {
		if v != elem {
			tgt = append(tgt, v)
		}
	}
	return tgt
}
