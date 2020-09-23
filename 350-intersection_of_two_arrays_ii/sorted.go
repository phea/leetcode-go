package lc

import "sort"

// Time: O(n logn)
// Benchmark: 4ms 2.8mb | 83% 97%

func intersect(nums1 []int, nums2 []int) []int {
	if len(nums2) > len(nums1) {
		nums1, nums2 = nums2, nums1
	}

	sort.Ints(nums1)
	sort.Ints(nums2)
	var pos int
	resp := []int{}
	for i := 0; i < len(nums1); i++ {
		for j := pos; j < len(nums2); j++ {
			if nums1[i] == nums2[j] {
				resp = append(resp, nums1[i])
				pos = j + 1
				break
			}
		}
	}

	return resp
}
