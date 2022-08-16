package main

type MyCircularDeque struct {
	queue                []int
	head, tail, count, k int
}

func Constructor(k int) MyCircularDeque {
	return MyCircularDeque{k: k, queue: make([]int, k)}
}

func (this *MyCircularDeque) InsertFront(value int) bool {
	if this.count == this.k {
		return false
	}
	this.count++
	this.queue[this.head] = value
	this.head++
	if this.head == this.k {
		this.head = 0
	}
	return true
}

func (this *MyCircularDeque) InsertLast(value int) bool {
	if this.count == this.k {
		return false
	}
	this.count++
	this.tail--
	if this.tail < 0 {
		this.tail = this.k - 1
	}
	this.queue[this.tail] = value
	return true
}

func (this *MyCircularDeque) DeleteFront() bool {
	if this.count == 0 {
		return false
	}
	this.count--
	this.head--
	if this.head < 0 {
		this.head = this.k - 1
	}
	return true
}

func (this *MyCircularDeque) DeleteLast() bool {
	if this.count == 0 {
		return false
	}
	this.count--
	this.tail++
	if this.tail == this.k {
		this.tail = 0
	}
	return true

}

func (this *MyCircularDeque) GetFront() int {
	if this.count == 0 {
		return -1
	}
	idx := this.head - 1
	if idx < 0 {
		idx = this.k - 1
	}
	return this.queue[idx]
}

func (this *MyCircularDeque) GetRear() int {
	if this.count == 0 {
		return -1
	}
	return this.queue[this.tail]
}

func (this *MyCircularDeque) IsEmpty() bool {
	return this.count == 0
}

func (this *MyCircularDeque) IsFull() bool {
	return this.count == this.k
}

/**
 * Your MyCircularDeque object will be instantiated and called as such:
 * obj := Constructor(k);
 * param_1 := obj.InsertFront(value);
 * param_2 := obj.InsertLast(value);
 * param_3 := obj.DeleteFront();
 * param_4 := obj.DeleteLast();
 * param_5 := obj.GetFront();
 * param_6 := obj.GetRear();
 * param_7 := obj.IsEmpty();
 * param_8 := obj.IsFull();
 */
