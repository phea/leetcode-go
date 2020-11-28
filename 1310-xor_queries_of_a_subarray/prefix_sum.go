package lc

// Time: O(n)
// Benchmark: 60ms 8.4mb | 96% 69%

// Just like we can use a prefix sum to quickly add a contigious section of a subarray, we can do
// the same for xor sum.
func xorQueries(arr []int, queries [][]int) []int {
	prefix := make([]int, len(arr)+1)
	prefix[0] = 0
	for i := 1; i < len(arr)+1; i++ {
		prefix[i] = prefix[i-1] ^ arr[i-1]
	}

	res := make([]int, len(queries))
	for i, q := range queries {
		res[i] = prefix[q[0]] ^ prefix[q[1]+1]
	}

	return res
}
