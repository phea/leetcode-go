package lc

// Time: O(n)
// Benchmark: 20ms 6mb | 80% 38%

func pivotIndex(nums []int) int {
	if len(nums) == 0 {
		return -1
	}

	prefix := make([]int, len(nums))
	prefix[0] = nums[0]
	for i := 1; i < len(nums); i++ {
		prefix[i] = prefix[i-1] + nums[i]
	}

	if prefix[len(nums)-1]-prefix[0] == 0 {
		return 0
	}

	for i := 1; i < len(nums); i++ {
		if prefix[i-1] == prefix[len(nums)-1]-prefix[i] {
			return i
		}
	}
	return -1
}
