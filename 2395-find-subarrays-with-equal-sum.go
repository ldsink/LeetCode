package main

func findSubarrays(nums []int) bool {
	exists := make(map[int]bool)
	for i := 0; i < len(nums)-1; i++ {
		if j := nums[i] + nums[i+1]; exists[j] {
			return true
		} else {
			exists[j] = true
		}
	}
	return false
}
