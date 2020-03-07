package main

func search(nums []int, target int) int {
	for left, right := 0, len(nums)-1; left <= right; {
		if left == right {
			if nums[left] == target {
				return left
			}
			return -1
		}
		middle := (left + right) >> 1
		// left is sorted, right is mixed
		if nums[left] <= nums[middle] {
			if nums[left] <= target && target <= nums[middle] {
				right = middle
				continue
			} else {
				left = middle + 1
				continue
			}
		} else { // left is mixed, right is sorted
			if nums[middle] < target && target <= nums[right] {
				left = middle + 1
				continue
			} else {
				right = middle
				continue
			}
		}
	}
	return -1
}
