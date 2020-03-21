package main

func placeQueen(q, n int, v []bool, t2b, b2t map[int]bool, boardPtr *[][]rune) int {
	board := *boardPtr
	if q == n {
		return 1
	}

	var result int
	for i := 0; i < n; i++ {
		// exclude unused position
		if v[i] {
			continue
		}
		if t2b[q-i] {
			continue
		}
		if b2t[q+i] {
			continue
		}

		board[q][i] = 'Q'
		v[i] = true
		t2b[q-i] = true
		b2t[q+i] = true

		r := placeQueen(q+1, n, v, t2b, b2t, boardPtr)
		result += r

		board[q][i] = '.'
		v[i] = false
		t2b[q-i] = false
		b2t[q+i] = false
	}

	return result
}

func totalNQueens(n int) int {
	vertical := make([]bool, n)
	leftTopToRightBottom := make(map[int]bool)
	leftBottomToRightTop := make(map[int]bool)

	var board [][]rune
	for i := 0; i < n; i++ {
		row := make([]rune, n)
		for j := 0; j < n; j++ {
			row[j] = '.'
		}
		board = append(board, row)
	}

	return placeQueen(0, n, vertical, leftTopToRightBottom, leftBottomToRightTop, &board)
}
