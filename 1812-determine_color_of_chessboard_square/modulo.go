package lc

// Time: O(1)
// Benchmark: 0ms 2mb | 100%

func squareIsWhite(coordinates string) bool {
	x := coordinates[0] - 'a'
	y := coordinates[1] - '1'

	return (x%2 == 0 && y%2 == 1) || (x%2 == 1 && y%2 == 0)
}
