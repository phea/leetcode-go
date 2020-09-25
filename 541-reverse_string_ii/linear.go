package lc

import "bytes"

// Time: O(n)
// Benchmark: 0ms 3.8mb | 100% 50%

func reverseStr(s string, k int) string {
	var buf bytes.Buffer
	pos := 0
	for buf.Len() < len(s) {
		for j := pos + k - 1; j >= pos; j-- {
			if j > len(s)-1 {
				continue
			}
			buf.WriteByte(s[j])
		}

		pos += k
		if pos <= len(s)-1 && pos+k < len(s)-1 {
			buf.WriteString(s[pos : pos+k])
		} else if pos <= len(s)-1 {
			buf.WriteString(s[pos:])
		}
		pos += k
	}

	return buf.String()
}
