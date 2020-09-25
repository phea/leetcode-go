package lc

// Time: O(n)
// Benchmark: 0ms 2.2mb | 100%

func subsets(nums []int) [][]int {
	subs := [][]int{}
	var search func(set []int, k int)
	search = func(set []int, k int) {
		if k == len(nums) {
			subs = append(subs, append([]int{}, set...))
			return
		}
		set = append(set, nums[k])
		search(set, k+1)
		set = set[:len(set)-1]
		search(set, k+1)
	}
	search([]int{}, 0)

	return subs
}
