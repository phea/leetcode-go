package lc

// Time: O(n)
// Benchmark: 12ms 5.7mb | 90% 76%

// This solution loops size/2 of matrix, and adds the first and second half row in one go.
func diagonalSumB(mat [][]int) int {
	var total int
	var pos int
	size := len(mat)
	for i := 0; i < size/2; i++ {
		total += mat[i][pos] + mat[i][size-1-i]               // first half row
		total += mat[size-1-i][pos] + mat[size-1-i][size-1-i] // second half row
		pos++
	}

	// If odd number of rows.
	if size%2 == 1 {
		total += mat[size/2][size/2]
	}
	return total
}
