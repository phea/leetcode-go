package lc

// Time: O(n)
// Benchmark: 12ms 5.8mb |  90% 78%

func maxArea(height []int) int {
	var current, max int
	l, r := 0, len(height)-1

	for l < r {
		if height[l] < height[r] {
			current = height[l] * (r - l)
			l++
		} else {
			current = height[r] * (r - l)
			r--
		}

		if current > max {
			max = current
		}
	}
	return max
}
