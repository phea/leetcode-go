package lc

// Time: O(n)
// Benchmark: 16ms 6.2mb | 97% 43%

func minSteps(s string, t string) int {
	charTally := make([]int, 26)
	for _, ch := range s {
		charTally[ch-'a']++
	}

	for _, ch := range t {
		charTally[ch-'a']--
	}

	var total int
	for _, n := range charTally {
		if n > 0 {
			total += n
		}
	}

	return total
}
