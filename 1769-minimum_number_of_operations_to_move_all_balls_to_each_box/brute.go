package lc

// Time: O(n^2)
// Benchmark: 88ms 5.3mb | 23% 69%

func minOperations(boxes string) []int {
	resp := make([]int, len(boxes))
	for i := 0; i < len(resp); i++ {
		var dist int
		for j := 0; j < len(boxes); j++ {
			if boxes[j] == '1' {
				if j-i < 0 {
					dist += i - j
				} else {
					dist += j - i
				}
			}
		}
		resp[i] = dist
	}

	return resp
}
