package main

func countSubarrays(nums []int, k int) int {
	var upperCount, lowerCount, idx int
	for idx = 0; idx < len(nums) && nums[idx] != k; idx++ {
	}
	over := make(map[int]int)
	over[0] = 1
	for i := idx - 1; i >= 0; i-- {
		if nums[i] > k {
			upperCount++
		} else {
			lowerCount++
		}
		over[upperCount-lowerCount]++
	}
	result := over[0] + over[1]
	upperCount, lowerCount = 0, 0
	for i := idx + 1; i < len(nums); i++ {
		if nums[i] > k {
			upperCount++
		} else {
			lowerCount++
		}
		diff := lowerCount - upperCount
		result += over[diff] + over[diff+1]
	}
	return result
}
