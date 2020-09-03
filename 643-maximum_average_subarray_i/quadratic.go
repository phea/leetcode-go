package lc

import "math"

// Time: O(n^2)
// For each num we get a sum by iterating k items and store the maximum sum in 'max'.

func findMaxAverageB(nums []int, k int) float64 {
	max := math.MinInt32
	for i := 0; i < len(nums)-k+1; i++ {
		sum := 0
		for j := i; j < i+k; j++ {
			sum += nums[j]
		}

		if sum > max {
			max = sum
		}
	}

	return float64(max) / float64(k)
}
