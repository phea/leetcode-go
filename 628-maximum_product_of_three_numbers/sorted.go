package lc

import "sort"

// Time: O(n logn)
// Benchmark: 80ms 6.3mb | 9% 16%

func maximumProductB(nums []int) int {
	sort.Ints(nums)
	left := nums[0] * nums[1] * nums[len(nums)-1]
	right := nums[len(nums)-1] * nums[len(nums)-2] * nums[len(nums)-3]

	if left > right {
		return left
	}
	return right
}
