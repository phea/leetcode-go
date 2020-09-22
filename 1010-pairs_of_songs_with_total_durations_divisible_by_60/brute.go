package lc

// Time: O(n^2)
// Benchmark: 3424ms 6.1mb | 40% 90%

func numPairsDivisibleBy60B(time []int) int {
	var count int
	for i := 0; i < len(time)-1; i++ {
		for j := i + 1; j < len(time); j++ {
			if (time[i]+time[j])%60 == 0 {
				count++
			}
		}
	}

	return count
}
