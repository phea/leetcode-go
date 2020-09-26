package lc

// Time: O(n+m)
// Benchmark:  8ms 6.2mb | 100% 62%

type Employee struct {
	Id           int
	Importance   int
	Subordinates []int
}

func getImportance(employees []*Employee, id int) int {
	var total int
	var search func(id int)
	search = func(id int) {
		for _, e := range employees {
			if e.Id == id {
				total += e.Importance
				for _, subId := range e.Subordinates {
					search(subId)
				}
			}
		}
	}
	search(id)

	return total
}
