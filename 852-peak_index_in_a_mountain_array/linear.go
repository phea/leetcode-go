package lc

// Time: O(n)
// Benchmark: 8ms 4.8mb | 92% 100%

func peakIndexInMountainArray(A []int) int {
	for i := 0; i < len(A)-1; i++ {
		if A[i+1] < A[i] {
			return i
		}
	}

	return 0
}
