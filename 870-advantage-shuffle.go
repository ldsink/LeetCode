package main

import "sort"

func advantageCount(nums1 []int, nums2 []int) []int {
	type node struct {
		val, idx int
	}
	var nodes []node
	for i := 0; i < len(nums2); i++ {
		nodes = append(nodes, node{val: nums2[i], idx: i})
	}
	sort.Slice(nodes, func(i, j int) bool {
		return nodes[i].val < nodes[j].val
	})
	sort.Ints(nums1)
	result := make([]int, len(nums1))
	for i := 0; i < len(nums1); i++ {
		result[i] = -1
	}
	var rest []int
	for i, j := 0, 0; i < len(nums1); i++ {
		if nums1[i] > nodes[j].val {
			result[nodes[j].idx] = nums1[i]
			j++
		} else {
			rest = append(rest, nums1[i])
		}
	}
	for i := 0; i < len(nums1); i++ {
		if result[i] == -1 {
			result[i] = rest[0]
			rest = rest[1:]
		}
	}
	return result
}
