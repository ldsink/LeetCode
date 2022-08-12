package main

import (
	"sort"
)

func maxChunksToSorted(arr []int) int {
	sorted := make([]int, len(arr))
	for idx, num := range arr {
		sorted[idx] = num
	}
	sort.Ints(sorted)
	var result, a, b int
	for i := 0; i < len(arr); i++ {
		a ^= 1 << arr[i]
		b ^= 1 << sorted[i]
		if a == b {
			result++
		}
	}
	return result
}
