package lc

// Time: O(n)
// Benchmark: 0ms 2mb | 100%

func minTimeToType(word string) int {
	var total int
	var pos, d1, d2 int
	for _, ch := range word {
		v := int(ch - 'a')
		if v >= pos {
			d1 = v - pos
			d2 = pos + 26 - v
		} else {
			d1 = 26 - pos + v
			d2 = pos - v
		}

		if d1 < d2 {
			total += d1 + 1
		} else {
			total += d2 + 1
		}
		pos = v
	}
	return total
}
