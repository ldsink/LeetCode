package main

func rotate(matrix [][]int) {
	n := len(matrix)
	var originMatrix [][]int
	for i := 0; i < n; i++ {
		row := make([]int, n)
		for j := 0; j < n; j++ {
			row[j] = matrix[i][j]
		}
		originMatrix = append(originMatrix, row)
	}
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			matrix[j][n-i-1] = originMatrix[i][j]
		}
	}
}
