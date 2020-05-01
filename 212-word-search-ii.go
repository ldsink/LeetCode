package main

func searchWord(word string, x, y int, board [][]byte, used [][]bool) bool {
	if len(word) == 1 {
		return true
	}

	used[x][y] = true
	success := false
	if x-1 >= 0 && board[x-1][y] == word[1] && !used[x-1][y] && searchWord(word[1:], x-1, y, board, used) {
		success = true
	} else if y-1 >= 0 && board[x][y-1] == word[1] && !used[x][y-1] && searchWord(word[1:], x, y-1, board, used) {
		success = true
	} else if x+1 < len(board) && board[x+1][y] == word[1] && !used[x+1][y] && searchWord(word[1:], x+1, y, board, used) {
		success = true
	} else if y+1 < len(board[0]) && board[x][y+1] == word[1] && !used[x][y+1] && searchWord(word[1:], x, y+1, board, used) {
		success = true
	}
	used[x][y] = false
	return success
}

func findWords(board [][]byte, words []string) []string {
	var m, n int
	if m = len(board); m == 0 {
		return []string{}
	} else if n = len(board[0]); n == 0 {
		return []string{}
	}

	position := make(map[byte][][]int)
	used := make([][]bool, m)
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			position[board[i][j]] = append(position[board[i][j]], []int{i, j})
		}
		used[i] = make([]bool, n)
	}

	var result []string
	for _, word := range words {
		for _, p := range position[word[0]] {
			if searchWord(word, p[0], p[1], board, used) {
				result = append(result, word)
				break
			}
		}
	}
	return result
}
