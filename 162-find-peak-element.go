package main

func findPeakElement(nums []int) int {
	if len(nums) == 1 {
		return 0
	}
	middle := len(nums) >> 1
	if nums[middle-1] < nums[middle] && (middle+1 == len(nums) || (middle+1 < len(nums) && nums[middle] < nums[middle+1])) {
		return middle + findPeakElement(nums[middle:])
	}
	if nums[middle-1] > nums[middle] && (middle+1 == len(nums) || (middle+1 < len(nums) && nums[middle] > nums[middle+1])) {
		return findPeakElement(nums[:middle])
	}
	if nums[middle-1] < nums[middle] && (middle+1 == len(nums) || (middle+1 < len(nums) && nums[middle] > nums[middle+1])) {
		return middle
	}
	return findPeakElement(nums[:middle])
}
