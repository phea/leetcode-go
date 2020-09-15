package lc

// Time: O(n)
// Benchmark:  36ms 6.4mb | 100%

// We track the column status using an array:
// -1 - invalid
// 0 - free
// 1 - occupied by special
func numSpecial(mat [][]int) int {
	columnStatus := make([]int, len(mat[0]))

	for i := 0; i < len(mat); i++ {
		colPos := []int{}
		for j := 0; j < len(mat[0]); j++ {
			if mat[i][j] == 1 {
				colPos = append(colPos, j)
			}
		}

		if len(colPos) == 1 && columnStatus[colPos[0]] == 0 {
			columnStatus[colPos[0]] = 1
		} else {
			for _, pos := range colPos {
				columnStatus[pos] = -1
			}
		}
	}

	var count int
	for _, tally := range columnStatus {
		if tally == 1 {
			count++
		}
	}

	return count
}
