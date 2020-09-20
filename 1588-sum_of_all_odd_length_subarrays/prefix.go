package lc

// Time: O(n)
// Benchmark: 0ms 2.2mb | 100%

func sumOddLengthSubarrays(arr []int) int {
	if len(arr) == 0 {
		return 0
	}
	prefix := make([]int, len(arr))
	prefix[0] = arr[0]
	for i := 1; i < len(prefix); i++ {
		prefix[i] = prefix[i-1] + arr[i]
	}

	var sum int
	for i := 1; i <= len(arr); i += 2 {
		for j := 0; i+j <= len(arr); j++ {
			if j == 0 {
				sum += prefix[j+i-1]
			} else {
				sum += prefix[j+i-1] - prefix[j-1]
			}
		}
	}

	return sum
}
