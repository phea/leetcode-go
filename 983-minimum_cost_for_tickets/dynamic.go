package lc

import (
	"math"
	"sort"
)

// Time: O(n*m)
// Benchmark: 0ms 2.3mb | 100%

func mincostTickets(days []int, costs []int) int {
	min := func(x, y int) int {
		if x < y {
			return x
		}
		return y
	}

	isTravelDay := func(day int) bool {
		idx := sort.SearchInts(days, day)
		return days[idx] == day
	}

	sums := make([]int, days[len(days)-1]+1)
	var lastDay, tally int
	for i := 1; i <= days[len(days)-1]; i++ {
		if !isTravelDay(i) {
			sums[i] = tally
			continue
		}
		sums[i] = math.MaxInt32
		for j := 0; j < 3; j++ {
			switch j {
			case 0:
				lastDay = i - 1
			case 1:
				lastDay = i - 7
			case 2:
				lastDay = i - 30
			}

			if lastDay < 0 {
				lastDay = 0
			}

			sums[i] = min(sums[i], sums[lastDay]+costs[j])
			tally = sums[i]
		}
	}

	return sums[len(sums)-1]
}
