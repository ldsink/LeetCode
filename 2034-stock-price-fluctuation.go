package main

import (
	"container/heap"
)

type MinHeap [][2]int

func (h MinHeap) Len() int           { return len(h) }
func (h MinHeap) Less(i, j int) bool { return h[i][1] < h[j][1] }
func (h MinHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h *MinHeap) Push(x interface{}) {
	*h = append(*h, x.([2]int))
}
func (h *MinHeap) Pop() interface{} {
	old := *h
	x := old[len(old)-1]
	*h = old[0 : len(old)-1]
	return x
}

type MaxHeap [][2]int

func (h MaxHeap) Len() int           { return len(h) }
func (h MaxHeap) Less(i, j int) bool { return h[i][1] > h[j][1] }
func (h MaxHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h *MaxHeap) Push(x interface{}) {
	*h = append(*h, x.([2]int))
}
func (h *MaxHeap) Pop() interface{} {
	old := *h
	x := old[len(old)-1]
	*h = old[0 : len(old)-1]
	return x
}

type StockPrice struct {
	maxTimestamp int
	prices       map[int]int
	maxPriceHeap MaxHeap
	minPriceHeap MinHeap
}

func Constructor() StockPrice {
	return StockPrice{prices: make(map[int]int), maxPriceHeap: MaxHeap{}, minPriceHeap: MinHeap{}}
}

func (this *StockPrice) Update(timestamp int, price int) {
	if this.maxTimestamp < timestamp {
		this.maxTimestamp = timestamp
	}
	heap.Push(&this.maxPriceHeap, [2]int{timestamp, price})
	heap.Push(&this.minPriceHeap, [2]int{timestamp, price})
	this.prices[timestamp] = price
}

func (this *StockPrice) Current() int {
	return this.prices[this.maxTimestamp]
}

func (this *StockPrice) Maximum() int {
	n := heap.Pop(&this.maxPriceHeap).([2]int)
	if this.prices[n[0]] != n[1] { // 已经失效的节点
		return this.Maximum()
	}
	heap.Push(&this.maxPriceHeap, n)
	return n[1]
}

func (this *StockPrice) Minimum() int {
	n := heap.Pop(&this.minPriceHeap).([2]int)
	if this.prices[n[0]] != n[1] { // 已经失效的节点
		return this.Minimum()
	}
	heap.Push(&this.minPriceHeap, n)
	return n[1]
}

/**
 * Your StockPrice object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Update(timestamp,price);
 * param_2 := obj.Current();
 * param_3 := obj.Maximum();
 * param_4 := obj.Minimum();
 */
