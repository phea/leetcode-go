package lc

import "bytes"

// Time: O(n)
// Benchmark: 0ms 2mb | 100%

func modifyString(s string) string {
	var buf bytes.Buffer
	prev := '0'
	var next rune
	for i, ch := range s {
		if i < len(s)-1 {
			next = rune(s[i+1])
		} else {
			next = '0'
		}
		if ch == '?' {
			ch = 'a'
			for ch == prev || ch == next {
				ch++
				if ch == '{' {
					ch = 'a'
				}
			}
		}
		prev = ch
		buf.WriteRune(ch)
	}

	return buf.String()
}
