package main

import "sort"

func kthLargestValue(matrix [][]int, k int) int {
	var xor [][]int
	for i := 0; i < len(matrix); i++ {
		xor = append(xor, make([]int, len(matrix[0])))
	}
	xor[0][0] = matrix[0][0]
	for i := 1; i < len(matrix[0]); i++ {
		xor[0][i] = xor[0][i-1] ^ matrix[0][i]
	}
	for i := 1; i < len(matrix); i++ {
		xor[i][0] = xor[i-1][0] ^ matrix[i][0]
		for j := 1; j < len(matrix[0]); j++ {
			xor[i][j] = xor[i-1][j-1] ^ xor[i][j-1] ^ xor[i-1][j] ^ matrix[i][j]
		}
	}
	var result []int
	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[0]); j++ {
			result = append(result, xor[i][j])
		}
	}
	sort.Ints(result)
	return result[len(result)-k]
}
