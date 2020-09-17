package lc

// Time: O(n)
// Benchmark: 4ms 2.3mb | 18% 25%

func dominantIndex(nums []int) int {
	var second, max, maxIdx int
	for i, n := range nums {
		if n > max {
			second, max = max, n
			maxIdx = i
			continue
		}

		if n > second {
			second = n
		}
	}

	if max >= second*2 {
		return maxIdx
	}

	return -1
}
