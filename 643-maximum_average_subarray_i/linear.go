package lc

// Time: O(n)
// Keeps a running tally of the last 'k' nums and store the highest tally in 'max'.

func findMaxAverage(nums []int, k int) float64 {
	var tally int
	for i := 0; i < k; i++ {
		tally += nums[i]
	}

	max := tally
	for i := k; i < len(nums); i++ {
		tally = tally + nums[i] - nums[i-k]

		if tally > max {
			max = tally
		}
	}

	return float64(max) / float64(k)
}
