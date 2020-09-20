package lc

// Time: O(n^2)
// Benchmark: 72ms 5.8mb | 23% 97%

// Optimizations include:
// - skipping over current height * (length-i) is less than max
// - breaking second for loop if we encounter height that is equal or taller.
func maxAreaB(height []int) int {
	var max int
	for i := 0; i < len(height)-1; i++ {
		if height[i]*(len(height)-1-i) < max {
			continue
		}

		var currentMax int
		for j := len(height) - 1; j >= i; j-- {
			if height[j] >= height[i] {
				if height[i]*(j-i) > currentMax {
					currentMax = height[i] * (j - i)
				}
				break
			}

			if height[j]*(j-i) > currentMax {
				currentMax = height[j] * (j - i)
			}
		}

		if currentMax > max {
			max = currentMax
		}
	}
	return max
}
