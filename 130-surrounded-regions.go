package main

func solve(board [][]byte) {
	var m, n int
	if m = len(board); m <= 1 {
		return
	}
	if n = len(board[0]); n <= 1 {
		return
	}

	type point struct{ x, y int }
	var stack []point
	var onBoard, onStack [][]bool
	var O byte = 'O'
	for i := 0; i < m; i++ {
		onBoard = append(onBoard, make([]bool, n))
		onStack = append(onStack, make([]bool, n))
	}
	for i := 0; i < n; i++ {
		if board[0][i] == O {
			onBoard[0][i] = true
			onStack[0][i] = true
			stack = append(stack, point{x: 0, y: i})
		}
		if board[m-1][i] == O {
			onBoard[m-1][i] = true
			onStack[m-1][i] = true
			stack = append(stack, point{x: m - 1, y: i})
		}
	}
	for i := 1; i < m-1; i++ {
		if board[i][0] == O {
			onBoard[i][0] = true
			onStack[i][0] = true
			stack = append(stack, point{x: i, y: 0})
		}
		if board[i][n-1] == O {
			onBoard[i][n-1] = true
			onStack[i][n-1] = true
			stack = append(stack, point{x: i, y: n - 1})
		}
	}
	for i := 0; i < len(stack); i++ {
		if x, y := stack[i].x-1, stack[i].y; 0 <= x && board[x][y] == O && !onStack[x][y] {
			onBoard[x][y] = true
			onStack[x][y] = true
			stack = append(stack, point{x: x, y: y})
		}
		if x, y := stack[i].x, stack[i].y+1; y < n && board[x][y] == O && !onStack[x][y] {
			onBoard[x][y] = true
			onStack[x][y] = true
			stack = append(stack, point{x: x, y: y})
		}
		if x, y := stack[i].x+1, stack[i].y; x < m && board[x][y] == O && !onStack[x][y] {
			onBoard[x][y] = true
			onStack[x][y] = true
			stack = append(stack, point{x: x, y: y})
		}
		if x, y := stack[i].x, stack[i].y-1; 0 <= y && board[x][y] == O && !onStack[x][y] {
			onBoard[x][y] = true
			onStack[x][y] = true
			stack = append(stack, point{x: x, y: y})
		}
	}

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if board[i][j] == O && !onBoard[i][j] {
				board[i][j] = 'X'
			}
		}
	}
}
