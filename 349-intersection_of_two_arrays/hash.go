package lc

// Time: O(n)
// Benchmark: 0ms 3mb | 100% 79%

func intersection(nums1 []int, nums2 []int) []int {
	add := func(set *[]int, num int) {
		for _, v := range *set {
			if v == num {
				return
			}
		}
		*set = append(*set, num)
	}

	nums := make(map[int]bool)
	for _, num := range nums1 {
		nums[num] = true
	}

	set := []int{}
	for _, num := range nums2 {
		if nums[num] {
			add(&set, num)
		}
	}

	return set
}
