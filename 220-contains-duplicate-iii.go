package main

import (
	"sort"
)

type ByNum [][2]int

func (a ByNum) Len() int           { return len(a) }
func (a ByNum) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a ByNum) Less(i, j int) bool { return a[i][0] < a[j][0] }

func abs(n int) int {
	if n >= 0 {
		return n
	}
	return -n
}

func containsNearbyAlmostDuplicate(nums []int, k int, t int) bool {
	var list ByNum
	for i := 0; i < len(nums); i++ {
		list = append(list, [2]int{nums[i], i})
	}
	sort.Sort(list)
	for i := 0; i < len(list); i++ {
		for j := i + 1; j < len(list) && list[j][0]-list[i][0] <= t; j++ {
			if abs(list[j][1]-list[i][1]) <= k {
				return true
			}
		}
	}
	return false
}
