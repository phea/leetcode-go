package lc

import "sort"

// Time: O(n logn)
// Benchmark: 28ms 6.2mb |  76% 60%

func findUnsortedSubarray(nums []int) int {
	sorted := make([]int, len(nums))
	copy(sorted, nums)
	sort.Ints(sorted)

	start, end := -1, -1
	for i := 0; i < len(nums); i++ {
		if nums[i] != sorted[i] && start < 0 {
			start = i
		}

		if nums[len(nums)-i-1] != sorted[len(nums)-i-1] && end < 0 {
			end = len(nums) - i - 1
		}

		if start >= 0 && end >= 0 {
			break
		}
	}

	if start < 0 {
		return 0
	}
	return end - start + 1
}
