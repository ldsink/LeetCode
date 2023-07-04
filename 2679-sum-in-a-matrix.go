package main

import "sort"

func matrixSum(nums [][]int) int {
	for i := 0; i < len(nums); i++ {
		sort.Ints(nums[i])
	}
	var result int
	for i := 0; i < len(nums[0]); i++ {
		j := nums[0][i]
		for k := 1; k < len(nums); k++ {
			if j < nums[k][i] {
				j = nums[k][i]
			}
		}
		result += j
	}
	return result
}
