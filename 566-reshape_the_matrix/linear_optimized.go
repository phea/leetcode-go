package lc

// Time: O(n)
// Benchmark: 12ms 6mb | 97%  53%

// We avoid the slow modulu operation by tracking the row and column position of the original
// matrix.
func matrixReshape(nums [][]int, r int, c int) [][]int {
	if (r * c) > (len(nums) * len(nums[0])) {
		return nums
	}

	m2 := make([][]int, r)
	cols := len(nums[0])
	var cRow, cCol int
	for i := 0; i < r; i++ {
		row := make([]int, c)
		for j := 0; j < c; j++ {
			row[j] = nums[cRow][cCol]
			cCol++
			if cCol >= cols {
				cCol = 0
				cRow++
			}
		}
		m2[i] = row
	}

	return m2
}
