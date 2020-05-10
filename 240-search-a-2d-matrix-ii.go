package main

func searchMatrix(matrix [][]int, target int) bool {
	var m, n, x, y int
	if m = len(matrix); m == 0 {
		return false
	} else if n = len(matrix[0]); n == 0 {
		return false
	}

	m, n = m-1, n-1
	for ; x <= m && y <= n; {
		if matrix[x][n] < target {
			x++
			continue
		}
		if matrix[m][y] < target {
			y++
			continue
		}
		if target < matrix[x][n] {
			n--
			continue
		}
		if target < matrix[m][y] {
			m--
			continue
		}
		return matrix[x][y] == target || matrix[m][y] == target || matrix[x][m] == target || matrix[m][n] == target
	}
	return false
}
