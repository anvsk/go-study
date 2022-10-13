package leetcode

import (
	"log"
	"testing"
)

func integerBreak2(n int) int {
	if n < 2 {
		return 0
	}
	if n == 2 {
		return 1
	}
	if n == 3 {
		return 2
	}
	// int [] dp = new int[n+1];
	dp := make([]int, n+1)
	dp[2] = 2
	dp[3] = 3
	// for(int i=2;i+2<=n;i++){
	// 	dp[i+2] = Math.max(dp[i+2], 2 * dp[i]);
	// }
	for i := 2; i+2 <= n; i++ {
		dp[i+2] = max(dp[i+2], 2*dp[i])
	}
	for i := 2; i+3 <= n; i++ {
		dp[i+2] = max(dp[i+2], 2*dp[i])
	}
	// for(int i=2;i+3<=n;i++){
	// 	dp[i+3] = Math.max(dp[i+3], 3 * dp[i]);
	// }
	return dp[n]
}

// 和 70 题相似，但需要检查并嵌套两层循环
// 注意 bottom-up 向上存储值时，要取 maxP 和 i 的最大值
func integerBreak3(n int) int {
	ps := make([]int, n+1)
	ps[1] = 1 // 边界值

	for i := 2; i <= n; i++ {
		var maxP int
		for j := 1; j <= i/2; j++ {
			maxP = max(maxP, ps[j]*ps[i-j]) // 状态转移
		}
		if i == n {
			return maxP // bingo
		}

		ps[i] = max(maxP, i) // 最优子结构 // 存储中间计算结果
	}

	return -1 // 代码不会执行到这
}

func TestXxx22222(t *testing.T) {
	log.Println(integerBreak2(10))
	log.Println(cm(10))
}

// func CalcMaxMul(n int) int {
// 	po := make(map[int]int)
// 	po[1] = 1
// 	maxi := 0
// 	for i := 2; i <= n; i++ { //把前面的计算出来
// 		for j := 1; j <= i/2; j++ { //两两乘积
// 			maxi = max(maxi, po[j]*po[i-j])
// 		}
// 		po[i] = max(i, maxi)
// 	}
// 	return maxi
// }

func cm(n int) int {
	po := make(map[int]int)
	po[1] = 1
	po[2] = 1
	maxi := 1
	for i := 3; i <= n; i++ {
		for j := 1; j <= i/2; j++ {
			maxi = max(maxi, po[j]*po[i-j])
		}
		po[i] = max(maxi, i)
	}
	return maxi
}
