package lc

// Time: O(n)
// Benchmark: 40ms 4.6mb | 5% 32%

// Naive approach by removing duplicate one at a time.
func removeDuplicatesB(nums []int) int {
	for i := 0; i < len(nums)-1; i++ {
		if nums[i] == nums[i+1] {
			nums = append(nums[:i], nums[i+1:]...)
			i--
		}
	}

	return len(nums)
}
