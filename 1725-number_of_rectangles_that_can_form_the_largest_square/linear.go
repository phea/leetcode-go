package lc

// Time: O(n)
// Benchmark: 20ms 6.7mb | 96% 19%

func countGoodRectangles3(rectangles [][]int) int {
	min := func(r []int) int {
		if r[0] < r[1] {
			return r[0]
		}
		return r[1]
	}

	var maxmin, c int
	for _, r := range rectangles {
		rm := min(r)
		if rm > maxmin {
			maxmin = rm
			c = 0
		}

		if rm == maxmin {
			c++
		}
	}

	return c
}
