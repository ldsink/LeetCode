package main

import "sort"

func maximumUnits(boxTypes [][]int, truckSize int) int {
	sort.Slice(boxTypes, func(i, j int) bool {
		return boxTypes[i][1] > boxTypes[j][1]
	})
	min := func(a, b int) int {
		if a > b {
			return b
		}
		return a
	}
	var result int
	for i := 0; i < len(boxTypes) && truckSize > 0; i++ {
		num := min(truckSize, boxTypes[i][0])
		result += num * boxTypes[i][1]
		truckSize -= num
	}
	return result
}
