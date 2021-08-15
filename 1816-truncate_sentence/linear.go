package lc

// Time: O(n)
// Benchmark: 2ms 2.3mb | 100%

func truncateSentence(s string, k int) string {
	var idx, count int
	for i := 0; i < len(s); i++ {
		if s[i] == ' ' {
			count++
			if k == count {
				idx = i
				break
			}
		}
	}

	if count < k {
		return s
	}
	return s[:idx]
}
