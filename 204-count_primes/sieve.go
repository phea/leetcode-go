package lc

// Time:  O(n log log n)
// Benchmark: 8ms 4.9mb | 95% 83%

func countPrimes(n int) int {
	if n <= 2 {
		return 0
	}

	sieve := make([]byte, n+1)
	count := 1
	for i := 3; i < n; i += 2 {
		if sieve[i] == 1 {
			continue
		}

		for j := 2 * i; j < n; j += i {
			sieve[j] = 1
		}

		count++
	}

	return count
}
