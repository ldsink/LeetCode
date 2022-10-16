package main

func possibleBipartition(n int, dislikes [][]int) bool {
	color := make([]int, n+1)
	linked := make(map[int][]int)
	for _, edge := range dislikes {
		a, b := edge[0], edge[1]
		linked[a] = append(linked[a], b)
		linked[b] = append(linked[b], a)
	}
	for i := 1; i <= n; i++ {
		if color[i] != 0 {
			continue
		}
		queue := []int{i}
		color[i] = i // 以某个圈子中最小的编号作为颜色，以后如果要列举出一种合法的可能性，也能做
		for ; len(queue) > 0; queue = queue[1:] {
			for _, next := range linked[queue[0]] {
				if color[next] == 0 { // 还不在队列中，添加并染色
					color[next] = -color[queue[0]]
					queue = append(queue, next)
				} else if color[next] == color[queue[0]] { // 已经在队列中，并且已经有了冲突
					return false
				}
			}
		}
	}
	return true
}
