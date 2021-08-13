package lc

// Time: O(n)
// Benchmark: 0ms 2mb | 100%

func countGoodSubstrings(s string) int {
	var total int
	for i := 0; i < len(s)-2; i++ {
		if s[i] == s[i+1] || s[i] == s[i+2] || s[i+1] == s[i+2] {
			continue
		}
		total++
	}
	return total
}
