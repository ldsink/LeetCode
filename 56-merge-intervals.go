package main

import (
	"sort"
)

type ByStart [][]int

func (a ByStart) Len() int           { return len(a) }
func (a ByStart) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a ByStart) Less(i, j int) bool { return a[i][0] < a[j][0] }

func merge(intervals [][]int) [][]int {
	var result [][]int
	sort.Sort(ByStart(intervals))
	if len(intervals) == 0 {
		return result
	}
	var current []int
	for i := 0; i < len(intervals); i++ {
		if len(current) == 0 {
			current = []int{intervals[i][0], intervals[i][1]}
			continue
		}
		if intervals[i][0] <= current[1] {
			if current[1] < intervals[i][1] {
				current[1] = intervals[i][1]
			}
		} else {
			result = append(result, current)
			current = []int{intervals[i][0], intervals[i][1]}
		}
	}
	return append(result, current)
}
