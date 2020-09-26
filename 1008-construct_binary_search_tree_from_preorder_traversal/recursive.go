package lc

// Time: O(n)
// Benchmark: 0ms 3.2mb | 100% 68%

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func bstFromPreorder(preorder []int) *TreeNode {
	head := &TreeNode{preorder[0], nil, nil}

	var add func(node *TreeNode, val int)
	add = func(node *TreeNode, val int) {
		if val < node.Val {
			if node.Left == nil {
				node.Left = &TreeNode{val, nil, nil}
			} else {
				add(node.Left, val)
			}
		} else {
			if node.Right == nil {
				node.Right = &TreeNode{val, nil, nil}
			} else {
				add(node.Right, val)
			}
		}

	}

	for i := 1; i < len(preorder); i++ {
		add(head, preorder[i])
	}

	return head
}
