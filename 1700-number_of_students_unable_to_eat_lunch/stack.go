package lc

// Time: O(n)
// Benchmark: 0ms 2.4mb | 100% 33%

func countStudents(students []int, sandwiches []int) int {
	f := true
	var ptr int
	for f {
		f = false
		round := len(students)
		for i := 0; i < round; i++ {
			if students[0] == sandwiches[ptr] {
				students = students[1:]
				ptr++
				f = true
			} else {
				students = append(students[1:], students[0])
			}
		}
	}

	return len(students)
}
