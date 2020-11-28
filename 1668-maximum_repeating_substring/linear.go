package lc

// Time: O(n)
// Benchmark: 0ms 2mb | 100%

func maxRepeating(sequence string, word string) int {
	size := len(word)
	var matches, max int
	var isSequence bool
	for i := 0; i+size <= len(sequence); i++ {
		if sequence[i:i+size] == word {
			matches++
			i += size - 1
			isSequence = true
		} else {
			if isSequence {
				if matches > max {
					max = matches
				}
				matches = 0
				isSequence = false
			}
		}
	}

	if matches > max {
		return matches
	}
	return max
}
