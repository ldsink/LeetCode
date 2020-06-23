package main

func findDuplicate(nums []int) int {
	var result int
	for i := uint(0); i < 31; i++ {
		bit := 1 << i
		expected, actual := 0, 0
		for j := 0; j < len(nums); j++ {
			if j&bit > 0 {
				expected++
			}
			if nums[j]&bit > 0 {
				actual++
			}
		}
		if actual > expected {
			result ^= bit
		}
	}
	return result
}
