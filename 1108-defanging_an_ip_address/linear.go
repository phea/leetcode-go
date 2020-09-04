package lc

import "bytes"

// Time: O(n)
// Space: O(n)

func defangIPaddr(address string) string {
	var s bytes.Buffer
	for _, c := range address {
		if c == '.' {
			s.WriteString("[.]")
		} else {
			s.WriteRune(c)
		}
	}
	return s.String()
}
