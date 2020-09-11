package lc

// Time: O(n)
// Benchmark: 0ms 2.2mb | 100% 75%

func containsPattern(arr []int, m int, k int) bool {
	equal := func(a, b []int) bool {
		for i, val := range a {
			if val != b[i] {
				return false
			}
		}

		return true
	}

	var pattern []int
	for i := 0; i+(m*k) <= len(arr); i++ {
		count := 1
		pattern = arr[i : i+m]
		for j := i + m; j < len(arr); j += m {
			if !equal(pattern, arr[j:j+m]) {
				break
			}

			count++
		}
		if count == k {
			return true
		}
	}

	return false
}
