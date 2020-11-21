package main

import (
	"sort"
)

func minPatches(nums []int, n int) int {
	var count, idx, edge int
	sort.Ints(nums)
	for edge < n {
		if idx == len(nums) {
			edge += edge + 1
			count++
			continue
		}
		if nums[idx] <= edge+1 {
			edge += nums[idx]
			idx++
		} else {
			edge += edge + 1
			count++
		}
	}
	return count
}
