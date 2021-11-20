package main

type Edge [2]int

func dfs(node, parent, depth int, visited []int, nodeEdges [][]int, edges map[Edge]bool) int {
	// 已经访问过
	if visited[node] >= 0 {
		return visited[node]
	}
	visited[node] = depth
	minDepth := len(nodeEdges)
	for _, child := range nodeEdges[node] {
		if child == parent {
			continue
		}
		childDepth := dfs(child, node, depth+1, visited, nodeEdges, edges)
		// 子节点深度更小，已成环
		if childDepth <= depth {
			a, b := node, child
			if a > b {
				a, b = b, a
			}
			delete(edges, Edge{a, b})
		}
		if minDepth > childDepth {
			minDepth = childDepth
		}
	}
	return minDepth
}

func criticalConnections(n int, connections [][]int) [][]int {
	nodeEdges := make([][]int, n) // 节点的边
	edges := make(map[Edge]bool)
	for _, conn := range connections {
		// 确保边按照顺序排序
		a, b := conn[0], conn[1]
		if a > b {
			a, b = b, a
		}
		nodeEdges[a] = append(nodeEdges[a], b)
		nodeEdges[b] = append(nodeEdges[b], a)
		edges[Edge{a, b}] = true
	}

	visited := make([]int, n)
	for i := 0; i < n; i++ {
		visited[i] = -1
	}
	dfs(0, -1, 0, visited, nodeEdges, edges)

	var criticalEdges [][]int
	for edge, _ := range edges {
		criticalEdges = append(criticalEdges, []int{edge[0], edge[1]})
	}
	return criticalEdges
}
