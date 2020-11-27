package lc

// Time: O(n)
// Benchmark: 4ms 3.5mb | 100%

func carPooling(trips [][]int, capacity int) bool {
	arr := []int{}
	for _, t := range trips {
		// this append is expensive.
		if t[2] > len(arr) {
			arr2 := make([]int, t[2]-len(arr)+1)
			arr = append(arr, arr2...)
		}

		v := t[0]
		for i := t[1]; i < t[2]; i++ {
			arr[i] += v
			if arr[i] > capacity {
				return false
			}
		}
	}

	return true
}
