package lc

import (
	"strconv"
	"strings"
)

// Time: O(n)
// Benchmark: 0ms 2mb | 100%

func compareVersion(version1 string, version2 string) int {
	v1 := strings.Split(version1, ".")
	v2 := strings.Split(version2, ".")

	var d1, d2 int
	for i, digit := range v1 {
		d1, _ = strconv.Atoi(digit)
		if i > len(v2)-1 {
			d2 = 0
		} else {
			d2, _ = strconv.Atoi(v2[i])
		}

		if d1 < d2 {
			return -1
		}

		if d1 > d2 {
			return 1
		}
	}

	if len(v2) > len(v1) {
		for i := len(v1); i < len(v2); i++ {
			d2, _ = strconv.Atoi(v2[i])
			if d2 > 0 {
				return -1
			}
		}
	}

	return 0
}
