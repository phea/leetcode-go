package lc

import "sort"

// Time: 0(n^2)
// Benchmark: 2040ms 7.5mb | 6%

func minMovesB(nums []int) int {
	if len(nums) <= 1 {
		return 0
	}
	sort.Ints(nums)
	var total int
	for nums[0] != nums[len(nums)-1] {
		d := nums[len(nums)-1] - nums[0]
		total += d
		for i := 0; i < len(nums)-1; i++ {
			nums[i] += d
		}
		nums = append([]int{nums[0], nums[len(nums)-1]}, nums[1:len(nums)-1]...)
	}

	return total
}
