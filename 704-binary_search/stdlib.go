package lc

import "sort"

// Time: O(logn)
// Benchmark: 32ms 6.5mb | 96% 73%

func searchB(nums []int, target int) int {
	idx := sort.SearchInts(nums, target)
	if idx >= len(nums) || nums[idx] != target {
		return -1
	}
	return idx
}
