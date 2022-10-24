package main

import (
	"sort"
)

func partitionDisjoint(nums []int) int {
	numsZipped := make([][2]int, len(nums))
	for i := 0; i < len(nums); i++ {
		numsZipped[i][0], numsZipped[i][1] = nums[i], i
	}
	sort.Slice(numsZipped, func(i, j int) bool {
		return numsZipped[i][0] < numsZipped[j][0] || (numsZipped[i][0] == numsZipped[j][0] && numsZipped[i][1] < numsZipped[j][1])
	})
	max, idx := nums[0], 1
	for j := 0; max > numsZipped[j][0]; j++ {
		if idx > numsZipped[j][1] {
			continue
		}
		for i := idx; i <= numsZipped[j][1]; i++ {
			if max < nums[i] {
				max = nums[i]
			}
		}
		idx = numsZipped[j][1] + 1
	}
	return idx
}
