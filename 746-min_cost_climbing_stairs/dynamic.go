package lc

// Time: O(n)
// Benchmark: 4ms 3.1mb | 85%

func minCostClimbingStairs(cost []int) int {
	n := len(cost)
	sums := make([]int, n)
	for i, c := range cost {
		if i >= 2 {
			if sums[i-1] < sums[i-2] {
				sums[i] = sums[i-1] + c
			} else {
				sums[i] = sums[i-2] + c
			}
		} else {
			sums[i] = c
		}
	}

	if sums[n-2] < sums[n-1] {
		return sums[n-2]
	}

	return sums[n-1]
}
