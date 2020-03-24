package main

func maximalRectangle(matrix [][]byte) int {
	var area, m, n int
	if m = len(matrix); m == 0 {
		return 0
	}
	if n = len(matrix[0]); n == 0 {
		return 0
	}

	var vertical, horizon [][]int
	for i := 0; i < m; i++ {
		v := make([]int, n)
		vertical = append(vertical, v)
		h := make([]int, n)
		horizon = append(horizon, h)
	}

	for i := 0; i < m; i++ {
		for j := n - 1; j >= 0; j-- {
			if matrix[i][j] == '0' {
				continue
			}
			if j+1 < n && matrix[i][j+1] == '1' {
				horizon[i][j] = horizon[i][j+1] + 1
				continue
			}
		}
	}
	for j := 0; j < n; j++ {
		for i := m - 1; i >= 0; i-- {
			if matrix[i][j] == '0' {
				continue
			}
			if i+1 < m && matrix[i+1][j] == '1' {
				vertical[i][j] = vertical[i+1][j] + 1
				continue
			}
		}
	}

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if matrix[i][j] == '0' {
				continue
			}
			if (horizon[i][j]+1)*(vertical[i][j]+1) < area {
				continue
			}
			for l, w, v := i, horizon[i][j], i+vertical[i][j]; l <= v; l++ {
				if w > horizon[l][j] {
					w = horizon[l][j]
				}
				if a := (w + 1) * (l - i + 1); a > area {
					area = a
				}
			}
		}
	}
	return area
}
