package lc

// Time: O(n)
// Benchmark: 0ms 2.0mb | 100%

func checkRecord(s string) bool {
	var aCount int
	for i, ch := range s {
		if ch == 'A' {
			aCount++
			if aCount > 1 {
				return false
			}
		} else if ch == 'L' {
			if i+2 < len(s) && s[i+1] == 'L' && s[i+2] == 'L' {
				return false
			}
		}
	}

	return true
}
