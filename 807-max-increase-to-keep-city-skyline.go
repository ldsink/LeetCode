package main

func maxIncreaseKeepingSkyline(grid [][]int) int {
	n := len(grid)
	vertical := make([]int, n)
	horizon := make([]int, n)
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if horizon[i] < grid[i][j] {
				horizon[i] = grid[i][j]
			}
			if vertical[j] < grid[i][j] {
				vertical[j] = grid[i][j]
			}
		}
	}
	var result int
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			max := horizon[i]
			if max > vertical[j] {
				max = vertical[j]
			}
			if max > grid[i][j] {
				result += max - grid[i][j]
			}
		}
	}
	return result
}
