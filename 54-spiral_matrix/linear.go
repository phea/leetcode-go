package lc

// Time: O(mn)
// Benchmark: 0ms 2mb | 100% 73%

const (
	EAST = iota
	SOUTH
	WEST
	NORTH
)

func spiralOrder(matrix [][]int) []int {
	if len(matrix) == 0 {
		return []int{}
	}

	m := len(matrix)
	n := len(matrix[0])
	nums := make([]int, m*n)
	bounds := []int{n, m, 0, 1}
	direction := EAST
	pos := []int{0, 0}
	for i := 0; i < len(nums); i++ {
		nums[i] = matrix[pos[0]][pos[1]]
		if direction == EAST {
			if pos[1] == bounds[EAST]-1 {
				direction = SOUTH
				pos[0]++
				bounds[EAST]--
			} else {
				pos[1]++
			}
		} else if direction == SOUTH {
			if pos[0] == bounds[SOUTH]-1 {
				direction = WEST
				pos[1]--
				bounds[SOUTH]--
			} else {
				pos[0]++
			}
		} else if direction == WEST {
			if pos[1] == bounds[WEST] {
				direction = NORTH
				pos[0]--
				bounds[WEST]++
			} else {
				pos[1]--
			}
		} else {
			if pos[0] == bounds[NORTH] {
				direction = EAST
				pos[1]++
				bounds[NORTH]++
			} else {
				pos[0]--
			}
		}
	}
	return nums
}
