package lc

// Time: O(n^2)
// Benchmark: 16ms 4.3mb | 24% 53%

func numSmallerByFrequencyC(queries []string, words []string) []int {
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

	// pre-calculate frequency for queries
	queryFreq := make([]int, len(queries))
	for i, q := range queries {
		queryFreq[i] = smallestWordFreq(q)
	}

	freq := make([]int, len(queries))
	for i, qFreq := range queryFreq {
		for _, wFreq := range wordsFreq {
			if qFreq < wFreq {
				freq[i]++
			}
		}
	}

	return freq
}
