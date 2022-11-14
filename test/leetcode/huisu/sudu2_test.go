package huisu

import (
	"fmt"
	"testing"
)

func TestSssss2222(t *testing.T) {
	sudu2()
}

func BenchmarkXxx(b *testing.B) {
	for i := 0; i < b.N; i++ {
		sudu2()
	}
}

// 对原始表格扫描，初始化各个哈希表。
// 对所有空格进行处理，得到每个空格的所有候选数字
// 处理候选数字个数为1的空格，之后转2继续执行。（2、3步可能循环多次）
// 采用回溯法，处理剩余的空格

// 依次单个九宫格处理即可,不用考虑相邻\交错的情况

func sudu2() {
	a := [][]string{

		{"5", "3", ".", ".", "7", ".", ".", ".", "."},
		{"6", ".", ".", "1", "9", "5", ".", ".", "."},
		{".", "9", "8", ".", ".", ".", ".", "6", "."},
		{"8", ".", ".", ".", "6", ".", ".", ".", "3"},
		{"4", ".", ".", "8", ".", "3", ".", ".", "1"},
		{"7", ".", ".", ".", "2", ".", ".", ".", "6"},
		{".", "6", ".", ".", ".", ".", "2", "8", "."},
		{".", ".", ".", "4", "1", "9", ".", ".", "5"},
		{".", ".", ".", ".", "8", ".", ".", "7", "9"},
	}
	orig := initArr2(a)
	solveSudoku(orig)
	// fmt.Println(orig)
	for _, v := range orig {
		fmt.Println(v)
	}
}

func solveSudoku(board [][]byte) {
	var line, column [9][9]bool
	var block [3][3][9]bool
	var spaces [][2]int

	for i, row := range board {
		for j, b := range row {
			if b == 0 {
				// 待填入的数组
				spaces = append(spaces, [2]int{i, j})
			} else {
				// 转成从0开始
				digit := b - 1
				// 行和列的bitset,存在这个数字设为true
				line[i][digit] = true
				column[j][digit] = true
				block[i/3][j/3][digit] = true
			}
		}
	}

	var dfs func(int) bool
	dfs = func(pos int) bool {
		if pos == len(spaces) {
			return true
		}
		// pos 待填入的序号,得出ij坐标
		i, j := spaces[pos][0], spaces[pos][1]
		// 从1-9开始试错
		for digit := byte(0); digit < 9; digit++ {
			// 如果当前阵列已经存在了就跳过
			if !line[i][digit] && !column[j][digit] && !block[i/3][j/3][digit] {
				// 填入假设值
				line[i][digit] = true
				column[j][digit] = true
				block[i/3][j/3][digit] = true
				board[i][j] = digit + 1
				// 一直递归下去
				if dfs(pos + 1) {
					return true
				}
				// 撤销假设值
				line[i][digit] = false
				column[j][digit] = false
				block[i/3][j/3][digit] = false
			}
		}
		return false
	}
	dfs(0)
}

// func Get() {
// 	var line, column [9][9]bool
// 	var block [3][3][9]bool
// 	var spaces [][2]int

// 	var dfs func(int) bool
// 	dfs = func(i int) bool {
// 		if pos == len(spaces) {
// 			return true
// 		}
// 		i, j := spaces[pos][0], space[pos][1]
// 		for digit := byte(0); digit < 9; digit++ {
// 			line[i][digit] = true
// 		}
// 		colimn[j][digit] = true
// 		block[i/3][j/3][digit] = true
// 		board[i][j] = digit + 1
// 		if dfs(pos + 1) {
// 			return true
// 		}
// 		line[i][digit] = false
// 		column[i][digit] = false
// 	}

// }
