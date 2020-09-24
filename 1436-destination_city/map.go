package lc

// Time: O(n)
// Benchmark: 4ms 4.3mb | 93% 32%

func destCity(paths [][]string) string {
	edges := make(map[string][2]int)
	for _, p := range paths {
		dep := edges[p[0]]
		dep[0]++
		edges[p[0]] = dep
		arr := edges[p[1]]
		arr[1]++
		edges[p[1]] = arr
	}

	for k, edge := range edges {
		if edge[0] == 0 {
			return k
		}
	}

	return ""
}
