package main

func maxProductPath(grid [][]int) int {
	m := len(grid)
	n := len(grid[0])
	mod := 1000000000 + 7
	nonNegative := make([]int, m*n) // 某个位置最大的非负数
	negative := make([]int, m*n)    // 某个位置最小的负数
	for i := 0; i < len(nonNegative); i++ {
		nonNegative[i] = -1
	}
	// 初始化第一行
	for i := 0; i < n; i++ {
		if i == 0 {
			if grid[0][0] >= 0 {
				nonNegative[0] = grid[0][0]
			} else {
				negative[0] = grid[0][0]
			}
		} else {
			if grid[0][i] >= 0 {
				if nonNegative[i-1] >= 0 {
					nonNegative[i] = nonNegative[i-1] * grid[0][i]
				}
				if negative[i-1] < 0 {
					negative[i] = negative[i-1] * grid[0][i]
				}
			} else {
				if nonNegative[i-1] >= 0 {
					negative[i] = nonNegative[i-1] * grid[0][i]
				}
				if negative[i-1] < 0 {
					nonNegative[i] = negative[i-1] * grid[0][i]
				}
			}
		}
	}
	// 后面的每一行数据
	for i := 1; i < m; i++ {
		for j := 0; j < n; j++ {
			idx := i*n + j
			if grid[i][j] > 0 {
				// 从上一行更新
				if nonNegative[idx-n] >= 0 {
					nonNegative[idx] = nonNegative[idx-n] * grid[i][j]
				}
				if negative[idx-n] < 0 {
					negative[idx] = negative[idx-n] * grid[i][j]
				}
				// 从左侧更新
				if j > 0 {
					if nonNegative[idx-1] >= 0 && nonNegative[idx] < nonNegative[idx-1]*grid[i][j] {
						nonNegative[idx] = nonNegative[idx-1] * grid[i][j]
					}
					if negative[idx-1] < 0 && negative[idx] > negative[idx-1]*grid[i][j] {
						negative[idx] = negative[idx-1] * grid[i][j]
					}
				}
			} else if grid[i][j] < 0 {
				if nonNegative[idx-n] >= 0 {
					negative[idx] = nonNegative[idx-n] * grid[i][j]
				}
				if negative[idx-n] < 0 {
					nonNegative[idx] = negative[idx-n] * grid[i][j]
				}
				if j > 0 {
					if nonNegative[idx-1] >= 0 && negative[idx] > nonNegative[idx-1]*grid[i][j] {
						negative[idx] = nonNegative[idx-1] * grid[i][j]
					}
					if negative[idx-1] < 0 && nonNegative[idx] < negative[idx-1]*grid[i][j] {
						nonNegative[idx] = negative[idx-1] * grid[i][j]
					}
				}
			} else {
				nonNegative[idx] = 0
			}
		}
	}
	if nonNegative[len(nonNegative)-1] != -1 { // 有非负积，需要取模
		nonNegative[len(nonNegative)-1] %= mod
	} else { // 判断数字中是否包含 0, 这样如果找不到最优解，至少有个 0
		hasZero := false
		for i := 0; i < m && !hasZero; i++ {
			for j := 0; j < n && !hasZero; j++ {
				if grid[i][j] == 0 {
					hasZero = true
				}
			}
		}
		if hasZero {
			nonNegative[len(nonNegative)-1] = 0
		}
	}
	return nonNegative[len(nonNegative)-1]
}
