package main

import "sort"

func maximumGap(nums []int) int {
	if len(nums) < 2 {
		return 0
	}
	sort.Ints(nums)
	var result int
	for i := 0; i < len(nums)-1; i++ {
		if r := nums[i+1] - nums[i]; result < r {
			result = r
		}
	}
	return result
}
