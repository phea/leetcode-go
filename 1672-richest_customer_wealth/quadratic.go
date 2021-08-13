package lc

// Time: O(n^2)
// Benchmark: 8ms 3.4mb | 100% 7%

func maximumWealth(accounts [][]int) int {
	var max int
	for _, account := range accounts {
		var total int
		for _, n := range account {
			total += n
		}

		if total > max {
			max = total
		}
	}

	return max
}
