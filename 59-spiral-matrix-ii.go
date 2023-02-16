package main

func generateMatrix(n int) [][]int {
	var matrix [][]int
	for i := 0; i < n; i++ {
		row := make([]int, n)
		matrix = append(matrix, row)
	}
	l, r, t, b := 0, n-1, 0, n-1
	count := 1
	for l <= r && t <= b {
		for i := l; i <= r; i++ {
			matrix[t][i] = count
			count++
		}
		t++
		if t > b {
			break
		}

		for i := t; i <= b; i++ {
			matrix[i][r] = count
			count++
		}
		r--
		if r < l {
			break
		}

		for i := r; i >= l; i-- {
			matrix[b][i] = count
			count++
		}
		b--
		if t > b {
			break
		}

		for i := b; i >= t; i-- {
			matrix[i][l] = count
			count++
		}
		l++
	}
	return matrix
}
