package lc

// Time: O(n)
// Benchmark: 0ms 2mb | 100%

func isSumEqual(firstWord string, secondWord string, targetWord string) bool {
	sum := func(s string) int {
		var total int
		for _, ch := range s {
			total = total*10 + int(ch-'a')
		}
		return total
	}

	return sum(firstWord)+sum(secondWord) == sum(targetWord)
}
