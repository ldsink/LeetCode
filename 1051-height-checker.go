package main

import "sort"

func heightChecker(heights []int) int {
	var sorted []int
	sorted = append(sorted, heights...)
	sort.Ints(sorted)
	var result int
	for idx, height := range heights {
		if height != sorted[idx] {
			result++
		}
	}
	return result
}
