package main

type Node struct {
	equation string
	val      float64
}

func dfs(src, dst string, graph map[string][]Node, visited map[string]bool) float64 {
	for _, node := range graph[src] {
		if node.equation == dst {
			return node.val
		}
		if visited[node.equation] {
			continue
		}
		visited[node.equation] = true
		if r := dfs(node.equation, dst, graph, visited); r != 0 {
			return node.val * r
		}
	}
	return 0
}

func calcEquation(equations [][]string, values []float64, queries [][]string) []float64 {
	graph := make(map[string][]Node)
	for i := 0; i < len(equations); i++ {
		a, b := equations[i][0], equations[i][1]
		graph[a] = append(graph[a], Node{equation: b, val: values[i]})
		graph[b] = append(graph[b], Node{equation: a, val: 1 / values[i]})
	}

	var result []float64
	for _, query := range queries {
		a, b := query[0], query[1]
		if len(graph[a]) == 0 || len(graph[b]) == 0 {
			result = append(result, -1)
			continue
		}
		if a == b {
			result = append(result, 1)
			continue
		}
		visited := make(map[string]bool)
		visited[a] = true
		r := dfs(a, b, graph, visited)
		if r == 0 {
			result = append(result, -1)
		} else {
			result = append(result, r)
		}
	}
	return result
}
