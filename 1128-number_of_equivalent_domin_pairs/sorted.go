package lc

import "sort"

// Time: O(n logn)
// Benchmark: 72ms  7.2mb | 23% 97%

func numEquivDominoPairs(dominoes [][]int) int {
	sort.Slice(dominoes, func(i, j int) bool {
		if dominoes[i][0] > dominoes[i][1] {
			dominoes[i][0], dominoes[i][1] = dominoes[i][1], dominoes[i][0]
		}

		if dominoes[j][0] > dominoes[j][1] {
			dominoes[j][0], dominoes[j][1] = dominoes[j][1], dominoes[j][0]
		}

		if dominoes[i][0] == dominoes[j][0] {
			return dominoes[i][1] < dominoes[j][1]
		}
		return dominoes[i][0] < dominoes[j][0]
	})

	var total int
	for i := 0; i < len(dominoes)-1; i++ {
		for j := i + 1; j < len(dominoes); j++ {
			if dominoes[i][0] < dominoes[j][0] || dominoes[i][1] < dominoes[j][1] {
				break
			}

			total++
		}
	}
	return total
}
