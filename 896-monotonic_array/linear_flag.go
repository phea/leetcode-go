package lc

// Time: 0(n)
// This solution uses a flag to determine the direction of the array.

func isMonotonicB(A []int) bool {
	if len(A) <= 1 {
		return true
	}

	var dir int
	for i := 1; i < len(A); i++ {
		if dir == 0 {
			if A[i] > A[i-1] {
				dir = 1
			}

			if A[i] < A[i-1] {
				dir = -1
			}
		}

		if dir > 0 && A[i] < A[i-1] {
			return false
		}

		if dir < 0 && A[i] > A[i-1] {
			return false
		}
	}
	return true
}
