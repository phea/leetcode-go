package lc

import "math"

// Time: O(n)
// Benchmark 4ms 3.1mb | 100%

func findNumbers(nums []int) int {
	var c, digits int
	for _, n := range nums {
		digits = int(math.Log10(float64(n))) + 1
		if digits%2 == 0 {
			c++
		}
	}
	return c
}
