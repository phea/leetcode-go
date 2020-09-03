package lc

// Time: O(n)
// Space: O(n)

func runningSumB(nums []int) []int {
	if len(nums) <= 1 {
		return nums
	}

	sum := make([]int, len(nums))
	sum[0] = nums[0]

	for i := 1; i < len(nums); i++ {
		sum[i] = nums[i] + sum[i-1]
	}

	return sum
}
