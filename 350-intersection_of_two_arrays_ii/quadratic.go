package lc

// Time: O(n^2)
// Benchmark: 4ms 2.8mb |  83% 86%

func intersectB(nums1 []int, nums2 []int) []int {
	if len(nums2) > len(nums1) {
		nums1, nums2 = nums2, nums1
	}

	taken := make([]bool, len(nums2))
	resp := []int{}
	for i := 0; i < len(nums1); i++ {
		for j := 0; j < len(nums2); j++ {
			if nums1[i] == nums2[j] && !taken[j] {
				resp = append(resp, nums1[i])
				taken[j] = true
				break
			}
		}
	}

	return resp
}
