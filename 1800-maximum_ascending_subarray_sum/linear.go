package lc

// Time: O(n)
// Benchmark: 4ms 2.3mb | 100%

func maxAscendingSum(nums []int) int {
	max, tally := nums[0], nums[0]
	for i := 1; i < len(nums); i++ {
		if nums[i] <= nums[i-1] {
			tally = nums[i]
		} else {
			tally += nums[i]
		}

		if tally > max {
			max = tally
		}
	}
	return max
}
