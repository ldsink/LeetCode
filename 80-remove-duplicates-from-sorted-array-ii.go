package main

func removeDuplicates(nums []int) int {
	var length, count int
	for i := 0; i < len(nums); i++ {
		if length == 0 {
			count = 1
			length++
		} else if nums[i] != nums[length-1] {
			nums[length] = nums[i]
			count = 1
			length++
		} else if count < 2 {
			nums[length] = nums[i]
			count++
			length++
		}
	}
	return length
}
