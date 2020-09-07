package lc

import "strings"

// Time: O(n)
// Space: O(n)
// Benchmark: 0ms 2mb | 100%

func wordPattern(pattern string, str string) bool {
	codes := make(map[string]byte)
	used := make([]bool, 26)
	words := strings.Split(str, " ")
	if len(pattern) != len(words) {
		return false
	}

	for i, word := range words {
		code, ok := codes[word]
		if !ok {
			// check if character was already used
			if used[pattern[i]-'a'] {
				return false
			}
			used[pattern[i]-'a'] = true
			codes[word] = pattern[i]
			continue
		}

		if code != pattern[i] {
			return false
		}

	}

	return true
}
