package main

func setZeroes(matrix [][]int) {
	var m, n int
	if m = len(matrix); m == 0 {
		return
	}
	if n = len(matrix[0]); n == 0 {
		return
	}
	x := make([]bool, m)
	y := make([]bool, n)
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if matrix[i][j] == 0 {
				x[i], y[j] = true, true
			}
		}
	}
	for i := 0; i < m; i++ {
		if x[i] {
			for j := 0; j < n; j++ {
				matrix[i][j] = 0
			}
		}
	}
	for j := 0; j < n; j++ {
		if y[j] {
			for i := 0; i < m; i++ {
				matrix[i][j] = 0
			}
		}
	}
}
