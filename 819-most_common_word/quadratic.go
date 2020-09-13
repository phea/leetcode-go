package lc

import (
	"strings"
	"unicode"
)

// Time: O(n^2)
// Benchmark: 4ms 3.2mb | 45% 27%

func mostCommonWordB(paragraph string, banned []string) string {
	words := strings.FieldsFunc(paragraph, func(c rune) bool {
		return !unicode.IsLetter(c) && !unicode.IsNumber(c)
	})

	isBanned := func(word string) bool {
		for i := 0; i < len(banned); i++ {
			if word == banned[i] {
				return true
			}
		}
		return false
	}

	wordCount := make(map[string]int)
	for _, w := range words {
		w = strings.ToLower(w)
		if !isBanned(w) {
			wordCount[w]++
		}
	}

	var maxWord string
	var maxCount int
	for w, freq := range wordCount {
		if freq > maxCount {
			maxCount = freq
			maxWord = w
		}
	}

	return maxWord
}
