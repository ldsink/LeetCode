package main

func searchMatrix(matrix [][]int, target int) bool {
	m := len(matrix)
	if m == 0 {
		return false
	}
	n := len(matrix[0])
	if n == 0 {
		return false
	}
	for i := 0; i < m; i++ {
		for ; n > 0 && target < matrix[i][n-1]; n-- {
		}
		for j := n - 1; j >= 0; j-- {
			if matrix[i][j] == target {
				return true
			} else if matrix[i][j] < target {
				break
			}
		}
	}
	return false
}
