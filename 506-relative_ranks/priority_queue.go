package lc

import (
	"container/heap"
	"strconv"
)

// Time: O(n)
// Benchmark: 12ms 5.4mb | 100% 96%

type elem struct {
	value, idx int
}

type queue []*elem

func (q queue) Len() int           { return len(q) }
func (q queue) Less(i, j int) bool { return q[i].value > q[j].value }
func (q queue) Swap(i, j int)      { q[i], q[j] = q[j], q[i] }
func (q *queue) Push(x interface{}) {
	item := x.(*elem)
	*q = append(*q, item)
}

func (q *queue) Pop() interface{} {
	old := *q
	n := len(old)
	item := old[n-1]
	old[n-1] = nil // avoid memory leak
	*q = old[0 : n-1]
	return item
}

func findRelativeRanks(nums []int) []string {
	q := make(queue, len(nums))
	for i, n := range nums {
		q[i] = &elem{n, i}
	}
	heap.Init(&q)

	rank := make([]string, len(nums))
	for i := 0; i < len(nums); i++ {
		r := heap.Pop(&q).(*elem)
		if i > 2 {
			rank[r.idx] = strconv.Itoa(i + 1)
		} else if i == 0 {
			rank[r.idx] = "Gold Medal"
		} else if i == 1 {
			rank[r.idx] = "Silver Medal"
		} else {
			rank[r.idx] = "Bronze Medal"
		}
	}

	return rank
}
