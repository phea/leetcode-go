package lc

import "sort"

// Time: O(n logn)
// Benchmark: 20ms 6.3mb | 91% 28%

func maximumUnits(boxTypes [][]int, truckSize int) int {
	sort.Slice(boxTypes, func(i, j int) bool {
		return boxTypes[i][1] > boxTypes[j][1]
	})

	var total int
	for _, box := range boxTypes {
		if box[0] >= truckSize {
			total += truckSize * box[1]
			break
		}
		total += box[0] * box[1]
		truckSize -= box[0]
	}
	return total
}
