package main

func maximumDifference(nums []int) int {
	if len(nums) < 2 {
		return -1
	}
	prevMax := nums[len(nums)-1]
	maxDiff := -1
	for i := len(nums) - 2; i >= 0; i-- {
		if prevMax < nums[i] {
			prevMax = nums[i]
		} else if prevMax > nums[i] {
			if diff := prevMax - nums[i]; maxDiff < diff {
				maxDiff = diff
			}
		}
	}
	return maxDiff
}
