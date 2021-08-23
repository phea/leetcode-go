package lc

import "math"

// Time: O(n)
// Benchmark: 4ms 3.1mb | 100%

func findGCD(nums []int) int {
	a := math.MinInt32
	b := math.MaxInt32

	for _, n := range nums {
		if n > a {
			a = n
		}
		if n < b {
			b = n
		}
	}

	for b != 0 {
		t := b
		b = a % b
		a = t
	}

	return a
}
