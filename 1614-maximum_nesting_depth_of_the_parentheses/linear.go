package lc

// Time: O(n)
// Benchmark: 0ms 2mb | 100%

func maxDepth(s string) int {
	var depth, max int
	for _, ch := range s {
		if ch == '(' {
			depth++
			if depth > max {
				max = depth
			}
		} else if ch == ')' {
			depth--
		}
	}

	return max
}
