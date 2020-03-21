package main

func maxSubArray(nums []int) int {
	result := nums[0]
	current := result
	for i, j := 1, len(nums); i < j; i++ {
		if current < 0 {
			current = nums[i]
		} else if current+nums[i] >= 0 {
			current += nums[i]
		} else {
			current = nums[i]
		}
		if result < current {
			result = current
		}
	}
	return result
}
