package lc

// Time: O(n)
// Benchmark: 4ms 3.1mb | 89% 13%

func minTimeToVisitAllPoints(points [][]int) int {
	max := func(x, y int) int {
		if x > y {
			return x
		}
		return y
	}

	abs := func(x int) int {
		if x < 0 {
			return x * -1
		}
		return x
	}

	var dist int
	for i := 0; i < len(points)-1; i++ {
		x1 := abs(points[i+1][0] - points[i][0])
		y1 := abs(points[i+1][1] - points[i][1])
		dist += max(x1, y1)
	}
	return dist
}
