package lc

// Time: O(n)
// Benchmark: 108ms 7.4mb | 100%

type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeInBetween(list1 *ListNode, a int, b int, list2 *ListNode) *ListNode {
	n := list1
	var startRef, endRef *ListNode
	for i := 0; i <= b; i++ {
		if i == a-1 {
			startRef = n
		}

		if i == b {
			endRef = n
		}

		n = n.Next
	}

	startRef.Next = list2
	n = list2
	for n.Next != nil {
		n = n.Next
	}

	n.Next = endRef.Next
	return list1
}
