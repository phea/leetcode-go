package main

// Time: O(1)
// Benchmark: 0ms 1.9mb | 100% 87%
// Check if x and y coordinates fall within range, and if both X and Y overlap then return true.
// This is a basic collision detector code you find in games.

func isRectangleOverlap(rec1 []int, rec2 []int) bool {
	xOverlap := (rec2[0] >= rec1[0] && rec2[0] < rec1[2]) || (rec1[0] >= rec2[0] && rec1[0] < rec2[2])
	yOverlap := (rec2[1] >= rec1[1] && rec2[1] < rec1[3]) || (rec1[1] >= rec2[1] && rec1[1] < rec2[3])

	return xOverlap && yOverlap
}