package main

func checkMove(board [][]byte, rMove int, cMove int, color byte) bool {
	if board[rMove][cMove] != '.' {
		return false
	}
	var middle byte // 中间的颜色
	if color == 'W' {
		middle = 'B'
	} else {
		middle = 'W'
	}
	directions := [][2]int{{-1, -1}, {-1, 0}, {-1, 1}, {0, -1}, {0, 1}, {1, -1}, {1, 0}, {1, 1}}
	for _, direction := range directions {
		var i int
		for i = 1; 0 <= rMove+direction[0]*i && rMove+direction[0]*i < len(board) && 0 <= cMove+direction[1]*i && cMove+direction[1]*i < len(board[0]) && board[rMove+direction[0]*i][cMove+direction[1]*i] == middle; i++ {
		}
		if i > 1 && 0 <= rMove+direction[0]*i && rMove+direction[0]*i < len(board) && 0 <= cMove+direction[1]*i && cMove+direction[1]*i < len(board[0]) && board[rMove+direction[0]*i][cMove+direction[1]*i] == color {
			return true
		}
	}
	return false
}
