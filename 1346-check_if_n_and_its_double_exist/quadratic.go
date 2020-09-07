package lc

// Time: O(n^2)
// Benchmark: 4ms 3mb | 91% 99%
// Max array size for this question is 500 so a quadratic solution is plenty fast.

func checkIfExist(arr []int) bool {
	for i := 0; i < len(arr); i++ {
		if arr[i]%2 == 0 {
			for j := 0; j < len(arr); j++ {
				if j == i {
					continue
				}

				if arr[i]/2 == arr[j] {
					return true
				}
			}
		}
	}

	return false
}
