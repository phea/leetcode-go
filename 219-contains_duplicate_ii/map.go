package lc

// Time: (n logn)
// Benchmark: 20ms 8mb | 38% 17%

// Using a map to track the last index for n.
func containsNearbyDuplicate(nums []int, k int) bool {
	// records index of last seen for n.
	seen := make(map[int]int)
	for i := 0; i < len(nums); i++ {
		idx, ok := seen[nums[i]]
		if ok && i-idx <= k {
			return true
		}
		seen[nums[i]] = i
	}

	return false
}
