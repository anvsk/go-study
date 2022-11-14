package huisu

import (
	"fmt"
	"testing"
)

/**
给定两个整数 n 和 k，返回 1 ... n 中所有可能的 k 个数的组合。

示例:
输入: n = 4, k = 2
输出:
[
[2,4],
[3,4],
[2,3],
[1,2],
[1,3],
[1,4],
]

k==3
[
[2,4],
[3,4],
[2,3],
[1,2,3],
[1,2,4],
[1,3,4],
[1,3,2],
[1,4,2],
[1,4,3],
]

resolveColunSizes
#

n=3,k=2

1: 1|2|3
2: 2|3 1|3 1|2
3: 1,2+1,3 2,1+2,3 3,1+3,2
**/

// 0-n范围里,k个数的所有排列组合

func TestZuhe1(t *testing.T) {
	n := 15
	k := 10
	fmt.Println(len(combine(n, k)))
	// fmt.Println((combine(n, k)))
	return
}

var res [][]int

func combine(n int, k int) [][]int {
	res = [][]int{}
	if n <= 0 || k <= 0 || k > n {
		return res
	}
	backtrack(n, k, 1, []int{})
	return res
}

func backtrack(n, k, start int, track []int) {
	if len(track) == k {
		temp := make([]int, k)
		copy(temp, track)
		res = append(res, temp)
		return
	}
	for i := start; i <= n; i++ {
		if len(track)+n-start+1 < k {
			continue
		}
		// 依次加进去
		track = append(track, i)
		// base这个track进行递归
		backtrack(n, k, i+1, track)
		// 撤销,用下一个试
		track = track[:len(track)-1]
	}
	// fmt.Println("go out!")
}

// []int len: 2, cap: 2, [1,3]
