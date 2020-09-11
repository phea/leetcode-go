package lc

func maxProduct(nums []int) int {
	if len(nums) == 1 {
		return nums[0]
	}

	var neg, pos int
	max := nums[0]
	for i := 0; i < len(nums); i++ {
		if nums[i] > 0 {
			if pos == 0 && neg == 0 {
				pos = nums[i]
				continue
			}

			if pos == 0 {
				pos, neg = nums[i], neg*nums[i]
			} else {
				pos, neg = pos*nums[i], neg*nums[i]
			}
		} else if nums[i] < 0 {
			if pos == 0 && neg == 0 {
				neg = nums[i]
				continue
			}

			if pos > max {
				max = pos
			}

			if pos*nums[i] < nums[i] {
				pos, neg = neg*nums[i], pos*nums[i]
			} else {
				pos, neg = neg*nums[i], nums[i]
			}
		} else {

			if pos > max {
				max = pos
			}
			neg, pos = 0, 0
		}
	}

	if pos > max {
		return pos
	}
	return max
}
