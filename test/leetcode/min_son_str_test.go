package leetcode

import (
	"log"
	"testing"
)

// 给你一个字符串 s 、一个字符串 t 。返回 s 中涵盖 t 所有字符的最小子串。如果 s 中不存在涵盖 t 所有字符的子串，则返回空字符串 “” 。

// 注意：如果 s 中存在这样的子串，我们保证它是唯一的答案。

// 示例 1：

// 输入：s = "ADOBECODEBANC", t = "ABC"
// 输出："BANC"

func TestMinCoverSonStr(t *testing.T) {
	s := MinCoverSonStr("ADOBECODEBANC", "ABC")
	log.Println(s)

}

func MinCoverSonStr(s, t string) string {
	match := make(map[byte]byte)
	for _, v := range []byte(t) {
		match[v] = v
	}
	minStrSli := []byte{}
	for kk, vv := range []byte(s) {
		if _, ok := match[vv]; !ok {
			continue
		}
		log.Println("begin match with", vv)
		matchedStrSli := []byte{vv}
		todoSearchMap := make(map[byte]byte)
		for _, v := range match {
			todoSearchMap[v] = v
		}
		delete(todoSearchMap, vv)
		for i := kk + 1; i < len(s); i++ {
			v := s[i]
			if _, ok := todoSearchMap[v]; ok {
				delete(todoSearchMap, v)
			}
			matchedStrSli = append(matchedStrSli, v)
			if len(todoSearchMap) == 0 {
				log.Println("find all ", matchedStrSli, "next loop")
				if len(minStrSli) == 0 || len(matchedStrSli) < len(minStrSli) {
					minStrSli = matchedStrSli
				}
				goto NEXT_LOOP
			}

		}
	NEXT_LOOP:
	}
	return string(minStrSli)
}
