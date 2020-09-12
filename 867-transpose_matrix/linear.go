package lc

// Time: O(n)
// Benchmark: 8ms 5.8mb | 91% 69%

func transpose(A [][]int) [][]int {
	m2 := make([][]int, len(A[0]))
	for i := 0; i < len(A[0]); i++ {
		row := make([]int, len(A))
		for j := 0; j < len(A); j++ {
			row[j] = A[j][i]
		}
		m2[i] = row
	}

	return m2
}
