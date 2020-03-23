package main

func searchWord(i, j, k, m, n int, board [][]byte, used [][]bool, word []byte) bool {
	if k == len(word)-1 {
		return board[i][j] == word[k]
	}
	if board[i][j] != word[k] {
		return false
	}
	used[i][j] = true
	if i-1 >= 0 && !used[i-1][j] {
		if ok := searchWord(i-1, j, k+1, m, n, board, used, word); ok {
			return true
		}
	}
	if j+1 < n && !used[i][j+1] {
		if ok := searchWord(i, j+1, k+1, m, n, board, used, word); ok {
			return true
		}
	}
	if i+1 < m && !used[i+1][j] {
		if ok := searchWord(i+1, j, k+1, m, n, board, used, word); ok {
			return true
		}
	}
	if j-1 >= 0 && !used[i][j-1] {
		if ok := searchWord(i, j-1, k+1, m, n, board, used, word); ok {
			return true
		}
	}
	used[i][j] = false
	return false
}

func exist(board [][]byte, word string) bool {
	var used [][]bool
	m := len(board)
	n := len(board[0])
	for i := 0; i < m; i++ {
		used = append(used, make([]bool, n))
	}
	wordBytes := []byte(word)
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if board[i][j] == wordBytes[0] {
				if ok := searchWord(i, j, 0, m, n, board, used, wordBytes); ok {
					return true
				}
			}
		}
	}
	return false
}
