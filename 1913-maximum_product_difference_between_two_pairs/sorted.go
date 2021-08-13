package lc

import "sort"

// Time: O(n logn)
// Benchmark: 32ms 6.5mb | 66% 97%

func maxProductDifference(nums []int) int {
	sort.Ints(nums)
	return (nums[len(nums)-1] * nums[len(nums)-2]) - (nums[0] * nums[1])
}
