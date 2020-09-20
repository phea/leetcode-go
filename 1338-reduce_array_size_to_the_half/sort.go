package lc

import "sort"

// Time: (n logn)
// Benchmark: 140ms 9.5mb | 39% 100%

func minSetSize(arr []int) int {
	sort.Ints(arr)

	counts := []int{}
	cur := 0
	for i := 0; i < len(arr)-1; i++ {
		cur++
		if arr[i] != arr[i+1] {
			counts = append(counts, cur)
			cur = 0
		}
	}

	if arr[len(arr)-1] != arr[len(arr)-2] {
		counts = append(counts, 1)
	} else {
		cur++
		counts = append(counts, cur)
	}

	sort.Ints(counts)
	var sum int
	var c int
	i := len(counts) - 1
	for sum < len(arr)/2 {
		sum += counts[i]
		c++
		i--
	}

	return c
}
