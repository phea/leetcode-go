package lc

// Time: O(n^2)
// Benchmark: 0ms 2.1mb | 100% 100%

func partitionLabels(S string) []int {
	res := []int{}
	startPos, endPos := 0, 0
	for i, ch := range S {
		for j := len(S) - 1; j >= 0; j-- {
			if byte(ch) == S[j] {
				if j > endPos {
					endPos = j
				}
				break
			}
		}

		if i == endPos {
			res = append(res, endPos+1-startPos)
			startPos = i + 1
		}
	}

	return res
}
