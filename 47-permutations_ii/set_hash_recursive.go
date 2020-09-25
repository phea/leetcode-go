package lc

// Time: O(n)
// Benchmark:  4ms 3.7mb | 75% 55%

func permuteUnique(nums []int) [][]int {
	permuts := [][]int{}

	var permFn func(set []int, k int)
	permFn = func(set []int, k int) {
		if k == len(set) {
			permuts = append(permuts, (append([]int{}, set...)))
			return
		}

		positions := make(map[int]bool)
		for i := k; i < len(nums); i++ {
			_, ok := positions[set[i]]
			if ok {
				continue
			}
			positions[set[i]] = true
			set[k], set[i] = set[i], set[k]
			permFn(set, k+1)
			set[k], set[i] = set[i], set[k]
		}
	}
	permFn(nums, 0)

	return permuts
}
