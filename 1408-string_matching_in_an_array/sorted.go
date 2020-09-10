package lc

import (
	"sort"
	"strings"
)

// Time: O(n logn)
// Benchmark: 0ms 2.3mb | 100%

// We sort words by length so substrings only search strings of equal or higher length.
func stringMatching(words []string) []string {
	sort.Slice(words, func(i, j int) bool {
		return len(words[i]) < len(words[j])
	})

	subs := []string{}
	for i := 0; i < len(words); i++ {
		for j := i + 1; j < len(words); j++ {
			if strings.Contains(words[j], words[i]) {
				subs = append(subs, words[i])
				break
			}
		}
	}
	return subs
}
