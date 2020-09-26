package lc

import "sort"

// Time: O(n^2)
// Benchmark: 552ms 6.4mb | 9% 88%

func minMoves(nums []int) int {
	if len(nums) <= 1 {
		return 0
	}

	sort.Ints(nums)
	var total int
	pos := len(nums) - 1
	for nums[0] != nums[pos] {
		d := nums[pos] - nums[0]
		total += d
		for i := 0; i < pos; i++ {
			nums[i] += d
		}
		pos--
	}

	return total
}
