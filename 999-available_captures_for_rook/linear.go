package lc

// Time: O(n)
// Benchmark: 0ms 2mb | 100%

func numRookCaptures(board [][]byte) int {
	var row, col int
	for i := 0; i < 8; i++ {
		for j := 0; j < 8; j++ {
			if board[i][j] == byte('R') {
				row, col = i, j
			}
		}
	}

	var count int
	// north
	for i := row - 1; i >= 0; i-- {
		if board[i][col] == byte('p') {
			count++
			break
		}

		if board[i][col] == byte('B') {
			break
		}
	}

	//south
	for i := row + 1; i < 8; i++ {
		if board[i][col] == byte('p') {
			count++
			break
		}

		if board[i][col] == byte('B') {
			break
		}
	}

	// east
	for j := col + 1; j < 8; j++ {
		if board[row][j] == byte('p') {
			count++
			break
		}

		if board[row][j] == byte('B') {
			break
		}
	}

	// west
	for j := col - 1; j >= 0; j-- {
		if board[row][j] == byte('p') {
			count++
			break
		}

		if board[row][j] == byte('B') {
			break
		}
	}

	return count
}
