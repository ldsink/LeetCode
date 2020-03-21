package main

func placeQueen(q, n int, v []bool, t2b, b2t map[int]bool, boardPtr *[][]rune) [][]string {
	board := *boardPtr
	if q == n {
		var snap []string
		for _, row := range board {
			snap = append(snap, string(row))
		}
		return [][]string{snap}
	}

	var result [][]string
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
		if len(r) > 0 {
			result = append(result, r...)
		}

		board[q][i] = '.'
		v[i] = false
		t2b[q-i] = false
		b2t[q+i] = false
	}

	return result
}

func solveNQueens(n int) [][]string {
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
