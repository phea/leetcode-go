package lc

import "math"

func canCompleteCircuit(gas []int, cost []int) int {
	if len(gas) == 0 {
		return -1
	}

	trip := make([]int, len(gas))
	trip[0] = gas[0] - cost[0]
	for i := 1; i < len(trip); i++ {
		trip[i] = gas[i] - cost[i] + trip[i-1]
	}

	if trip[len(trip)-1] < 0 {
		return -1
	}

	min := math.MaxInt32
	var minIdx int
	for i := 0; i < len(trip); i++ {
		if trip[i] < min {
			min = trip[i]
			minIdx = i
		}
	}

	if minIdx == len(trip)-1 {
		return 0
	}

	return minIdx + 1
}
