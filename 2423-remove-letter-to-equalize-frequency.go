package main

import "sort"

func equalFrequency(word string) bool {
	count := make(map[rune]int)
	for _, r := range word {
		count[r]++
	}
	if len(count) == 1 {
		return true
	}
	var nums []int
	for _, c := range count {
		nums = append(nums, c)
	}
	sort.Ints(nums)
	last := len(nums) - 1
	return (nums[0] == nums[last-1] && nums[last-1]+1 == nums[last]) || (nums[0] == 1 && nums[1] == nums[last])
}
