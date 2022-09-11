package main

import (
	"container/heap"
	"math"
	"sort"
)

type worker struct {
	quality, wage int
}
type maxHeap []worker

func (h *maxHeap) Len() int           { return len(*h) }
func (h *maxHeap) Less(i, j int) bool { return (*h)[i].quality > (*h)[j].quality }
func (h *maxHeap) Swap(i, j int)      { (*h)[i], (*h)[j] = (*h)[j], (*h)[i] }
func (h *maxHeap) Push(x interface{}) { *h = append(*h, x.(worker)) }
func (h *maxHeap) Pop() interface{} {
	x := (*h)[len(*h)-1]
	*h = (*h)[0 : len(*h)-1]
	return x
}

func mincostToHireWorkers(quality []int, wage []int, k int) float64 {
	var workers []worker
	for i := 0; i < len(quality); i++ {
		workers = append(workers, worker{quality: quality[i], wage: wage[i]})
	}
	getRate := func(w worker) float64 {
		return float64(w.wage) / float64(w.quality)
	}
	calcCost := func(quality int, rate float64) float64 {
		return float64(quality) * rate
	}
	sort.Slice(workers, func(i, j int) bool {
		return getRate(workers[i]) < getRate(workers[j])
	})

	var workerHeap maxHeap
	var sumQ int
	for i := 0; i < k-1; i++ {
		heap.Push(&workerHeap, workers[i])
		sumQ += workers[i].quality
	}
	result := math.MaxFloat64
	for i := k - 1; i < len(quality); i++ {
		sumQ += workers[i].quality
		if cost := calcCost(sumQ, getRate(workers[i])); result > cost {
			result = cost
		}
		workerHeap.Push(workers[i])
		sumQ -= heap.Pop(&workerHeap).(worker).quality
	}
	return result
}
