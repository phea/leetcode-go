package lc

// Time: O(n)

func numIdenticalPairs(nums []int) int {
	var c int
	seen := make(map[int]int, len(nums)-1)
	for _, v := range nums {
		n, ok := seen[v]
		if ok {
			c += n
		}
		seen[v] = n + 1
	}

	return c
}
