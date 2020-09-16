package lc

import (
	"math"
	"strings"
)

// Time: O(n)
// Benchmark: 0ms 2.3mb | 100%

func myAtoi(str string) int {
	str = strings.TrimSpace(str)
	if len(str) == 0 {
		return 0
	}

	if !(str[0] >= '0' && str[0] <= '9') && str[0] != '-' && str[0] != '+' {
		return 0
	}

	var num int
	sign := 1
	i := 0
	if str[0] == '-' {
		i = 1
		sign = -1
	}

	if str[0] == '+' {
		i = 1
	}

	for ; i < len(str); i++ {
		if !(str[i] >= '0' && str[i] <= '9') {
			break
		}

		num = (num * 10) + (int(str[i]) - 48)
		if sign == -1 && num*sign <= math.MinInt32 {
			return math.MinInt32
		}

		if sign == 1 && num >= math.MaxInt32 {
			return math.MaxInt32
		}

	}
	return num * sign
}
