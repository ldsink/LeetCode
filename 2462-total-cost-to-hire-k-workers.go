package main

import (
	"container/heap"
)

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

func totalCost(costs []int, k int, candidates int) int64 {
	leftHeap, rightHeap := &MinHeap{}, &MinHeap{}
	var left, right int
	for left = 0; left < candidates; left++ {
		heap.Push(leftHeap, costs[left])
	}
	right = len(costs) - 1
	for i := 1; i <= candidates; i++ {
		if right < left {
			break
		}
		heap.Push(rightHeap, costs[right])
		right--
	}
	var result int
	for i := 0; i < k; i++ {
		if leftHeap.Len() != 0 && (rightHeap.Len() == 0 || (*leftHeap)[0] <= (*rightHeap)[0]) {
			result += heap.Pop(leftHeap).(int)
			if left <= right {
				heap.Push(leftHeap, costs[left])
				left++
			}
		} else {
			result += heap.Pop(rightHeap).(int)
			if left <= right {
				heap.Push(rightHeap, costs[right])
				right--
			}
		}
	}
	return int64(result)
}
