package lc

// Time: O(n)
// Benchmark: 4ms 3.3mb | 88% 100%

type ListNode struct {
	Val  int
	Next *ListNode
}

func oddEvenList(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}

	start := head
	odd := head
	var even, evenHead *ListNode
	i := 0
	head = head.Next
	for head != nil {
		if i == 0 {
			if even == nil {
				even = head
				evenHead = head
			} else {
				even.Next = head
				even = even.Next
			}
		} else {
			odd.Next = head
			odd = odd.Next
		}
		i++
		if i >= 2 {
			i = 0
		}
		head = head.Next
	}

	if even != nil {
		even.Next = nil
	}

	odd.Next = evenHead
	return start
}
