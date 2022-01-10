package main

func isEscapePossible(blocked [][]int, source []int, target []int) bool {
	barriers := make(map[[2]int]bool)
	for _, point := range blocked {
		barriers[[2]int{point[0], point[1]}] = true
	}

	// 删除非连续的落单的点，这类点没法起到阻拦的作用
	// 连续的点的定义为靠边或者相邻 8 个方向存在其他的点
	var points [][2]int
	directions := [][2]int{{-1, -1}, {-1, 0}, {-1, 1}, {0, -1}, {0, 1}, {1, -1}, {1, 0}, {1, 1}}
	maxN := 1000000 - 1
	for _, point := range blocked {
		x, y := point[0], point[1]
		added := false
		if x == 0 || y == 0 || x == maxN || y == maxN { // 靠边的点
			added = true
		} else {
			// 检测 8 个方向，是否有其他的点，存在说明连续
			for _, direction := range directions {
				p := [2]int{x + direction[0], y + direction[1]}
				if barriers[p] {
					added = true
					break
				}
			}
		}
		if added {
			points = append(points, [2]int{x, y})
		} else {
			delete(barriers, [2]int{x, y}) // 独立的点保存没有意义，直接删除
		}
	}

	if len(points) == 0 { // 没有连续的点，那肯定封锁不死。直接返回 true
		return true
	}

	// 假设 N 个点可以围成的最大面积，猜的大一点都可以。下面这个 40000 足够了，实际上应该围不出来。
	area := len(points) * len(points)

	directions = [][2]int{{-1, 0}, {0, -1}, {0, 1}, {1, 0}} // 只能往四个相邻的方向上移动

	// 先处理起点，看起点能不能走出这么多步骤
	steps := [][2]int{{source[0], source[1]}}
	visited := make(map[[2]int]bool)
	visited[[2]int{source[0], source[1]}] = true
	i := 0
	for i = 0; i < len(steps) && i < area; i++ {
		x, y := steps[i][0], steps[i][1]
		for _, direction := range directions {
			if 0 <= x+direction[0] && x+direction[0] <= maxN && 0 <= y+direction[1] && y+direction[1] <= maxN {
				point := [2]int{x + direction[0], y + direction[1]}
				if point[0] == target[0] && point[1] == target[1] { // 直接走到了终点，返回 true
					return true
				}
				if !barriers[point] && !visited[point] {
					visited[point] = true
					steps = append(steps, point)
				}
			}
		}
	}
	if i != area { // 没有到最大步数，说明起点在一个封闭的环里面，并且不包含终点
		return false
	}

	// 再看终点在不在一个封闭的环里面
	steps = [][2]int{{target[0], target[1]}}
	visited = make(map[[2]int]bool)
	visited[[2]int{target[0], target[1]}] = true
	for i = 0; i < len(steps) && i < area; i++ {
		x, y := steps[i][0], steps[i][1]
		for _, direction := range directions {
			if 0 <= x+direction[0] && x+direction[0] <= maxN && 0 <= y+direction[1] && y+direction[1] <= maxN {
				// 这里不用在检查起点了，因为如果封闭在一起，起点遍历的时候肯定可以到终点
				point := [2]int{x + direction[0], y + direction[1]}
				if !barriers[point] && !visited[point] {
					visited[point] = true
					steps = append(steps, point)
				}
			}
		}
	}
	if i != area { // 没有到最大步数，说明终点在一个封闭的环里面，并且不包含起点
		return false
	}
	return true
}
