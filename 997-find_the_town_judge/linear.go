package lc

// Time: O(n)
// Benchmark: 100ms 7.1mb | 89% 77%

func findJudge(N int, trust [][]int) int {
	if len(trust) == 0 && N == 1 {
		return 1
	}

	// out, in
	edges := make([][2]int, N+1)
	candidates := []int{}
	for _, t := range trust {
		edges[t[0]][0]++
		edges[t[1]][1]++
		if edges[t[1]][1] == N-1 {
			candidates = append(candidates, t[1])
		}
	}

	for _, c := range candidates {
		if edges[c][0] == 0 {
			return c
		}
	}
	return -1
}
