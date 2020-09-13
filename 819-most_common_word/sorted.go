package lc

import (
	"sort"
	"strings"
	"unicode"
)

// Time: O(n logn)
// Benchmark: 0ms 3mb | 100% 55%

// We sort the banned list so we can do a binary search when looking up if a word is banned.
func mostCommonWord(paragraph string, banned []string) string {
	words := strings.FieldsFunc(paragraph, func(c rune) bool {
		return !unicode.IsLetter(c) && !unicode.IsNumber(c)
	})

	sort.Strings(banned)
	isBanned := func(word string) bool {
		idx := sort.SearchStrings(banned, word)
		if idx >= len(banned) {
			return false
		}
		return banned[idx] == word
	}

	wordCount := make(map[string]int)
	var maxWord string
	var maxCount int
	for _, w := range words {
		w = strings.ToLower(w)
		if isBanned(w) {
			continue
		}

		wordCount[w]++
		freq := wordCount[w]
		if freq > maxCount {
			maxCount = freq
			maxWord = w
		}
	}
	return maxWord
}
