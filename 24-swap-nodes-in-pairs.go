package main

func swapPairs(head *ListNode) *ListNode {
	var result, previous *ListNode
	var node1, node2 *ListNode
	for current := head; current != nil; current = (*current).Next {
		if node1 == nil {
			node1 = current
			continue
		}
		node2 = current
		(*node1).Next = (*node2).Next
		(*node2).Next = node1
		current = node1
		node1 = nil
		if result == nil {
			result = node2
		}
		if previous == nil {
			previous = current
		} else {
			(*previous).Next = node2
			previous = current
		}
	}
	if result == nil {
		result = head
	}
	return result
}
