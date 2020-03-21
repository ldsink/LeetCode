package main

import (
	"sort"
)

func _permute(nums []int) [][]int {
	var result [][]int
	if len(nums) == 0 {
		return result
	}
	if len(nums) == 1 {
		return append(result, []int{nums[0]})
	}

	for i := 0; i < len(nums); i++ {
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		num := nums[i]
		newNums := append([]int{}, nums[:i]...)
		newNums = append(newNums, nums[i+1:]...)
		for _, r := range _permute(newNums) {
			row := append([]int{num}, r...)
			result = append(result, row)
		}
	}
	return result
}

func permuteUnique(nums []int) [][]int {
	sort.Ints(nums)
	return _permute(nums)
}
