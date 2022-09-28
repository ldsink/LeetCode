package main

import (
	"container/heap"
)

type UglyNumber struct {
	num, base, idx int
}
type MinHeap []UglyNumber

func (h *MinHeap) Len() int           { return len(*h) }
func (h *MinHeap) Less(i, j int) bool { return (*h)[i].num < (*h)[j].num }
func (h *MinHeap) Swap(i, j int)      { (*h)[i], (*h)[j] = (*h)[j], (*h)[i] }
func (h *MinHeap) Push(x interface{}) {
	*h = append(*h, x.(UglyNumber))
}
func (h *MinHeap) Pop() interface{} {
	x := (*h)[len(*h)-1]
	*h = (*h)[0 : len(*h)-1]
	return x
}

func nthSuperUglyNumber(n int, primes []int) int {
	if n == 1 {
		return 1
	}
	inHeap := make(map[int]bool)
	h := &MinHeap{}
	heap.Push(h, UglyNumber{num: primes[0], base: 1, idx: 0})
	for i := 2; ; {
		curr := heap.Pop(h).(UglyNumber)
		if inHeap[curr.num] {
			if curr.idx+1 < len(primes) {
				heap.Push(h, UglyNumber{num: curr.base * primes[curr.idx+1], base: curr.base, idx: curr.idx + 1})
			}
			continue
		}
		if i == n {
			return curr.num
		}
		inHeap[curr.num] = true
		heap.Push(h, UglyNumber{num: curr.num * primes[0], base: curr.num, idx: 0})
		if curr.idx+1 < len(primes) {
			heap.Push(h, UglyNumber{num: curr.base * primes[curr.idx+1], base: curr.base, idx: curr.idx + 1})
		}
		i++
	}
}
