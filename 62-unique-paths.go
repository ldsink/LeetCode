package main

func uniquePaths(m int, n int) int {
	counts := [2][]int{
		make([]int, m),
		make([]int, m),
	}
	for i, j := 0, (n-1)%2; i < m; i++ {
		counts[j][i] = 1
	}
	for i := n - 2; i >= 0; i-- {
		current := i % 2
		previous := (i + 1) % 2
		counts[current][m-1] = 1
		for j := m - 2; j >= 0; j-- {
			counts[current][j] = counts[current][j+1] + counts[previous][j]
		}
	}
	return counts[0][0]
}
