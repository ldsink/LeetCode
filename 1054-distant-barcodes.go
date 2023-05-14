package main

import (
	"container/heap"
)

type MaxHeap [][2]int

func (h *MaxHeap) Len() int           { return len(*h) }
func (h *MaxHeap) Less(i, j int) bool { return (*h)[i][0] > (*h)[j][0] }
func (h *MaxHeap) Swap(i, j int)      { (*h)[i], (*h)[j] = (*h)[j], (*h)[i] }
func (h *MaxHeap) Push(x interface{}) {
	*h = append(*h, x.([2]int))
}
func (h *MaxHeap) Pop() interface{} {
	x := (*h)[len(*h)-1]
	*h = (*h)[0 : len(*h)-1]
	return x
}

func rearrangeBarcodes(barcodes []int) []int {
	count := make(map[int]int)
	for _, barcode := range barcodes {
		count[barcode]++
	}
	h := &MaxHeap{}
	for b, c := range count {
		heap.Push(h, [2]int{c, b})
	}
	rePush := func(x [2]int) {
		x[0]--
		if x[0] > 0 {
			heap.Push(h, x)
		}
	}
	result := make([]int, len(barcodes))
	n := heap.Pop(h).([2]int)
	result[0] = n[1]
	rePush(n)
	for i := 1; i < len(barcodes); i++ {
		if a := heap.Pop(h).([2]int); a[1] == result[i-1] {
			b := heap.Pop(h).([2]int)
			result[i] = b[1]
			rePush(b)
			heap.Push(h, a)
		} else {
			result[i] = a[1]
			rePush(a)
		}
	}
	return result
}
