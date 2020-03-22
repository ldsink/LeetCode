package main

func searchMatrix(matrix [][]int, target int) bool {
	var m, n int
	if m = len(matrix); m == 0 {
		return false
	}
	if n = len(matrix[0]); n == 0 {
		return false
	}
	for i, j := 0, m*n-1; i <= j; {
		if i == j {
			return matrix[i/n][i%n] == target
		}
		middle := (i + j) >> 1
		if target <= matrix[middle/n][middle%n] {
			j = middle
		} else {
			i = middle + 1
		}
	}
	return false
}
