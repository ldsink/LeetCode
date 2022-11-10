package main

func shortestPathAllKeys(grid []string) int {
	g := make([][]rune, len(grid))
	for i := 0; i < len(grid); i++ {
		g[i] = []rune(grid[i])
	}
	const EMPTY = '.'
	const WALL = '#'
	const START = '@'
	// 获得起点、钥匙及锁的位置
	var x, y int
	keys, locks := make(map[rune][2]int), make(map[rune][2]int)
	hasKey := make(map[rune]bool)
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[0]); j++ {
			if g[i][j] == START {
				x, y, g[i][j] = i, j, EMPTY // 重置起点为 EMPTY
			} else if 'a' <= g[i][j] && g[i][j] <= 'z' {
				keys[g[i][j]-'a'] = [2]int{i, j}
			} else if 'A' <= g[i][j] && g[i][j] <= 'Z' {
				locks[g[i][j]-'A'] = [2]int{i, j}
			}
		}
	}
	directions := [4][2]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}
	// 某个位置的点是否是可以通过的
	canWalking := func(r rune) bool {
		if r == WALL { // 墙不可通过
			return false
		} else if 'A' <= r && r <= 'Z' && !hasKey[r-'A'] { // 是锁且当前没有钥匙
			return false
		}
		return true
	}
	// 获得某个字符的最短距离
	getMinSteps := func(x, y int, r rune) int {
		visited := make(map[[2]int]bool)
		queue := [][3]int{{x, y, 0}}
		visited[[2]int{x, y}] = true
		for len(queue) > 0 {
			x, y, step := queue[0][0], queue[0][1], queue[0][2]
			queue = queue[1:]
			for _, d := range directions {
				i, j := x+d[0], y+d[1]
				if i < 0 || i == len(grid) || j < 0 || j == len(grid[0]) || visited[[2]int{i, j}] {
					continue
				}
				if g[i][j] == r {
					return step + 1
				} else if canWalking(g[i][j]) {
					visited[[2]int{i, j}] = true
					queue = append(queue, [3]int{i, j, step + 1})
				}
			}
		}
		return -1
	}
	// 从起点开始，求最短路径
	var search func(int, int) int
	search = func(x, y int) int {
		var array []rune
		for k := range keys {
			array = append(array, k)
		}
		if len(array) == 1 { // 只剩最后一把钥匙的时候，只需要获得钥匙
			return getMinSteps(x, y, 'a'+array[0])
		}
		steps := -1
		for _, k := range array {
			if r := getMinSteps(x, y, 'a'+k); r != -1 {
				if steps != -1 && steps <= r { // 当前解法一定不是最优解，直接跳过
					continue
				}
				hasKey[k] = true
				keyPosition := keys[k]
				delete(keys, k)
				g[keyPosition[0]][keyPosition[1]] = EMPTY
				if next := search(keyPosition[0], keyPosition[1]); next != -1 {
					if steps == -1 || steps > next+r { // 更新最优解
						steps = next + r
					}
				}
				delete(hasKey, k)
				keys[k] = keyPosition
				g[keyPosition[0]][keyPosition[1]] = 'a' + k
			}
		}
		return steps
	}
	return search(x, y)
}
