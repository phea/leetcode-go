package lc

// Time: O(n^2)
// Benchmark: 220ms 9.9mb | 97% 79%

func rearrangeArray(nums []int) []int {
	swapped := true
	for swapped {
		swapped = false
		for i := len(nums) - 2; i > 0; i-- {
			if (nums[i-1]+nums[i+1])%2 != 0 {
				continue
			}

			if (nums[i-1]+nums[i+1])/2 == nums[i] {
				if nums[i+1] > nums[i-1] {
					nums[i], nums[i-1], nums[i+1] = nums[i+1], nums[i], nums[i-1]
				} else {
					nums[i], nums[i-1], nums[i+1] = nums[i-1], nums[i+1], nums[i]
				}

				swapped = true
			}
		}
	}

	return nums
}
