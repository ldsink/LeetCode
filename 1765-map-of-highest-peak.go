package main

func highestPeak(isWater [][]int) [][]int {
	m := len(isWater)
	n := len(isWater[0])

	const undefined = -1

	// init height
	var height [][]int
	var queue [][2]int
	for i := 0; i < m; i++ {
		height = append(height, make([]int, n))
		for j := 0; j < n; j++ {
			if isWater[i][j] == 1 {
				queue = append(queue, [2]int{i, j})
			} else {
				height[i][j] = undefined
			}
		}
	}

	for len(queue) > 0 {
		x, y := queue[0][0], queue[0][1]
		queue = queue[1:]
		// top
		if x-1 >= 0 && height[x-1][y] == undefined {
			height[x-1][y] = height[x][y] + 1
			queue = append(queue, [2]int{x - 1, y})
		}
		// right
		if y+1 < n && height[x][y+1] == undefined {
			height[x][y+1] = height[x][y] + 1
			queue = append(queue, [2]int{x, y + 1})
		}
		// bottom
		if x+1 < m && height[x+1][y] == undefined {
			height[x+1][y] = height[x][y] + 1
			queue = append(queue, [2]int{x + 1, y})
		}
		// left
		if y-1 >= 0 && height[x][y-1] == undefined {
			height[x][y-1] = height[x][y] + 1
			queue = append(queue, [2]int{x, y - 1})
		}
	}
	return height
}
