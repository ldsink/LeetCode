package main

import (
	"container/heap"
	"sort"
)

type Arrive [][2]int

func (h Arrive) Len() int { return len(h) }
func (h Arrive) Less(i, j int) bool {
	return h[i][0] < h[j][0]
}
func (h Arrive) Swap(i, j int) { h[i], h[j] = h[j], h[i] }
func (h *Arrive) Push(x interface{}) {
	*h = append(*h, x.([2]int))
}
func (h *Arrive) Pop() interface{} {
	x := (*h)[len(*h)-1]
	*h = (*h)[0 : len(*h)-1]
	return x
}

func carPooling(trips [][]int, capacity int) bool {
	sort.Slice(trips, func(i, j int) bool {
		return trips[i][1] < trips[j][1]
	})
	var passenger int
	h := &Arrive{}
	for _, trip := range trips {
		num, src, dst := trip[0], trip[1], trip[2]
		for h.Len() > 0 && (*h)[0][0] <= src { // 有乘客下车
			data := heap.Pop(h).([2]int)
			passenger -= data[1]
		}
		if passenger+num > capacity {
			return false
		}
		passenger += num
		heap.Push(h, [2]int{dst, num})
	}
	return true
}
