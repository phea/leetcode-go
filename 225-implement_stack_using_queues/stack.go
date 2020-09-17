package lc

// Benchmark: 0ms 2.1mb | 100%

type MyStack struct {
	head  int
	stack []int
}

/** Initialize your data structure here. */
func Constructor() MyStack {
	return MyStack{head: 0, stack: []int{}}
}

/** Push element x onto stack. */
func (this *MyStack) Push(x int) {
	this.stack = append(this.stack, x)
	this.head++
}

/** Removes the element on top of the stack and returns that element. */
func (this *MyStack) Pop() int {
	elem := this.stack[this.head-1]
	this.stack = this.stack[:this.head-1]
	this.head--
	return elem
}

/** Get the top element. */
func (this *MyStack) Top() int {
	return this.stack[this.head-1]
}

/** Returns whether the stack is empty. */
func (this *MyStack) Empty() bool {
	return this.head == 0
}

/**
* Your MyStack object will be instantiated and called as such:
* obj := Constructor();
* obj.Push(x);
* param_2 := obj.Pop();
* param_3 := obj.Top();
* param_4 := obj.Empty();
 */
