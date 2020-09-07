package lc

// Time: O(n)
// Benchmark: 0ms 2mb | 100%

func numberOfLines(widths []int, S string) []int {
	var total int
	lines := 1
	for _, ch := range S {
		if total+widths[ch-'a'] > 100 {
			lines++
			total = widths[ch-'a']

		} else {
			total += widths[ch-'a']
		}
	}

	return []int{lines, total}
}
