package lc

// Time: O(n)
// Benchmark: 0ms 2mb | 100%

func checkZeroOnes(s string) bool {
	var oneC, zeroC int
	var oneMax, zeroMax int

	for _, ch := range s {
		if ch == '0' {
			oneC = 0
			zeroC++
			if zeroC > zeroMax {
				zeroMax = zeroC
			}
		} else {
			zeroC = 0
			oneC++
			if oneC > oneMax {
				oneMax = oneC
			}
		}
	}

	return oneMax > zeroMax
}
