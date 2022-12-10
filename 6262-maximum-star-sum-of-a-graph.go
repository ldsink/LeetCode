package main

import (
	"math"
	"sort"
)

func maxStarSum(vals []int, edges [][]int, k int) int {
	linked := make([][]int, len(vals))
	for _, edge := range edges {
		a, b := edge[0], edge[1]
		linked[a] = append(linked[a], b)
		linked[b] = append(linked[b], a)
	}
	result := -math.MaxInt
	for i := 0; i < len(vals); i++ {
		array := linked[i]
		sort.Slice(array, func(i, j int) bool {
			return vals[array[i]] > vals[array[j]]
		})
		val := vals[i]
		for j := 0; j < len(array) && j < k && vals[array[j]] > 0; j++ {
			val += vals[array[j]]
		}
		if result < val {
			result = val
		}
	}
	return result
}
