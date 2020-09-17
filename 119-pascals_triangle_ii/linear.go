package lc

// Time: O(n)
// Benchmark: 0ms 2.1mb | 100%

func getRow(rowIndex int) []int {
	row := make([]int, rowIndex+1)
	for i := 0; i < len(row); i++ {
		if i == 0 {
			row[i] = 1
		} else if i == 1 {
			row[i] = rowIndex
		} else {
			row[i] = (row[i-1] * (rowIndex - i + 1)) / i
		}
	}

	return row
}
