package main

import "sort"

func countPairs(nums []int, target int) int {
	sort.Ints(nums)
	var result int
	for i, j := 0, len(nums)-1; i < len(nums); i++ {
		for ; j > i && nums[i]+nums[j] >= target; j-- {
		}
		if j == i {
			break
		}
		result += j - i
	}
	return result
}
