package main

import (
	"container/heap"
)

type BuyHeap [][2]int

func (h *BuyHeap) Len() int           { return len(*h) }
func (h *BuyHeap) Less(i, j int) bool { return (*h)[i][0] > (*h)[j][0] }
func (h *BuyHeap) Swap(i, j int)      { (*h)[i], (*h)[j] = (*h)[j], (*h)[i] }
func (h *BuyHeap) Push(x interface{}) {
	*h = append(*h, x.([2]int))
}
func (h *BuyHeap) Pop() interface{} {
	x := (*h)[len(*h)-1]
	*h = (*h)[0 : len(*h)-1]
	return x
}

type SellHeap [][2]int

func (h *SellHeap) Len() int           { return len(*h) }
func (h *SellHeap) Less(i, j int) bool { return (*h)[i][0] < (*h)[j][0] }
func (h *SellHeap) Swap(i, j int)      { (*h)[i], (*h)[j] = (*h)[j], (*h)[i] }
func (h *SellHeap) Push(x interface{}) {
	*h = append(*h, x.([2]int))
}
func (h *SellHeap) Pop() interface{} {
	x := (*h)[len(*h)-1]
	*h = (*h)[0 : len(*h)-1]
	return x
}

func getNumberOfBacklogOrders(orders [][]int) int {
	buyHeap := &BuyHeap{}
	sellHeap := &SellHeap{}
	const buy = 0
	const sell = 1

	var processOrder func(int, int, int)
	processOrder = func(price, amount, orderType int) {
		if amount == 0 {
			return
		}
		if orderType == buy {
			if (*sellHeap).Len() == 0 || (*sellHeap)[0][0] > price {
				heap.Push(buyHeap, [2]int{price, amount})
			} else {
				sellOrder := heap.Pop(sellHeap).([2]int)
				if sellOrder[1] >= amount {
					processOrder(sellOrder[0], sellOrder[1]-amount, sell)
				} else {
					processOrder(price, amount-sellOrder[1], buy)
				}
			}
		} else { // sell
			if (*buyHeap).Len() == 0 || (*buyHeap)[0][0] < price {
				heap.Push(sellHeap, [2]int{price, amount})
			} else {
				buyOrder := heap.Pop(buyHeap).([2]int)
				if buyOrder[1] >= amount {
					processOrder(buyOrder[0], buyOrder[1]-amount, buy)
				} else {
					processOrder(price, amount-buyOrder[1], sell)
				}
			}
		}
	}

	for _, order := range orders {
		price, amount, orderType := order[0], order[1], order[2]
		processOrder(price, amount, orderType)
	}
	var result int
	const mod = 1e9 + 7
	for i := 0; i < (*sellHeap).Len(); i++ {
		result += (*sellHeap)[i][1]
		result %= mod
	}
	for i := 0; i < (*buyHeap).Len(); i++ {
		result += (*buyHeap)[i][1]
		result %= mod
	}
	return result
}
