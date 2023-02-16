package main

func spiralOrder(matrix [][]int) []int {
	var result []int
	m := len(matrix)
	if m == 0 {
		return result
	}
	n := len(matrix[0])
	if n == 0 {
		return result
	}

	l, r, t, b := 0, n-1, 0, m-1
	for l <= r && t <= b {
		for i := l; i <= r; i++ {
			result = append(result, matrix[t][i])
		}
		t++
		if t > b {
			break
		}

		for i := t; i <= b; i++ {
			result = append(result, matrix[i][r])
		}
		r--
		if r < l {
			break
		}

		for i := r; i >= l; i-- {
			result = append(result, matrix[b][i])
		}
		b--
		if t > b {
			break
		}

		for i := b; i >= t; i-- {
			result = append(result, matrix[i][l])
		}
		l++
	}
	return result
}
