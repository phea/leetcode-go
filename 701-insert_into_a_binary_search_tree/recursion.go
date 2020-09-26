package lc

// Time: O(nlog n)
// Benchmark: 480ms 169.4mb | 88% 15%

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func insert(node *TreeNode, val int) {
	if node == nil {
		node = &TreeNode{val, nil, nil}
	}

	if node.Left != nil && node.Val > val {
		insert(node.Left, val)
	} else if node.Right != nil && node.Val < val {
		insert(node.Right, val)
	} else if node.Val > val {
		node.Left = &TreeNode{val, nil, nil}
	} else {
		node.Right = &TreeNode{val, nil, nil}
	}

}

func insertIntoBST(root *TreeNode, val int) *TreeNode {
	if root == nil {
		return &TreeNode{val, nil, nil}
	}
	insert(root, val)
	return root
}
