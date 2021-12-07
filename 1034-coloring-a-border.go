package main

func colorBorder(grid [][]int, row int, col int, color int) [][]int {
	if len(grid) == 0 && len(grid[0]) == 0 {
		return grid
	}

	directions := [4][2]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}

	// 找出全部的连通分量
	visited := make(map[[2]int]bool)
	var queue [][2]int
	queue = append(queue, [2]int{row, col})
	visited[[2]int{row, col}] = true
	for idx := 0; idx < len(queue); idx++ {
		x, y := queue[idx][0], queue[idx][1]
		for _, d := range directions {
			i, j := x+d[0], y+d[1]
			if i < 0 || i == len(grid) || j < 0 || j == len(grid[0]) {
				continue
			}
			if visited[[2]int{i, j}] {
				continue
			}
			if grid[i][j] == grid[x][y] {
				visited[[2]int{i, j}] = true
				queue = append(queue, [2]int{i, j})
			}
		}
	}

	// 对连通分量的边界进行染色
	for idx := 0; idx < len(queue); idx++ {
		x, y := queue[idx][0], queue[idx][1]
		// 在网格的边界上（第一行/列或最后一行/列）
		if x == 0 || x == len(grid)-1 || y == 0 || y == len(grid[0])-1 {
			grid[x][y] = color
			continue
		}
		// 在上、下、左、右四个方向上与不属于同一连通分量的网格块相邻
		all := true
		for _, d := range directions {
			i, j := x+d[0], y+d[1]
			if !visited[[2]int{i, j}] {
				all = false
				break
			}
		}
		if !all {
			grid[x][y] = color
		}
	}

	return grid
}
