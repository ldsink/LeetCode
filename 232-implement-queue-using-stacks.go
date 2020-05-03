package main

type MyQueue struct {
	stack []int
}

/** Initialize your data structure here. */
func Constructor() MyQueue {
	return MyQueue{}
}

/** Push element x to the back of queue. */
func (this *MyQueue) Push(x int) {
	this.stack = append(this.stack, x)
}

/** Removes the element from in front of queue and returns that element. */
func (this *MyQueue) Pop() int {
	var tempStack []int
	for ; len(this.stack) > 1; this.stack = this.stack[:len(this.stack)-1] {
		tempStack = append(tempStack, this.stack[len(this.stack)-1])
	}
	x := this.stack[0]
	this.stack = []int{}
	for ; len(tempStack) > 0; tempStack = tempStack[:len(tempStack)-1] {
		this.stack = append(this.stack, tempStack[len(tempStack)-1])
	}
	return x
}

/** Get the front element. */
func (this *MyQueue) Peek() int {
	var tempStack []int
	for ; len(this.stack) > 1; this.stack = this.stack[:len(this.stack)-1] {
		tempStack = append(tempStack, this.stack[len(this.stack)-1])
	}
	x := this.stack[0]
	for ; len(tempStack) > 0; tempStack = tempStack[:len(tempStack)-1] {
		this.stack = append(this.stack, tempStack[len(tempStack)-1])
	}
	return x
}

/** Returns whether the queue is empty. */
func (this *MyQueue) Empty() bool {
	return len(this.stack) == 0
}

/**
 * Your MyQueue object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * param_2 := obj.Pop();
 * param_3 := obj.Peek();
 * param_4 := obj.Empty();
 */
