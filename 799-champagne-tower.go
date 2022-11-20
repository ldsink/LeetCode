package main

func champagneTower(poured int, query_row int, query_glass int) float64 {
	tower := make([][]float64, query_row+1)
	for i := 0; i <= query_row; i++ {
		tower[i] = make([]float64, i+1)
	}
	tower[0][0] = float64(poured)
	queue := [][2]int{{0, 0}}
	for len(queue) > 0 {
		x, y := queue[0][0], queue[0][1]
		queue = queue[1:]
		var next float64
		if tower[x][y] >= 1 {
			next = tower[x][y] - 1
			tower[x][y] = 1
		}
		if next == 0 || x == query_row {
			continue
		}
		tower[x+1][y] += next / 2
		queue = append(queue, [2]int{x + 1, y})
		tower[x+1][y+1] += next / 2
		queue = append(queue, [2]int{x + 1, y + 1})
	}
	if tower[query_row][query_glass] >= 1 {
		return 1
	}
	return tower[query_row][query_glass]
}
