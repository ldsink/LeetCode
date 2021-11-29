package main

import (
	"container/heap"
)

type SmallestHeap struct {
	arr  []int
	arrF []float64
	data [][2]int
}

func (h SmallestHeap) Len() int { return len(h.data) }
func (h SmallestHeap) Less(i, j int) bool {
	return h.arrF[h.data[i][0]]/h.arrF[h.data[i][1]] < h.arrF[h.data[j][0]]/h.arrF[h.data[j][1]]
}
func (h SmallestHeap) Swap(i, j int) { h.data[i], h.data[j] = h.data[j], h.data[i] }
func (h *SmallestHeap) Push(x interface{}) {
	(*h).data = append((*h).data, x.([2]int))
}
func (h *SmallestHeap) Pop() interface{} {
	x := (*h).data[len((*h).data)-1]
	(*h).data = (*h).data[0 : len((*h).data)-1]
	return x
}

func kthSmallestPrimeFraction(arr []int, k int) []int {
	var arrF []float64
	for _, num := range arr {
		arrF = append(arrF, float64(num))
	}
	// 构造一个最小堆，每个节点为以 arr[i] 开头的最小值的情况。
	// 堆中的节点 [2]int。第一个元素为分子所在的索引，不变。第二个元素为分母所在的索引，会一直下降，直到第一个元素的值。
	h := &SmallestHeap{arr: arr, arrF: arrF}
	for i := 0; i < len(arr)-1; i++ {
		node := [2]int{i, len(arr) - 1}
		heap.Push(h, node)
	}
	// 取 k-1 次
	for i := 1; i < k; i++ {
		node := heap.Pop(h).([2]int)
		node[1]--
		if node[0] < node[1] {
			heap.Push(h, node)
		}
	}
	// 堆顶的元素就是答案的索引
	idx := (*h).data[0]
	return []int{arr[idx[0]], arr[idx[1]]}
}
