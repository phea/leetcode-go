package lc

// Time: O(n)
// Benchmark: 12ms 4.5mb | 61% 69%

func isToeplitzMatrixB(matrix [][]int) bool {
	r := len(matrix)
	c := len(matrix[0])
	for i := 0; i < (r-1)*c; i++ {
		row := i / c
		col := i % c
		if col == c-1 {
			continue
		}
		if matrix[row][col] != matrix[row+1][col+1] {
			return false
		}
	}

	return true
}
