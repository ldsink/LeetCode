package main

func getIntersectionNode(headA, headB *ListNode) *ListNode {
	nodes := make(map[*ListNode]bool)
	for node := headA; node != nil; node = node.Next {
		nodes[node] = true
	}
	for node := headB; node != nil; node = node.Next {
		if _, ok := nodes[node]; ok {
			return node
		}
	}
	return nil
}
