package main

func canMeasureWater(x int, y int, z int) bool {
	if z > x+y {
		return false
	}

	searched := make(map[[2]int]bool)
	searched[[2]int{0, 0}] = true
	queue := [][2]int{{0, 0}}
	for i := 0; i < len(queue); i++ {
		if queue[i][0]+queue[i][1] == z {
			return true
		}
		// 装满任意一个水壶
		if queue[i][0] == 0 {
			state := [2]int{x, queue[i][1]}
			if _, ok := searched[state]; !ok {
				searched[state] = true
				queue = append(queue, state)
			}
		}
		if queue[i][1] == 0 {
			state := [2]int{queue[i][0], y}
			if _, ok := searched[state]; !ok {
				searched[state] = true
				queue = append(queue, state)
			}
		}
		// 清空任意一个水壶
		if queue[i][0] != 0 {
			state := [2]int{0, queue[i][1]}
			if _, ok := searched[state]; !ok {
				searched[state] = true
				queue = append(queue, state)
			}
		}
		if queue[i][1] != 0 {
			state := [2]int{queue[i][0], 0}
			if _, ok := searched[state]; !ok {
				searched[state] = true
				queue = append(queue, state)
			}
		}
		// 从一个水壶向另外一个水壶倒水，直到装满或者倒空
		// 从 x 往 y 倒水
		if queue[i][0] != 0 {
			water := queue[i][0]
			if water > y-queue[i][1] {
				water = y - queue[i][1]
			}
			state := [2]int{queue[i][0] - water, queue[i][1] + water}
			if _, ok := searched[state]; !ok {
				searched[state] = true
				queue = append(queue, state)
			}
		}
		// 从 y 往 x 倒水
		if queue[i][1] != 0 {
			water := queue[i][1]
			if water > x-queue[i][0] {
				water = x - queue[i][0]
			}
			state := [2]int{queue[i][0] + water, queue[i][1] - water}
			if _, ok := searched[state]; !ok {
				searched[state] = true
				queue = append(queue, state)
			}
		}
	}
	return false
}
