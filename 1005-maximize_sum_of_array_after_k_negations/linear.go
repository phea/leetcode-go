package lc

import "sort"

// Time: O(n)
// Benchmark: 4ms 2.7mb

func largestSumAfterKNegations(A []int, K int) int {
	sort.Ints(A)
	var pos int
	for K > 0 {
		if A[pos] == 0 {
			break
		}

		if pos < len(A)-1 {
			if (A[pos] < 0 && A[pos+1] < 0) || A[pos]*-1 > A[pos+1] {
				A[pos] *= -1
				pos++
			} else {
				A[pos] *= -1
			}
		} else {
			A[pos] *= -1
		}

		K--
	}

	var total int
	for _, n := range A {
		total += n
	}

	return total

}
