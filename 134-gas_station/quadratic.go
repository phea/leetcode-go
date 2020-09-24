package lc

// Time: O(n^2)
// Benchmark: 208ms 3mb | 9% 100%

func canCompleteCircuitB(gas []int, cost []int) int {
	for i := 0; i < len(gas); i++ {
		if cost[i] > gas[i] {
			continue
		}
		pos := i
		var tank int
		for tank >= 0 {
			tank += gas[pos]
			tank -= cost[pos]
			pos = (pos + 1) % len(gas)
			if pos == i && tank >= 0 {
				return i
			}
		}
	}

	return -1
}
