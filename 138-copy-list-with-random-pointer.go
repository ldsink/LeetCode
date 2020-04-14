package main

func copyRandomList(head *Node) *Node {
	if head == nil {
		return nil
	}
	nodeMap := make(map[*Node]int)
	var random []*Node
	var value []int
	for node := head; node != nil; node = node.Next {
		value = append(value, node.Val)
		random = append(random, node.Random)
		nodeMap[node] = len(value) - 1
	}

	var copyNodes []Node
	for i := 0; i < len(value); i++ {
		copyNodes = append(copyNodes, Node{Val: value[i]})
	}
	for i := 0; i < len(value)-1; i++ {
		copyNodes[i].Next = &copyNodes[i+1]
		if random[i] != nil {
			copyNodes[i].Random = &copyNodes[nodeMap[random[i]]]
		}
	}

	if random[len(random)-1] != nil {
		copyNodes[len(random)-1].Random = &copyNodes[nodeMap[random[len(random)-1]]]
	}

	return &copyNodes[0]
}
