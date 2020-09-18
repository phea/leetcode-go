package lc

// Time: O(n^2)
// Benchmark: 0ms 2mb | 100%

func combinationSum3(k int, n int) [][]int {
	var rc func(nums []int, start int, target int, valid *[][]int)
	rc = func(nums []int, start int, target int, valid *[][]int) {
		if len(nums) == k {
			if target == 0 {
				*valid = append(*valid, append([]int{}, nums...))
			}
			return
		}

		for i := start; i < 10; i++ {
			if i > target {
				break
			}
			rc(append(nums, i), i+1, target-i, valid)
		}
	}

	valid := [][]int{}
	rc([]int{}, 1, n, &valid)
	return valid
}
