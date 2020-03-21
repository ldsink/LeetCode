package main

func minPathSum(grid [][]int) int {
	var m, n int
	if n = len(grid); n == 0 {
		return 0
	}
	if m = len(grid[0]); m == 0 {
		return 0
	}
	minSum := [2][]int{make([]int, m), make([]int, m)}

	current := (n - 1) % 2
	minSum[current][m-1] = grid[n-1][m-1]
	for i := m - 2; i >= 0; i-- {
		minSum[current][i] = minSum[current][i+1] + grid[n-1][i]
	}
	for i := n - 2; i >= 0; i-- {
		current := i % 2
		previous := (i + 1) % 2
		minSum[current][m-1] = minSum[previous][m-1] + grid[i][m-1]
		for j := m - 2; j >= 0; j-- {
			var value int
			if minSum[current][j+1] > minSum[previous][j] {
				value = minSum[previous][j]
			} else {
				value = minSum[current][j+1]
			}
			minSum[current][j] = grid[i][j] + value
		}
	}
	return minSum[0][0]
}
