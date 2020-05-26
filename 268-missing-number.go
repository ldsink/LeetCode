package main

func missingNumber(nums []int) int {
	var result int
	nums = append(nums, len(nums)+1)
	for i := 0; i < len(nums); i++ {
		if i == nums[i] {
		} else if len(nums) == nums[i] {
			result = i
		} else {
			nums[nums[i]], nums[i] = nums[i], nums[nums[i]]
			i-- // continue swap current position until match index or flag
		}
	}
	return result
}
