package lc

// Time: O(n^2)
// Benchmark: 4ms 2.8mb | 86% 100%

func intersectionC(nums1 []int, nums2 []int) []int {
	set := []int{}

	add := func(set *[]int, num int) {
		for _, v := range *set {
			if v == num {
				return
			}
		}
		*set = append(*set, num)
	}

	for i := 0; i < len(nums1); i++ {
		for j := 0; j < len(nums2); j++ {
			if nums1[i] == nums2[j] {
				add(&set, nums1[i])
				break
			}
		}
	}

	return set
}
