package lc

import (
	"bytes"
	"strconv"
)

// Time: O(n)
// Benchmark: 0ms 2.2mb | 100%

func getLucky(s string, k int) int {
	var buf bytes.Buffer
	for _, ch := range s {
		buf.WriteString(strconv.Itoa(int(ch - 'a' + 1)))
	}

	s2 := buf.String()
	for i := 0; i < k; i++ {
		var sum int
		for _, ch := range s2 {
			sum += int(ch - '0')
		}
		s2 = strconv.Itoa(sum)
	}

	n, _ := strconv.Atoi(s2)
	return n
}
