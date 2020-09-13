package lc

// Time: O(n)
// Benchmark: 0ms 2.8mb | 100% 28%

func longestPalindrome(s string) int {
	tally := make([]int, 52)
	for _, ch := range s {
		if ch >= 'a' && ch <= 'z' {
			tally[ch-'a']++
		} else {
			tally[ch-'A'+26]++
		}
	}

	var sum, oddCount int
	for _, n := range tally {
		if n%2 != 0 {
			oddCount++
		}
		sum += n
	}

	if oddCount > 0 {
		return sum - oddCount + 1
	}
	return sum
}
