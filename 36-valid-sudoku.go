package main

func isValidSudoku(board [][]byte) bool {
	dot := byte('.')
	// valid horizon
	for i := 0; i < 9; i++ {
		nums := make(map[byte]int)
		for j := 0; j < 9; j++ {
			if board[i][j] == dot {
				continue
			}
			nums[board[i][j]] += 1
			if nums[board[i][j]] > 1 {
				return false
			}
		}
	}
	// valid vertical
	for i := 0; i < 9; i++ {
		nums := make(map[byte]int)
		for j := 0; j < 9; j++ {
			if board[j][i] == dot {
				continue
			}
			nums[board[j][i]] += 1
			if nums[board[j][i]] > 1 {
				return false
			}
		}
	}
	// valid gird
	for i := 0; i < 9; i += 3 {
		for j := 0; j < 9; j += 3 {
			nums := make(map[byte]int)
			for k := i; k < i+3; k++ {
				for l := j; l < j+3; l++ {
					if board[k][l] == dot {
						continue
					}
					nums[board[k][l]] += 1
					if nums[board[k][l]] > 1 {
						return false
					}
				}
			}
		}
	}
	return true
}
