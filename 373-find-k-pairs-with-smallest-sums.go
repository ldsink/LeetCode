package main

import (
	"container/heap"
)

type MinHeap struct {
	nums1, nums2 []int
	nodes        [][2]int
}

func (h MinHeap) Len() int { return len(h.nodes) }
func (h MinHeap) Less(i, j int) bool {
	return h.nums1[h.nodes[i][0]]+h.nums2[h.nodes[i][1]] < h.nums1[h.nodes[j][0]]+h.nums2[h.nodes[j][1]]
}
func (h MinHeap) Swap(i, j int) { h.nodes[i], h.nodes[j] = h.nodes[j], h.nodes[i] }
func (h *MinHeap) Push(x interface{}) {
	(*h).nodes = append((*h).nodes, x.([2]int))
}
func (h *MinHeap) Pop() interface{} {
	x := (*h).nodes[len((*h).nodes)-1]
	(*h).nodes = (*h).nodes[0 : len((*h).nodes)-1]
	return x
}

func kSmallestPairs(nums1 []int, nums2 []int, k int) [][]int {
	var result [][]int
	if len(nums1) == 0 || len(nums2) == 0 {
		return result
	}
	h := &MinHeap{nums1: nums1, nums2: nums2}
	h.Push([2]int{0, 0}) // 初始化，添加第一个最小的
	for i := 0; i < k && h.Len() > 0; i++ {
		node := heap.Pop(h).([2]int)
		result = append(result, []int{nums1[node[0]], nums2[node[1]]})
		if node[1] == 0 && node[0]+1 < len(nums1) { // 刚好是 nums1 某个元素的第一行，把下一个元素也带进去
			heap.Push(h, [2]int{node[0] + 1, 0})
		}
		if node[1]+1 < len(nums2) {
			heap.Push(h, [2]int{node[0], node[1] + 1})
		}
	}
	return result
}
