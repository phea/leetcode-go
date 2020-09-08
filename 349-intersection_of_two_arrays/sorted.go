package lc

import "sort"

// Time: O(n logn)
// Benchmark: 4ms 2.8mb | 86% 97%

func intersectionB(nums1 []int, nums2 []int) []int {
	set := []int{}

	add := func(set *[]int, num int) {
		for _, v := range *set {
			if v == num {
				return
			}
		}
		*set = append(*set, num)
	}

	sort.Ints(nums2)
	for i := 0; i < len(nums1); i++ {
		idx := sort.SearchInts(nums2, nums1[i])

		if idx < len(nums2) && nums2[idx] == nums1[i] {
			add(&set, nums1[i])
		}
	}

	return set
}
