package lc

// Time: O(n)
// Benchmark: 0ms 2mb | 100% 51%

func isLongPressedName(name string, typed string) bool {
	pos, count, required := 0, 0, 1
	for i := 0; i < len(name); i++ {
		if i < len(name)-1 && name[i+1] == name[i] {
			required++
			continue
		}

		for j := pos; j < len(typed); j++ {
			pos = j
			if name[i] != typed[j] {
				break
			}
			count++
		}

		if count < required {
			return false
		}

		count, required = 0, 1
	}

	if pos < len(typed)-1 || typed[pos] != name[len(name)-1] {
		return false
	}

	return true
}
