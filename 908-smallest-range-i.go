package main

import "math"

func smallestRangeI(nums []int, k int) int {
	max, min := math.MinInt, math.MaxInt
	for _, num := range nums {
		if max < num {
			max = num
		}
		if min > num {
			min = num
		}
	}
	if result := max - min - 2*k; result > 0 {
		return result
	}
	return 0
}
