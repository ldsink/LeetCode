package main

func maximumElementAfterDecrementingAndRearranging(arr []int) int {
	nums := make([]int, len(arr)+1)
	for i := 0; i < len(arr); i++ {
		if arr[i] < len(nums) {
			nums[arr[i]]++
		} else {
			nums[len(nums)-1]++
		}
	}
	over := 0  // 多出的数字
	waste := 0 // 浪费的数字
	for i := len(nums) - 1; i > 0; i-- {
		if over+nums[i] > 0 {
			over += nums[i] - 1
		} else {
			waste++
		}
	}
	return len(arr) - waste
}
