package main

type dict struct {
	word     string
	count    int
	children [26]*dict
}

func makeTrie(words []string) *dict {
	root := &dict{}
	for _, word := range words {
		root := root
		for _, r := range word {
			key := r - 'a'
			if root.children[key] == nil {
				root.children[key] = &dict{}
				root.count++
			}
			root = root.children[key]
		}
		root.word = word
	}
	return root
}

var directions = [4][2]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}

func search(x, y int, board [][]byte, root *dict, result *[]string) {
	r := board[x][y]
	node := root.children[r-'a']
	if node == nil {
		return
	}

	// add current word
	if node.word != "" {
		*result = append(*result, node.word)
		node.word = ""
		// prune
		if node.count == 0 {
			root.children[r-'a'] = nil
			root.count--
			return
		}
	}

	board[x][y] = '.'
	for _, d := range directions {
		i, j := x+d[0], y+d[1]
		if i < 0 || i == len(board) || j < 0 || j == len(board[0]) {
			continue
		}
		if board[i][j] == '.' {
			continue
		}
		search(i, j, board, node, result)
		if node.count == 0 {
			root.children[r-'a'] = nil
			root.count--
			break
		}
		if root.count == 0 {
			break
		}
	}
	board[x][y] = r
}

func findWords(board [][]byte, words []string) []string {
	var result []string
	if len(words) == 0 || len(board) == 0 || len(board[0]) == 0 {
		return result
	}
	root := makeTrie(words)
	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[0]); j++ {
			search(i, j, board, root, &result)
			if root.count == 0 {
				return result
			}
		}
	}
	return result
}
