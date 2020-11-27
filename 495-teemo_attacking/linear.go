package lc

// Time: O(n)
// Benchmark: 32ms 6.4mb | 100% 59%

func findPoisonedDuration(timeSeries []int, duration int) int {
	var total int
	var endPoison int
	for _, t := range timeSeries {
		if t < endPoison {
			total += t + duration - endPoison
			endPoison = t + duration
		} else {
			endPoison = t + duration
			total += duration
		}
	}

	return total
}
