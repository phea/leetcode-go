package lc

// Time: O(n)
// Benchmark: 0ms 4.7mb | 100%

func insert(intervals [][]int, newInterval []int) [][]int {
	if len(intervals) == 0 {
		return [][]int{newInterval}
	}

	lowIdx, highIdx := 0, 0
	for i, interval := range intervals {
		if newInterval[0] > interval[0] {
			lowIdx = i
		}

		if newInterval[0] > interval[1] {
			lowIdx = i + 1
		}

		if newInterval[1] > interval[0] {
			highIdx = i
		}

		if newInterval[1] > interval[1] {
			highIdx = i + 1
		}
	}

	if lowIdx == highIdx {
		if lowIdx == 0 {
			if intervals[lowIdx][0] > newInterval[1] {
				intervals = append(intervals[:lowIdx], append([][]int{newInterval}, intervals[highIdx:]...)...)
			} else {
				if intervals[lowIdx][0] > newInterval[0] {
					intervals[lowIdx][0] = newInterval[0]
				}
			}
			return intervals
		}

		if highIdx == len(intervals) {
			intervals = append(intervals[:lowIdx], append([][]int{newInterval}, intervals[highIdx:]...)...)
			return intervals
		}

		if intervals[lowIdx-1][1] < newInterval[0] && intervals[highIdx][0] > newInterval[1] {
			intervals = append(intervals[:lowIdx], append([][]int{newInterval}, intervals[highIdx:]...)...)
			return intervals
		}
	}

	if highIdx < len(intervals) && newInterval[1] >= intervals[highIdx][0] {
		low := intervals[lowIdx][0]
		if newInterval[0] < low {
			low = newInterval[0]
		}

		high := intervals[highIdx][1]
		if newInterval[1] > high {
			high = newInterval[1]
		}
		intervals = append(intervals[:lowIdx], intervals[highIdx:]...)
		intervals[lowIdx] = []int{low, high}
	} else {
		originalLow := intervals[lowIdx][0]
		intervals = append(intervals[:lowIdx], intervals[highIdx-1:]...)
		if newInterval[0] < originalLow {
			intervals[lowIdx][0] = newInterval[0]
		} else {
			intervals[lowIdx][0] = originalLow
		}
		intervals[highIdx-(highIdx-lowIdx)][1] = newInterval[1]
	}
	return intervals
}
