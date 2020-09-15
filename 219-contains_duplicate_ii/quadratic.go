package lc

// Time: O(n^2)
// Benchmark: 768ms 5.1mb | 13% 93%

func containsNearbyDuplicateB(nums []int, k int) bool {
	for i := 0; i < len(nums)-1; i++ {
		for j := i + 1; j <= i+k; j++ {
			if j >= len(nums) {
				break
			}

			if nums[j] == nums[i] {
				return true
			}
		}
	}

	return false
}
