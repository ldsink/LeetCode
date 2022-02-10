package main

import "sort"

func minimumDifference(nums []int, k int) int {
	sort.Ints(nums)
	result := 100000 + 1
	for i := 0; i < len(nums)-k+1; i++ {
		if j := nums[i+k-1] - nums[i]; result > j {
			result = j
		}
	}
	return result
}
