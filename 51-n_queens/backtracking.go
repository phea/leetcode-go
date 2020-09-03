package lc

// Time: O(n!)
// This is the classic backtracking example that shows how factorial growing algorithms can quickly get computational
// expensive.
//
// n    |  solutions
//------------------
//  4   |          2
//  8   |         92
// 16   | 14,772,512

// Helper function that returns the string representation of the board.
func printBoard(places []int, size int) []string {
	board := make([]string, size)
	for i := range board {
		buf := make([]byte, size)
		for j := 0; j < size; j++ {
			if j == places[i] {
				buf[j] = byte('Q')
			} else {
				buf[j] = byte('.')
			}
		}

		board[i] = string(buf)
	}

	return board
}

func solveNQueens(n int) [][]string {
	boards := [][]string{}

	col := make([]int, n)
	diag := make([]int, n*2)
	diag2 := make([]int, n*2)
	places := make([]int, n)

	// recursive backtracking function.
	var search func(y int)
	search = func(y int) {
		if y == n {
			boards = append(boards, printBoard(places, n))
			return
		}

		for x := 0; x < n; x++ {
			if col[x] != 0 || diag[x+y] != 0 || diag2[x-y+n-1] != 0 {
				continue
			}
			col[x], diag[x+y], diag2[x-y+n-1], places[x] = 1, 1, 1, y
			search(y + 1)
			col[x], diag[x+y], diag2[x-y+n-1], places[x] = 0, 0, 0, 0
		}
	}

	search(0)
	return boards
}
