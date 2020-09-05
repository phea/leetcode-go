package lc

import (
	"sync"

	"github.com/phea/leetcode-go/types"
)

// Time: O(n)

func walkB(node *types.TreeNode, stack *[]int) {
	if node == nil {
		return
	}

	walk(node.Left, stack)
	*stack = append(*stack, node.Val)
	walk(node.Right, stack)
}

func getAllElementsB(root1 *types.TreeNode, root2 *types.TreeNode) []int {
	stack1, stack2 := []int{}, []int{}
	var wg sync.WaitGroup

	wg.Add(2)
	go func(wg *sync.WaitGroup) {
		walkB(root1, &stack1)
		wg.Done()
	}(&wg)

	go func(wg *sync.WaitGroup) {
		walkB(root2, &stack2)
		wg.Done()
	}(&wg)

	wg.Wait()

	arr := make([]int, len(stack1)+len(stack2))
	pos1, pos2 := 0, 0
	for i := 0; i < len(arr); i++ {
		if pos1 < len(stack1) && pos2 < len(stack2) {
			if stack1[pos1] <= stack2[pos2] {
				arr[i] = stack1[pos1]
				pos1++
			} else {
				arr[i] = stack2[pos2]
				pos2++
			}
		} else if pos1 < len(stack1) {
			arr[i] = stack1[pos1]
			pos1++
		} else {
			arr[i] = stack2[pos2]
			pos2++
		}
	}
	return arr
}
