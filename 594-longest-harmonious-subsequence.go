package main

import "sort"

func findLHS(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	sort.Ints(nums)
	var result, prevNum, prevCount int
	currNum := nums[0]
	currCount := 1
	for i := 1; i < len(nums); i++ {
		if nums[i] == currNum {
			currCount++
			continue
		}
		if prevCount > 0 && prevNum == currNum-1 {
			if result < prevCount+currCount {
				result = prevCount + currCount
			}
		}
		prevNum = currNum
		prevCount = currCount
		currNum = nums[i]
		currCount = 1
	}
	if prevCount > 0 && prevNum == currNum-1 {
		if result < prevCount+currCount {
			result = prevCount + currCount
		}
	}
	return result
}
