package lc

// Time: O(n)
// Benchmark: 176ms 17.5mb | 100% 92%

func findSmallestSetOfVertices(n int, edges [][]int) []int {
	hasIncoming := make([]bool, n)
	for _, e := range edges {
		hasIncoming[e[1]] = true
	}

	resp := []int{}
	for i, v := range hasIncoming {
		if !v {
			resp = append(resp, i)
		}
	}

	return resp
}
