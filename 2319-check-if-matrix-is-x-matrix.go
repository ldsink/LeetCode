package main

func checkXMatrix(grid [][]int) bool {
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid); j++ {
			if i == j || i+j+1 == len(grid) {
				if grid[i][j] == 0 {
					return false
				}
			} else {
				if grid[i][j] != 0 {
					return false
				}
			}
		}
	}
	return true
}
