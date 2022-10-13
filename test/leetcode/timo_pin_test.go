package leetcode

import (
	"fmt"
	"testing"
)

// 输入: [1,2], 2
// 输出: 3
// 原因: 第 1 秒初，提莫开始对艾希进行攻击并使其立即中毒。中毒状态会维持 2 秒钟，直到第 2 秒末结束。
// 但是第 2 秒初，提莫再次攻击了已经处于中毒状态的艾希。
// 由于中毒状态不可叠加，提莫在第 2 秒初的这次攻击会在第 3 秒末结束。
// 所以最终输出 3 。
func TestXX(t *testing.T) {
	n := TimoPinSeconds([]int{1, 2, 3, 7, 8, 99}, 2)
	fmt.Println(n)
}

func TimoPinSeconds(pinTiming []int, rate int) int {
	sumSeconds := 0
	for i := 0; i < len(pinTiming); i++ {
		if i+1 < len(pinTiming) {
			nextSecond := pinTiming[i+1]
			if pinTiming[i]+rate > nextSecond {
				sumSeconds += nextSecond - pinTiming[i]
				continue
			}
		}
		sumSeconds += rate

	}
	return sumSeconds
}
