package main

func firstMissingPositive(nums []int) int {
	for i, j := 0, len(nums); i < j; i++ {
		k := nums[i] - 1
		if 0 <= k && k < j && nums[k] != nums[i] {
			nums[i], nums[k] = nums[k], nums[i]
			i--
		}
	}
	for i, j := 0, len(nums); i < j; i++ {
		if nums[i] != i+1 {
			return i + 1
		}
	}
	return len(nums) + 1
}
