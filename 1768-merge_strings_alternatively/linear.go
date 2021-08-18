package lc

import "bytes"

// Time: O(n)
// Benchmark: 0ms 2.1mb | 100%

func mergeAlternately(word1 string, word2 string) string {
	var buf bytes.Buffer
	for i := range word1 {
		buf.WriteByte(word1[i])
		if i < len(word2) {
			buf.WriteByte(word2[i])
		}
	}

	if len(word1) < len(word2) {
		buf.WriteString(word2[len(word1):])
	}
	return buf.String()
}
