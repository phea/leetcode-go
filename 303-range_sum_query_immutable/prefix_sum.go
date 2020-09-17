package lc

// Time: O(1) for SumRange()
// Benchmark: 56ms 9.6mb | 37% 17%

type NumArray struct {
	sums []int
}

func Constructor(nums []int) NumArray {
	if len(nums) == 0 {
		return NumArray{[]int{}}
	}

	// precalculate the sums.
	sums := make([]int, len(nums))
	sums[0] = nums[0]
	for i := 1; i < len(nums); i++ {
		sums[i] = sums[i-1] + nums[i]
	}
	return NumArray{sums}
}

func (this *NumArray) SumRange(i int, j int) int {
	if i == 0 {
		return this.sums[j]
	}
	return this.sums[j] - this.sums[i-1]
}
