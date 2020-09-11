package lc

import "sort"

// Time: O(n logn)
// Benchmark: 0ms 2.1mb | 100%

func lastStoneWeight(stones []int) int {
	sort.Ints(stones)

	for len(stones) >= 2 {
		stone := stones[len(stones)-1] - stones[len(stones)-2]
		stones = stones[:len(stones)-2]
		if stone > 0 {
			// get the index to push stone on slice.
			idx := sort.SearchInts(stones, stone)
			if idx >= len(stones) {
				stones = append(stones, stone)
			} else {
				stones = append(stones[:idx], append([]int{stone}, stones[idx:]...)...)
			}
		}
	}

	if len(stones) == 0 {
		return 0
	}
	return stones[0]
}
