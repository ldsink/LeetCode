package main

func closedIsland(grid [][]int) int {
	directions := [4][2]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}
	color := func(x, y, c int) {
		s := grid[x][y]
		grid[x][y] = c
		for queue := [][2]int{{x, y}}; len(queue) > 0; queue = queue[1:] {
			x, y := queue[0][0], queue[0][1]
			for _, direction := range directions {
				a, b := direction[0], direction[1]
				if 0 <= x+a && x+a < len(grid) && 0 <= y+b && y+b < len(grid[0]) && grid[x+a][y+b] == s {
					grid[x+a][y+b] = c
					queue = append(queue, [2]int{x + a, y + b})
				}
			}
		}
	}
	// 首先把靠边的都标记成水
	for i := 0; i < len(grid[0]); i++ {
		if grid[0][i] == 0 {
			color(0, i, 1)
		}
		if grid[len(grid)-1][i] == 0 {
			color(len(grid)-1, i, 1)
		}
	}
	for i := 0; i < len(grid); i++ {
		if grid[i][0] == 0 {
			color(i, 0, 1)
		}
		if grid[i][len(grid[0])-1] == 0 {
			color(i, len(grid[0])-1, 1)
		}
	}
	// 找到土地，开始染色
	var count int
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[0]); j++ {
			if grid[i][j] == 0 {
				count++
				color(i, j, -count)
			}
		}
	}
	return count
}
