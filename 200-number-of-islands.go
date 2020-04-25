package main

func numIslands(grid [][]byte) int {
	var m, n int
	if m = len(grid); m < 1 {
		return 0
	} else if n = len(grid[0]); n < 1 {
		return 0
	}

	type Point struct{ x, y int }
	var island int
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] != '1' {
				continue
			}
			island++
			grid[i][j] = '2'
			stack := []Point{{i, j}}
			for ; len(stack) > 0; stack = stack[1:] {
				if stack[0].y+1 < n && grid[stack[0].x][stack[0].y+1] == '1' {
					grid[stack[0].x][stack[0].y+1] = '2'
					stack = append(stack, Point{stack[0].x, stack[0].y + 1})
				}
				if stack[0].x+1 < m && grid[stack[0].x+1][stack[0].y] == '1' {
					grid[stack[0].x+1][stack[0].y] = '2'
					stack = append(stack, Point{stack[0].x + 1, stack[0].y})
				}
				if stack[0].x-1 >= 0 && grid[stack[0].x-1][stack[0].y] == '1' {
					grid[stack[0].x-1][stack[0].y] = '2'
					stack = append(stack, Point{stack[0].x - 1, stack[0].y})
				}
				if stack[0].y-1 >= 0 && grid[stack[0].x][stack[0].y-1] == '1' {
					grid[stack[0].x][stack[0].y-1] = '2'
					stack = append(stack, Point{stack[0].x, stack[0].y - 1})
				}
			}
		}
	}
	return island
}
