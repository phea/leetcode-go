package lc

// Time: O(n)
// Benchmark: 12ms 5.2mb | 85%

// Implementation of Boyer-Moore algorithm to determine majority element.
func majorityElement(nums []int) []int {
	var c1, c1Total int
	var c2, c2Total int

	for _, n := range nums {
		if c1 == n {
			c1Total++
		} else if c2 == n {
			c2Total++
		} else if c1Total == 0 {
			c1 = n
			c1Total = 1
		} else if c2Total == 0 {
			c2 = n
			c2Total = 1
		} else {
			c1Total--
			c2Total--
		}
	}

	c1Total, c2Total = 0, 0
	for _, n := range nums {
		if n == c1 {
			c1Total++
		}

		if n == c2 {
			c2Total++
		}
	}

	resp := []int{}
	if c1Total > len(nums)/3 {
		resp = append(resp, c1)
	}

	if c2 != c1 && c2Total > len(nums)/3 {
		resp = append(resp, c2)
	}

	return resp
}
