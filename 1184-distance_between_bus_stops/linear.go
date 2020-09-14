package lc

// Time: O(n)
// Benchmark: 4ms 2.8mb |  40% 27%

// Two passes, forwards and backwards through array.
func distanceBetweenBusStops(distance []int, start int, destination int) int {
	var forward int
	for i := start; i != destination; i = (i + 1) % len(distance) {
		forward += distance[i]
	}

	var backward int
	for i := start - 1; i != destination-1; {
		if i < 0 {
			i = len(distance) - 1
		}
		backward += distance[i]
		i--
	}

	if forward < backward {
		return forward
	}
	return backward
}
