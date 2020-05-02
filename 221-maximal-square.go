package main

func maximalSquare(matrix [][]byte) int {
	var m, n int
	if m = len(matrix); m == 0 {
		return 0
	} else if n = len(matrix[0]); n == 0 {
		return 0
	}

	var down, line [2][]int
	for i := 0; i < 2; i++ {
		down[i] = make([]int, n)
		line[i] = make([]int, n)
	}

	var result, curr, prev, right, prevRight int
	for i := m - 1; i >= 0; i-- {
		curr = i % 2
		prev = curr ^ 0x1
		for j := n - 1; j >= 0; j-- {
			if matrix[i][j] != '1' {
				down[curr][j] = 0
				line[curr][j] = 0
				continue
			}
			right = 1
			if j+1 < n && matrix[i][j+1] == '1' {
				right += prevRight
			}
			prevRight = right
			down[curr][j] = 1
			if i+1 < m {
				down[curr][j] += down[prev][j]
			}
			line[curr][j] = 1
			if i+1 < m && j+1 < n {
				line[curr][j] += line[prev][j+1]
			}
			if line[curr][j] > right {
				line[curr][j] = right
			}
			if line[curr][j] > down[curr][j] {
				line[curr][j] = down[curr][j]
			}
			if result < line[curr][j] {
				result = line[curr][j]
			}
		}
	}
	return result * result
}
