package main

type MyCircularDeque struct {
	queue                []int
	head, tail, count, k int
	getIndex             func(int) int
}

func Constructor(k int) MyCircularDeque {
	var getIndex func(int) int
	getIndex = func(i int) int {
		if i >= 0 {
			return i % k
		}
		return getIndex(i + (1+i/k)*k)
	}
	return MyCircularDeque{k: k, queue: make([]int, k), getIndex: getIndex}
}

func (this *MyCircularDeque) InsertFront(value int) bool {
	if this.count == this.k {
		return false
	}
	this.count++
	this.queue[this.getIndex(this.head)] = value
	this.head++
	return true
}

func (this *MyCircularDeque) InsertLast(value int) bool {
	if this.count == this.k {
		return false
	}
	this.count++
	this.tail--
	this.queue[this.getIndex(this.tail)] = value
	return true
}

func (this *MyCircularDeque) DeleteFront() bool {
	if this.count == 0 {
		return false
	}
	this.count--
	this.head--
	return true
}

func (this *MyCircularDeque) DeleteLast() bool {
	if this.count == 0 {
		return false
	}
	this.count--
	this.tail++
	return true

}

func (this *MyCircularDeque) GetFront() int {
	if this.count == 0 {
		return -1
	}
	return this.queue[this.getIndex(this.head-1)]
}

func (this *MyCircularDeque) GetRear() int {
	if this.count == 0 {
		return -1
	}
	return this.queue[this.getIndex(this.tail)]
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
