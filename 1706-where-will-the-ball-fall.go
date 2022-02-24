package main

func rollBall(index, level int, grid [][]int) int {
	if level == len(grid) {
		return index
	}
	if grid[level][index] == 1 && index+1 < len(grid[level]) && grid[level][index+1] == 1 {
		return rollBall(index+1, level+1, grid)
	} else if grid[level][index] == -1 && index-1 >= 0 && grid[level][index-1] == -1 {
		return rollBall(index-1, level+1, grid)
	}
	return -1
}

func findBall(grid [][]int) []int {
	var result []int
	for i := 0; i < len(grid[0]); i++ {
		result = append(result, rollBall(i, 0, grid))
	}
	return result
}
