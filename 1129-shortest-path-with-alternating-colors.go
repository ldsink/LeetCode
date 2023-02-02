package main

func shortestAlternatingPaths(n int, redEdges [][]int, blueEdges [][]int) []int {
	redNext, blueNext := make(map[int][]int), make(map[int][]int)
	for _, edge := range redEdges {
		redNext[edge[0]] = append(redNext[edge[0]], edge[1])
	}
	for _, edge := range blueEdges {
		blueNext[edge[0]] = append(blueNext[edge[0]], edge[1])
	}
	blueDistance, redDistance := make([]int, n), make([]int, n)
	for i := 1; i < n; i++ {
		blueDistance[i], redDistance[i] = -1, -1
	}
	const BLUE = 1
	const RED = 0
	// 使用整数的最后一位代表预期颜色，从哪条线来当前点，认为当前点是什么颜色
	queue := []int{0, 1}
	for inQueue := make(map[int]bool); len(queue) > 0; queue = queue[1:] {
		node := queue[0]
		delete(inQueue, queue[0])
		color := node & 1
		node >>= 1
		next, distance, nextColor, nextDistance := blueNext, redDistance, BLUE, blueDistance
		if color == BLUE {
			next, distance, nextColor, nextDistance = redNext, blueDistance, RED, redDistance
		}
		for _, n := range next[node] {
			if nextDistance[n] == -1 || nextDistance[n] > distance[node]+1 {
				nextDistance[n] = distance[node] + 1
				key := (n << 1) ^ nextColor
				if !inQueue[key] {
					inQueue[key] = true
					queue = append(queue, key)
				}
			}
		}
	}
	for i := 1; i < n; i++ {
		if blueDistance[i] == -1 || (redDistance[i] != -1 && blueDistance[i] > redDistance[i]) {
			blueDistance[i] = redDistance[i]
		}
	}
	return blueDistance
}
