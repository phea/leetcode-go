package lc

// Time: O(n logn)
// Benchmark: 524ms 97mb |  39% 68%

func insertIntoBSTB(root *TreeNode, val int) *TreeNode {
	if root == nil {
		return &TreeNode{val, nil, nil}
	}

	node := root
	for node != nil {
		if node.Left != nil && node.Val > val {
			node = node.Left
		} else if node.Right != nil && node.Val < val {
			node = node.Right
		} else if node.Val > val {
			node.Left = &TreeNode{val, nil, nil}
			break
		} else {
			node.Right = &TreeNode{val, nil, nil}
			break
		}
	}

	return root
}
