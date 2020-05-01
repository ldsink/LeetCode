package main

func minSubArrayLen(s int, nums []int) int {
	result := len(nums) + 1
	for i, j, v := 0, 0, 0; i < len(nums); i++ {
		v += nums[i]
		if v >= s {
			for ; j < i; j++ {
				if v-nums[j] < s {
					break
				}
				v -= nums[j]
			}
			if result > i-j+1 {
				result = i - j + 1
			}
		}
	}
	if result == len(nums)+1 {
		result = 0
	}
	return result
}
