package main

var nums = [9]byte{'1', '2', '3', '4', '5', '6', '7', '8', '9'}

func searchSudoku(board *[][]byte, verticalUsed, horizonUsed, gridUsed *[9]map[byte]bool, filled int, needFilled [][2]int) bool {
	if filled == len(needFilled) {
		return true
	}

	i, j := needFilled[filled][0], needFilled[filled][1]
	used := make(map[byte]bool)
	for k, _ := range (*horizonUsed)[i] {
		used[k] = true
	}
	for k, _ := range (*verticalUsed)[j] {
		used[k] = true
	}
	gridIndex := (j/3)*3 + (i / 3)
	for k, _ := range (*gridUsed)[gridIndex] {
		used[k] = true
	}
	if len(used) == 9 {
		return false
	}
	for _, num := range nums {
		if used[num] {
			continue
		}
		// mark num as used
		(*board)[i][j] = num
		(*horizonUsed)[i][num] = true
		(*verticalUsed)[j][num] = true
		(*gridUsed)[gridIndex][num] = true
		// try search result
		success := searchSudoku(board, verticalUsed, horizonUsed, gridUsed, filled+1, needFilled)
		if success {
			return success
		}
		// reset if not success
		(*board)[i][j] = byte('.')
		delete((*horizonUsed)[i], num)
		delete((*verticalUsed)[j], num)
		delete((*gridUsed)[gridIndex], num)
	}
	return false
}

func solveSudoku(board [][]byte) {
	var verticalNums, horizonNums, gridNums [9]map[byte]bool
	var needFilled [][2]int
	for i := 0; i < 9; i++ {
		verticalNums[i] = make(map[byte]bool)
		horizonNums[i] = make(map[byte]bool)
		gridNums[i] = make(map[byte]bool)
	}
	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			if board[i][j] == byte('.') {
				needFilled = append(needFilled, [2]int{i, j})
				continue
			}
			horizonNums[i][board[i][j]] = true
			verticalNums[j][board[i][j]] = true
			gridIndex := (j/3)*3 + (i / 3)
			gridNums[gridIndex][board[i][j]] = true
		}
	}

	// search other result
	searchSudoku(&board, &verticalNums, &horizonNums, &gridNums, 0, needFilled)
}
