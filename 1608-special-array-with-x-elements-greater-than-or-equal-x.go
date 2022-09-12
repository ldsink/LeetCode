package main

import "sort"

func specialArray(nums []int) int {
	sort.Ints(nums)
	// 因为是非负整数组，一定不会有 0 个元素大于等于 0。所以 x 在 0 的情况下一定不成立。
	for x := len(nums); x > 0; x-- {
		idx := len(nums) - x
		// 恰好有 x 个元素 大于或者等于 x
		if nums[idx] < x || (idx > 0 && nums[idx-1] >= x) {
			continue
		}
		return x
	}
	return -1
}
