package main

func minOperations(nums []int) int {
	var result int
	for i := 1; i < len(nums); i++ {
		if nums[i] <= nums[i-1] {
			result += nums[i-1] + 1 - nums[i]
			nums[i] = nums[i-1] + 1
		}
	}
	return result
}
