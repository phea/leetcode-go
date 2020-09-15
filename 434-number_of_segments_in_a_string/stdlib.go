package lc

import "strings"

// Time: O(n)
// Benchmark: 0ms 2mb | 100%

func countSegments(s string) int {
	if s == "" {
		return 0
	}

	f := func(c rune) bool {
		return c == ' '
	}

	return len(strings.FieldsFunc(s, f))
}
