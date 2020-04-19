package main

func findMin(nums []int) int {
	if len(nums) == 1 {
		return nums[0]
	}

	middle := (len(nums) - 1) >> 1
	if nums[0] > nums[middle] || nums[middle+1] < nums[len(nums)-1] {
		// left mixed
		if left := findMin(nums[:middle+1]); left > nums[middle+1] {
			return nums[middle+1]
		} else {
			return left
		}
	} else if nums[0] < nums[middle] || nums[middle+1] > nums[len(nums)-1] {
		// right mixed
		if right := findMin(nums[middle+1:]); right > nums[0] {
			return nums[0]
		} else {
			return right
		}
	}
	if left, right := findMin(nums[:middle+1]), findMin(nums[middle+1:]); left > right {
		return right
	} else {
		return left
	}
}
