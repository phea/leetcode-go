package lc

// Time: O(n)
// Benchmark: 0ms 2mb | 100% 36%

func totalMoney(n int) int {
	var sum int
	earnings := 1
	for i := 0; i < n; i++ {
		if i%7 == 0 {
			earnings = (i / 7) + 1
		}
		sum += earnings
		earnings++
	}
	return sum
}
