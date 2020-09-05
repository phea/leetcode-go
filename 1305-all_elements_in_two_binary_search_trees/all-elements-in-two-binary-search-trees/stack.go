package lc

import "github.com/phea/leetcode-go/types"

// Time: O(n)
// Benchmark: 108ms 7.5mb | 90%

func walk(node *types.TreeNode, stack *[]int) {
	if node == nil {
		return
	}

	walk(node.Left, stack)
	*stack = append(*stack, node.Val)
	walk(node.Right, stack)
}

func getAllElements(root1 *types.TreeNode, root2 *types.TreeNode) []int {
	stack1, stack2 := []int{}, []int{}
	walk(root1, &stack1)
	walk(root2, &stack2)

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
