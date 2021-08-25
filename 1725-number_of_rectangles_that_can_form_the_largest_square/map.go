package lc

// Time: O(n)
// Benchmark: 24ms 6.8mb | 81% 19%

func countGoodRectangles2(rectangles [][]int) int {
	min := func(r []int) int {
		if r[0] < r[1] {
			return r[0]
		}
		return r[1]
	}

	m := make(map[int]int, len(rectangles))
	var maxmin int
	for _, r := range rectangles {
		rm := min(r)
		m[rm]++
		if rm > maxmin {
			maxmin = rm
		}
	}

	return m[maxmin]
}
