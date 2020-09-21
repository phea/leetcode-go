package lc

import "sort"

// Time: O(n logn)
// Benchmark: 228ms 6.7mb | 33% 85%

func distributeCandiesB(candies []int) int {
	sort.Ints(candies)

	count := 1
	for i := 0; i < len(candies); i++ {
		for j := i + 1; j < len(candies); j++ {
			if candies[i] != candies[j] {
				i = j - 1
				count++
				break

			}
		}

		if count > len(candies)/2 {
			return len(candies) / 2
		}
	}
	return count
}
