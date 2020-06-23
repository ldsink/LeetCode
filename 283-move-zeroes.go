package main

func moveZeroes(nums []int) {
	if len(nums) == 0 {
		return
	}
	var start int
	for i := 0; i < len(nums); i++ {
		if nums[i] == 0 {
			continue
		}
		nums[start] = nums[i]
		start++
	}
	for i := start; i < len(nums); i++ {
		nums[i] = 0
	}
}
