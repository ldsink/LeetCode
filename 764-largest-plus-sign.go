package main

import (
	"math"
)

func orderOfLargestPlusSign(n int, mines [][]int) int {
	if len(mines) == n*n {
		return 0
	}
	var grid [][]bool
	for i := 0; i < n; i++ {
		grid = append(grid, make([]bool, n))
	}
	for _, mine := range mines {
		x, y := mine[0], mine[1]
		grid[x][y] = true
	}
	var count [][][4]int // top, right, bottom, left
	for i := 0; i < n; i++ {
		count = append(count, make([][4]int, n))
	}
	for i := 0; i < n; i++ {
		// left to right
		for j, k := 0, 0; j < n; j++ {
			count[j][i][3] = k
			if grid[j][i] {
				k = 0
			} else {
				k++
			}
		}
		// right to left
		for j, k := n-1, 0; j >= 0; j-- {
			count[j][i][1] = k
			if grid[j][i] {
				k = 0
			} else {
				k++
			}
		}
		// top to bottom
		for j, k := 0, 0; j < n; j++ {
			count[i][j][0] = k
			if grid[i][j] {
				k = 0
			} else {
				k++
			}
		}
		// bottom to top
		for j, k := n-1, 0; j >= 0; j-- {
			count[i][j][2] = k
			if grid[i][j] {
				k = 0
			} else {
				k++
			}
		}
	}
	minEdge := func(c [4]int) int {
		r := math.MaxInt
		for i := 0; i < 4; i++ {
			if r > c[i] {
				r = c[i]
			}
		}
		return r
	}
	var edge int
	for i := edge; i < n-edge; i++ {
		for j := edge; j < n-edge; j++ {
			if !grid[i][j] {
				if r := minEdge(count[i][j]); edge < r {
					edge = r
				}
			}
		}
	}
	return edge + 1
}
