package main

func luckyNumbers(matrix [][]int) []int {
	m := len(matrix)
	n := len(matrix[0])

	// 同一列的所有元素中最大
	column := make([]int, n)
	for i := 0; i < n; i++ {
		column[i] = matrix[0][i]
	}
	for i := 1; i < m; i++ {
		for j := 0; j < n; j++ {
			if column[j] < matrix[i][j] {
				column[j] = matrix[i][j]
			}
		}
	}

	var result []int
	for i := 0; i < m; i++ {
		// 在同一行的所有元素中最小
		num, idx := matrix[i][0], 0
		for j := 1; j < n; j++ {
			if num > matrix[i][j] {
				num, idx = matrix[i][j], j
			}
		}
		if column[idx] == num {
			result = append(result, num)
		}
	}
	return result
}
