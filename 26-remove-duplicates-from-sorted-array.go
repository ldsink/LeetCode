package main

func removeDuplicates(nums []int) int {
	var i, j int
	for i, j = 0, 0; i < len(nums); i++ {
		if nums[j] != nums[i] {
			j++
			nums[j] = nums[i]
		}
	}
	return j + 1
}
