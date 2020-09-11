package lc

import "time"

// Time: O(1)
// Benchmark: 0ms 2mb | 100% 60%

func daysBetweenDates(date1 string, date2 string) int {
	const dateFormat = "2006-01-02"
	d1, _ := time.Parse(dateFormat, date1)
	d2, _ := time.Parse(dateFormat, date2)

	days := int(d2.Sub(d1).Hours()) / 24
	if days > 0 {
		return days
	}

	return days * -1
}
