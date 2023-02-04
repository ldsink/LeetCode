package main

import "sort"

func getMaximumConsecutive(coins []int) int {
	sort.Ints(coins)
	var x int
	for _, coin := range coins {
		if x+1 >= coin {
			x += coin
		} else {
			break
		}
	}
	return x + 1
}
