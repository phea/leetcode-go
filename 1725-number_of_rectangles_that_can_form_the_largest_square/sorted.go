package lc

import "sort"

// Time: O(n logn)
// Benchmark: 28ms 6.8mb | 80% 19%

func countGoodRectangles(rectangles [][]int) int {
	min := func(r []int) int {
		if r[0] < r[1] {
			return r[0]
		}
		return r[1]
	}

	sort.Slice(rectangles, func(i, j int) bool {
		r1 := min(rectangles[i])
		r2 := min(rectangles[j])

		return r1 > r2
	})

	count := 1
	for i := 0; i < len(rectangles)-1; i++ {
		r1 := min(rectangles[i])
		r2 := min(rectangles[i+1])
		if r1 > r2 {
			break
		}
		count++
	}
	return count
}
