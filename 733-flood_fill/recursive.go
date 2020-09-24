package lc

// Time: O(n)
// Benchmark: 8ms 4.2mb | 73% 75%

func floodFill(image [][]int, sr int, sc int, newColor int) [][]int {
	m := len(image)
	n := len(image[0])
	origColor := image[sr][sc]
	if origColor == newColor {
		return image
	}

	var fill func(row, col int)
	fill = func(row, col int) {
		image[row][col] = newColor
		if row != 0 && image[row-1][col] == origColor {
			fill(row-1, col)
		}
		if row != m-1 && image[row+1][col] == origColor {
			fill(row+1, col)
		}
		if col != 0 && image[row][col-1] == origColor {
			fill(row, col-1)
		}
		if col != n-1 && image[row][col+1] == origColor {
			fill(row, col+1)
		}
	}
	fill(sr, sc)
	return image
}
