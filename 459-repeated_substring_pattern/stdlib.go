package lc

import "strings"

// Time: O(n)
// Standard library uses Rabin-Karp algorithm for string searching which is roughly linear.
// https://en.wikipedia.org/wiki/Rabin%E2%80%93Karp_algorithm

func repeatedSubstringPattern(s string) bool {
	if len(s) == 0 {
		return false
	}

	ss := (s + s)[1 : len(s)*2-1]
	return strings.Contains(ss, s)
}
