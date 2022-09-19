package main

import "sort"

func frequencySort(nums []int) []int {
	count := make(map[int]int)
	for _, num := range nums {
		count[num]++
	}
	sort.Slice(nums, func(i, j int) bool {
		return count[nums[i]] < count[nums[j]] || (count[nums[i]] == count[nums[j]] && nums[i] > nums[j])
	})
	return nums
}
