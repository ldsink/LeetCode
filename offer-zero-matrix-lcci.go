package main

func setZeroes(matrix [][]int) {
	m, n := len(matrix), len(matrix[0])
	row, col := make([]bool, m), make([]bool, n)
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if matrix[i][j] == 0 {
				row[i], col[j] = true, true
			}
		}
	}
	for i := 0; i < m; i++ {
		if row[i] {
			for j := 0; j < n; j++ {
				matrix[i][j] = 0
			}
		}
	}
	for i := 0; i < n; i++ {
		if col[i] {
			for j := 0; j < m; j++ {
				matrix[j][i] = 0
			}
		}
	}
}
