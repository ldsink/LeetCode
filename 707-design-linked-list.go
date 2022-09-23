package main

type MyLinkedList struct {
	arr []int
}

func Constructor() MyLinkedList {
	return MyLinkedList{}
}

func (this *MyLinkedList) Get(index int) int {
	if index >= len(this.arr) {
		return -1
	}
	return this.arr[index]
}

func (this *MyLinkedList) AddAtHead(val int) {
	this.arr = append([]int{val}, this.arr...)
}

func (this *MyLinkedList) AddAtTail(val int) {
	this.arr = append(this.arr, val)
}

func (this *MyLinkedList) AddAtIndex(index int, val int) {
	if index < 0 {
		this.AddAtHead(val)
	} else if index > len(this.arr) {
		return
	} else if index == len(this.arr) {
		this.AddAtTail(val)
	} else {
		newArr := append([]int{}, this.arr[:index]...)
		newArr = append(newArr, val)
		newArr = append(newArr, this.arr[index:]...)
		this.arr = newArr
	}
}

func (this *MyLinkedList) DeleteAtIndex(index int) {
	if index < 0 || index >= len(this.arr) {
		return
	}
	this.arr = append(this.arr[:index], this.arr[index+1:]...)
}

/**
 * Your MyLinkedList object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Get(index);
 * obj.AddAtHead(val);
 * obj.AddAtTail(val);
 * obj.AddAtIndex(index,val);
 * obj.DeleteAtIndex(index);
 */
