package main

func possibleBipartition(n int, dislikes [][]int) bool {
	color := make([]int, n+1)
	linked := make([][]int, n+1)
	for _, edge := range dislikes {
		linked[edge[0]] = append(linked[edge[0]], edge[1])
		linked[edge[1]] = append(linked[edge[1]], edge[0])
	}
	for people := 1; people <= n; people++ {
		if color[people] != 0 {
			continue
		}
		queue := []int{people}
		color[people] = people // 以某个圈子中最小的编号作为颜色，以后如果要列举出一种合法的可能性，也能做
		for i := 0; i < len(queue); i++ {
			for _, next := range linked[queue[i]] {
				if color[next] == 0 { // 还不在队列中，添加并染色
					color[next] = -color[queue[i]]
					queue = append(queue, next)
				} else if color[next] == color[queue[i]] { // 已经在队列中，并且已经有了冲突
					return false
				}
			}
		}
	}
	return true
}
