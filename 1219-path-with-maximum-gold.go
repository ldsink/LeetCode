package main

func collectGold(x, y int, grid [][]int, collected map[[2]int]bool) int {
	collected[[2]int{x, y}] = true
	var gold int

	for _, d := range [4][2]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}} {
		i, j := x+d[0], y+d[1]
		if i < 0 || i == len(grid) || j < 0 || j == len(grid[0]) {
			continue
		}
		if grid[i][j] == 0 || collected[[2]int{i, j}] {
			continue
		}
		if temp := collectGold(i, j, grid, collected); gold < temp {
			gold = temp
		}
	}
	delete(collected, [2]int{x, y})
	return gold + grid[x][y]
}

func getMaximumGold(grid [][]int) int {
	// 最多 25 个单元格中有黄金。
	collected := make(map[[2]int]bool)
	m := len(grid)
	n := len(grid[1])
	var result int
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] > 0 {
				if gold := collectGold(i, j, grid, collected); result < gold {
					result = gold
				}
			}
		}
	}
	return result
}
