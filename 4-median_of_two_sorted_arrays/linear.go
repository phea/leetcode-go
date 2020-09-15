package lc

// Time: O(n)
// Benchmark: 16ms 6mb | 64% 21%

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	sorted := make([]int, len(nums1)+len(nums2))
	var pos1, pos2 int
	for i := 0; i < len(sorted); i++ {
		if pos2 >= len(nums2) {
			sorted[i] = nums1[pos1]
			pos1++
			continue
		}

		if pos1 >= len(nums1) {
			sorted[i] = nums2[pos2]
			pos2++
			continue
		}

		if nums1[pos1] <= nums2[pos2] {
			sorted[i] = nums1[pos1]
			pos1++
		} else {
			sorted[i] = nums2[pos2]
			pos2++
		}
	}

	median := len(sorted) / 2
	if len(sorted)%2 != 0 {
		return float64(sorted[median])
	}

	return float64(sorted[median-1]+sorted[median]) / 2
}
