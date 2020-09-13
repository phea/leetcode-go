package lc

// Time: O(n)
// Benchmark: 32ms 4.6 | 7% 20%

// We remove all duplicates in a single append.
func removeDuplicates(nums []int) int {
	for i := 0; i < len(nums)-1; i++ {
		if nums[i] == nums[i+1] {
			j := i + 1
			for j < len(nums) {
				if nums[i] != nums[j] {
					break
				}
				j++
			}
			nums = append(nums[:i], nums[j-1:]...)
			i--
		}
	}

	return len(nums)
}
