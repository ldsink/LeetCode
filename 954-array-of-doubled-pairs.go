package main

import (
	"sort"
)

func abs(x int) int {
	if x >= 0 {
		return x
	}
	return -x
}

func canReorderDoubled(arr []int) bool {
	count := make(map[int]int)
	for _, num := range arr {
		count[num]++
	}
	if count[0]%2 == 1 {
		return false
	}

	var nums []int
	for num := range count {
		nums = append(nums, num)
	}
	sort.Slice(nums, func(i, j int) bool { return abs(nums[i]) < abs(nums[j]) })

	for _, num := range nums {
		if count[2*num] < count[num] {
			return false
		}
		count[2*num] -= count[num]
	}
	return true
}
