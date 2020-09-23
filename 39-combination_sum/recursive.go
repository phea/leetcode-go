package lc

import "sort"

// Time: O(n)
// Benchmark: 4ms 3.1mb | 100%

func combinationSum(candidates []int, target int) [][]int {
	var rc func(nums []int, idx int, target int, valid *[][]int)
	rc = func(nums []int, idx int, target int, valid *[][]int) {
		if target <= 0 {
			if target == 0 {
				*valid = append(*valid, append([]int{}, nums...))
			}
			return
		}

		for i := idx; i < len(candidates); i++ {
			if candidates[i] > target {
				continue
			}

			rc(append(nums, candidates[i]), i, target-candidates[i], valid)
		}
	}
	sort.Ints(candidates)
	valid := [][]int{}
	rc([]int{}, 0, target, &valid)
	return valid
}
