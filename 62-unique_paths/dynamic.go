package lc

// Time: O(mn)
// Benchmark: 0ms 2mb | 100%

func uniquePaths(m int, n int) int {
	grid := make([][]int, m)
	for i := 0; i < m; i++ {
		row := make([]int, n)
		for j := 0; j < n; j++ {
			if i == 0 {
				row[j] = 1
			} else {
				if j == 0 {
					row[j] = 1
				} else {
					row[j] = row[j-1] + grid[i-1][j]
				}
			}
		}
		grid[i] = row
	}

	return grid[m-1][n-1]
}
