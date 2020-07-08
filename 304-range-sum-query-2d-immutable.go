package main

type NumMatrix struct {
	sumMatrix [][]int
}

func Constructor(matrix [][]int) NumMatrix {
	var sumMatrix [][]int

	var m, n int
	if m = len(matrix); m == 0 {
		return NumMatrix{sumMatrix: sumMatrix}
	} else if n = len(matrix[0]); n == 0 {
		return NumMatrix{sumMatrix: sumMatrix}
	}

	sumMatrix = make([][]int, m+1)
	for i := 0; i < m+1; i++ {
		sumMatrix[i] = make([]int, n+1)
	}
	for i := m - 1; i >= 0; i-- {
		for j := n - 1; j >= 0; j-- {
			sumMatrix[i][j] = matrix[i][j] + sumMatrix[i][j+1] + sumMatrix[i+1][j] - sumMatrix[i+1][j+1]
		}
	}
	return NumMatrix{sumMatrix: sumMatrix}
}

func (this *NumMatrix) SumRegion(row1 int, col1 int, row2 int, col2 int) int {
	return this.sumMatrix[row1][col1] + this.sumMatrix[row2+1][col2+1] - this.sumMatrix[row2+1][col1] - this.sumMatrix[row1][col2+1]
}
