package main

import (
	"sort"
)

func scheduleCourse(courses [][]int) int {
	if len(courses) == 0 {
		return 0
	}
	// 按照结束时间顺序排序
	sort.Slice(courses, func(i, j int) bool {
		return courses[i][1] < courses[j][1]
	})
	// 动态规划，result[i] 表示第 i 天起可以上的最多的课程。课程从后往前递归处理。
	result := make([]int, courses[len(courses)-1][1]+2)
	for idx := len(courses) - 1; idx >= 0; idx-- {
		duration, lastDay := courses[idx][0], courses[idx][1]
		for i, j := 1, lastDay-duration+1; i <= j; i++ {
			if result[i] < 1+result[i+duration] {
				result[i] = 1 + result[i+duration]
			}
		}
	}
	return result[1]
}
