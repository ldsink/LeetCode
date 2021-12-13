package main

import (
	"sort"
)

func scheduleCourse(courses [][]int) int {
	if len(courses) == 0 {
		return 0
	}
	// 按照结束时间逆序排列，如果结束时间相同，把课时短的放前面
	sort.Slice(courses, func(i, j int) bool {
		return courses[i][1] > courses[j][1] || (courses[i][1] == courses[j][1] && courses[i][0] < courses[j][0])
	})
	// 动态规划，result[i] 表示第 i 天起可以上的最多的课程
	result := make([]int, courses[0][1]+2)
	for _, course := range courses {
		for i, j := 1, course[1]-course[0]+1; i <= j; i++ {
			if result[i] < 1+result[i+course[0]] {
				result[i] = 1 + result[i+course[0]]
			}
		}
	}
	return result[1]
}
