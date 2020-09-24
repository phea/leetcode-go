package lc

import (
	"fmt"
	"strconv"
)

// Time: O(n)
// Benchmark: 12ms 9.3mb | 57% 19%

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func walk(node *TreeNode) string {
	if node == nil {
		return ""
	}

	if node.Left != nil && node.Right != nil {
		return fmt.Sprintf("%d(%s)(%s)", node.Val, walk(node.Left), walk(node.Right))
	} else if node.Left == nil && node.Right == nil {
		return strconv.Itoa(node.Val)
	} else if node.Left != nil {
		return fmt.Sprintf("%d(%s)", node.Val, walk(node.Left))
	} else {
		return fmt.Sprintf("%d()(%s)", node.Val, walk(node.Right))
	}
}

func tree2str(t *TreeNode) string {
	return walk(t)
}
