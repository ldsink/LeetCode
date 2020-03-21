package main

func uniquePathsWithObstacles(obstacleGrid [][]int) int {
	var m, n int
	if n = len(obstacleGrid); n == 0 {
		return 0
	}
	if m = len(obstacleGrid[0]); m == 0 {
		return 0
	}

	counts := [2][]int{make([]int, m), make([]int, m)}

	current := (n - 1) % 2
	if obstacleGrid[n-1][m-1] == 1 {
		counts[current][m-1] = 0
	} else {
		counts[current][m-1] = 1
	}
	for i := m - 2; i >= 0; i-- {
		if obstacleGrid[n-1][i] == 1 {
			counts[current][i] = 0
		} else {
			counts[current][i] = counts[current][i+1]
		}
	}

	for i := n - 2; i >= 0; i-- {
		current := i % 2
		previous := (i + 1) % 2
		if obstacleGrid[i][m-1] == 1 {
			counts[current][m-1] = 0
		} else {
			counts[current][m-1] = counts[previous][m-1]
		}
		for j := m - 2; j >= 0; j-- {
			if obstacleGrid[i][j] == 1 {
				counts[current][j] = 0
			} else {
				counts[current][j] = counts[current][j+1] + counts[previous][j]
			}
		}
	}
	return counts[0][0]
}
