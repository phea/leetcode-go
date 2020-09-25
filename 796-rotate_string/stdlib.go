package lc

import "strings"

func rotateString(A string, B string) bool {
	if len(B) != len(A) {
		return false
	}

	return strings.Contains(B+B, A)
}
