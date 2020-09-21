package lc

// Time: O(n logn)
// Benchmark: 160ms 6.7mb | 94% 85%

func distributeCandies(candies []int) int {
	sis := make(map[int]int8)
	var count int
	for _, v := range candies {
		_, ok := sis[v]
		if !ok {
			sis[v] = 1
			count++
			if count > len(candies)/2 {
				return len(candies) / 2
			}
		}
	}

	return count
}
