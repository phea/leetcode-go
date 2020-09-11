package lc

import "bytes"

// Time: O(n)
// Benchmark: 0ms 2mb | 100% 67%

// We treat the letter slice like a stack.
func reverseOnlyLetters(S string) string {
	chars := []rune{}
	isLetter := func(ch rune) bool {
		return (ch >= 'a' && ch <= 'z') || (ch >= 'A' && ch <= 'Z')
	}

	for _, ch := range S {
		if isLetter(ch) {
			chars = append(chars, ch)
		}
	}

	var buf bytes.Buffer
	pos := len(chars) - 1
	for _, ch := range S {
		if isLetter(ch) {
			buf.WriteRune(chars[pos])
			pos--
		} else {
			buf.WriteRune(ch)
		}
	}

	return buf.String()
}
