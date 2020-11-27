package lc

import "math"

// Time: O(n)
// Benchmark: 56ms 8.8mb | 99% 88%

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func constructMaximumBinaryTree(nums []int) *TreeNode {
	if len(nums) == 0 {
		return nil
	}

	max := math.MinInt32
	var idx int
	for i, n := range nums {
		if n > max {
			max = n
			idx = i
		}
	}

	return &TreeNode{Val: max, Left: constructMaximumBinaryTree(nums[:idx]), Right: constructMaximumBinaryTree(nums[idx+1:])}
}
