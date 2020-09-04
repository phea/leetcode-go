package lc

// Time: O(n)
// Space: O(n)

func shuffle(nums []int, n int) []int {
	n2 := make([]int, len(nums))
	for i := 0; i < n; i++ {
		n2[i*2] = nums[i]
		n2[i*2+1] = nums[i+n]
	}
	return n2
}
