package lc

import (
	"bytes"
	"strings"
)

// Time: O(n)
// Benchmark: 0ms 2.2mb | 100%

func numDifferentIntegers(word string) int {
	nums := make(map[string]bool)
	var counting bool
	var buf bytes.Buffer
	for _, ch := range word {
		if ch >= 'a' {
			if counting {
				s := strings.TrimLeft(buf.String(), "0")
				nums[s] = true
				buf.Reset()
			}
			counting = false
		} else {
			buf.WriteRune(ch)
			counting = true
		}
	}

	if counting {
		s := strings.TrimLeft(buf.String(), "0")
		nums[s] = true
	}

	return len(nums)
}
