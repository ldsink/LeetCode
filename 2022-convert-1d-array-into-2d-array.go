package main

func construct2DArray(original []int, m int, n int) [][]int {
	var result [][]int
	if len(original) != m*n {
		return result
	}
	for i := 0; i < m; i++ {
		result = append(result, original[i*n:(i+1)*n])
	}
	return result
}
