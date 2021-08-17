package lc

// Time: O(n)
// Benchmark: 4ms 5.1 | 91% 86%

func minOperations2(boxes string) []int {
	var ln, rn int
	resp := make([]int, len(boxes))
	for i, ch := range boxes {
		if ch == '1' {
			resp[0] += i
			rn++
		}
	}

	for i := 1; i < len(resp); i++ {
		if boxes[i-1] == '1' {
			rn--
			ln++
		}

		resp[i] = resp[i-1] - rn + ln
	}
	return resp
}
