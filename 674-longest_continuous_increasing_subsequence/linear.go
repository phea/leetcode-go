package lc

// Time: O(n)
// Benchmark: 8ms 4.5mb | 81%

func findLengthOfLCIS(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	score := 1
	max := 1
	for i := 1; i < len(nums); i++ {
		if nums[i] > nums[i-1] {
			score++
		} else {
			score = 1
		}

		if score > max {
			max = score
		}
	}

	return max
}
