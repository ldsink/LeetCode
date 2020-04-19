package main

type MinStack struct {
	stack    []int
	minStack []int
}

/** initialize your data structure here. */
func Constructor() MinStack {
	return MinStack{}
}

func (this *MinStack) Push(x int) {
	this.stack = append(this.stack, x)
	if len(this.minStack) == 0 || this.stack[this.minStack[len(this.minStack)-1]] > x {
		this.minStack = append(this.minStack, len(this.stack)-1)
	}
}

func (this *MinStack) Pop() {
	end := len(this.stack) - 1
	if end == this.minStack[len(this.minStack)-1] {
		this.minStack = this.minStack[:len(this.minStack)-1]
	}
	this.stack = this.stack[:end]
}

func (this *MinStack) Top() int {
	return this.stack[len(this.stack)-1]
}

func (this *MinStack) GetMin() int {
	return this.stack[this.minStack[len(this.minStack)-1]]
}
