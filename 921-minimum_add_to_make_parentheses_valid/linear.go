package lc

// Time: O(n)
// Benchmark: 0ms 2mb | 100%

func minAddToMakeValid(S string) int {
	var total int
	var open, close int
	for _, ch := range S {
		if ch == '(' {
			if close > 0 {
				total += close
				close = 0
			}
			open++
		} else {
			// If I have open brackets to 'close'
			if open > 0 {
				open--
			} else {
				close++
			}
		}
	}

	return total + open + close
}
