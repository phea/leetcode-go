package lc

// Time: O(n)
// Benchmark: 0ms 2.6mb | 100% 29%

type TreeNode struct {
	Val         int
	Left, Right *TreeNode
}

func walk(node *TreeNode, ch chan int) {
	if node == nil {
		return
	}
	walk(node.Left, ch)
	if node.Left == nil && node.Right == nil {
		ch <- node.Val
	}
	walk(node.Right, ch)
}

func leafSimilar(root1 *TreeNode, root2 *TreeNode) bool {
	ch1 := make(chan int)
	go func() {
		defer close(ch1)
		walk(root1, ch1)
	}()

	ch2 := make(chan int)
	go func() {
		defer close(ch2)
		walk(root2, ch2)
	}()

	for n1 := range ch1 {
		n2, ok := <-ch2
		if n2 != n1 || !ok {
			return false
		}
	}

	_, ok := <-ch2
	if ok {
		return false
	}

	return true
}
