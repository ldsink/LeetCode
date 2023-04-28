package main

import "container/heap"

type MinHeap []int

func (h *MinHeap) Len() int           { return len(*h) }
func (h *MinHeap) Less(i, j int) bool { return (*h)[i] < (*h)[j] }
func (h *MinHeap) Swap(i, j int) {
	(*h)[i], (*h)[j] = (*h)[j], (*h)[i]
}
func (h *MinHeap) Push(x interface{}) {
	*h = append(*h, x.(int))
}
func (h *MinHeap) Pop() interface{} {
	x := (*h)[len(*h)-1]
	*h = (*h)[:len(*h)-1]
	return x
}

type DinnerPlates struct {
	plates     [][]int
	capacity   int
	max        int // 当前所在的最右边的位置
	candidates MinHeap
}

func Constructor(capacity int) DinnerPlates {
	return DinnerPlates{capacity: capacity, plates: [][]int{{}}}
}

func (this *DinnerPlates) getPushIndex() int {
	idx := -1
	for len(this.candidates) > 0 {
		idx = heap.Pop(&this.candidates).(int)
		// 失效的情况：
		// 1. >= max
		// 2. 这个位置已经满了
		if idx >= this.max || len(this.plates[idx]) == this.capacity {
			idx = -1
			continue
		}
		break
	}
	return idx
}

func (this *DinnerPlates) Push(val int) {
	// 从候选中尝试取出一个空位置
	if idx := this.getPushIndex(); idx != -1 {
		this.plates[idx] = append(this.plates[idx], val)
		return
	}
	// 前面都是满的，需要往最后加
	if len(this.plates[this.max]) >= this.capacity {
		this.max++
		if len(this.plates) <= this.max {
			this.plates = append(this.plates, []int{})
		}
	}
	this.plates[this.max] = append(this.plates[this.max], val)
}

func (this *DinnerPlates) Pop() int {
	if this.max == 0 && len(this.plates[0]) == 0 {
		return -1
	}
	if len(this.plates[this.max]) == 0 {
		this.max--
		return this.Pop()
	}
	val := this.plates[this.max][len(this.plates[this.max])-1]
	this.plates[this.max] = this.plates[this.max][:len(this.plates[this.max])-1]
	return val
}

func (this *DinnerPlates) PopAtStack(index int) int {
	if len(this.plates) <= index || len(this.plates[index]) == 0 {
		return -1
	}
	val := this.plates[index][len(this.plates[index])-1]
	this.plates[index] = this.plates[index][:len(this.plates[index])-1]
	// 把 index 增加到 candidates 里面
	heap.Push(&this.candidates, index)
	return val
}

/**
 * Your DinnerPlates object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * obj.Push(val);
 * param_2 := obj.Pop();
 * param_3 := obj.PopAtStack(index);
 */
