package main

func reorderList(head *ListNode) {
	var nodes []*ListNode
	for node := head; node != nil; node = node.Next {
		nodes = append(nodes, node)
	}
	for i, j := 0, len(nodes)-1; i <= j; i, j = i+1, j-1 {
		if i+1 >= j {
			nodes[j].Next = nil
			break
		}
		nodes[i].Next, nodes[j].Next = nodes[j], nodes[i].Next
	}
}
