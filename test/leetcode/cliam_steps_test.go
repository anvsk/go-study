package leetcode

import (
	"log"
	"sort"
	"testing"
)

func TestCliamStep(t *testing.T) {
	log.Println(climbStairs3(6))
	log.Println(climbStairs3(5))
	log.Println(climbStairs3(4))
	log.Println(climbStairs3(3))
}

func TestInterBreak(t *testing.T) {
	log.Println(integerBreak(10))
}

// 和 70 题相似，但需要检查并嵌套两层循环
// 注意 bottom-up 向上存储值时，要取 maxP 和 i 的最大值
func integerBreak(n int) int {
	ps := make([]int, n+1)
	ps[1] = 1 // 边界值

	times := 0

	for i := 2; i <= n; i++ {
		var maxP int
		for j := 1; j <= i/2; j++ {
			maxP = max(maxP, ps[j]*ps[i-j]) // 状态转移
			times++
		}
		if i == n {
			// return times
			return maxP // bingo
		}

		ps[i] = max(maxP, i) // 最优子结构 // 存储中间计算结果
	}

	return -1 // 代码不会执行到这
}

func max(n, m int) int {
	if n > m {
		return n
	}
	return m
}

func climbStairs3(n int) int {
	if n == 1 {
		return 1
	}

	i, j := 1, 2
	for n > 2 {
		i, j = j, i+j // bottom -> up
		n--
	}
	return j
}

func topKFrequent(nums []int, k int) []int {
	mmap := map[int]int{}
	for _, v := range nums {
		mmap[v]++
	}
	type aaaa struct {
		n    int
		nums int
	}
	con := make([]aaaa, 0)
	for k, v := range mmap {
		con = append(con, aaaa{k, v})
	}

	sort.Slice(con, func(i, j int) bool {
		return con[i].nums > con[j].nums
	})

	res := make([]int, 0)
	for _, v := range con[:k] {
		res = append(res, v.n)
	}
	return res
}

func TestTopKFrequent(t *testing.T) {
	res := topKFrequent([]int{2, 2, 4444, 333, 333, 1, 1, 1, 1, 1, 3, 23, 45, 4, 5, 45, 45}, 4)
	log.Println(res)
}
