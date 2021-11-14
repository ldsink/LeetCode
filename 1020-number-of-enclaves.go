package main

func numEnclaves(grid [][]int) int {
	var queue [][2]int
	m := len(grid)
	n := len(grid[0])
	if m <= 2 || n <= 2 {
		return 0
	}
	for i := 0; i < n; i++ {
		if grid[0][i] == 1 {
			queue = append(queue, [2]int{0, i})
			grid[0][i] = 2
		}
		if grid[m-1][i] == 1 {
			queue = append(queue, [2]int{m - 1, i})
			grid[m-1][i] = 2
		}
	}
	for i := 1; i < m-1; i++ {
		if grid[i][0] == 1 {
			queue = append(queue, [2]int{i, 0})
			grid[i][0] = 2
		}
		if grid[i][n-1] == 1 {
			queue = append(queue, [2]int{i, n - 1})
			grid[i][n-1] = 2
		}
	}
	for len(queue) > 0 {
		x, y := queue[0][0], queue[0][1]
		queue = queue[1:]
		// top
		if x-1 >= 0 && grid[x-1][y] == 1 {
			grid[x-1][y] = 2
			queue = append(queue, [2]int{x - 1, y})
		}
		// right
		if y+1 < n && grid[x][y+1] == 1 {
			grid[x][y+1] = 2
			queue = append(queue, [2]int{x, y + 1})
		}
		// bottom
		if x+1 < m && grid[x+1][y] == 1 {
			grid[x+1][y] = 2
			queue = append(queue, [2]int{x + 1, y})
		}
		// left
		if y-1 >= 0 && grid[x][y-1] == 1 {
			grid[x][y-1] = 2
			queue = append(queue, [2]int{x, y - 1})
		}
	}

	var result int
	for i := 1; i < m-1; i++ {
		for j := 1; j < n-1; j++ {
			if grid[i][j] == 1 {
				result++
			}
		}
	}
	return result
}
