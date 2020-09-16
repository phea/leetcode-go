package lc

import "sort"

// Time: O(n^2)
// Benchmark: 44ms 5.9mb | 27% 100%

func findContentChildren(g []int, s []int) int {
	sort.Ints(g)
	sort.Ints(s)

	var child, cookie int
	for child < len(g) && cookie < len(s) {
		if s[cookie] >= g[child] {
			child++
		}
		cookie++
	}

	return child
}
