package lc

import "strings"

// Time: O(n) - n = letters
// Benchmark: 4ms 4.1mb | 100%

func shortestCompletingWord(licensePlate string, words []string) string {
	var count int
	letters := make([]int, 26)
	for _, ch := range licensePlate {
		if ch >= 'a' && ch <= 'z' {
			letters[ch-'a']++
			count++
		}

		if ch >= 'A' && ch <= 'Z' {
			letters[ch-'A']++
			count++
		}
	}

	hasLetters := func(word string) bool {
		for ch, total := range letters {
			if total == 0 {
				continue
			}

			if strings.Count(word, string(rune(ch+'a'))) < total {
				return false
			}

		}

		return true
	}

	var word string
	for _, w := range words {
		if hasLetters(w) && (len(w) < len(word) || word == "") {
			word = w
		}

	}
	return word
}
