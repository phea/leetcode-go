package lc

import "sort"

// Time: O(n^2)
// Benchmark: 0ms 3.3mb | 100%

func combinationSum2(candidates []int, target int) [][]int {
	valid := [][]int{}
	push := func(nums ...int) {
		for _, v := range valid {
			if len(v) != len(nums) {
				continue
			}

			var matched int
			for i := 0; i < len(v); i++ {
				if v[i] != nums[i] {
					break
				}
				matched++
			}

			if matched == len(v) {
				return
			}
		}

		valid = append(valid, nums)
	}

	var rc func(nums []int, idx int, target int)
	rc = func(nums []int, idx int, target int) {
		if target <= 0 {
			if target == 0 {
				push(nums...)
			}
			return
		}

		for i := idx; i < len(candidates); i++ {
			if candidates[i] > target {
				continue
			}

			rc(append(nums, candidates[i]), i+1, target-candidates[i])
		}
	}
	sort.Ints(candidates)
	rc([]int{}, 0, target)
	if candidates[len(candidates)-1] == target {
		push(target)
	}
	return valid
}
