package lc

// Time: O(n)
// Benchmark: 4ms 3.3mb | 90% 6%

func arraySign(nums []int) int {
	c := 1
	for _, n := range nums {
		if n == 0 {
			return 0
		} else if n < 0 {
			c *= -1
		}
	}

	if c < 0 {
		return -1
	}
	return 1
}
