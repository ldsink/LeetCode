package main

import (
	"sort"
)

func dp(nums []int, target int, cache []int) int {
	if target == 0 {
		return 1
	}
	if cache[target] != -1 {
		return cache[target]
	}

	var count int
	for i := 0; i < len(nums); i++ {
		if nums[i] > target {
			break
		}
		count += dp(nums, target-nums[i], cache)
	}
	cache[target] = count
	return count
}

func combinationSum4(nums []int, target int) int {
	// 1 <= nums.length <= 200
	// 1 <= target <= 1000
	sort.Ints(nums)
	result := make([]int, 1001)
	for i := 0; i <= 1000; i++ {
		result[i] = -1
	}
	dp(nums, target, result)
	if result[target] == -1 {
		return 0
	}
	return result[target]
}
