package lc

// Time: O(n)
// Benchmark: 8ms 6.4mb | 89% 69%

func getConcatenation(nums []int) []int {
	return append(nums, nums...)
}
