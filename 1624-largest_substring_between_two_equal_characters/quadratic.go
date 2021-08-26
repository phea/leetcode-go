package lc

// Time: O(n^2)
// Benchmark: 0ms 2mb | 100%

func maxLengthBetweenEqualCharacters(s string) int {
	max := -1
	for i := 0; i < len(s)/2; i++ {
		for j := len(s) - 1; j >= len(s)/2; j-- {
			if s[i] == s[j] {
				if j-i-1 > max {
					max = j - i - 1
				}
				break
			}
		}
	}

	return max
}
