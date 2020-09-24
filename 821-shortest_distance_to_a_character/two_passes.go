package lc

// Time: O(n)
// Benchmark: 0ms 2.4mb | 100%

// Two passes, first to find the indexes of C, and then a second to calculate the distances.
func shortestToChar(S string, C byte) []int {
	min := func(x, y int) int {
		if x < y {
			return x
		}
		return y
	}

	idx := []int{}
	for i := 0; i < len(S); i++ {
		if S[i] == C {
			idx = append(idx, i)
		}
	}

	distance := make([]int, len(S))
	for i := 0; i < len(distance); i++ {
		if S[i] == C {
			distance[i] = 0
			continue
		}

		lower, upper := -1, 0
		for k := 0; k < len(idx); k++ {
			if idx[k] > i {
				upper = idx[k]
				break
			}

			if idx[k] < i {
				lower = idx[k]
			}
		}
		// no lower bounds
		if lower == -1 {
			distance[i] = upper - i
		} else if upper == 0 { // no upper bounds
			distance[i] = i - lower
		} else {
			distance[i] = min(upper-i, i-lower)
		}
	}

	return distance
}
