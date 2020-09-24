package lc

// Time: O(logn)
// Benchmark: 0ms 1.9mb | 100% 81%

func isBadVersion(version int) bool

func firstBadVersion(n int) int {
	min, max := 1, n
	var guess int
	for min < max {
		guess = min + (max-min)/2
		if isBadVersion(guess) {
			max = guess
		} else {
			min = guess + 1
		}
	}
	return min
}
