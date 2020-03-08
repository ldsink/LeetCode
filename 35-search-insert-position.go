package main

func searchInsert(nums []int, target int) int {
	for left, right := 0, len(nums)-1; left <= right; {
		if left == right {
			if target <= nums[left] {
				return left
			}
			return left + 1
		}
		middle := (left + right) >> 1
		if target <= nums[middle] {
			right = middle
		} else {
			left = middle + 1
		}
	}
	return 0
}
