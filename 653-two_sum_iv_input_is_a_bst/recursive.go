package lc

// Time: O(n)
// Benchmark: 24ms 6.3mb | 68% 30%

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func walk(node *TreeNode, k int, vals *[]int) bool {
	if node == nil {
		return false
	}

	for _, v := range *vals {
		if v == k-node.Val {
			return true
		}
	}

	*vals = append(*vals, node.Val)
	return walk(node.Left, k, vals) || walk(node.Right, k, vals)

}

func findTarget(root *TreeNode, k int) bool {
	var vals []int
	return walk(root, k, &vals)
}
