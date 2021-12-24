package main

import (
	"container/heap"
)

type AppleNode struct {
	apple  int
	rotDay int
}

type MinHeap []AppleNode

func (h MinHeap) Len() int           { return len(h) }
func (h MinHeap) Less(i, j int) bool { return h[i].rotDay < h[j].rotDay }
func (h MinHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h *MinHeap) Push(x interface{}) {
	*h = append(*h, x.(AppleNode))
}
func (h *MinHeap) Pop() interface{} {
	old := *h
	x := old[len(old)-1]
	*h = old[0 : len(old)-1]
	return x
}

func eatenApples(apples []int, days []int) int {
	var result int
	minHeap := MinHeap{}
	for i := 0; i < len(apples); i++ {
		if apples[i] > 0 {
			appleNode := AppleNode{apple: apples[i], rotDay: i + days[i] - 1}
			heap.Push(&minHeap, appleNode)
		}
		for minHeap.Len() > 0 {
			appleNode := heap.Pop(&minHeap).(AppleNode)
			// 已经腐败了，这天的苹果不能吃了
			if appleNode.rotDay < i {
				continue
			}
			result++
			appleNode.apple -= 1
			if appleNode.apple > 0 {
				heap.Push(&minHeap, appleNode)
			}
			break
		}
	}
	for i := len(apples); minHeap.Len() > 0; i++ {
		for minHeap.Len() > 0 {
			appleNode := heap.Pop(&minHeap).(AppleNode)
			// 已经腐败了，这天的苹果不能吃了
			if appleNode.rotDay < i {
				continue
			}
			// 直接尽可能吃多的数量
			num := appleNode.rotDay - i + 1
			if num > appleNode.apple {
				num = appleNode.apple
			}
			result += num
			appleNode.apple -= num
			i += num - 1
			if appleNode.apple > 0 {
				heap.Push(&minHeap, appleNode)
			}
			break
		}
	}
	return result
}
