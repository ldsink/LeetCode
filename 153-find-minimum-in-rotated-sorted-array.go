package main

func findMin(nums []int) int {
	if len(nums) == 1 {
		return nums[0]
	} else if len(nums) == 2 {
		if nums[0] > nums[1] {
			return nums[1]
		} else {
			return nums[0]
		}
	}

	if nums[0] < nums[len(nums)-1] {
		return nums[0]
	}
	middle := len(nums) >> 1
	if nums[middle] < nums[len(nums)-1] {
		if left := findMin(nums[:middle]); left > nums[middle] {
			return nums[middle]
		} else {
			return left
		}
	}
	if right := findMin(nums[middle:]); right > nums[0] {
		return nums[0]
	} else {
		return right
	}
}
