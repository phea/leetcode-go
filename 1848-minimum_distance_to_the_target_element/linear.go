package lc

// Time: O(n)
// Benchmark: 4ms 3.3mb | 76% 29%

func getMinDistance(nums []int, target int, start int) int {
	rounds := start
	if len(nums)-start > start {
		rounds = len(nums) - start
	}

	for i := 0; i <= rounds; i++ {
		if start-i >= 0 && nums[start-i] == target {
			return i
		}

		if start+i < len(nums) && nums[start+i] == target {
			return i
		}
	}

	return 0
}
