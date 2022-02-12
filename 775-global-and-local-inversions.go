package main

func isIdealPermutation(nums []int) bool {
	minNum := nums[len(nums)-1]           // 数组后面部分的最小值
	for i := len(nums) - 3; i >= 0; i-- { // 从后往前遍历
		if nums[i] > minNum { // 检查是否有多余的全局倒置
			return false
		}
		if minNum > nums[i+1] { // 更新后面部分的最小值
			minNum = nums[i+1]
		}
	}
	return true
}
