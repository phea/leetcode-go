package lc

// Time: O(n)
// Benchmark 0ms 2mb | 100%

func xorOperation(n int, start int) int {
	var t int
	for i := 0; i < n; i++ {
		t ^= start + (2 * i)
	}
	return t
}
