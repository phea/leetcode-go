package lc

// Time: O(n)
// Benchmark: 0ms 2mb | 100%

func checkOnesSegment(s string) bool {
	var found bool
	for i := 1; i < len(s); i++ {
		if s[i] == '0' && !found {
			found = true
		} else if s[i] == '1' && found {
			return false
		}
	}

	return true
}
