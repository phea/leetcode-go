package lc

// Time: O(n^2)
// Benchmark: 144ms 6.3mb | 9% 44%
// We optimized the quadratic algorithm by checking inversions starting at the tail.

func findUnsortedSubarrayB(nums []int) int {
	start, end := -2, -1
	for i := 0; i < len(nums); i++ {
		for j := len(nums) - 1; j > i; j-- {
			if j < end {
				break
			}

			if nums[i] > nums[j] {
				if start < 0 {
					start = i
				}

				end = j
			}
		}
	}

	if start < 0 {
		return 0
	}

	return end - start + 1
}
