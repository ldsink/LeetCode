package main

import "sort"

func maxIceCream(costs []int, coins int) int {
	sort.Ints(costs)
	var result int
	for i := 0; i < len(costs); i++ {
		if coins < costs[i] {
			break
		}
		result += 1
		coins -= costs[i]
	}
	return result
}
