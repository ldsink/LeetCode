package main

func shortestPathAllKeys(grid []string) int {
	const EMPTY = '.'
	const WALL = '#'
	const START = '@'
	// 获得起点、钥匙及锁的位置
	var x, y int
	keys := make(map[uint8][2]int)
	hasKey := make(map[uint8]bool)
	var keyOrder []uint8
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[0]); j++ {
			if grid[i][j] == START {
				x, y = i, j // 重置起点为 EMPTY
			} else if 'a' <= grid[i][j] && grid[i][j] <= 'z' {
				keys[grid[i][j]-'a'] = [2]int{i, j}
				keyOrder = append(keyOrder, grid[i][j]-'a')
			}
		}
	}
	directions := [4][2]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}
	// 某个位置的点是否是可以通过的
	canWalking := func(r uint8) bool {
		if r == WALL { // 墙不可通过
			return false
		} else if 'A' <= r && r <= 'Z' && !hasKey[r-'A'] { // 是锁且当前没有钥匙
			return false
		}
		return true
	}
	// 获得某个字符的最短距离
	getMinSteps := func(x, y int, r uint8) int {
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
				if grid[i][j] == r {
					return step + 1
				} else if canWalking(grid[i][j]) {
					visited[[2]int{i, j}] = true
					queue = append(queue, [3]int{i, j, step + 1})
				}
			}
		}
		return -1
	}
	// 从起点开始，求最短路径
	var search func(int, int, int) int
	search = func(x, y, idx int) int {
		if idx == len(keyOrder)-1 {
			return getMinSteps(x, y, 'a'+keyOrder[idx])
		}
		steps := -1
		for i := idx; i < len(keyOrder); i++ {
			// 设置当前层级第一优先寻找的 key
			keyOrder[i], keyOrder[idx] = keyOrder[idx], keyOrder[i]
			k := keyOrder[idx]
			if r := getMinSteps(x, y, 'a'+k); r != -1 {
				if steps == -1 || steps > r { // 当前解法一定不是最优解，直接跳过
					hasKey[k] = true
					if next := search(keys[k][0], keys[k][1], idx+1); next != -1 {
						if steps == -1 || steps > next+r { // 更新最优解
							steps = next + r
						}
					}
					delete(hasKey, k)
				}
			}
			// 还原当前层级第一优先寻找的 key
			keyOrder[i], keyOrder[idx] = keyOrder[idx], keyOrder[i]
		}
		return steps
	}
	return search(x, y, 0)
}
