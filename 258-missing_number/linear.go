package lc

// Time: O(n)
// If we have distinct numbers from 0,1..n then that means we can XOR the index with the number.

func missingNumber(nums []int) int {
	res := len(nums)
	for i, n := range nums {
		res = res ^ i ^ n
	}

	return res
}
