package lc

// Time: O(n)
// Benchmark: 28ms 6.2mb | 9% 35%

func matrixReshapeB(nums [][]int, r int, c int) [][]int {
	if (r * c) > (len(nums) * len(nums[0])) {
		return nums
	}

	m2 := make([][]int, r)
	cols := len(nums[0])
	var pos int
	for i := 0; i < r; i++ {
		row := make([]int, c)
		for j := 0; j < c; j++ {
			row[j] = nums[pos/cols][pos%cols]
			pos++
		}
		m2[i] = row
	}

	return m2
}
