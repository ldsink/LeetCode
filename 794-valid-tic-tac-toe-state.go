package main

func staticResult(square uint8, xSuccess, oSuccess bool) (bool, bool) {
	if square == 'X' {
		xSuccess = true
	} else if square == 'O' {
		oSuccess = true
	}
	// 否则 square 上没有棋子，什么都不改变
	return xSuccess, oSuccess
}

func validTicTacToe(board []string) bool {
	// 棋子数目不对的情况
	var xCount, oCount int
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			if board[i][j] == 'X' {
				xCount++
			} else if board[i][j] == 'O' {
				oCount++
			}
		}
	}
	if !(xCount == oCount || xCount == oCount+1) {
		return false
	}
	// 胜利条件的情况
	var xSuccess, oSuccess bool
	for i := 0; i < 3; i++ {
		if board[i][0] == board[i][1] && board[i][1] == board[i][2] {
			xSuccess, oSuccess = staticResult(board[i][0], xSuccess, oSuccess)
		}
		if board[0][i] == board[1][i] && board[1][i] == board[2][i] {
			xSuccess, oSuccess = staticResult(board[0][i], xSuccess, oSuccess)
		}
	}
	if board[0][0] == board[1][1] && board[1][1] == board[2][2] {
		xSuccess, oSuccess = staticResult(board[1][1], xSuccess, oSuccess)
	}
	if board[0][2] == board[1][1] && board[1][1] == board[2][0] {
		xSuccess, oSuccess = staticResult(board[1][1], xSuccess, oSuccess)
	}
	// 如果后手赢了，考虑各种异常情况。此时先手肯定不能赢，并且棋子数目不能超过后者。
	if oSuccess {
		if xSuccess {
			return false
		} else if xCount != oCount {
			return false
		}
	}
	// 再考虑先手赢的情况，棋子数一定大于后者
	if xSuccess && xCount == oCount {
		return false
	}
	return true
}
