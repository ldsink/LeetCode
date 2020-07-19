package main

import (
	"container/heap"
	"sort"
)

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

func nthSuperUglyNumber(n int, primes []int) int {
	if n == 1 {
		return 1
	}
	sort.Ints(primes)
	const max int = 1 << 31
	inHeap := make(map[int]bool)
	minHeap := MinHeap{}
	for _, num := range primes {
		heap.Push(&minHeap, num)
		inHeap[num] = true
	}
	for count := 2; ; count++ {
		num := heap.Pop(&minHeap).(int)
		if count == n {
			return num
		}
		for _, p := range primes {
			v := p * num
			if inHeap[v] {
				continue
			}
			inHeap[v] = true
			heap.Push(&minHeap, v)
		}
	}
}
