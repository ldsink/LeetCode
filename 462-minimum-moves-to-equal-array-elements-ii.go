package main

import "sort"

func abs(a int) int {
	if a >= 0 {
		return a
	}
	return -a
}

func minMoves2(nums []int) int {
	sort.Ints(nums)
	x := nums[len(nums)/2]
	var result int
	for _, n := range nums {
		result += abs(n - x)
	}
	return result
}
