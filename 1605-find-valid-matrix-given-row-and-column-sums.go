package main

func restoreMatrix(rowSum []int, colSum []int) [][]int {
	min := func(a, b int) int {
		if a > b {
			return b
		}
		return a
	}
	result := make([][]int, len(rowSum))
	for i := 0; i < len(rowSum); i++ {
		result[i] = make([]int, len(colSum))
	}
	for i := 0; i < len(rowSum); i++ {
		for j := 0; j < len(colSum) && rowSum[i] != 0; j++ {
			val := min(rowSum[i], colSum[j])
			result[i][j] = val
			rowSum[i] -= val
			colSum[j] -= val
		}
	}
	return result
}
