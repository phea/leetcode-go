package lc

// Time: O(n^2)
// Benchmark: 1348ms 9.4mb | 9% 32%

func xorQueriesB(arr []int, queries [][]int) []int {
	res := make([]int, len(queries))
	for i, q := range queries {
		sum := 0
		for k := q[0]; k <= q[1]; k++ {
			sum ^= arr[k]
		}
		res[i] = sum
	}

	return res
}
