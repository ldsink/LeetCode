package main

func removeElement(nums []int, val int) int {
	var i, j int
	for i, j = 0, 0; i < len(nums); i++ {
		if nums[i] != val {
			nums[j] = nums[i]
			j++
		}
	}
	return j
}
