package lc

// Time: O(n)
// Benchmark: 0ms 2.3mb | 100%

func rob(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	if len(nums) == 1 {
		return nums[0]
	}

	max := func(x, y int) int {
		if x > y {
			return x
		}
		return y
	}

	sums := make([]int, len(nums))
	sums[0] = nums[0]
	sums[1] = nums[1]
	if len(nums) >= 3 {
		sums[2] = nums[0] + nums[2]
	}

	for i := 3; i < len(nums); i++ {
		sums[i] = max(sums[i-3], sums[i-2]) + nums[i]
	}

	if sums[len(nums)-2] > sums[len(nums)-1] {
		return sums[len(nums)-2]
	}
	return sums[len(nums)-1]
}
