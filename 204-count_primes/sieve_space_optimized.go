package lc

// Time: O(n log log n)
// Benchmark: 12ms 2.3mb | 67% 85%

// We trade speed for space by using bitwise operations, using the bit position to represent numbers.
func countPrimesB(n int) int {
	if n <= 2 {
		return 0
	}

	sieve := make([]int64, n/64+1)
	count := 1
	for i := 3; i < n; i += 2 {
		if (sieve[i/64]>>(i%64))&1 == 1 {
			continue
		}

		for j := 2 * i; j < n; j += i {
			sieve[j/64] |= 1 << (j % 64)
		}

		count++
	}

	return count
}
