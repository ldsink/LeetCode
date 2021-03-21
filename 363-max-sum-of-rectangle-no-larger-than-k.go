package main

func maxSumSubmatrix(matrix [][]int, k int) int {
	m := len(matrix)
	n := len(matrix[0])
	top := k

	sum := make([][]int, m+1)
	for i := 0; i <= m; i++ {
		sum[i] = make([]int, n+1)
	}

	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {
			sum[i][j] = sum[i-1][j] + sum[i][j-1] - sum[i-1][j-1] + matrix[i-1][j-1]
		}
	}

	diff := 5000000
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			for k := i + 1; k <= m; k++ {
				for l := j + 1; l <= n; l++ {
					s := sum[k][l] - (sum[k][j] + sum[i][l] - sum[i][j])
					if s <= top && diff > top-s {
						diff = top - s
						if diff == 0 {
							return top
						}
					}
				}
			}
		}
	}

	return top - diff
}
