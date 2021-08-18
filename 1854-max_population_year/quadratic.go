package lc

// Time: O(n^2)
// Benchmark: 0ms 2.5mb | 100%

func maximumPopulation(logs [][]int) int {
	b := logs[0][0]
	d := logs[0][1]
	for _, l := range logs {
		if l[0] < b {
			b = l[0]
		}

		if l[1] > d {
			d = l[1]
		}
	}

	var max int
	minYear := 2051
	for i := b; i < d; i++ {
		var count int
		for _, l := range logs {
			if i >= l[0] && i < l[1] {
				count++
			}
		}

		if count > max {
			max = count
			minYear = i
		}
	}

	return minYear
}
