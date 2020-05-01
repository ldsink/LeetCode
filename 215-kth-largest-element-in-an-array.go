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

func findKthLargest(nums []int, k int) int {
	mode := true // true for Kth largest, false for Kth smallest
	if k > len(nums)>>1 {
		k = len(nums) + 1 - k
		mode = !mode
	}

	if mode {
		h := &MinHeap{}
		for i := 0; i < k; i++ {
			heap.Push(h, nums[i])
		}
		for i := k; i < len(nums); i++ {
			if nums[i] > (*h)[0] {
				(*h)[0] = nums[i]
				heap.Fix(h, 0)
			}
		}
		return (*h)[0]
	} else {
		h := &MaxHeap{}
		for i := 0; i < k; i++ {
			heap.Push(h, nums[i])
		}
		for i := k; i < len(nums); i++ {
			if nums[i] < (*h)[0] {
				(*h)[0] = nums[i]
				heap.Fix(h, 0)
			}
		}
		return (*h)[0]
	}
}
