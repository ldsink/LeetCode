package main

import "sort"

func largestIsland(grid [][]int) int {
	n := len(grid)
	area := []int{0, 0} // 岛的面积，从 2 开始编号
	directions := [4][2]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}
	// 染个色
	var colorIsland func(int, int, int)
	colorIsland = func(x, y, color int) {
		grid[x][y] = color
		area[color]++
		for _, d := range directions {
			i, j := x+d[0], y+d[1]
			if 0 <= i && i < n && 0 <= j && j < n && grid[i][j] == 1 {
				colorIsland(i, j, color)
			}
		}
	}
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] == 1 {
				color := len(area)
				area = append(area, 0)
				colorIsland(i, j, color)
			}
		}
	}
	// 先用岛屿做初始化的结果
	var result int
	for i := 2; i < len(area); i++ {
		if result < area[i] {
			result = area[i]
		}
	}
	// 是否可以通过连接创造更大的岛
	getMax := func(x, y int) int {
		var islands []int
		for _, d := range directions {
			i, j := x+d[0], y+d[1]
			if 0 <= i && i < n && 0 <= j && j < n && grid[i][j] != 0 {
				added := false
				for _, island := range islands {
					if island == grid[i][j] {
						added = true
						break
					}
				}
				if !added {
					islands = append(islands, grid[i][j])
				}
			}
		}
		sort.Slice(islands, func(i, j int) bool {
			return area[islands[i]] > area[islands[j]]
		})
		result := 1
		for i := 0; i < len(islands); i++ {
			result += area[islands[i]]
		}
		return result
	}
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] == 0 {
				if a := getMax(i, j); result < a {
					result = a
				}
			}
		}
	}
	return result
}
