package main

import "sort"

func findMissingAndRepeatedValues(grid [][]int) []int {
	var total, a int
	var nums []int
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid); j++ {
			total += grid[i][j]
			nums = append(nums, grid[i][j])
		}
	}
	n := len(grid) * len(grid)
	delta := ((n * (n + 1)) >> 1) - total
	sort.Ints(nums)
	for i := 0; i < len(nums); i++ {
		if nums[i] == nums[i+1] {
			a = nums[i]
			break
		}
	}
	return []int{a, a + delta}
}
