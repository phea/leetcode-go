package lc

// Time: O(n)
// Solution does not use a flag, instead checks the direction of the array by checking the first and last
// elements.

func isMonotonic(A []int) bool {
	if len(A) == 1 {
		return true
	}

	if A[0] >= A[len(A)-1] { // decreasing array
		for i := 1; i < len(A); i++ {
			if A[i] > A[i-1] {
				return false
			}
		}
	} else {
		for i := 1; i < len(A); i++ {
			if A[i] < A[i-1] {
				return false
			}
		}
	}
	return true
}
