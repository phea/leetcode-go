package lc

import (
	"math/bits"
	"sort"
)

// Time: (n logn)
// The sort.Slice function of the Standard Library uses a quicksort algorithm.

func sortByBits(arr []int) []int {
	sort.Slice(arr, func(i, j int) bool {
		x, y := bits.OnesCount(uint(arr[i])), bits.OnesCount(uint(arr[j]))
		if x == y {
			return arr[i] < arr[j]
		}
		return x < y
	})
	return arr
}
