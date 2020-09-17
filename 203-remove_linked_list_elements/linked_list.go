package lc

type ListNode struct {
	Val  int
	Next *ListNode
}

// Benchmark: 12ms

func removeElements(head *ListNode, val int) *ListNode {
	orig := head
	var prev *ListNode
	for head != nil {
		if head.Val == val {
			if prev != nil {
				prev.Next = head.Next
			} else {
				if head.Next != nil {
					orig = head.Next
				} else {
					orig = nil
				}
			}
			head = head.Next
		} else {
			prev = head
			head = head.Next
		}

	}

	return orig
}
