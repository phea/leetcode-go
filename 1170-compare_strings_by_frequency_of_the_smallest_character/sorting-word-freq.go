package lc

import "sort"

// Time: O(n^2)
// Benchmark: 12ms 4.6mb | 47% 24%

func numSmallerByFrequencyB(queries []string, words []string) []int {
	// returns the frequency of the smallest character
	var smallestWordFreq = func(s string) int {
		if len(s) == 0 {
			return 0
		}

		count := 1
		minChar := s[0]
		for i := 1; i < len(s); i++ {
			if s[i] < minChar {
				minChar = s[i]
				count = 1
				continue
			}

			if minChar == s[i] {
				count++
			}
		}

		return count
	}

	// pre-calculate frequency for words
	wordsFreq := make([]int, len(words))
	for i, word := range words {
		wordsFreq[i] = smallestWordFreq(word)
	}
	sort.Slice(wordsFreq, func(i, j int) bool {
		return wordsFreq[i] > wordsFreq[j]
	})

	// pre-calculate frequency for queries
	queryFreq := make([]int, len(queries))
	for i, q := range queries {
		queryFreq[i] = smallestWordFreq(q)
	}

	freq := make([]int, len(queries))
	for i, qFreq := range queryFreq {
		for _, wFreq := range wordsFreq {
			if qFreq >= wFreq {
				break
			}
			freq[i]++
		}
	}

	return freq
}
