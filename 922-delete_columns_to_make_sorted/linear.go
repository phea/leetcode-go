package lc

// Time: O(n) = n is number of characters
// Benchmark: 8ms 6.1mb | 93%

func minDeletionSize(A []string) int {
	if len(A) <= 1 {
		return 0
	}

	var count int
	for i := range A[0] {
		for j := 0; j <= len(A)-1; j++ {
			if A[j][i] >= A[j+1][i] 
				count++
				break
			}
		}
	}

	return count
}