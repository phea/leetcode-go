package lc

// Time: O(n)
// Benchmark: 0ms 2.7mb | 100%

func largeGroupPositions(s string) [][]int {
	resp := [][]int{}
	start, c := 0, 1
	for i := 1; i < len(s); i++ {
		if s[i] == s[i-1] {
			c++
		} else {
			if c >= 3 {
				resp = append(resp, []int{start, i - 1})
			}
			start = i
			c = 1
		}
	}

	if c >= 3 {
		resp = append(resp, []int{start, len(s) - 1})
	}

	return resp
}
