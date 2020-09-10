package lc

import "fmt"

// Time: O(n)
// Benchmark: 0ms 2.2mb | 100%

func getHint(secret string, guess string) string {
	tally := make([]int, 10)
	for _, ch := range guess {
		tally[ch-'0']++
	}

	var bulls, cows int
	for i, ch := range secret {
		if ch == rune(guess[i]) {
			bulls++
		}

		if tally[ch-'0'] > 0 {
			tally[ch-'0']--
			cows++
		}
	}

	return fmt.Sprintf("%dA%dB", bulls, cows-bulls)
}
