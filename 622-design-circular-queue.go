package main

type MyCircularQueue struct {
	queue                []int
	head, tail, count, k int
}

func Constructor(k int) MyCircularQueue {
	return MyCircularQueue{k: k, queue: make([]int, k)}
}

func (this *MyCircularQueue) EnQueue(value int) bool {
	if this.count == this.k {
		return false
	}
	this.count++
	this.queue[this.tail] = value
	this.tail++
	if this.tail == this.k {
		this.tail = 0
	}
	return true
}

func (this *MyCircularQueue) DeQueue() bool {
	if this.count == 0 {
		return false
	}
	this.count--
	this.head++
	if this.head == this.k {
		this.head = 0
	}
	return true
}

func (this *MyCircularQueue) Front() int {
	if this.count == 0 {
		return -1
	}
	return this.queue[this.head]
}

func (this *MyCircularQueue) Rear() int {
	if this.count == 0 {
		return -1
	}
	idx := this.tail - 1
	if idx < 0 {
		idx = this.k - 1
	}
	return this.queue[idx]
}

func (this *MyCircularQueue) IsEmpty() bool {
	return this.count == 0
}

func (this *MyCircularQueue) IsFull() bool {
	return this.count == this.k
}

/**
 * Your MyCircularQueue object will be instantiated and called as such:
 * obj := Constructor(k);
 * param_1 := obj.EnQueue(value);
 * param_2 := obj.DeQueue();
 * param_3 := obj.Front();
 * param_4 := obj.Rear();
 * param_5 := obj.IsEmpty();
 * param_6 := obj.IsFull();
 */
