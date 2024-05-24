package main

import (
	"sort"
)

func mostCompetitive(nums []int, k int) []int {
	type Node struct {
		val, idx int
	}
	var arr []Node
	for idx, num := range nums {
		arr = append(arr, Node{idx: idx, val: num})
	}
	sort.Slice(arr, func(i, j int) bool {
		return arr[i].val < arr[j].val || (arr[i].val == arr[j].val && arr[i].idx < arr[j].idx)
	})
	var result []int
	start, prevIdx := 0, -1
	for ; k > 0; k-- {
		for ; arr[start].idx <= prevIdx; start++ {
		}
		var target int
		for i := start; i < len(arr); i++ {
			if arr[i].idx > prevIdx && arr[i].idx+k <= len(nums) {
				target = i
				break
			}
		}
		// 当剩余长度刚好构成长度为 k 的数组
		if arr[target].idx+k == len(nums) {
			for i := arr[target].idx; i < len(nums); i++ {
				result = append(result, nums[i])
			}
			break
		}
		// 否则增加当前位置的最优解
		result = append(result, arr[target].val)
		prevIdx = arr[target].idx
	}
	return result
}
