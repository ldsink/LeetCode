package main

import (
	"container/heap"
	"fmt"
)

type Val struct {
	val, idx int
}
type MaxHeap []Val

func (h MaxHeap) Len() int { return len(h) }
func (h MaxHeap) Less(i, j int) bool {
	return h[i].val > h[j].val || (h[i].val == h[j].val && h[i].idx < h[j].idx)
}
func (h MaxHeap) Swap(i, j int) { h[i], h[j] = h[j], h[i] }
func (h *MaxHeap) Push(x interface{}) {
	*h = append(*h, x.(Val))
}
func (h *MaxHeap) Pop() interface{} {
	x := (*h)[len((*h))-1]
	*h = (*h)[0 : len((*h))-1]
	return x
}

func maxSumAfterPartitioning(arr []int, k int) int {
	var maxHeaps []*MaxHeap
	for i := 0; i <= k; i++ {
		maxHeaps = append(maxHeaps, &MaxHeap{})
	}
	maxSum := make([]int, len(arr))
	getMaxVal := func(h *MaxHeap, i int) int { // 获得满足这个 i 限制下的最大的值，同时清理 heap
		for (*h)[0].idx <= i {
			heap.Pop(h)
		}
		return (*h)[0].val
	}
	for idx, val := range arr {
		v := Val{val: val, idx: idx}
		for j := 1; j <= k; j++ {
			heap.Push(maxHeaps[j], v) // 在这个长度里面，当前点加进去
		}
		for j := 1; j <= k && idx-j >= -1; j++ {
			prevIdx := idx - j
			r := j * getMaxVal(maxHeaps[j], prevIdx)
			if prevIdx >= 0 {
				r += maxSum[prevIdx]
			}
			if maxSum[idx] < r {
				maxSum[idx] = r
			}
		}
	}
	return maxSum[len(arr)-1]
}

func main() {
	fmt.Println(maxSumAfterPartitioning([]int{1, 15, 7, 9, 2, 5, 10}, 3))
}
