package lc

// Time: O(n)
// Benchmark: 0ms 2.1mb | 100%

func checkIfPangram(sentence string) bool {
	var mask int
	for _, ch := range sentence {
		mask |= 1 << (ch - 'a')
		if mask == 0x3FFFFFF {
			return true
		}
	}
	return false
}
