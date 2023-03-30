package main

import "sort"

func maxWidthOfVerticalArea(points [][]int) int {
	sort.Slice(points, func(i, j int) bool {
		return points[i][0] < points[j][0]
	})
	var result int
	for i := 0; i < len(points)-1; i++ {
		if r := points[i+1][0] - points[i][0]; result < r {
			result = r
		}
	}
	return result
}
