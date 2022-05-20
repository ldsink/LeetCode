package main

import (
	"sort"
)

func findStart(target int, starts []int) (index int) {
	index = -1
	left, right := 0, len(starts)-1
	for left < right {
		middle := (left + right) >> 1
		if target > starts[middle] {
			left = middle + 1
		} else {
			right = middle
		}
	}
	if starts[left] >= target {
		index = left
	}
	return
}

func findRightInterval(intervals [][]int) []int {
	startIndexMap := make(map[int]int)
	for index, interval := range intervals {
		startIndexMap[interval[0]] = index
	}
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})

	var starts []int
	for start := range startIndexMap {
		starts = append(starts, start)
	}
	sort.Ints(starts)

	result := make([]int, len(intervals))
	for _, interval := range intervals {
		start, end := interval[0], interval[1]
		index := findStart(end, starts)
		if index == -1 {
			result[startIndexMap[start]] = -1
		} else {
			result[startIndexMap[start]] = startIndexMap[starts[index]]
		}
	}
	return result
}
