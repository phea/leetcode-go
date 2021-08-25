package lc

// Time: O(n)
// Benchmark: 0ms 2.2mb | 100%

func minCostToMoveChips(position []int) int {
	var odds, evens int
	for _, p := range position {
		if p%2 == 0 {
			evens++
		} else {
			odds++
		}
	}

	if odds > evens {
		return evens
	}
	return odds
}
