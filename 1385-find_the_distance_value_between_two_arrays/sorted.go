package lc

import "sort"

// Time: O(n logn)
// Benchmark: 8ms 3.8mb | 92% 87%

func findTheDistanceValue(arr1 []int, arr2 []int, d int) int {
	sort.Ints(arr2)

	var count int
	for _, num := range arr1 {
		minIdx := sort.SearchInts(arr2, num-d)
		if minIdx < len(arr2) && arr2[minIdx] >= num-d && arr2[minIdx] <= num+d {
			continue
		}
		count++
	}

	return count
}
