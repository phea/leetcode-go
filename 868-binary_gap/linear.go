package lc

import "math"

// Time: O(n) n = number of bits of input number
// Benchmark: 0ms 2mb | 100% 93%

func binaryGap(N int) int {
	distance, max := 0, math.MinInt32

	for N > 0 {
		if N&1 == 0 {
			N = N >> 1
		} else {
			distance = 0
			N = N >> 1
			for N > 0 {
				distance++
				if N&1 == 1 {
					break
				}
				N = N >> 1
			}

			if distance > max {
				max = distance
			}
		}
	}

	return max
}
