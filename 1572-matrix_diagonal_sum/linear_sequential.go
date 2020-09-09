package lc

// Time: O(n)
// Benchmark: 4ms 5.7mb | 100%

// Accessing items sequentially is much faster than jumping around.
func diagonalSum(mat [][]int) int {
	var total int
	size := len(mat)
	for i := 0; i < size; i++ {
		if i < size/2 {
			total += mat[i][i] + mat[i][size-1-i]
		} else {
			total += mat[i][size-1-i] + mat[i][i]
		}
	}

	// If odd number of rows.
	if size%2 == 1 {
		total -= mat[size/2][size/2]
	}
	return total
}
