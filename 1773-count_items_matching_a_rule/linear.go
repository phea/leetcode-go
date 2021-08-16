package lc

// Time: O(n)
// Benchmark: 24ms 7.1mb | 100%

func countMatches(items [][]string, ruleKey string, ruleValue string) int {
	var i int
	switch ruleKey {
	case "type":
		i = 0
	case "color":
		i = 1
	case "name":
		i = 2
	}

	var total int
	for _, item := range items {
		if item[i] == ruleValue {
			total++
		}
	}
	return total
}
