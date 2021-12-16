package exam

import (
	"fmt"
	"testing"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func maxDepth(node *TreeNode) int {
	if node == nil {
		return 0
	}
	n1 := maxDepth(node.Left)
	n2 := maxDepth(node.Right)
	return max(n1, n2) + 1
}

func max(x, y int) int {
	if x < y {
		return y
	}
	return x
}

func TestTreeDeep(*testing.T) {
	node5 := &TreeNode{
		Val:   7,
		Left:  nil,
		Right: nil,
	}
	node4 := &TreeNode{
		Val:   15,
		Left:  nil,
		Right: nil,
	}
	node3 := &TreeNode{
		Val:   20,
		Left:  node4,
		Right: node5,
	}
	node2 := &TreeNode{
		Val:   9,
		Left:  nil,
		Right: nil,
	}
	node1 := &TreeNode{
		Val:   3,
		Left:  node2,
		Right: node3,
	}
	fmt.Println(maxDepth(node1))
}
