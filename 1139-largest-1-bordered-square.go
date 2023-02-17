package main

func largest1BorderedSquare(grid [][]int) int {
	var down, right [][]int
	for i := 0; i < len(grid); i++ {
		down = append(down, make([]int, len(grid[0])))
		right = append(right, make([]int, len(grid[0])))
	}
	for row := len(grid) - 1; row >= 0; row-- {
		for col := len(grid[0]) - 1; col >= 0; col-- {
			if grid[row][col] == 0 {
				continue
			}
			right[row][col] = 1
			down[row][col] = 1
			if col+1 < len(grid[0]) {
				right[row][col] += right[row][col+1]
			}
			if row+1 < len(grid) {
				down[row][col] += down[row+1][col]
			}
		}
	}
	max := func(a, b int) int {
		if b > a {
			return b
		}
		return a
	}
	var maxEdge int
	for row := 0; row < len(grid); row++ {
		for col := 0; col < len(grid[0]); col++ {
			for edge := max(len(grid)-row, len(grid[0])-col); edge > maxEdge; edge-- {
				if right[row][col] >= edge && down[row][col] >= edge && right[row+edge-1][col] >= edge && down[row][col+edge-1] >= edge {
					maxEdge = edge
					break
				}
			}
		}
	}
	return maxEdge * maxEdge
}
