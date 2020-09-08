package lc

import (
	"bytes"
	"strconv"
)

// Time: O(n)
// Benchmark: 0ms 2mb | 100%

func freqAlphabets(s string) string {
	var buf bytes.Buffer
	for i, ch := range s {
		if ch == '#' {
			buf.Truncate(buf.Len() - 2)
			d, _ := strconv.Atoi(s[i-2 : i])
			buf.WriteString(string(96 + d))
		} else {
			buf.WriteRune(ch + 48)
		}
	}

	return buf.String()
}
