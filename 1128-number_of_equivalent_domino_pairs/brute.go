package lc

// Time: O(n^2)
// Benchmark: 4676ms 7.2mb | 7% 67%

func numEquivDominoPairsB(dominoes [][]int) int {
	var total int
	for i := 0; i < len(dominoes)-1; i++ {
		for j := i + 1; j < len(dominoes); j++ {
			if (dominoes[i][0] == dominoes[j][0] && dominoes[i][1] == dominoes[j][1]) ||
				(dominoes[i][0] == dominoes[j][1] && dominoes[i][1] == dominoes[j][0]) {
				total++
			}
		}
	}

	return total
}
