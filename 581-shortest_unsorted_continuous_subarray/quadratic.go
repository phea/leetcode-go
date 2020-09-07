package lc

// Time O(n^2)
// Benchmark: 952ms 6.3mb | 6% 44%

func findUnsortedSubarrayC(nums []int) int {
	low, high := -1, -2
	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			if nums[i] > nums[j] {
				if low < 0 {
					low = i
				}

				if j > high {
					high = j
				}
			}
		}
	}

	return high - low + 1
}
