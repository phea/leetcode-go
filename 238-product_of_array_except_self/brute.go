package lc

// Time: O(n^2)
// Benchmark: 1908ms 6.3mb | 6% 84%

func productExceptSelfB(nums []int) []int {
	products := make([]int, len(nums))
	for i := 0; i < len(nums); i++ {
		sum := 1
		for j := 0; j < len(nums); j++ {
			if i == j {
				continue
			}
			sum *= nums[j]
		}

		products[i] = sum
	}

	return products
}
