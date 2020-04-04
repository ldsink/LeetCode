package main

func cloneGraph(node *Node) *Node {
	if node == nil {
		return nil
	}
	graph := make(map[int][]int)
	stack := []*Node{node}
	graph[node.Val] = []int{}
	for i := 0; i < len(stack); i++ {
		var neighbors []int
		for _, neighbor := range stack[i].Neighbors {
			neighbors = append(neighbors, neighbor.Val)
			if _, ok := graph[neighbor.Val]; !ok {
				stack = append(stack, neighbor)
				graph[neighbor.Val] = []int{}
			}
		}
		graph[stack[i].Val] = neighbors
	}

	// create nodes
	nodes := make([]Node, len(graph)+1)
	for i := 1; i <= len(graph); i++ {
		nodes[i] = Node{Val: i}
	}
	// add neighbors
	for i := 1; i <= len(graph); i++ {
		var neighbors []*Node
		for _, val := range graph[i] {
			neighbors = append(neighbors, &nodes[val])
		}
		nodes[i].Neighbors = neighbors
	}
	return &nodes[1]
}
