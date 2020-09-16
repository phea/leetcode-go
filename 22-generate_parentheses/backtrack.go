package lc

// Time: O(n^2)
// Benchmark: 0ms 3.2mb | 100%

func generateParenthesis(n int) []string {
	var backtrack func(s string, count int, valid *[]string)
	backtrack = func(s string, count int, valid *[]string) {
		if len(s) == n*2 {
			if count == 0 {
				*valid = append(*valid, s)
			}
			return
		}

		if count < 0 {
			return
		}

		backtrack(s+"(", count+1, valid)
		backtrack(s+")", count-1, valid)

	}

	brackets := []string{}
	backtrack("", 0, &brackets)
	return brackets
}
