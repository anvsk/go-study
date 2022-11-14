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

// n := 3
// k := 2

// 组合默写

func TestZuhedemo1(*testing.T) {
	n := 15
	k := 10
	res = [][]int{}
	diguidemo1(n, k, 1, []int{})
	fmt.Println(len(res))
}

func diguidemo1(n, k, start int, track []int) {
	if len(track) == k {
		tmp := make([]int, k)
		copy(tmp, track)
		res = append(res, tmp)
		return
	}

	for i := start; i < n+1; i++ {
		track = append(track, i)
		diguidemo1(n, k, i+1, track)
		track = track[:len(track)-1]
	}
}
