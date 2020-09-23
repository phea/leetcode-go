package lc

// Time: O(n)
// Benchmark: 124ms 7.9mb | 88% 61%

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func walk(node *TreeNode, level int, sums *[]int) {
	if node == nil {
		return
	}

	if level >= len(*sums) {
		*sums = append(*sums, 0)
	}

	walk(node.Left, level+1, sums)
	(*sums)[level] += node.Val
	walk(node.Right, level+1, sums)
}

func maxLevelSum(root *TreeNode) int {
	sums := []int{}
	walk(root, 0, &sums)

	var level, max int
	for i, total := range sums {
		if total > max {
			max = total
			level = i
		}
	}

	return level + 1
}
