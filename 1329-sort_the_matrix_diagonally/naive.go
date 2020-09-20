package lc

import "sort"

// Time: O(n logn)
// Benchmark: 16ms 6.1mb | 13%

func diagonalSort(mat [][]int) [][]int {
	m := len(mat)
	n := len(mat[0])
	// start 0 row
	for i := 0; i < n; i++ {
		diag := []int{}
		pointers := []*int{}
		for offset := 0; i+offset < n && offset < m; offset++ {
			diag = append(diag, mat[0+offset][i+offset])
			pointers = append(pointers, &mat[0+offset][i+offset])
		}

		sort.Ints(diag)
		for i, p := range pointers {
			*p = diag[i]
		}
	}

	// column 0
	for i := 1; i < m; i++ {
		diag := []int{}
		pointers := []*int{}
		for offset := 0; i+offset < m && offset < n; offset++ {
			diag = append(diag, mat[i+offset][offset])
			pointers = append(pointers, &mat[i+offset][offset])
		}

		sort.Ints(diag)
		for i, p := range pointers {
			*p = diag[i]
		}
	}

	return mat
}
