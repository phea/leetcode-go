package lc

import (
	"strconv"
	"strings"
)

// Time: O(n)
// Benchmark: 0ms 2.2mb | 100% 37%

func areNumbersAscending(s string) bool {
	var last int
	s2 := strings.Split(s, " ")
	for _, x := range s2 {
		if x[0] >= 'a' && x[0] <= 'z' {
			continue
		}

		n, _ := strconv.Atoi(x)
		if n <= last {
			return false
		}
		last = n
	}

	return true
}
