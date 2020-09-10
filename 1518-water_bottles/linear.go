package lc

// Time: O(n)
// Benchmark: 0ms 1.9mb | 100%

func numWaterBottles(numBottles int, numExchange int) int {
	total := numBottles
	empties := numBottles
	for empties >= numExchange {
		exchanged := empties / numExchange
		total += exchanged
		empties = exchanged + empties%numExchange

	}

	return total
}
