package lc

import "sort"

// Time: O(n logn)
// Benchmark: 0ms 2.2mb | 100%

func sumOfUnique(nums []int) int {
	n := len(nums)
	if n <= 1 {
		return nums[0]
	}

	if n == 2 {
		if nums[0] != nums[1] {
			return nums[0] + nums[1]
		}
		return 0
	}

	sort.Ints(nums)
	var total int
	for i := 1; i < n-1; i++ {
		if nums[i] != nums[i-1] && nums[i] != nums[i+1] {
			total += nums[i]
		}
	}

	if nums[0] != nums[1] {
		total += nums[0]
	}

	if nums[n-1] != nums[n-2] {
		total += nums[n-1]
	}

	return total
}
