package lc

import "math"

// Time: O(mn)
// Benchmark: 8ms 5.5mb | 98%

func minFallingPathSum(A [][]int) int {
	min := func(x, y int) int {
		if x < y {
			return x
		}
		return y
	}

	m := len(A)
	n := len(A[0])
	if n == 1 {
		var score int
		for i := 0; i < m; i++ {
			score += A[i][0]
		}
		return score
	}

	if m == 1 {
		var score int
		for _, n := range A[0] {
			score += n
		}
		return score
	}

	prevRow := A[0]
	for i := 1; i < m; i++ {
		row := make([]int, n)
		for j := 0; j < n; j++ {
			var low int
			if j == 0 {
				low = min(prevRow[0], prevRow[1])
			} else if j == n-1 {
				low = min(prevRow[j], prevRow[j-1])
			} else {
				low = min(min(prevRow[j-1], prevRow[j]), prevRow[j+1])
			}
			row[j] = low + A[i][j]
		}
		prevRow = row
	}

	low := math.MaxInt32
	for _, n := range prevRow {
		if n < low {
			low = n
		}
	}
	return low
}
