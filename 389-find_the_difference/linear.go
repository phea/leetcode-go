package lc

// Time: O(n)
// Benchmark: 0ms 2.2mb | 100%

func findTheDifference(s string, t string) byte {
	var missing byte
	for i := 0; i < len(s); i++ {
		missing ^= s[i] ^ t[i]
	}

	return missing ^ t[len(t)-1]
}
