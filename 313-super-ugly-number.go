package main

import (
	"container/heap"
)

type MinHeap [][3]int // num, base, index

func (h *MinHeap) Len() int           { return len(*h) }
func (h *MinHeap) Less(i, j int) bool { return (*h)[i][0] < (*h)[j][0] }
func (h *MinHeap) Swap(i, j int)      { (*h)[i], (*h)[j] = (*h)[j], (*h)[i] }
func (h *MinHeap) Push(x interface{}) {
	*h = append(*h, x.([3]int))
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
	heap.Push(h, [3]int{primes[0], 1, 0})
	for i := 2; ; {
		curr := heap.Pop(h).([3]int)
		if inHeap[curr[0]] {
			if curr[2]+1 < len(primes) {
				heap.Push(h, [3]int{curr[1] * primes[curr[2]+1], curr[1], curr[2] + 1})
			}
			continue
		}
		if i == n {
			return curr[0]
		}
		inHeap[curr[0]] = true
		heap.Push(h, [3]int{curr[0] * primes[0], curr[0], 0})
		if curr[2]+1 < len(primes) {
			heap.Push(h, [3]int{curr[1] * primes[curr[2]+1], curr[1], curr[2] + 1})
		}
		i++
	}
}
