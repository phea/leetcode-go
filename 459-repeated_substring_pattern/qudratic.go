package lc

// Time: O(n^2)
// Uses an increasing step size and compares the rest of the string by step size.
// Additional conditional check if step size does not mod to 0.

func repeatedSubstringPatternB(s string) bool {
	var sub string
	var found bool
	for step := 1; step <= len(s)/2; step++ {
		sub = s[:step]

		// step over the rest of the character, step size
		if len(s)%step == 0 {
			for window := step; window+step <= len(s); window += step {
				if sub != s[window:window+step] {
					found = false
					break
				}
				found = true
			}
		}

		if found {
			return true
		}
	}

	return false

}
