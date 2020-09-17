package lc

import "sort"

// Time: (n logn)
// Benchmark: 0ms 2.9mb | 100%

func nextGreatestLetter(letters []byte, target byte) byte {
	idx := sort.Search(len(letters), func(i int) bool { return letters[i] > target })
	if idx >= len(letters) {
		return letters[0]
	}

	if letters[idx] == target {
		if idx == len(letters)-1 {
			return letters[0]
		}
		return letters[idx+1]
	}

	return letters[idx]
}
