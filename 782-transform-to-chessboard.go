package main

import (
	"math/bits"
)

func movesToChessboard(board [][]int) int {
	// 2 <= n <= 30，将 board 压缩成 int32 存储
	n := len(board)
	var rowBoard, colBoard []uint32
	for i := 0; i < n; i++ {
		rowBoard = make([]uint32, n)
		colBoard = make([]uint32, n)
	}
	for j := 0; j < n; j++ {
		for i := 0; i < n; i++ {
			rowBoard[j] ^= uint32(board[j][i]) << i
			colBoard[j] ^= uint32(board[i][j]) << i
		}
	}

	abs := func(a int) int {
		if a >= 0 {
			return a
		}
		return -a
	}

	min := func(a, b int) int {
		if a > b {
			return b
		}
		return a
	}

	getReverse := func(x uint32) uint32 {
		return ^x & (1<<n - 1)
	}

	rowBits, colBits := rowBoard[0], colBoard[0] // 第一行第一列的目标情况
	reverseRowBits, reverseColBits := getReverse(rowBits), getReverse(colBits)

	// 检查不可能的情况，顺便确定第一列和第一行
	isImpossible := func() bool {
		// n 是偶数，0, 1 数量相同，否则相差 1，且多的作为 row / col，少的作为 reverse
		if n%2 == 0 {
			if bits.OnesCount32(rowBits) != n/2 || bits.OnesCount32(colBits) != n/2 {
				return true
			}
			var rowCount, colCount int
			for i := 0; i < n; i++ {
				// 行的情况
				if rowBoard[i] == rowBits {
					rowCount++
				} else if rowBoard[i] != reverseRowBits {
					return true
				}
				// 列的情况
				if colBoard[i] == colBits {
					colCount++
				} else if colBoard[i] != reverseColBits {
					return true
				}
			}
			if rowCount != n>>1 || colCount != n>>1 {
				return true
			}
		} else { // 奇数的情况，多的作为 row / col，少的作为 reverse
			if abs(bits.OnesCount32(rowBits)-n/2) > 1 || abs(bits.OnesCount32(colBits)-n/2) > 1 {
				return true
			}

			var rowCount, colCount int
			for i := 0; i < n; i++ {
				// 行的情况
				if rowBoard[i] == rowBits {
					rowCount++
				} else if rowBoard[i] != reverseRowBits {
					return true
				}
				// 列的情况
				if colBoard[i] == colBits {
					colCount++
				} else if colBoard[i] != reverseColBits {
					return true
				}
			}
			// 多的作为 row / col 的目标情况
			if rowCount == n>>1 {
				rowCount = n - rowCount
				rowBits, reverseRowBits = reverseRowBits, rowBits
			}
			if colCount == n>>1 {
				colCount = n - colCount
				colBits, reverseColBits = reverseColBits, colBits
			}
			if rowCount-1 != n>>1 || colCount-1 != n>>1 {
				return true
			}
		}
		return false
	}

	getMoves := func(target uint32) int {
		if n%2 == 0 {
			a := n>>1 - bits.OnesCount32(target&0xAAAAAAAA)
			b := n>>1 - bits.OnesCount32(target&0x55555555)
			return min(a, b)
		} else {
			if bits.OnesCount32(target) == n>>1 {
				return n>>1 - bits.OnesCount32(target&0xAAAAAAAA)
			} else {
				return (n+1)>>1 - bits.OnesCount32(target&0x55555555)
			}
		}
	}

	if isImpossible() {
		return -1
	}
	return getMoves(rowBits) + getMoves(colBits)
}
