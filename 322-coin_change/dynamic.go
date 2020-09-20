package lc

import "math"

// Time: O(n*m)
// Benchmark: 4ms 6mb | 99%

func coinChange(coins []int, amount int) int {
	min := func(x, y int) int {
		if x < y {
			return x
		}
		return y
	}

	sums := make([]int, amount+1)
	for i := 1; i <= amount; i++ {
		sums[i] = math.MaxInt32
		for j := 0; j < len(coins); j++ {
			if coins[j] <= i {
				sums[i] = min(sums[i], sums[i-coins[j]]+1)
			}
		}
	}

	if sums[amount] == math.MaxInt32 {
		return -1
	}

	return sums[amount]
}
