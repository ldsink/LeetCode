package main

func gridIllumination(n int, lamps [][]int, queries [][]int) []int {
	grid := make(map[[2]int]bool)
	row := make(map[int]int)
	col := make(map[int]int)
	leftTopRightBottom := make(map[int]int)
	leftBottomRightTop := make(map[int]int)

	// 初始化灯的点亮状态
	for _, lamp := range lamps {
		x, y := lamp[0], lamp[1]
		if !grid[[2]int{x, y}] {
			grid[[2]int{x, y}] = true
			row[x]++
			col[y]++
			leftTopRightBottom[x-y]++
			leftBottomRightTop[x+y]++
		}
	}

	// 开始查新并更新状态
	var result []int
	directions := [][2]int{{-1, -1}, {-1, 0}, {-1, 1}, {0, -1}, {0, 1}, {1, -1}, {1, 0}, {1, 1}, {0, 0}}
	for _, query := range queries {
		x, y := query[0], query[1]
		// 先确定灯是否是亮着的
		if row[x] > 0 || col[y] > 0 || leftTopRightBottom[x-y] > 0 || leftBottomRightTop[x+y] > 0 {
			result = append(result, 1)
		} else {
			result = append(result, 0)
		}
		// 关灯，并更新状态
		for _, direction := range directions {
			p := [2]int{x + direction[0], y + direction[1]}
			if grid[p] { // 如果这个位置灯是亮着的，关灯并且更新计数
				delete(grid, p) // 关灯
				row[p[0]]--
				col[p[1]]--
				leftTopRightBottom[p[0]-p[1]]--
				leftBottomRightTop[p[0]+p[1]]--
			}
		}
	}
	return result
}
