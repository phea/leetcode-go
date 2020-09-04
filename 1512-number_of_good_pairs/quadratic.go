package lc

// Time: O(n^2)
// This solution should be as quick or even quicker than using a hash table when 'nums' is small.

func numIdenticalPairsB(nums []int) int {
	var c int
	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			if nums[i] == nums[j] {
				c++
			}
		}
	}

	return c
}
