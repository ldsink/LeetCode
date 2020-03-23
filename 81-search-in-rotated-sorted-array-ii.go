package main

func searchInRange(left, right int, nums []int, target int) bool {
	if left == right {
		return nums[left] == target
	}
	middle := (left + right) >> 1
	var leftMix, rightMix bool
	if nums[left] < nums[middle] || nums[middle+1] > nums[right] {
		rightMix = true
	} else if nums[left] > nums[middle] || nums[middle+1] < nums[right] {
		leftMix = true
	}

	if leftMix {
		if nums[middle+1] <= target && target <= nums[right] {
			return searchInRange(middle+1, right, nums, target)
		} else {
			return searchInRange(left, middle, nums, target)
		}
	} else if rightMix {
		if nums[left] <= target && target <= nums[middle] {
			return searchInRange(left, middle, nums, target)
		} else {
			return searchInRange(middle+1, right, nums, target)
		}
	}
	if ok := searchInRange(left, middle, nums, target); ok {
		return true
	}
	if ok := searchInRange(middle+1, right, nums, target); ok {
		return true
	}
	return false
}

func search(nums []int, target int) bool {
	if len(nums) == 0 {
		return false
	}
	return searchInRange(0, len(nums)-1, nums, target)
}
