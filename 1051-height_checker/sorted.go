package lc

import "sort"

// Time: O(n logn)
// Benchmark: 0ms 2.3mb | 100% 18%

// We make a sorted copy of the original array and count the misplaced heights.
func heightChecker(heights []int) int {
	sorted := make([]int, len(heights))
	copy(sorted, heights)
	sort.Ints(sorted)

	var count int
	for i, n := range sorted {
		if n != heights[i] {
			count++
		}
	}

	return count
}
