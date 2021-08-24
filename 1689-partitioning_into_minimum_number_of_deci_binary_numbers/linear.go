package lc

// Time: O(n)
// Benchmark: 4ms 6.5mb | 100%

func minPartitions(n string) int {
	var max int
	for _, ch := range n {
		if int(ch-'0') > max {
			if int(ch-'0') == 9 {
				return 9
			}

			max = int(ch - '0')
		}
	}

	return max
}
