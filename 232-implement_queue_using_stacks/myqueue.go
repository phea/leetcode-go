package lc

type MyQueue struct {
  q []int
}

/** Initialize your data structure here. */
func Constructor() MyQueue {
  return MyQueue{q: []int{}}
}

/** Push element x to the back of queue. */
func (this *MyQueue) Push(x int)  {
  this.q = append(this.q, x)
}

/** Removes the element from in front of queue and returns that element. */
func (this *MyQueue) Pop() int {
  x := this.q[0]
  this.q = this.q[1:]
  return x
}

/** Get the front element. */
func (this *MyQueue) Peek() int {
  return this.q[0]
}

/** Returns whether the queue is empty. */
func (this *MyQueue) Empty() bool {
  return len(this.q) == 0
}