package main

func maxValue(grid [][]int) int {
	m, n := len(grid), len(grid[0])
	value := make([]int, m*n)
	queue := []int{0}
	value[0] = grid[0][0]
	for ; len(queue) > 0; queue = queue[1:] {
		idx := queue[0]
		if (idx+1)%n != 0 && idx+1 < len(value) && value[idx+1] < value[idx]+grid[idx/n][(idx+1)%n] {
			if value[idx+1] == 0 { // 没有礼物，表示没有进入队列，添加到队列
				queue = append(queue, idx+1)
			}
			value[idx+1] = value[idx] + grid[idx/n][(idx+1)%n]
		}
		if idx+n < len(value) && value[idx+n] < value[idx]+grid[idx/n+1][idx%n] {
			if value[idx+n] == 0 { // 没有礼物，表示没有进入队列，添加到队列
				queue = append(queue, idx+n)
			}
			value[idx+n] = value[idx] + grid[idx/n+1][idx%n]
		}
	}
	return value[len(value)-1]
}
