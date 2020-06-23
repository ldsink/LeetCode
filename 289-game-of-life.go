package main

func gameOfLife(board [][]int) {
	var m, n int
	if m = len(board); m == 0 {
		return
	} else if n = len(board[0]); n == 0 {
		return
	}

	count := make([][]int, m)
	for i := 0; i < m; i++ {
		count[i] = make([]int, n)
	}
	offset := [8][2]int{{-1, -1}, {-1, 0}, {-1, 1}, {0, -1}, {0, 1}, {1, -1}, {1, 0}, {1, 1}}
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			for k := 0; k < 8; k++ {
				if 0 <= i+offset[k][0] && i+offset[k][0] < m && 0 <= j+offset[k][1] && j+offset[k][1] < n {
					if board[i+offset[k][0]][j+offset[k][1]] == 1 {
						count[i][j]++
					}
				}
			}
		}
	}
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if count[i][j] < 2 || count[i][j] > 3 {
				board[i][j] = 0
			} else if count[i][j] == 3 {
				board[i][j] = 1
			}
		}
	}
}
