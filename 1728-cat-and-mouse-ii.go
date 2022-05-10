package main

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

const Wall = '#'
const Board = '.'

func canJump(grid [][]rune, a, b, c, d, maxJump int) bool {
	if !(a == c || b == d) { // 不在同样的 X、Y 轴上
		return false
	} else if abs(a-c+b-d) > maxJump { // 跳跃的距离超过 maxJump
		return false
	}
	if a == c {
		for i, j := min(b, d), max(b, d); i < j; i++ {
			if grid[a][i] == Wall {
				return false
			}
		}
		return true
	} else {
		for i, j := min(a, c), max(a, c); i < j; i++ {
			if grid[i][b] == Wall {
				return false
			}
		}
		return true
	}
}

func isMouseMustWin(grid [][]rune, catJump, mouseJump, turns, catX, catY, mouseX, mouseY, foodX, foodY int, cache map[int]bool) bool {
	// 如果老鼠不能在 1000 次操作以内到达食物，那么猫获胜。
	if turns > 66 { // 1000 会超时，没必要那么大。
		return false
	}
	// 如果猫跟老鼠处在相同的位置，那么猫获胜。
	if catX == mouseX && catY == mouseY {
		return false
	}

	key := (catX << 24) | (catY << 16) | (mouseX << 8) | mouseY
	key = (key << 10) | turns      // 2^10=1024 肯定不超过10位
	if win, ok := cache[key]; ok { // 检查缓存，避免重复搜索
		return win
	}
	cache[key] = false // 初始化值，避免死循环。先初始化成失败，再寻找成功的机会

	if turns&1 == 1 { // 猫的回合
		// 如果猫可以跳到老鼠或者食物，那么猫获胜。
		if canJump(grid, catX, catY, mouseX, mouseY, catJump) || canJump(grid, catX, catY, foodX, foodY, catJump) {
			return false
		}
		// 否则往各个可以跳的方向走，只要不是全都老鼠获胜，那就是猫获胜
		if !isMouseMustWin(grid, catJump, mouseJump, turns+1, catX, catY, mouseX, mouseY, foodX, foodY, cache) {
			return false
		}
		// 上边
		for i := catX - 1; i >= 0 && grid[i][catY] == Board && catX-i <= catJump; i-- {
			if !isMouseMustWin(grid, catJump, mouseJump, turns+1, i, catY, mouseX, mouseY, foodX, foodY, cache) {
				return false
			}
		}
		// 右边
		for i := catY + 1; i < len(grid[0]) && grid[catX][i] == Board && i-catY <= catJump; i++ {
			if !isMouseMustWin(grid, catJump, mouseJump, turns+1, catX, i, mouseX, mouseY, foodX, foodY, cache) {
				return false
			}
		}
		// 下边
		for i := catX + 1; i < len(grid) && grid[i][catY] == Board && i-catX <= catJump; i++ {
			if !isMouseMustWin(grid, catJump, mouseJump, turns+1, i, catY, mouseX, mouseY, foodX, foodY, cache) {
				return false
			}
		}
		// 左边
		for i := catY - 1; i >= 0 && grid[catX][i] == Board && catY-i <= catJump; i-- {
			if !isMouseMustWin(grid, catJump, mouseJump, turns+1, catX, i, mouseX, mouseY, foodX, foodY, cache) {
				return false
			}
		}
		cache[key] = true // 怎么走都是老鼠获胜，那就老鼠获胜
		return true
	}
	// 老鼠的回合
	// 如果老鼠先到达食物，那么老鼠获胜。
	if canJump(grid, mouseX, mouseY, foodX, foodY, mouseJump) {
		cache[key] = true
		return true
	}
	// 否则往各个可以跳的方向走，只要有一种必胜，那就是必胜
	if isMouseMustWin(grid, catJump, mouseJump, turns+1, catX, catY, mouseX, mouseY, foodX, foodY, cache) {
		cache[key] = true
		return true
	}
	// 上边
	for i := mouseX - 1; i >= 0 && grid[i][mouseY] == Board && mouseX-i <= mouseJump; i-- {
		if isMouseMustWin(grid, catJump, mouseJump, turns+1, catX, catY, i, mouseY, foodX, foodY, cache) {
			cache[key] = true
			return true
		}
	}
	// 右边
	for i := mouseY + 1; i < len(grid[0]) && grid[mouseX][i] == Board && i-mouseY <= mouseJump; i++ {
		if isMouseMustWin(grid, catJump, mouseJump, turns+1, catX, catY, mouseX, i, foodX, foodY, cache) {
			cache[key] = true
			return true
		}
	}
	// 下边
	for i := mouseX + 1; i < len(grid) && grid[i][mouseY] == Board && i-mouseX <= mouseJump; i++ {
		if isMouseMustWin(grid, catJump, mouseJump, turns+1, catX, catY, i, mouseY, foodX, foodY, cache) {
			cache[key] = true
			return true
		}
	}
	// 左边
	for i := mouseY - 1; i >= 0 && grid[mouseX][i] == Board && mouseY-i <= mouseJump; i-- {
		if isMouseMustWin(grid, catJump, mouseJump, turns+1, catX, catY, mouseX, i, foodX, foodY, cache) {
			cache[key] = true
			return true
		}
	}
	return false
}

func canMouseWin(grid []string, catJump int, mouseJump int) bool {
	var runeGrid [][]rune
	for _, s := range grid {
		runeGrid = append(runeGrid, []rune(s))
	}
	var catX, catY, mouseX, mouseY, foodX, foodY int
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[0]); j++ {
			if runeGrid[i][j] == 'C' {
				catX, catY = i, j
				runeGrid[i][j] = Board
			} else if runeGrid[i][j] == 'M' {
				mouseX, mouseY = i, j
				runeGrid[i][j] = Board
			} else if runeGrid[i][j] == 'F' {
				foodX, foodY = i, j
				runeGrid[i][j] = Board
			}
		}
	}
	cache := make(map[int]bool)
	return isMouseMustWin(runeGrid, catJump, mouseJump, 0, catX, catY, mouseX, mouseY, foodX, foodY, cache)
}
