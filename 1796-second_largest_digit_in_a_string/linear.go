package lc

// Time: O(n)
// Benchmark: 0ms 2.3mb | 100%

func secondHighest(s string) int {
	x, y := -1, -1
	for _, ch := range s {
		if ch >= 'a' {
			continue
		}

		n := int(ch - '0')
		if n > x {
			x, y = n, x
		} else if n > y && n != x {
			y = n
		}
	}

	return y
}
