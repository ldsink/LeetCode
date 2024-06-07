package main

func maxOperations(nums []int) int {
	result, score := 1, nums[0]+nums[1]
	for i := 2; i+1 < len(nums) && nums[i]+nums[i+1] == score; i, result = i+2, result+1 {
	}
	return result
}
