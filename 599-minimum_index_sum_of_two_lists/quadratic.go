package lc

import "math"

// Time: O(n^2)
// Benchmark: 28ms 6.4mb | 83% 98%

func findRestaurant(list1 []string, list2 []string) []string {
	idx := make([]int, len(list1))
	min := math.MaxInt32
	for i := 0; i < len(list1); i++ {
		score := math.MaxInt32
		for j := 0; j < len(list2); j++ {
			if i+j > min {
				break
			}

			if list1[i] == list2[j] {
				score = i + j
				break
			}
		}
		idx[i] = score
		if score < min {
			min = score
		}
	}

	resp := []string{}
	for i, n := range idx {
		if n != math.MaxInt32 && n == min {
			resp = append(resp, list1[i])
		}
	}

	return resp
}
