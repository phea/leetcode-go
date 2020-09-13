package lc

import "bytes"

// Time: O(n)
// Benchmark: 0ms 2mb | 100% 77%

func backspaceCompare(S string, T string) bool {
	var buf1 bytes.Buffer
	for _, ch := range S {
		if ch == '#' {
			if buf1.Len() > 0 {
				buf1.Truncate(buf1.Len() - 1)
			}
		} else {
			buf1.WriteRune(ch)
		}
	}

	var buf2 bytes.Buffer
	for _, ch := range T {
		if ch == '#' {
			if buf2.Len() > 0 {
				buf2.Truncate(buf2.Len() - 1)
			}
		} else {
			buf2.WriteRune(ch)
		}
	}

	return buf1.String() == buf2.String()
}
