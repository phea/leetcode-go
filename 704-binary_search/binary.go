package lc

// Time: O(logn)
// Benchmark: 40ms 6.5mb | 32%

func search(nums []int, target int) int {
	low, high := 0, len(nums)-1
	var mid int
	for low <= high {
		mid = (low + high) / 2
		if nums[mid] < target {
			low = mid + 1
		} else if nums[mid] > target {
			high = mid - 1
		} else {
			return mid
		}
	}

	return -1
}
