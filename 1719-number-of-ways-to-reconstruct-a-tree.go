package main

func getWays(root int, ancestors map[int]bool, linked map[int][]int, connected map[[2]int]bool) int {
	// 当前节点到祖先节点必须都是联通的
	for ancestor, _ := range ancestors {
		if !connected[[2]int{root, ancestor}] {
			return 0
		}
	}
	// 当前节点放到祖先节点里面
	ancestors[root] = true
	// 拆分成 N 个合法的子集。
	var subSets [][]int
	inSet := make(map[int]bool)
	ways := 1 // 假设至少可以凑一种
	for _, node := range linked[root] {
		// 祖先节点或者已经加入集合的节点跳过
		if ancestors[node] || inSet[node] {
			continue
		}
		// 基于当前点，划出一个集合
		currentSet := make(map[int]bool)
		currentSet[node], inSet[node] = true, true
		queue := []int{node}
		for len(queue) > 0 {
			curr := queue[0]
			queue = queue[1:]
			for _, a := range linked[curr] {
				if ancestors[a] || currentSet[a] { // 是祖先节点 或者 已经在当前集合，跳过
					continue
				}
				if inSet[a] { // 不是祖先节点，也不在当前集合。那说明访问到了其他节点的数据，这样成环了，失败，ways 置 0
					ways = 0
					break
				}
				currentSet[a], inSet[a] = true, true
				queue = append(queue, a)
			}
		}
		if ways == 0 { // 当前集合构造失败，停止检查
			break
		}
		var subSet []int
		for node, _ := range currentSet {
			subSet = append(subSet, node)
		}
		subSets = append(subSets, subSet)
	}
	// 前面分组的地方都没出错，那递归检查，每个子组是不是都合法的。
	if ways != 0 {
		if len(subSets) == 0 { // 当 root 就是叶子节点了，那连接数必须合法
			if len(linked[root]) != len(ancestors)-1 { // ancestors 包含自己，所以需要 -1
				ways = 0
			}
		} else { // 看不同子集合的情况
			for _, subSet := range subSets {
				subWays := 0 // 子组的方案数，每个组都需要检查
				// 找到集合中节点最大的连接数
				var max int
				for _, r := range subSet {
					if max < len(linked[r]) {
						max = len(linked[r])
					}
				}
				for _, r := range subSet {
					if len(linked[r]) != max { // 子节点只能是当前组连接最多的点
						continue
					}
					subWays += getWays(r, ancestors, linked, connected) // 以 r 为这个子组的根节点
					if subWays > 1 {
						break // 这一组已经超过 1 种方案，终止
					}
				}
				if subWays == 0 { // 这个组合没法构造方法，全局就无法构造方法
					ways = 0
					break
				}
				ways *= subWays
				if ways > 1 {
					ways = 2
				}
			}
		}
	}
	delete(ancestors, root)
	return ways
}

func checkWays(pairs [][]int) int {
	// 统计各个点的连接情况
	linked := make(map[int][]int)
	visited := make(map[int]bool)
	connected := make(map[[2]int]bool)
	var start int
	for _, pair := range pairs {
		a, b := pair[0], pair[1]
		visited[a], visited[b], start = true, true, b
		linked[a] = append(linked[a], b)
		linked[b] = append(linked[b], a)
		connected[[2]int{a, b}] = true
		connected[[2]int{b, a}] = true
	}

	// 先检查各个点一定可以连接到，连接不到的返回 0
	queue := []int{start}
	delete(visited, start)
	for len(queue) > 0 {
		node := queue[0]
		queue = queue[1:]
		for _, next := range linked[node] {
			if visited[next] {
				delete(visited, next)
				queue = append(queue, next)
			}
		}
	}
	if len(visited) != 0 {
		return 0
	}

	// 合法的根节点肯定是连接数最多的点
	var max int
	for _, edges := range linked {
		if max < len(edges) {
			max = len(edges)
		}
	}
	var ways int
	ancestors := make(map[int]bool)
	for root, edges := range linked {
		if len(edges) != max {
			continue
		}
		ways += getWays(root, ancestors, linked, connected) // 尝试基于该节点作为根节点进行构造，增加方法数量
		if ways > 1 {
			ways = 2 // 如果 ways > 1 ，返回 2
			break
		}
	}
	return ways
}
