package lc

import "github.com/phea/leetcode-go/types"

// Time: O(n)
// Benchmark: 0ms 3.1mb | 100% 95%

func walk(node *types.TreeNode, sum int) int {
	if node == nil {
		return 0
	}

	sum = (sum << 1) | node.Val
	if node.Left == nil && node.Right == nil {
		return sum
	}

	return walk(node.Left, sum) + walk(node.Right, sum)
}

func sumRootToLeaf(root *types.TreeNode) int {
	return walk(root, 0)
}
