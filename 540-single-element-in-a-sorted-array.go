package main

func singleNonDuplicate(nums []int) int {
	if len(nums) == 1 {
		return nums[0]
	}
	middle := len(nums) >> 1
	if nums[middle-1] == nums[middle] {
		if middle%2 == 0 {
			return singleNonDuplicate(nums[:middle-1])
		} else {
			return singleNonDuplicate(nums[middle+1:])
		}
	} else {
		if middle%2 == 0 {
			return singleNonDuplicate(nums[middle:])
		} else {
			return singleNonDuplicate(nums[:middle])
		}
	}
}
