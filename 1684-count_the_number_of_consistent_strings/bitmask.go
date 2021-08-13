package lc

// Time: O(n^2)
// Benchmark: 24ms 6.9mb | 100%

func countConsistentStrings(allowed string, words []string) int {
	var mask int32
	for _, ch := range allowed {
		mask |= 1 << (ch - 'a')
	}

	var total int
	for _, word := range words {
		var wMask int32
		for _, ch := range word {
			wMask |= 1 << (ch - 'a')
		}

		if (mask & wMask) == wMask {
			total++
		}
	}

	return total
}
