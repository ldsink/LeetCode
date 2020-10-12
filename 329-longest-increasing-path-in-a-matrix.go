package main

import (
	"sort"
)

func longestIncreasingPath(matrix [][]int) int {
	var m, n int
	if m = len(matrix); m == 0 {
		return 0
	} else if n = len(matrix[0]); n == 0 {
		return 0
	}

	var points [][3]int
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			points = append(points, [3]int{matrix[i][j], i, j})
		}
	}
	sort.Slice(points, func(i, j int) bool {
		return points[i][0] < points[j][0]
	})

	var maxCount [][]int
	for i := 0; i < m; i++ {
		maxCount = append(maxCount, make([]int, n))
	}

	directions := [4][2]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}
	for _, point := range points {
		for _, direction := range directions {
			if 0 <= point[1]+direction[0] && point[1]+direction[0] < m && 0 <= point[2]+direction[1] && point[2]+direction[1] < n && matrix[point[1]+direction[0]][point[2]+direction[1]] > point[0] && maxCount[point[1]+direction[0]][point[2]+direction[1]] < maxCount[point[1]][point[2]]+1 {
				maxCount[point[1]+direction[0]][point[2]+direction[1]] = maxCount[point[1]][point[2]] + 1
			}
		}
	}
	var max int
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if max < maxCount[i][j] {
				max = maxCount[i][j]
			}
		}
	}
	return max + 1
}
