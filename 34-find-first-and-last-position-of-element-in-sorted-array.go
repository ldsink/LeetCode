package main

func searchRange(nums []int, target int) []int {
	if len(nums) == 0 {
		return []int{-1, -1}
	}
	var left, right int
	// find left index
	for l, r := 0, len(nums)-1; l <= r; {
		if l == r {
			if nums[l] == target {
				left = l
				break
			}
			return []int{-1, -1}
		}
		m := (l + r) >> 1
		if target <= nums[m] {
			r = m
			continue
		} else {
			l = m + 1
			continue
		}
	}
	// find right index
	for l, r := left, len(nums)-1; l <= r; {
		if l == r {
			right = l
			break
		}
		m := (l + r) >> 1
		if target >= nums[m+1] {
			l = m + 1
			continue
		} else {
			r = m
			continue
		}
	}
	return []int{left, right}
}
