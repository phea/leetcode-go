package lc

import "github.com/phea/leetcode-go/types"

func maxDepth(root *types.TreeNode) int {
	if root == nil {
		return 0
	}

	left := maxDepth(root.Left)
	right := maxDepth(root.Right)

	if left > right {
		return left + 1
	}
	return right + 1
}
