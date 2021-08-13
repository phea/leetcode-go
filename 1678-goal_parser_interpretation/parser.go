package lc

import "bytes"

// Time: O(n)
// Benchmark: 0ms 2mb | 100% 58%

func interpret(command string) string {
	var buf bytes.Buffer
	for i := 0; i < len(command); i++ {
		ch := command[i]
		if ch == '(' {
			// peek ahead
			if command[i+1] == ')' {
				buf.WriteRune('o')
			} else {
				buf.WriteString("al")
				i += 2
			}
		} else if ch == 'G' {
			buf.WriteRune('G')
		}
	}

	return buf.String()
}
