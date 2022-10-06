package main

func maxAscendingSum(nums []int) int {
	result, curr := nums[0], nums[0]
	for i := 0; i < len(nums)-1; i++ {
		if nums[i] < nums[i+1] {
			curr += nums[i+1]
		} else {
			curr = nums[i+1]
		}
		if result < curr {
			result = curr
		}
	}
	return result
}
