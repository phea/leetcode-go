package lc

// Time: O(n)
// Benchmark: 12ms 5.1mb | 85%

func majorityElementB(nums []int) []int {
	tally := make(map[int]int)
	for _, n := range nums {
		tally[n]++
	}

	resp := []int{}
	for n, c := range tally {
		if c > len(nums)/3 {
			resp = append(resp, n)
		}
	}

	return resp
}
