package lc

// Time: O(n)
// 0ms 2.1mb | 100%

func subtractProductAndSum(n int) int {
	prod, sum := 1, 0
	for n > 0 {
		prod *= n % 10
		sum += n % 10
		n /= 10
	}

	return prod - sum
}
