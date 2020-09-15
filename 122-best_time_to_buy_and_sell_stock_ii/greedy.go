package lc

// Time: O(n)
// Benchmark: 0ms 3.1mb | 100%

func maxProfit(prices []int) int {
	var profit int
	for i := 0; i < len(prices)-1; i++ {
		if prices[i+1] > prices[i] {
			profit += prices[i+1] - prices[i]
		}
	}

	return profit
}
