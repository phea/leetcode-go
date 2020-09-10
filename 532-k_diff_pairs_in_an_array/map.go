package lc

import (
	"fmt"
	"sort"
)

// Time: O(n logn)
// Benchmark: 32ms 6.1mb | 39% 72%

func findPairs(nums []int, k int) int {
	sort.Ints(nums)

	m := make(map[string]bool)
	for i := 0; i < len(nums)-1; i++ {
		for j := i + 1; j < len(nums); j++ {
			if nums[j]-nums[i] == k {
				m[fmt.Sprintf("%d%d", nums[j], nums[i])] = true
			}

			if nums[j]-nums[i] > k {
				break
			}
		}
	}

	return len(m)
}
