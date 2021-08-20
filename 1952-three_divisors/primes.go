package lc

// Time: O(1)
// Benchmark: 0ms 2mb | 100%

func isThree(n int) bool {
	primes := []int{2, 3, 5, 7, 11, 13, 17, 19, 23, 29, 31, 37, 41, 43, 47, 53, 59, 61, 67, 71, 73, 79, 83, 89, 97}

	for _, p := range primes {
		if p*p == n {
			return true
		}
	}
	return false
}
