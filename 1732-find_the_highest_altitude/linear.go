package lc

// Time: O(n)
// Benchmark: 0ms 2.2mb | 100%

func largestAltitude(gain []int) int {
	var max, alt int
	for _, n := range gain {
		alt += n
		if alt > max {
			max = alt
		}
	}

	return max
}
