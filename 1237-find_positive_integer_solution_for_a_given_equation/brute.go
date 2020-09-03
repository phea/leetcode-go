package lc

// Time:    O(n^2)
// Memory:  O(1)

func findSolution(customFunction func(int, int) int, z int) [][]int {
	results := [][]int{}
	for x := 1; x <= z; x++ {
		for y := z; y >= 1; y-- {
			if customFunction(x, y) == z {
				results = append(results, []int{x, y})
			}
		}
	}

	return results
}
