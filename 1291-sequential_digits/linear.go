package lc

import (
	"fmt"
	"io/ioutil"
	"math"
)

// Time: O(n)
// Benchmark: 0ms 2.1mb | 100%

func sequentialDigits(low int, high int) []int {
	makeSequence := func(digits int) int {
		var n int
		for i := 0; i < digits; i++ {
			n = (n * 10) + i + 1
		}
		return n
	}

	digits, _ := fmt.Fprintf(ioutil.Discard, "%d", low)

	delta := 1
	for i := 0; i < digits-1; i++ {
		delta = (delta * 10) + 1
	}

	resp := []int{}
	n := makeSequence(digits)
	power := int(math.Pow10(digits - 1))
	for n <= high {
		if n >= low && n/power <= 10-digits {
			resp = append(resp, n)
		}
		n += delta
		if n/power >= 10 {
			digits++
			n = makeSequence(digits)
			power *= 10
			delta = (delta * 10) + 1
		}
	}

	return resp
}
