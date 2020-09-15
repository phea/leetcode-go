package lc

import "sort"

// Time: O(n)
// Benchmark: 0ms 2.7mb | 100% 5%

func merge(nums1 []int, m int, nums2 []int, n int) {
	if n == 0 {
		return
	}

	var pos1, pos2 int
	for i := 0; i < len(nums1); i++ {
		if pos1 < m {
			if nums1[pos1] <= nums2[pos2] {
				pos1++
			} else {
				idx := sort.SearchInts(nums2, nums1[pos1])
				if idx >= len(nums2) {
					nums2 = append(nums2, nums1[pos1])
				} else {
					nums2 = append(nums2[:idx], append([]int{nums1[pos1]}, nums2[idx:]...)...)
				}
				nums1[pos1] = nums2[0]
				nums2 = nums2[1:]
				pos1++
			}
		} else {
			nums1[i] = nums2[pos2]
			pos2++
		}
	}
}
