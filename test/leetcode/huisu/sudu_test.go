// 数独
// https://leetcode.cn/problems/sudoku-solver/
package huisu

import (
	"strconv"
	"testing"
)

func TestSssss(t *testing.T) {
	sudu()
}

// 对原始表格扫描，初始化各个哈希表。
// 对所有空格进行处理，得到每个空格的所有候选数字
// 处理候选数字个数为1的空格，之后转2继续执行。（2、3步可能循环多次）
// 采用回溯法，处理剩余的空格

// 依次单个九宫格处理即可,不用考虑相邻\交错的情况

var 

func sudu() {
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
	_ = a
	orig := initArr(a)
	// 可能的数量==>坐标
	h1 := map[int][2]int{}
	// 保存阵列每个位置可能出现的数字
	couldArr := [9][9][]int{}
	for rowk, row := range orig {
		for colk, num := range row {
			if num != 0 {
				continue
			}
			// 取可能出现的数字,存入hash表
			could := getNums(orig, rowk, colk)
			if len(could) == 0 {
				panic("grid error!")
			}
			// 只有一个时,直接写入
			if len(could) == 1 {
				orig[rowk][colk] = could[0]
			} else {
				couldArr[rowk][colk] = could
			}
			h1[len(could)] = [2]int{rowk, colk}
		}
	}

	// 回溯法
	i := 2
	for i < 9 {
		// todo

		i++
	}
}

func initArr(s [][]string) [][]int {
	res := [][]int{}
	for _, v := range s {
		tmp := []int{}
		for _, vv := range v {
			vvv := 0
			if vv != "." {
				vvv, _ = strconv.Atoi(vv)
			}
			tmp = append(tmp, (vvv))
		}
		res = append(res, tmp)
	}
	return res
}

func initArr2(s [][]string) [][]byte {
	res := [][]byte{}
	for _, v := range s {
		tmp := []byte{}
		for _, vv := range v {
			vvv := 0
			if vv != "." {
				vvv, _ = strconv.Atoi(vv)
			}
			tmp = append(tmp, byte(vvv))
		}
		res = append(res, tmp)
	}
	return res
}

// 求差集
func GetDiff(exist []int) (could []int) {
	base := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	for _, v := range base {
		for _, vv := range exist {
			if v == vv {
				goto OUT
			}
		}
		could = append(could, v)
	OUT:
	}
	return
}

// 获取9981阵列里,同行,同列,以及当前九宫格里的数字
func getNums(orig [][]int, row, col int) []int {
	// 定位9宫格 ,3X3的阵列
	gridRowIndex := locateStartIndex(row)
	gridColIndex := locateStartIndex(col)
	grid := []int{}
	for i := gridRowIndex; i < gridRowIndex+3; i++ {
		for j := gridColIndex; j < gridColIndex+3; i++ {
			if orig[i][j] == 0 {
				continue
			}
			grid = append(grid, orig[i][j])
		}
	}

	// 横向
	rowDiff := GetDiff(orig[row])

	// 纵向
	colDiff := []int{}
	for _, v := range orig {
		if v[col] == 0 {
			continue
		}
		colDiff = append(colDiff, v[col])
	}
	alls := append(append(grid, rowDiff...), colDiff...)
	return GetDiff(alls)
}

func locateStartIndex(i int) int {
	switch i {
	case 0, 1, 2:
		return 0
	case 3, 4, 5:
		return 3
	case 6, 7, 8:
		return 6
	}
	return 0
}
