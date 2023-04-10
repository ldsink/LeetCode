package main

import "container/heap"

type Node struct{ idx, val int }
type MinHeap []Node

func (h MinHeap) Len() int           { return len(h) }
func (h MinHeap) Less(i, j int) bool { return h[i].val < h[j].val }
func (h MinHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h *MinHeap) Push(x interface{}) {
	*h = append(*h, x.(Node))
}
func (h *MinHeap) Pop() interface{} {
	x := (*h)[len(*h)-1]
	*h = (*h)[0 : len(*h)-1]
	return x
}

func nextLargerNodes(head *ListNode) []int {
	var result []int
	h := &MinHeap{}
	for idx := 0; head != nil; head = head.Next {
		result = append(result, 0)
		for len(*h) > 0 && (*h)[0].val < head.Val {
			node := heap.Pop(h).(Node)
			result[node.idx] = head.Val
		}
		heap.Push(h, Node{idx: idx, val: head.Val})
		idx++
	}
	return result
}
