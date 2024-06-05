package main

func countPairsOfConnectableServers(edges [][]int, signalSpeed int) []int {
	n := len(edges) + 1
	type Server struct {
		node, weight int
	}
	linked := make(map[int][]Server, n)
	for _, edge := range edges {
		a, b, weight := edge[0], edge[1], edge[2]
		linked[a] = append(linked[a], Server{node: b, weight: weight})
		linked[b] = append(linked[b], Server{node: a, weight: weight})
	}

	result := make([]int, n)
	for i, next := range linked {
		if len(next) == 1 {
			continue
		}
		var count, sum int
		var dfs func(int, int, int)
		dfs = func(node, parent, weight int) {
			if weight%signalSpeed == 0 {
				count++
			}
			for _, server := range linked[node] {
				if server.node != parent {
					dfs(server.node, node, weight+server.weight)
				}
			}
		}
		for _, server := range next {
			count = 0
			dfs(server.node, i, server.weight)
			result[i] += count * sum
			sum += count
		}
	}
	return result
}
