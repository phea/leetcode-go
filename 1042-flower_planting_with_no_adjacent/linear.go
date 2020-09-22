package lc

// Benchmark: O(n)
// Time: 56ms 7.8mb | 95% 36%

func gardenNoAdj(N int, paths [][]int) []int {
	adjancy := make([][]int, N+1)
	for _, p := range paths {
		adjancy[p[0]-1] = append(adjancy[p[0]-1], p[1])
		adjancy[p[1]-1] = append(adjancy[p[1]-1], p[0])
	}

	garden := make([]int, N)
	for i := 0; i < N; i++ {
		var mask int
		for _, edge := range adjancy[i] {
			// mask numbers already taken
			mask |= 1 << garden[edge-1]
		}

		for k := 1; k <= 4; k++ {
			if (mask>>k)&1 == 0 {
				garden[i] = k
				break
			}
		}
	}

	return garden
}
