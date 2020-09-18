package lc

// Time: O(n)
// Benchmark: 0ms 2mb | 100%

func letterCombinations(digits string) []string {
	if digits == "" {
		return nil
	}

	mapping := [][]string{{"a", "b", "c"}, {"d", "e", "f"}, {"g", "h", "i"},
		{"j", "k", "l"}, {"m", "n", "o"}, {"p", "q", "r", "s"},
		{"t", "u", "v"}, {"w", "x", "y", "z"}}

	var rf func(s string, digits string, valid *[]string)
	rf = func(s string, digits string, valid *[]string) {
		if digits == "" {
			*valid = append(*valid, s)
			return
		}

		d := digits[0]
		for _, ch := range mapping[d-'2'] {
			rf(s+ch, digits[1:], valid)
		}
	}

	valid := []string{}
	rf("", digits, &valid)
	return valid
}
