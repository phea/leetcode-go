package lc

import "sort"

// Time: O(n)
// Benchmark: 12ms 3.2mb

func minSubsequence(nums []int) []int {
	sort.Slice(nums, func(i, j int) bool {
		return nums[i] > nums[j]
	})

	var total int
	for _, n := range nums {
		total += n
	}

	var subTotal, i int
	for i = 0; i < len(nums); i++ {
		subTotal += nums[i]
		if subTotal > total-subTotal {
			break
		}
	}

	return nums[:i+1]
}
