package lc

// Time: O(n^2)
// Benchmark: 0ms 2.9mb | 100%

func commonChars(A []string) []string {
	tally := make([]int, 26)
	for _, ch := range A[0] {
		tally[ch-'a']++
	}

	for i := 1; i < len(A); i++ {
		t2 := make([]int, 26)
		for _, ch := range A[i] {
			t2[ch-'a']++
		}

		for i, t := range tally {
			if t2[i] < t {
				tally[i] = t2[i]
			}
		}
	}

	common := []string{}
	for ch, total := range tally {
		for total > 0 {
			total--
			common = append(common, string('a'+ch))
		}
	}
	return common
}
