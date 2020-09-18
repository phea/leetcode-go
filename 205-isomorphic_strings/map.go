package lc

// Time: O(n)
// Benchmark: 0ms 2.7mb | 100%

func isIsomorphic(s string, t string) bool {
	mapping := make(map[rune]rune)
	used := []rune{}
	for i, ch := range s {
		m, ok := mapping[ch-'a']
		if !ok {
			for _, c := range used {
				if c == rune(t[i]) {
					return false
				}
			}
			mapping[ch-'a'] = rune(t[i])
			used = append(used, rune(t[i]))
		} else {
			if m != rune(t[i]) {
				return false
			}
		}
	}

	return true
}
