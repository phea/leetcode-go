package lc

// Time: O(n)
// Benchmark: 8ms 3.6mb | 32% 100%

func validateStackSequences(pushed []int, popped []int) bool {
	var p int
	for _, target := range popped {
		for i, n := range pushed {
			if n == target {
				if p-i > 0 {
					return false
				}
				p = i - 1

				pushed = append(pushed[:i], pushed[i+1:]...)
				break
			}
		}
	}

	return true
}
