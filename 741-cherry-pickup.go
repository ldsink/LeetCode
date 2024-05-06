package main

func cherryPickup(grid [][]int) int {
	const CHERRY = 1
	const THORN = -1 // 障碍物或者死路

	n := len(grid)
	if grid[n-1][n-1] == THORN || grid[0][0] == THORN {
		return 0
	}
	result := make(map[int]int)
	result[(n-1)<<24|(n-1)<<16|(n-1)<<8|(n-1)] = grid[n-1][n-1] // 初始化终点数据
	var pick func(int, int, int, int) int
	pick = func(x1, y1, x2, y2 int) int {
		k := (x1 << 24) | (y1 << 16) | (x2 << 8) | y2
		if v, ok := result[k]; ok {
			return v
		}
		var current, next int
		// 计算当前位置的樱桃数
		if grid[x1][y1] == CHERRY {
			current++
		}
		if grid[x2][y2] == CHERRY {
			current++
		}
		if x1 == x2 && y1 == y2 {
			current >>= 1 // 两个人在同一个位置，只能拿一次，除以 2
		}
		// 计算下个位置的樱桃数
		next = THORN
		var next1, next2 [][2]int
		for _, d := range [][2]int{{1, 0}, {0, 1}} {
			x1n, y1n := x1+d[0], y1+d[1]
			if x1n < n && y1n < n && grid[x1n][y1n] != THORN {
				next1 = append(next1, [2]int{x1n, y1n})
			}
			x2n, y2n := x2+d[0], y2+d[1]
			if x2n < n && y2n < n && grid[x2n][y2n] != THORN {
				next2 = append(next2, [2]int{x2n, y2n})
			}
		}
		for _, n1 := range next1 {
			for _, n2 := range next2 {
				if c := pick(n1[0], n1[1], n2[0], n2[1]); next < c {
					next = c
				}
			}
		}
		if next == THORN { // 死路
			result[k] = THORN
		} else {
			result[k] = current + next
		}
		return result[k]
	}
	if v := pick(0, 0, 0, 0); v != -1 {
		return v
	}
	return 0
}
