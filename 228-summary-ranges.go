package main

import (
	"fmt"
	"strconv"
)

func summaryRanges(nums []int) []string {
	var ranges []string
	var i, j int
	for i = 0; i < len(nums); i++ {
		for j = 1; i+j < len(nums) && nums[i]+j == nums[i+j]; j++ {
		}
		j--
		if j == 0 {
			ranges = append(ranges, strconv.Itoa(nums[i]))
		} else {
			ranges = append(ranges, fmt.Sprintf("%s->%s", strconv.Itoa(nums[i]), strconv.Itoa(nums[i]+j)))
			i += j
		}
	}
	return ranges
}
