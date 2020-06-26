package main

import "container/heap"

type MinHeap []int

func (h MinHeap) Len() int           { return len(h) }
func (h MinHeap) Less(i, j int) bool { return h[i] < h[j] }
func (h MinHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h *MinHeap) Push(x interface{}) {
	*h = append(*h, x.(int))
}
func (h *MinHeap) Pop() interface{} {
	old := *h
	x := old[len(old)-1]
	*h = old[0 : len(old)-1]
	return x
}

type MaxHeap []int

func (h MaxHeap) Len() int           { return len(h) }
func (h MaxHeap) Less(i, j int) bool { return h[i] > h[j] }
func (h MaxHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h *MaxHeap) Push(x interface{}) {
	*h = append(*h, x.(int))
}
func (h *MaxHeap) Pop() interface{} {
	old := *h
	x := old[len(old)-1]
	*h = old[0 : len(old)-1]
	return x
}

type MedianFinder struct {
	minHeap MinHeap
	maxHeap MaxHeap
}

/** initialize your data structure here. */
func Constructor() MedianFinder {
	return MedianFinder{minHeap: MinHeap{}, maxHeap: MaxHeap{}}
}

func (this *MedianFinder) AddNum(num int) {
	if len(this.maxHeap) == 0 || this.maxHeap[0] > num {
		heap.Push(&this.maxHeap, num)
	} else {
		heap.Push(&this.minHeap, num)
	}
	if len(this.maxHeap)-len(this.minHeap) > 1 {
		n := heap.Pop(&this.maxHeap)
		heap.Push(&this.minHeap, n)
	} else if len(this.minHeap)-len(this.maxHeap) > 1 {
		n := heap.Pop(&this.minHeap)
		heap.Push(&this.maxHeap, n)
	}
}

func (this *MedianFinder) FindMedian() float64 {
	if len(this.minHeap) == len(this.maxHeap) {
		return float64(this.minHeap[0]+this.maxHeap[0]) / 2
	} else if len(this.minHeap) > len(this.maxHeap) {
		return float64(this.minHeap[0])
	} else {
		return float64(this.maxHeap[0])
	}
}
