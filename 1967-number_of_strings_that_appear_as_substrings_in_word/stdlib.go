package lc

import "strings"

// Time: O(nm) or O(n+m)
// Stdlib uses brute force or Rabin-Karp depending on substring length.
// Benchmark: 0ms 2.7mb | 100%

func numOfStrings(patterns []string, word string) int {
	var total int
	for _, s := range patterns {
		if strings.Index(word, s) >= 0 {
			total++
		}
	}
	return total
}
