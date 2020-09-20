package lc

// Time: O(n)
// Benchmark: 8ms 4.4mb | 91% 44%

func minPathSum(grid [][]int) int {
	min := func(x, y int) int {
		if x < y {
			return x
		}
		return y
	}

	m := len(grid)    // y
	n := len(grid[0]) // x

	sums := make([][]int, m)
	for y := 0; y < m; y++ {
		row := make([]int, n)
		for x := 0; x < n; x++ {
			// If I'm on the top row
			if y == 0 {
				if x == 0 {
					row[x] = grid[y][x]
				} else {
					row[x] = row[x-1] + grid[y][x]
				}
			} else {
				if x == 0 {
					row[x] = sums[y-1][x] + grid[y][x]
				} else {
					row[x] = min(row[x-1], sums[y-1][x]) + grid[y][x]
				}
			}
		}
		sums[y] = row
	}

	return sums[m-1][n-1]
}
