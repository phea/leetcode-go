package lc

import "math"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func walk(node *TreeNode, ch chan int) {
	if node == nil {
		return
	}

	walk(node.Left, ch)
	ch <- node.Val
	walk(node.Right, ch)
}

func findSecondMinimumValue(root *TreeNode) int {
	ch := make(chan int)
	go func() {
		defer close(ch)
		walk(root, ch)
	}()

	min1, min2 := math.MaxInt64, math.MaxInt64
	for v := range ch {
		if v < min1 {
			min1, min2 = v, min1
		} else if v < min2 && v != min1 {
			min2 = v
		}
	}

	if min2 == math.MaxInt64 {
		return -1
	}
	return min2
}
