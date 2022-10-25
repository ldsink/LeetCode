package main

func shortestBridge(grid [][]int) int {
	getInit := func() [2]int {
		for i := 0; i < len(grid); i++ {
			for j := 0; j < len(grid[0]); j++ {
				if grid[i][j] == 1 {
					return [2]int{i, j}
				}
			}
		}
		return [2]int{}
	}
	start := getInit()
	directions := [4][2]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}
	islandQueue := [][2]int{start}
	var waterQueue [][3]int
	visited := make(map[[2]int]bool)
	visited[start] = true
	// 基于岛初始化沿海的水域
	for idx := 0; idx < len(islandQueue); idx++ {
		x, y := islandQueue[idx][0], islandQueue[idx][1]
		for _, d := range directions {
			i, j := x+d[0], y+d[1]
			if i < 0 || i == len(grid) || j < 0 || j == len(grid[0]) || visited[[2]int{i, j}] {
				continue
			}
			visited[[2]int{i, j}] = true
			if grid[i][j] == 1 {
				islandQueue = append(islandQueue, [2]int{i, j})
			} else {
				waterQueue = append(waterQueue, [3]int{i, j, 1})
			}
		}
	}
	// 基于水域广搜获取最短的岛
	for idx := 0; idx < len(waterQueue); idx++ {
		x, y := waterQueue[idx][0], waterQueue[idx][1]
		for _, d := range directions {
			i, j := x+d[0], y+d[1]
			if i < 0 || i == len(grid) || j < 0 || j == len(grid[0]) || visited[[2]int{i, j}] {
				continue
			}
			if grid[i][j] == 1 {
				return waterQueue[idx][2]
			}
			visited[[2]int{i, j}] = true
			waterQueue = append(waterQueue, [3]int{i, j, waterQueue[idx][2] + 1})
		}
	}
	return -1
}
