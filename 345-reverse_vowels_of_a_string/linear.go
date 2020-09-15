package lc

import "bytes"

func reverseVowels(s string) string {
	vowels := []rune{}
	for _, ch := range s {
		switch ch {
		case 'a', 'e', 'i', 'o', 'u':
			vowels = append(vowels, ch)
		case 'A', 'E', 'I', 'O', 'U':
			vowels = append(vowels, ch)
		}
	}

	var buf bytes.Buffer
	pos := len(vowels) - 1
	for _, ch := range s {
		switch ch {
		case 'a', 'e', 'i', 'o', 'u':
			buf.WriteRune(vowels[pos])
			pos--
		case 'A', 'E', 'I', 'O', 'U':
			buf.WriteRune(vowels[pos])
			pos--
		default:
			buf.WriteRune(ch)
		}
	}

	return buf.String()
}
