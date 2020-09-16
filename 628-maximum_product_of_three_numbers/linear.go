package lc

import "math"

// Time: O(n)
// Benchmark: 44ms 6.5mb | 51%

func maximumProduct(nums []int) int {
	high1, high2, high3 := math.MinInt32+2, math.MinInt32+1, math.MinInt32
	low1, low2 := math.MaxInt32, math.MaxInt32+1

	for _, n := range nums {
		if n > high1 {
			if n > high3 {
				high1, high2, high3 = high2, high3, n
			} else if n > high2 {
				high1, high2 = high2, n
			} else {
				high1 = n
			}
		}

		if n < low1 {
			low1, low2 = n, low1
		} else if n < low2 {
			low2 = n
		}
	}

	left := low1 * low2 * high3
	right := high1 * high2 * high3
	if left > right {
		return left
	}
	return right
}
