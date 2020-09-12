package lc

// Time: O(n)
// Linear: 8ms 5.8mb | 100% 75%

func prefixesDivBy5(A []int) []bool {
	resp := make([]bool, len(A))
	var bits int
	for i := 0; i < len(A); i++ {
		bits = ((bits << 1) | A[i]) % 5
		resp[i] = bits == 0
	}

	return resp
}
