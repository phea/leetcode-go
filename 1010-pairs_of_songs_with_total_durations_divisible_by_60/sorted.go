package lc

import "sort"

// Time: (mn logn)
// Benchmark:  136ms  6.1mb | 43% 100%

func numPairsDivisibleBy60(time []int) int {
	sort.Ints(time)

	max := time[len(time)-1]
	if len(time) >= 2 {
		max += time[len(time)-2]
	}

	var count int
	for i := 0; i < len(time)-1; i++ {
		for j := 60; j <= max+60; j += 60 {
			required := j - time[i]
			if required < 0 || required < time[i] {
				continue
			}

			idx := sort.SearchInts(time, required)
			if required == time[i] {
				idx = i + 1
			}

			for idx < len(time) && time[idx] == required {
				count++
				idx++
			}
		}
	}

	return count
}
