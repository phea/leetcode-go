package lc

import (
	"bytes"
	"strings"
)

// Time: O(n)
// Benchmark: 0ms 2.1mb | 100%

func reorderSpaces(text string) string {
	var spaceCount int
	for _, ch := range text {
		if ch == ' ' {
			spaceCount++
		}
	}

	words := strings.Fields(text)
	var spacesBetween, spacesAfter int
	if len(words) <= 1 {
		spacesAfter = spaceCount
	} else {
		spacesBetween = spaceCount / (len(words) - 1)
		spacesAfter = spaceCount % (len(words) - 1)
	}

	var buf bytes.Buffer
	for i, w := range words {
		buf.WriteString(w)
		if i == len(words)-1 {
			break
		}
		for j := 0; j < spacesBetween; j++ {
			buf.WriteString(" ")
		}
	}

	for i := 0; i < spacesAfter; i++ {
		buf.WriteString(" ")
	}
	return buf.String()
}
