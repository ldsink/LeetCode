package main

import "container/heap"

type MinHeap []int

func (h *MinHeap) Len() int           { return len(*h) }
func (h *MinHeap) Less(i, j int) bool { return (*h)[i] < (*h)[j] }
func (h *MinHeap) Swap(i, j int)      { (*h)[i], (*h)[j] = (*h)[j], (*h)[i] }
func (h *MinHeap) Push(x interface{}) {
	*h = append(*h, x.(int))
}
func (h *MinHeap) Pop() interface{} {
	x := (*h)[len(*h)-1]
	*h = (*h)[0 : len(*h)-1]
	return x
}

func getKthMagicNumber(k int) int {
	if k == 1 {
		return 1
	}

	h := &MinHeap{}
	multi := []int{3, 5, 7}
	for i := 0; i < 3; i++ {
		heap.Push(h, multi[i])
	}

	inHeap := make(map[int]bool)
	addToHeap := func(h *MinHeap, x int) {
		if inHeap[x] {
			return
		}
		inHeap[x] = true
		heap.Push(h, x)
	}

	var num int
	for i := 2; i <= k; i++ {
		num = heap.Pop(h).(int)
		for j := 0; j < 3; j++ {
			addToHeap(h, num*multi[j]) // 扩展下一个值
		}
	}
	return num
}
