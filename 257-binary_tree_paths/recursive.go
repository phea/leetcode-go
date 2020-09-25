package lc

import (
	"fmt"
	"strconv"
)

// Time: O(n)
// Benchmark: 0ms 2.5mb | 100% 23%

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func walk(node *TreeNode, s string, ch chan string) {
	if node == nil {
		return
	}

	if node.Left == nil && node.Right == nil {
		ch <- s + strconv.Itoa(node.Val)
	}

	s2 := fmt.Sprintf("%d->", node.Val)
	walk(node.Left, s+s2, ch)
	walk(node.Right, s+s2, ch)
}

func binaryTreePaths(root *TreeNode) []string {
	ch := make(chan string)
	go func() {
		defer close(ch)
		walk(root, "", ch)
	}()

	paths := []string{}
	for s := range ch {
		paths = append(paths, s)
	}
	return paths
}
