package main

func countBattleships(board [][]byte) int {
	m := len(board)
	n := len(board[0])
	ship := byte('X')
	var result int
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if board[i][j] != ship {
				continue
			}
			result++
			board[i][j] = '.'
			if j+1 < n && board[i][j+1] == ship {
				for k := j + 1; k < n && board[i][k] == ship; k++ {
					board[i][k] = '.'
				}
			} else if i+1 < m && board[i+1][j] == ship {
				for k := i + 1; k < m && board[k][j] == ship; k++ {
					board[k][j] = '.'
				}
			}
		}
	}
	return result
}
