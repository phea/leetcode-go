package lc

// Time: O(n)
// Benchmark: 24ms 6.2mb | 91% 68%

func validMountainArray(A []int) bool {
	if len(A) < 3 {
		return false
	}

	if A[0] > A[1] {
		return false
	}

	isIncreasing := true
	for i := 1; i < len(A); i++ {
		if A[i-1] == A[i] {
			return false
		}

		if isIncreasing && A[i-1] > A[i] {
			isIncreasing = false
			continue
		}

		if !isIncreasing && A[i-1] < A[i] {
			return false
		}
	}

	if isIncreasing {
		return false
	}
	return true
}
