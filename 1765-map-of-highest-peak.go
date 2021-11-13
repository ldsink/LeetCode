package main

func highestPeak(isWater [][]int) [][]int {
	m := len(isWater)
	n := len(isWater[0])

	const undefined = -1

	// init height
	var height [][]int
	var stack [][2]int
	for i := 0; i < m; i++ {
		height = append(height, make([]int, n))
		for j := 0; j < n; j++ {
			if isWater[i][j] == 1 {
				stack = append(stack, [2]int{i, j})
			} else {
				height[i][j] = undefined
			}
		}
	}

	for len(stack) > 0 {
		x, y := stack[0][0], stack[0][1]
		stack = stack[1:]
		// top
		if x-1 >= 0 && height[x-1][y] == undefined {
			height[x-1][y] = height[x][y] + 1
			stack = append(stack, [2]int{x - 1, y})
		}
		// right
		if y+1 < n && height[x][y+1] == undefined {
			height[x][y+1] = height[x][y] + 1
			stack = append(stack, [2]int{x, y + 1})
		}
		// bottom
		if x+1 < m && height[x+1][y] == undefined {
			height[x+1][y] = height[x][y] + 1
			stack = append(stack, [2]int{x + 1, y})
		}
		// left
		if y-1 >= 0 && height[x][y-1] == undefined {
			height[x][y-1] = height[x][y] + 1
			stack = append(stack, [2]int{x, y - 1})
		}
	}
	return height
}
