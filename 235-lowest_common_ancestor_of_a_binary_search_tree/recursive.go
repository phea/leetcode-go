package lc

// Time: O(logn)
// Benchmark: 20ms 6.7mb | 92% 81%

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	if p.Val > q.Val {
		p, q = q, p
	}

	if root.Val >= p.Val && root.Val <= q.Val {
		return root
	} else if root.Val > q.Val {
		return lowestCommonAncestor(root.Left, p, q)
	} else {
		return lowestCommonAncestor(root.Right, p, q)
	}
}
