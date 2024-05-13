package main

func orangesRotting(grid [][]int) int {
	m, n := len(grid), len(grid[0])
	offset := 8
	mask := (1 << offset) - 1
	oranges := make(map[int]bool)
	var current, next []int
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] == 1 {
				oranges[i<<offset|j] = true
			} else if grid[i][j] == 2 {
				next = append(next, i<<offset|j)
			}
		}
	}
	if len(oranges) == 0 {
		return 0
	}
	var minute int
	directions := [4][2]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}
	for minute = 0; len(next) != 0; minute++ {
		current = next
		next = []int{}
		for _, p := range current {
			x, y := p>>offset, p&mask
			for _, d := range directions {
				xt, yt := x+d[0], y+d[1]
				if 0 <= xt && xt <= m && 0 <= yt && yt <= n {
					pt := xt<<offset | yt
					if _, ok := oranges[pt]; ok {
						delete(oranges, pt)
						next = append(next, pt)
					}
				}
			}
		}
	}
	if len(oranges) > 0 {
		return -1
	}
	return minute - 1
}
