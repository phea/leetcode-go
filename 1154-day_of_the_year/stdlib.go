package lc

import "time"

// Benchmark: 0ms 2mb | 100% 60%

func dayOfYear(date string) int {
	t, _ := time.Parse("2006-01-02", date)
	start := time.Date(t.Year(), 1, 1, 0, 0, 0, 0, time.UTC)

	return int(t.Sub(start).Round(time.Hour).Hours())/24 + 1
}
