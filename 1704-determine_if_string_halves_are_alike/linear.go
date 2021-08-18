package lc

// Time: O(n)
// Benchmark: 0ms 2.2mb | 100%

func halvesAreAlike(s string) bool {
	var count int
	l := len(s)
	isVowel := func(ch byte) bool {
		if ch > 'Z' {
			ch = ch - 32
		}

		if ch == 'A' || ch == 'E' || ch == 'I' ||
			ch == 'O' || ch == 'U' {
			return true
		}
		return false
	}

	for i := 0; i < l/2; i++ {
		if isVowel(s[i]) {
			count++
		}

		if isVowel(s[l-i-1]) {
			count--
		}
	}

	return count == 0
}
