package main

func alternatingSubarray(nums []int) int {
	plus := make([]int, len(nums))
	minus := make([]int, len(nums))
	for i := len(nums) - 2; i >= 0; i-- {
		if nums[i]+1 == nums[i+1] {
			plus[i] = 1 + minus[i+1]
		} else if nums[i]-1 == nums[i+1] {
			minus[i] = 1 + plus[i+1]
		}
	}
	var result int
	for i := 0; i < len(nums); i++ {
		if result < plus[i] {
			result = plus[i]
		}
	}
	if result == 0 {
		return -1
	}
	return result + 1
}
