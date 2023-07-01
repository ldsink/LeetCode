package main

import "sort"

func twoSum(nums []int, target int) []int {
	type Number struct {
		val, idx int
	}
	var array []Number
	for idx, val := range nums {
		array = append(array, Number{idx: idx, val: val})
	}
	sort.Slice(array, func(i, j int) bool {
		return array[i].val < array[j].val || (array[i].val == array[j].val && array[i].idx < array[j].idx)
	})
	for start, end := 0, len(nums)-1; ; start++ {
		for ; array[start].val+array[end].val > target && start < end-1; end-- {
		}
		for ; array[start].val+array[end].val < target && end+1 < len(nums); end++ {
		}
		if array[start].val+array[end].val == target {
			return []int{array[start].idx, array[end].idx}
		}
	}
}
