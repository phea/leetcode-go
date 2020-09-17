package lc

// Time: O(n)
// Benchmark: 4ms 2.7mb | 40% 27%

func queensAttacktheKing(queens [][]int, king []int) [][]int {
	directions := [][]int{
		{0, 1}, {0, -1}, {1, 0}, {-1, 0},
		{1, 1}, {-1, -1}, {-1, 1}, {1, -1},
	}

	captured := [][]int{}
	for _, dir := range directions {
		cur := []int{king[0], king[1]}
		var hadCapture bool
		for cur[0] >= 0 && cur[0] <= 8 && cur[1] >= 0 && cur[1] <= 8 {
			cur[0] += dir[0]
			cur[1] += dir[1]

			for _, q := range queens {
				if q[0] == cur[0] && q[1] == cur[1] {
					captured = append(captured, q)
					hadCapture = true
					break
				}
			}

			if hadCapture {
				break
			}
		}
	}

	return captured
}
