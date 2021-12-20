package main

import (
	"sort"
)

func abs(a int) int {
	if a > 0 {
		return a
	}
	return -a
}

func findRadius(houses []int, heaters []int) int {
	sort.Ints(houses)
	sort.Ints(heaters)
	var idx, radius int
	for i := 0; i < len(houses); i++ {
		// 找到房子 i 最近的供暖器，idx 只增不减
		for ; idx+1 < len(heaters) && abs(heaters[idx+1]-houses[i]) <= abs(heaters[idx]-houses[i]); idx++ {
		}
		if r := abs(heaters[idx] - houses[i]); radius < r {
			radius = r
		}
	}
	return radius
}
