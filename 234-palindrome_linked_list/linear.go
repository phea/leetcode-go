package lc

// Time: O(n)
// Benchmark: 12ms 7.1mb | 93% 20%

type ListNode struct {
	Val  int
	Next *ListNode
}

func isPalindrome(head *ListNode) bool {
	orig := head
	nums := []int{}
	for orig != nil {
		nums = append(nums, orig.Val)
		orig = orig.Next
	}

	// reversing linked-list
	var tail *ListNode
	for head != nil {
		head.Next, tail, head = tail, head, head.Next
	}

	for _, n := range nums {
		if n != tail.Val {
			return false
		}
		tail = tail.Next
	}

	return true
}
