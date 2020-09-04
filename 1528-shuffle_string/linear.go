package lc

// Time: O(n)

func restoreString(s string, indices []int) string {
	b := make([]byte, len(s))

	for i, n := range indices {
		b[n] = s[i]
	}

	return string(b)
}
