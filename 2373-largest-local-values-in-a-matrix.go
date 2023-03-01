package main

func largestLocal(grid [][]int) [][]int {
	n := len(grid)
	var result [][]int
	for i := 1; i < n-1; i++ {
		var row []int
		for j := 1; j < n-1; j++ {
			v := grid[i][j]
			for x := -1; x <= 1; x++ {
				for y := -1; y <= 1; y++ {
					if v < grid[i+x][j+y] {
						v = grid[i+x][j+y]
					}
				}
			}
			row = append(row, v)
		}
		result = append(result, row)
	}
	return result
}
