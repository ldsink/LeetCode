package main

import "math"

func minSwap(nums1 []int, nums2 []int) int {
	type Condition struct {
		num, other, count int
	}
	var prev, curr []Condition
	prev = append(prev, Condition{num: nums1[0], other: nums2[0], count: 0})
	prev = append(prev, Condition{num: nums2[0], other: nums1[0], count: 1})
	for i := 1; i < len(nums1); i++ {
		curr = []Condition{}
		node := Condition{num: nums1[i], other: nums2[i], count: math.MaxInt}
		for _, cond := range prev {
			if node.num > cond.num && node.other > cond.other && node.count > cond.count {
				node.count = cond.count
			}
		}
		if node.count != math.MaxInt {
			curr = append(curr, node)
		}

		node = Condition{num: nums2[i], other: nums1[i], count: math.MaxInt}
		for _, cond := range prev {
			if node.num > cond.num && node.other > cond.other && node.count > cond.count {
				node.count = cond.count
			}
		}
		if node.count != math.MaxInt {
			node.count += 1 // 如果用 nums2 的值，意味着这里有一次交换
			curr = append(curr, node)
		}
		prev = curr
	}
	result := math.MaxInt
	for _, cond := range prev {
		if result > cond.count {
			result = cond.count
		}
	}
	return result
}
