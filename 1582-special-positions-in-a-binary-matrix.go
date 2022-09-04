package main

func numSpecial(mat [][]int) int {
	m, n := len(mat), len(mat[0])
	rowOnes, colOnes := make([]int, m), make([]int, n)
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if mat[i][j] == 1 {
				rowOnes[i]++
				colOnes[j]++
			}
		}
	}
	var result int
	for i := 0; i < m; i++ {
		if rowOnes[i] != 1 {
			continue
		}
		for j := 0; j < n; j++ {
			if mat[i][j] == 1 && colOnes[j] == 1 {
				result++
			}
		}
	}
	return result
}
