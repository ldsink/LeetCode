package main

func secondMinimum(n int, edges [][]int, time int, change int) int {
	distance := make([]int, n+1)
	secondDistance := make([]int, n+1)
	secondDistance[1] = -1
	for i := 2; i <= n; i++ {
		distance[i] = -1
		secondDistance[i] = -1
	}

	linked := make([][]int, n+1)
	for _, edge := range edges {
		linked[edge[0]] = append(linked[edge[0]], edge[1])
		linked[edge[1]] = append(linked[edge[1]], edge[0])
	}

	nodeSet := make(map[int]bool)
	nodeSet[1] = true // 增加起点
	queue := []int{1}
	for len(queue) > 0 {
		// 取出本次出发的节点
		curr := queue[0]
		queue = queue[1:]
		delete(nodeSet, curr)
		// 计算当前节点第一次可以出发的时间
		first := distance[curr]
		if (first/change)%2 == 1 { // 当前是红灯，需要等到绿灯
			first += change - (first % change)
		}
		// 计算当前节点第二次可以出发的时间
		second := -1
		if secondDistance[curr] != -1 {
			second = secondDistance[curr]
			if (second/change)%2 == 1 { // 当前是红灯，需要等到绿灯
				second += change - (second % change)
			}
		}
		// 尝试更新相邻的节点
		for _, next := range linked[curr] {
			// 尝试用第一次的距离更新下一个节点
			if distance[next] == -1 || distance[next] > first+time { // 相邻的节点可以更新最短距离
				secondDistance[next] = distance[next]
				distance[next] = first + time
				if !nodeSet[next] {
					nodeSet[next] = true
					queue = append(queue, next)
				}
			} else if distance[next] != first+time && (secondDistance[next] == -1 || secondDistance[next] > first+time) { // 相邻的节点可以更新第二短距离
				secondDistance[next] = first + time
				if !nodeSet[next] {
					nodeSet[next] = true
					queue = append(queue, next)
				}
			}
			// 尝试用第二次的距离更新下一个节点
			if second != -1 && distance[next] != second+time && (secondDistance[next] == -1 || secondDistance[next] > second+time) { // 相邻的节点可以更新第二短距离
				secondDistance[next] = second + time
				if !nodeSet[next] {
					nodeSet[next] = true
					queue = append(queue, next)
				}
			}
		}
	}
	return secondDistance[n]
}
