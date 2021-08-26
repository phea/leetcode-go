package lc

import "sort"

// Time: O(n logn)
// Benchmark: 0ms 2.8mb | 100% 14%

func isCovered(ranges [][]int, left int, right int) bool {
	sort.Slice(ranges, func(i, j int) bool {
		if ranges[i][0] == ranges[j][0] {
			return ranges[i][1] < ranges[j][1]
		}

		return ranges[i][0] < ranges[j][0]
	})

	var idx int
	for i := left; i <= right; i++ {
		var found bool
		for k := idx; k < len(ranges); k++ {
			if i >= ranges[k][0] && i <= ranges[k][1] {
				idx = k
				found = true
				i = ranges[k][1]
				break
			}
		}

		if !found {
			return false
		}
	}

	return true
}
