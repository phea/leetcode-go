package lc

import "bytes"

// Time: O(n)
// Benchmark: 0ms 2.6mb | 100%

func isPrefixString(s string, words []string) bool {
	var idx int
	var buf bytes.Buffer
	for buf.Len() < len(s) {
		if idx+1 > len(words) {
			return false
		}

		buf.WriteString(words[idx])
		idx++
	}

	return s == buf.String()
}
