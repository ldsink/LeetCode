package main

func reachableNodes(n int, edges [][]int, restricted []int) int {
	linked := make([][]int, n)
	for _, edge := range edges {
		a, b := edge[0], edge[1]
		linked[a] = append(linked[a], b)
		linked[b] = append(linked[b], a)
	}
	visited := make([]bool, n)
	visited[0] = true
	for _, node := range restricted {
		visited[node] = true
	}
	result := 1
	queue := []int{0}
	for i := 0; i < len(queue); i++ {
		for _, node := range linked[queue[i]] {
			if visited[node] {
				continue
			}
			visited[node] = true
			queue = append(queue, node)
			result++
		}
	}
	return result
}
