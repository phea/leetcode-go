package lc

// Time: O(n^2)
// Benchmark: 0ms 3.1mb | 100% 82%

// Use a map to keep track of greater number.
func nextGreaterElement(nums1 []int, nums2 []int) []int {
	greater := make(map[int]int, len(nums2))
	for i, num := range nums2 {
		for j := i + 1; j < len(nums2); j++ {
			if nums2[j] > num {
				greater[num] = nums2[j]
				break
			}
		}
	}

	ans := make([]int, len(nums1))
	for i, num := range nums1 {
		val, ok := greater[num]
		if !ok {
			ans[i] = -1
		} else {
			ans[i] = val
		}
	}

	return ans
}
