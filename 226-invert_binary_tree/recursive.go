package lc

import "github.com/phea/leetcode-go/types"

// Time: O(n)

func invertTree(root *types.TreeNode) *types.TreeNode {
	if root == nil {
		return root
	}

	root.Left, root.Right = invertTree(root.Right), invertTree(root.Left)
	return root
}
