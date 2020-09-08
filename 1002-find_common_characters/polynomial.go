package lc

import "strings"

// Time: O(n^3)
// Benchmark: 8ms  5.6mb | 43% 5%

func commonCharsB(A []string) []string {
	common := strings.Split(A[0], "")

	for i := 1; i < len(A); i++ {
		word := strings.Split(A[i], "")
		for j := 0; j < len(common); j++ {
			var found bool
			for k := 0; k < len(word); k++ {
				if common[j] == word[k] {
					word[k] = "0" // mark as viewed
					found = true
					break
				}
			}

			if !found {
				common = append(common[:j], common[j+1:]...)
				j--
			}
		}
	}

	return common
}
