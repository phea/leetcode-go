package lc

// Time: O(n!)
// Question is an easier version of #51 N-Queens.

func totalNQueens(n int) int {
	col := make([]int, n)
	diag := make([]int, n*2)
	diag2 := make([]int, n*2)
	var count int
	var search func(y int)
	search = func(y int) {
		if y == n {
			count++
			return
		}

		for x := 0; x < n; x++ {
			if col[x] != 0 || diag[x+y] != 0 || diag2[x-y+n-1] != 0 {
				continue
			}
			col[x], diag[x+y], diag2[x-y+n-1] = 1, 1, 1
			search(y + 1)
			col[x], diag[x+y], diag2[x-y+n-1] = 0, 0, 0
		}
	}

	search(0)
	return count
}
