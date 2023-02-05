package main

func minimumMoves(grid [][]int) int {
	// 使用 15 个 bit 位表示一个状态，前 14 位表示头在的格子 (row, col)，最后 1 位表示当前姿态
	const bitMask = 1<<7 - 1
	const horizontal = 0
	const vertical = 1
	const emptyCell = 0
	// 处理状态转化的几个工具函数
	getState := func(x, y, h int) int {
		return x<<8 ^ y<<1 ^ h
	}
	getXYH := func(s int) (int, int, int) {
		return (s >> 8) & bitMask, (s >> 1) & bitMask, s & 1
	}
	// 求最小的移动次数
	source := getState(0, 1, horizontal)
	target := getState(len(grid)-1, len(grid)-1, horizontal)
	moves := make(map[int]int)
	getMove := func(s int) int {
		if move, ok := moves[s]; ok {
			return move
		}
		return -1
	}
	moves[source] = 0
	inQueue := make(map[int]bool)
	queue := []int{source}
	extendState := func(s, base int) {
		if m := getMove(s); m == -1 || m > moves[base]+1 {
			moves[s] = moves[base] + 1
			if !inQueue[s] {
				inQueue[s] = true
				queue = append(queue, s)
			}
		}
	}
	for ; len(queue) > 0; queue = queue[1:] {
		current := queue[0]
		delete(inQueue, current)
		row, col, position := getXYH(current)
		if position == horizontal {
			if col+1 < len(grid) && grid[row][col+1] == emptyCell {
				extendState(getState(row, col+1, position), current)
			}
			if row+1 < len(grid) && grid[row+1][col-1] == emptyCell && grid[row+1][col] == emptyCell {
				extendState(getState(row+1, col, horizontal), current) // 水平往下
				extendState(getState(row+1, col-1, vertical), current) // 水平转垂直
			}
		} else {
			if row+1 < len(grid) && grid[row+1][col] == emptyCell {
				extendState(getState(row+1, col, position), current)
			}
			if col+1 < len(grid) && grid[row-1][col+1] == emptyCell && grid[row][col+1] == emptyCell {
				extendState(getState(row, col+1, vertical), current)     // 垂直往右
				extendState(getState(row-1, col+1, horizontal), current) // 垂直转水平
			}
		}
	}
	return getMove(target)
}
