package lc

// Time: O(n)
// Benchmark: 8ms 5.6mb | 100%

func minOperations(nums []int) int {
	var total int
	for i := 1; i < len(nums); i++ {
		if nums[i] > nums[i-1] {
			continue
		}
		total += nums[i-1] - nums[i] + 1
		nums[i] = nums[i-1] + 1
	}
	return total
}
